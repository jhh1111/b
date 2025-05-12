package core

import (
	"sync"
)

// SimpleQueue 是Queue接口的简单实现，使用切片存储URL
type SimpleQueue struct {
	urls []*URL
	mu   sync.Mutex
}

// NewSimpleQueue 创建一个新的简单队列
func NewSimpleQueue() *SimpleQueue {
	return &SimpleQueue{
		urls: make([]*URL, 0),
	}
}

// Push 添加URL到队列
func (q *SimpleQueue) Push(url *URL) {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.urls = append(q.urls, url)
}

// Pop 从队列获取下一个URL
func (q *SimpleQueue) Pop() (*URL, bool) {
	q.mu.Lock()
	defer q.mu.Unlock()

	if len(q.urls) == 0 {
		return nil, false
	}

	// 从队列头部获取URL
	url := q.urls[0]
	q.urls = q.urls[1:]

	return url, true
}

// Len 返回队列长度
func (q *SimpleQueue) Len() int {
	q.mu.Lock()
	defer q.mu.Unlock()

	return len(q.urls)
}

// Clear 清空队列
func (q *SimpleQueue) Clear() {
	q.mu.Lock()
	defer q.mu.Unlock()

	q.urls = make([]*URL, 0)
}
