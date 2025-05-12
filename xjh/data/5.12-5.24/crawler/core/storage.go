package core

import (
	"encoding/json"
	"fmt"
	"os"
	"sync"
)

// MemoryStorage 是一个内存存储实现
type MemoryStorage struct {
	data map[string][]Result
	mu   sync.RWMutex
}

// NewMemoryStorage 创建一个新的内存存储
func NewMemoryStorage() *MemoryStorage {
	return &MemoryStorage{
		data: make(map[string][]Result),
	}
}

// Store 存储URL及其关联的结果
func (s *MemoryStorage) Store(url string, results []Result) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data[url] = results
	return nil
}

// Get 获取URL对应的结果
func (s *MemoryStorage) Get(url string) ([]Result, bool) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	results, ok := s.data[url]
	return results, ok
}

// GetAll 获取所有结果
func (s *MemoryStorage) GetAll() map[string][]Result {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// 创建一个副本
	result := make(map[string][]Result)
	for url, data := range s.data {
		result[url] = data
	}

	return result
}

// Clear 清空存储
func (s *MemoryStorage) Clear() {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.data = make(map[string][]Result)
}

// SaveToFile 将结果保存到文件
func (s *MemoryStorage) SaveToFile(filename string) error {
	s.mu.RLock()
	defer s.mu.RUnlock()

	// 创建文件
	file, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("创建文件失败: %w", err)
	}
	defer file.Close()

	// 将结果编码为JSON
	encoder := json.NewEncoder(file)
	encoder.SetIndent("", "  ")
	err = encoder.Encode(s.data)
	if err != nil {
		return fmt.Errorf("编码数据失败: %w", err)
	}

	return nil
}

// LoadFromFile 从文件加载结果
func (s *MemoryStorage) LoadFromFile(filename string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	// 打开文件
	file, err := os.Open(filename)
	if err != nil {
		return fmt.Errorf("打开文件失败: %w", err)
	}
	defer file.Close()

	// 解码JSON
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&s.data)
	if err != nil {
		return fmt.Errorf("解码数据失败: %w", err)
	}

	return nil
}

// SimpleChecker 是一个简单的URL去重器实现
type SimpleChecker struct {
	urls map[string]bool
	mu   sync.RWMutex
}

// NewSimpleChecker 创建一个新的简单去重器
func NewSimpleChecker() *SimpleChecker {
	return &SimpleChecker{
		urls: make(map[string]bool),
	}
}

// IsDuplicate 检查URL是否已存在
func (c *SimpleChecker) IsDuplicate(url string) bool {
	c.mu.RLock()
	defer c.mu.RUnlock()

	return c.urls[url]
}

// MarkAsDuplicate 标记URL为已存在
func (c *SimpleChecker) MarkAsDuplicate(url string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.urls[url] = true
}

// Clear 清空去重器
func (c *SimpleChecker) Clear() {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.urls = make(map[string]bool)
}
