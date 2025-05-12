package main

import (
	"example.com/m/xjh/data/5.12-5.24/crawler/core"
	"flag"
	"fmt"
	"log"
	"time"
)

var (
	// 命令行参数
	startURL    = flag.String("url", "https://go.dev", "起始URL")
	depth       = flag.Int("depth", 2, "最大爬取深度")
	concurrency = flag.Int("concurrency", 5, "并发数")
	timeout     = flag.Int("timeout", 30, "总超时时间(秒)")
	reqTimeout  = flag.Int("req-timeout", 10, "请求超时时间(秒)")
	reqDelay    = flag.Int("delay", 100, "请求间隔(毫秒)")
	outputFile  = flag.String("output", "results.json", "输出文件")
	robotsTxt   = flag.Bool("robots", true, "是否遵守robots.txt")
)

// CrawlerResults 存储爬虫结果的结构体
type CrawlerResults struct {
	StartTime    time.Time
	EndTime      time.Time
	VisitedPages int
	TotalLinks   int
}

func main() {
	// 解析命令行参数
	flag.Parse()

	// 创建爬虫选项
	options := &core.Options{
		MaxDepth:         *depth,
		Concurrency:      *concurrency,
		Timeout:          time.Duration(*timeout) * time.Second,
		RequestTimeout:   time.Duration(*reqTimeout) * time.Second,
		RequestDelay:     time.Duration(*reqDelay) * time.Millisecond,
		RespectRobotsTxt: *robotsTxt,
		Headers: map[string]string{
			"User-Agent": "GoCrawler/1.0 (https://example.com/bot)",
		},
	}

	// 创建爬虫引擎
	crawler := core.NewEngine(options)

	// 添加起始URL
	crawler.AddURL(*startURL)

	// 打印爬虫配置信息
	fmt.Println("=== Go爬虫启动 ===")
	fmt.Printf("起始URL: %s\n", *startURL)
	fmt.Printf("最大深度: %d\n", *depth)
	fmt.Printf("并发数: %d\n", *concurrency)
	fmt.Printf("总超时: %d秒\n", *timeout)
	fmt.Printf("请求间隔: %d毫秒\n", *reqDelay)
	fmt.Printf("输出文件: %s\n", *outputFile)
	fmt.Println("==================")

	// 启动爬虫
	fmt.Println("爬虫开始运行...")
	startTime := time.Now()

	err := crawler.Start()
	if err != nil {
		log.Fatalf("爬虫运行出错: %v", err)
	}

	// 获取统计信息
	stats := crawler.GetStats()
	duration := time.Since(startTime)

	// 打印结果
	fmt.Println("\n爬虫运行完成!")
	fmt.Printf("运行时间: %v\n", duration)
	fmt.Printf("处理的URL数: %d\n", stats.URLsProcessed)
	fmt.Printf("成功页面数: %d\n", stats.PagesSucceeded)
	fmt.Printf("失败页面数: %d\n", stats.PagesFailed)
	fmt.Printf("发现的URL数: %d\n", stats.URLsFound)

	// 保存结果
	storage := crawler.GetStorage()
	if memStorage, ok := storage.(*core.MemoryStorage); ok {
		err := memStorage.SaveToFile(*outputFile)
		if err != nil {
			log.Printf("保存结果出错: %v", err)
		} else {
			fmt.Printf("结果已保存到 %s\n", *outputFile)
		}
	}
}
