package core

import (
	"regexp"
	"strings"
)

// DefaultParser 是一个基本的HTML解析器
type DefaultParser struct {
	// 链接提取正则表达式
	linkRegex *regexp.Regexp

	// 标题提取正则表达式
	titleRegex *regexp.Regexp

	// 忽略的URL后缀
	ignoreSuffixes []string
}

// NewDefaultParser 创建一个新的默认HTML解析器
func NewDefaultParser() *DefaultParser {
	return &DefaultParser{
		linkRegex:  regexp.MustCompile(`<a\s+[^>]*href="([^"]+)"[^>]*>`),
		titleRegex: regexp.MustCompile(`<title[^>]*>(.*?)</title>`),
		ignoreSuffixes: []string{
			".jpg", ".jpeg", ".png", ".gif", ".pdf", ".zip", ".tar.gz",
			".css", ".js", ".xml", ".json", ".mp3", ".mp4", ".avi", ".mov",
		},
	}
}

// Parse 实现Parser接口，解析HTML页面内容
func (p *DefaultParser) Parse(page *Page) ([]Result, []string) {
	if page == nil || len(page.Content) == 0 {
		return nil, nil
	}

	// 提取链接
	links := p.extractLinks(page)

	// 提取结果
	results := p.extractResults(page)

	return results, links
}

// extractLinks 从页面中提取链接
func (p *DefaultParser) extractLinks(page *Page) []string {
	matches := p.linkRegex.FindAllSubmatch(page.Content, -1)
	uniqueLinks := make(map[string]bool)

	baseURL := page.URL

	for _, match := range matches {
		if len(match) < 2 {
			continue
		}

		href := string(match[1])

		// 忽略空链接
		if href == "" || href == "#" || strings.HasPrefix(href, "javascript:") {
			continue
		}

		// 忽略指定后缀的链接
		if p.hasSuffix(href, p.ignoreSuffixes) {
			continue
		}

		// 处理相对链接
		absURL := p.resolveURL(baseURL, href)

		// 将URL标准化
		absURL = p.normalizeURL(absURL)

		// 保存唯一链接
		uniqueLinks[absURL] = true
	}

	// 转换为切片
	links := make([]string, 0, len(uniqueLinks))
	for link := range uniqueLinks {
		links = append(links, link)
	}

	return links
}

// extractResults 从页面中提取结果
func (p *DefaultParser) extractResults(page *Page) []Result {
	var results []Result

	// 提取标题
	title := page.Title
	if title == "" {
		// 如果页面对象中没有标题，尝试从内容中提取
		titleMatches := p.titleRegex.FindSubmatch(page.Content)
		if len(titleMatches) > 1 {
			title = string(titleMatches[1])
		}
	}

	// 创建基本结果
	if title != "" {
		result := Result{
			Type: "page",
			Data: map[string]interface{}{
				"url":       page.URL,
				"title":     title,
				"length":    len(page.Content),
				"timestamp": page.Timestamp,
			},
		}
		results = append(results, result)
	}

	return results
}

// hasSuffix 检查URL是否有指定的后缀
func (p *DefaultParser) hasSuffix(url string, suffixes []string) bool {
	urlLower := strings.ToLower(url)
	for _, suffix := range suffixes {
		if strings.HasSuffix(urlLower, suffix) {
			return true
		}
	}
	return false
}

// resolveURL 解析相对URL
func (p *DefaultParser) resolveURL(baseURL, href string) string {
	// 如果是绝对URL，直接返回
	if strings.HasPrefix(href, "http://") || strings.HasPrefix(href, "https://") {
		return href
	}

	// 处理锚点
	if strings.HasPrefix(href, "#") {
		return baseURL
	}

	// 确保baseURL是绝对URL
	if !strings.HasPrefix(baseURL, "http://") && !strings.HasPrefix(baseURL, "https://") {
		return href
	}

	// 解析baseURL
	parts := strings.SplitN(baseURL, "://", 2)
	if len(parts) != 2 {
		return href
	}

	scheme := parts[0]
	hostPath := parts[1]

	// 获取主机部分
	hostEnd := strings.IndexAny(hostPath, "/?#")
	var host, path string
	if hostEnd == -1 {
		host = hostPath
		path = "/"
	} else {
		host = hostPath[:hostEnd]
		path = hostPath[hostEnd:]
		if !strings.HasPrefix(path, "/") {
			path = "/" + path
		}
	}

	// 处理不同类型的相对链接
	if strings.HasPrefix(href, "/") {
		// 从根目录开始的相对链接
		return scheme + "://" + host + href
	} else {
		// 从当前路径开始的相对链接
		lastSlash := strings.LastIndex(path, "/")
		if lastSlash != -1 {
			path = path[:lastSlash+1]
		} else {
			path = "/"
		}
		return scheme + "://" + host + path + href
	}
}

// normalizeURL 标准化URL
func (p *DefaultParser) normalizeURL(url string) string {
	// 移除URL中的片段标识符
	fragmentIndex := strings.LastIndex(url, "#")
	if fragmentIndex != -1 {
		url = url[:fragmentIndex]
	}

	// 确保URL不以斜杠结尾
	if strings.HasSuffix(url, "/") {
		url = url[:len(url)-1]
	}

	return url
}
