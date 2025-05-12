package core

import (
	"context"
	"fmt"
	"sync"
	"time"
)

// Engine 爬虫引擎，负责协调整个爬取过程
type Engine struct {
	// 配置选项
	options *Options

	// URL队列
	queue Queue

	// 页面抓取器
	fetcher Fetcher

	// 页面解析器
	parser Parser

	// 结果存储
	storage Storage

	// 去重器
	duplicateChecker DuplicateChecker

	// 计数器
	stats *Stats

	// 控制
	ctx    context.Context
	cancel context.CancelFunc
	wg     sync.WaitGroup
}

// Options 爬虫引擎配置选项
type Options struct {
	// 最大爬取深度
	MaxDepth int

	// 最大并发数
	Concurrency int

	// 爬取超时时间
	Timeout time.Duration

	// 每个请求的超时时间
	RequestTimeout time.Duration

	// 请求间隔
	RequestDelay time.Duration

	// 是否遵守robots.txt规则
	RespectRobotsTxt bool

	// 自定义请求头
	Headers map[string]string

	// 是否启用Cookie
	EnableCookies bool
}

// Stats 爬虫统计信息
type Stats struct {
	// 爬取的URL数量
	URLsProcessed int64

	// 成功爬取的页面数
	PagesSucceeded int64

	// 失败的页面数
	PagesFailed int64

	// 发现的URL总数
	URLsFound int64

	// 最近的错误
	LastError error

	// 锁，保护上述字段
	mu sync.RWMutex
}

// NewEngine 创建一个新的爬虫引擎
func NewEngine(options *Options) *Engine {
	ctx, cancel := context.WithCancel(context.Background())

	// 使用默认配置
	if options == nil {
		options = &Options{
			MaxDepth:         5,
			Concurrency:      10,
			Timeout:          10 * time.Minute,
			RequestTimeout:   30 * time.Second,
			RequestDelay:     100 * time.Millisecond,
			RespectRobotsTxt: true,
			Headers:          map[string]string{},
			EnableCookies:    true,
		}
	}

	return &Engine{
		options:          options,
		queue:            NewSimpleQueue(),
		fetcher:          NewHTTPFetcher(options.RequestTimeout),
		parser:           NewDefaultParser(),
		storage:          NewMemoryStorage(),
		duplicateChecker: NewSimpleChecker(),
		stats:            &Stats{},
		ctx:              ctx,
		cancel:           cancel,
	}
}

// SetFetcher 设置自定义的页面抓取器
func (e *Engine) SetFetcher(fetcher Fetcher) {
	e.fetcher = fetcher
}

// SetParser 设置自定义的页面解析器
func (e *Engine) SetParser(parser Parser) {
	e.parser = parser
}

// SetQueue 设置自定义的URL队列
func (e *Engine) SetQueue(queue Queue) {
	e.queue = queue
}

// SetStorage 设置自定义的结果存储
func (e *Engine) SetStorage(storage Storage) {
	e.storage = storage
}

// AddURL 添加URL到爬取队列
func (e *Engine) AddURL(url string) {
	e.queue.Push(&URL{
		Address: url,
		Depth:   0,
	})
}

// Start 启动爬虫引擎
func (e *Engine) Start() error {
	// 设置全局超时
	var ctx context.Context
	var cancel context.CancelFunc

	if e.options.Timeout > 0 {
		ctx, cancel = context.WithTimeout(e.ctx, e.options.Timeout)
		defer cancel()
	} else {
		ctx = e.ctx
	}

	// 并发控制
	semaphore := make(chan struct{}, e.options.Concurrency)

	fmt.Printf("爬虫启动，并发数: %d，最大深度: %d\n",
		e.options.Concurrency, e.options.MaxDepth)

	// 记录活动的worker数量
	activeWorkers := 0
	var activeMu sync.Mutex

	// 主循环
	for {
		select {
		case <-ctx.Done():
			fmt.Println("爬虫已超时或被取消")
			e.wg.Wait()
			return ctx.Err()
		default:
			// 获取下一个URL
			url, ok := e.queue.Pop()
			if !ok {
				// 队列为空，但可能有工作尚未完成
				activeMu.Lock()
				active := activeWorkers
				activeMu.Unlock()

				if active == 0 {
					fmt.Println("队列为空，爬取完成")
					return nil
				}
				// 等待工作完成或新的URL
				time.Sleep(100 * time.Millisecond)
				continue
			}

			// 检查深度
			if url.Depth > e.options.MaxDepth {
				continue
			}

			// 检查是否已访问过
			if e.duplicateChecker.IsDuplicate(url.Address) {
				continue
			}

			// 请求限速
			if e.options.RequestDelay > 0 {
				time.Sleep(e.options.RequestDelay)
			}

			// 并发控制
			semaphore <- struct{}{}
			e.wg.Add(1)

			// 更新活动worker计数
			activeMu.Lock()
			activeWorkers++
			activeMu.Unlock()

			// 启动工作协程
			go func(u *URL) {
				defer func() {
					<-semaphore
					e.wg.Done()

					// 减少活动worker计数
					activeMu.Lock()
					activeWorkers--
					activeMu.Unlock()
				}()

				// 标记为已访问
				e.duplicateChecker.MarkAsDuplicate(u.Address)

				// 处理URL
				e.processURL(ctx, u)
			}(url)
		}
	}
}

// Stop 停止爬虫引擎
func (e *Engine) Stop() {
	e.cancel()
	// 等待所有工作完成
	e.wg.Wait()
}

// GetStats 获取当前统计信息
func (e *Engine) GetStats() Stats {
	e.stats.mu.RLock()
	defer e.stats.mu.RUnlock()

	return *e.stats
}

// processURL 处理单个URL
func (e *Engine) processURL(ctx context.Context, url *URL) {
	// 更新统计信息
	e.stats.mu.Lock()
	e.stats.URLsProcessed++
	e.stats.mu.Unlock()

	fmt.Printf("正在处理 [%d] %s\n", url.Depth, url.Address)

	// 抓取页面
	page, err := e.fetcher.Fetch(ctx, url.Address)
	if err != nil {
		e.stats.mu.Lock()
		e.stats.PagesFailed++
		e.stats.LastError = err
		e.stats.mu.Unlock()

		fmt.Printf("抓取失败 %s: %v\n", url.Address, err)
		return
	}

	// 更新统计信息
	e.stats.mu.Lock()
	e.stats.PagesSucceeded++
	e.stats.mu.Unlock()

	// 解析页面
	results, links := e.parser.Parse(page)

	// 存储结果
	if len(results) > 0 {
		e.storage.Store(url.Address, results)
	}

	// 更新统计信息
	e.stats.mu.Lock()
	e.stats.URLsFound += int64(len(links))
	e.stats.mu.Unlock()

	// 将新的链接添加到队列
	newDepth := url.Depth + 1
	if newDepth <= e.options.MaxDepth {
		for _, link := range links {
			e.queue.Push(&URL{
				Address: link,
				Depth:   newDepth,
				Parent:  url.Address,
			})
		}
	}
}

// GetStorage 获取结果存储组件
func (e *Engine) GetStorage() Storage {
	return e.storage
}

// getActiveWorkers 获取当前活动的worker数量
func (e *Engine) getActiveWorkers() int {
	// 这是一个简化的实现，实际上应该使用原子计数器来跟踪活动的worker数量
	return 0 // 此处简化处理，返回0表示没有活动的worker
}
