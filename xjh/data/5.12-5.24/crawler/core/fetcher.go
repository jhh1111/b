package core

import (
	"context"
	"fmt"
	"io"
	"net/http"
	"time"
)

// HTTPFetcher 是一个基于HTTP的页面抓取器
type HTTPFetcher struct {
	client  *http.Client
	headers map[string]string
}

// NewHTTPFetcher 创建一个新的HTTP抓取器
func NewHTTPFetcher(timeout time.Duration) *HTTPFetcher {
	if timeout <= 0 {
		timeout = 30 * time.Second
	}

	client := &http.Client{
		Timeout: timeout,
	}

	return &HTTPFetcher{
		client:  client,
		headers: make(map[string]string),
	}
}

// SetHeaders 设置HTTP请求头
func (f *HTTPFetcher) SetHeaders(headers map[string]string) {
	f.headers = headers
}

// Fetch 实现Fetcher接口，抓取指定URL的页面
func (f *HTTPFetcher) Fetch(ctx context.Context, url string) (*Page, error) {
	// 创建HTTP请求
	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("创建请求失败: %w", err)
	}

	// 设置通用头信息
	req.Header.Set("User-Agent", "GoCrawler/1.0")

	// 设置自定义头信息
	for k, v := range f.headers {
		req.Header.Set(k, v)
	}

	// 发送请求
	resp, err := f.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("请求失败: %w", err)
	}
	defer resp.Body.Close()

	// 检查状态码
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("HTTP状态码异常: %d", resp.StatusCode)
	}

	// 读取响应体
	content, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("读取响应失败: %w", err)
	}

	// 提取页面标题
	title := extractTitle(content)

	// 获取响应头
	headers := make(map[string]string)
	for k, v := range resp.Header {
		if len(v) > 0 {
			headers[k] = v[0]
		}
	}

	// 创建页面对象
	page := &Page{
		URL:        url,
		Title:      title,
		Content:    content,
		StatusCode: resp.StatusCode,
		Headers:    headers,
		Charset:    detectCharset(resp.Header, content),
		Timestamp:  time.Now().Unix(),
	}

	return page, nil
}

// extractTitle 从HTML内容中提取标题
func extractTitle(content []byte) string {
	// 简单实现，实际应使用HTML解析库
	titleStart := []byte("<title>")
	titleEnd := []byte("</title>")

	startIdx := bytesIndex(content, titleStart)
	if startIdx == -1 {
		return ""
	}
	startIdx += len(titleStart)

	endIdx := bytesIndex(content[startIdx:], titleEnd)
	if endIdx == -1 {
		return ""
	}

	return string(content[startIdx : startIdx+endIdx])
}

// bytesIndex 查找子切片在切片中的位置
func bytesIndex(s, sep []byte) int {
	n := len(sep)
	if n == 0 {
		return 0
	}
	if n > len(s) {
		return -1
	}

	for i := 0; i <= len(s)-n; i++ {
		if bytesEqual(s[i:i+n], sep) {
			return i
		}
	}
	return -1
}

// bytesEqual 比较两个字节切片是否相等
func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

// detectCharset 检测页面编码
func detectCharset(header http.Header, content []byte) string {
	// 从HTTP响应头中获取
	contentType := header.Get("Content-Type")
	if contentType != "" {
		// 简单实现，实际应使用更复杂的解析
		if bytesContains([]byte(contentType), []byte("charset=")) {
			start := bytesIndex([]byte(contentType), []byte("charset=")) + 8
			end := len(contentType)
			for i := start; i < len(contentType); i++ {
				if contentType[i] == ';' || contentType[i] == ' ' {
					end = i
					break
				}
			}
			return contentType[start:end]
		}
	}

	// 从HTML内容中获取
	// 简单实现，实际应使用HTML解析库
	meta := []byte("charset=")
	idx := bytesIndex(content, meta)
	if idx != -1 {
		start := idx + len(meta)
		end := start
		for i := start; i < len(content) && i < start+20; i++ {
			if content[i] == '"' || content[i] == '\'' || content[i] == ' ' {
				end = i
				break
			}
		}
		if end > start {
			return string(content[start:end])
		}
	}

	// 默认UTF-8
	return "utf-8"
}

// bytesContains 检查切片是否包含子切片
func bytesContains(s, substr []byte) bool {
	return bytesIndex(s, substr) >= 0
}
