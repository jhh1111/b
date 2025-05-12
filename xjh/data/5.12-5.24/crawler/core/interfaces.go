package core

import (
	"context"
)

// URL 表示一个待爬取的URL
type URL struct {
	// URL地址
	Address string

	// 爬取深度
	Depth int

	// 父URL
	Parent string

	// 元数据，可存储额外信息
	Metadata map[string]interface{}
}

// Page 表示一个已爬取的页面
type Page struct {
	// URL地址
	URL string

	// 页面标题
	Title string

	// 页面内容
	Content []byte

	// 页面HTTP状态码
	StatusCode int

	// HTTP头信息
	Headers map[string]string

	// 页面编码
	Charset string

	// 爬取时间戳
	Timestamp int64
}

// Result 表示从页面中提取的结果项
type Result struct {
	// 结果类型
	Type string

	// 结果数据
	Data map[string]interface{}
}

// Queue 表示URL队列的接口
type Queue interface {
	// 添加URL到队列
	Push(*URL)

	// 从队列获取下一个URL，如果队列为空，返回nil和false
	Pop() (*URL, bool)

	// 返回队列长度
	Len() int

	// 清空队列
	Clear()
}

// Fetcher 表示页面抓取器的接口
type Fetcher interface {
	// 抓取指定URL的页面
	Fetch(ctx context.Context, url string) (*Page, error)
}

// Parser 表示页面解析器的接口
type Parser interface {
	// 解析页面，返回提取的结果和发现的链接
	Parse(page *Page) ([]Result, []string)
}

// Storage 表示结果存储的接口
type Storage interface {
	// 存储URL及其关联的结果
	Store(url string, results []Result) error

	// 获取指定URL的结果
	Get(url string) ([]Result, bool)

	// 获取所有结果
	GetAll() map[string][]Result

	// 清空存储
	Clear()
}

// DuplicateChecker 表示URL去重器的接口
type DuplicateChecker interface {
	// 检查URL是否已经爬取过
	IsDuplicate(url string) bool

	// 标记URL为已爬取
	MarkAsDuplicate(url string)

	// 清空去重器
	Clear()
}
