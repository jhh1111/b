# Go爬虫项目

这是一个使用Go语言实现的网页爬虫系统，专注于高效、可扩展的网页数据采集与处理。

## 项目结构

- **core/** - 爬虫核心组件
  - 爬虫引擎
  - URL队列管理
  - 调度器实现
  
- **utils/** - 工具函数和辅助组件
  - HTTP工具
  - 并发控制
  - 错误处理
  
- **plugins/** - 可插拔组件
  - 自定义Fetcher实现
  - 自定义Parser实现
  - 结果处理器
  
- **examples/** - 示例和演示
  - 基本爬虫示例
  - 本地测试服务器

## 主要功能

- 多协程并发抓取
- 可配置的抓取深度和广度
- 智能URL去重
- 支持自定义的页面解析器
- 定制化数据提取规则
- 灵活的数据存储方式

## 如何运行

### 命令行参数

爬虫程序支持以下命令行参数：

```
  -concurrency int
        并发数 (default 5)
  -delay int
        请求间隔(毫秒) (default 100)
  -depth int
        最大爬取深度 (default 2)
  -output string
        输出文件 (default "results.json")
  -req-timeout int
        请求超时时间(秒) (default 10)
  -robots
        是否遵守robots.txt (default true)
  -timeout int
        总超时时间(秒) (default 30)
  -url string
        起始URL (default "https://go.dev")
```

### 常用命令

```bash
# 基本使用（使用默认参数）
go run main.go

# 指定URL和爬取深度
go run main.go -url=https://example.com -depth=3

# 高并发爬取
go run main.go -url=https://example.com -concurrency=20 -timeout=60

# 运行演示功能
go run main.go demo

# 使用本地测试服务器
cd examples/localserver
go run main.go
# 然后根据提示运行爬虫
```

### 本地测试

项目提供了一个本地测试服务器，可以用来测试爬虫功能而不需要访问外部网站：

1. 启动本地服务器：
   ```bash
   cd examples/localserver
   go run main.go
   ```

2. 服务器会给出运行爬虫的命令，通常是：
   ```bash
   go run ../../main.go -url=http://localhost:8080/ -depth=2 -concurrency=1
   ```

3. 在另一个终端窗口运行该命令进行测试

## 技术特点

- 基于接口的模块化设计
- 高性能并发控制
- 内存友好的数据处理
- 可插拔的组件系统
- 完善的错误处理和重试机制

## 示例代码

### 基本使用

```go
// crawler := NewCrawler(options)
// crawler.AddURL("https://example.com")
// crawler.Start()
```

### 自定义解析器

```go
// parser := func(page *Page) []string {
//    // 解析逻辑，返回新的URL列表
// }
// crawler.SetParser(parser)
```

## 开发路线图

### 阶段1：核心功能实现
- **TODO**: 实现基本的URL队列 (`core/queue.go`)
- **TODO**: 实现爬虫引擎 (`core/engine.go`)
- **TODO**: 实现HTTP抓取功能 (`core/fetcher.go`)
- **TODO**: 实现默认HTML解析器 (`core/parser.go`)

### 阶段2：功能增强
- **TODO**: 添加去重机制 (`utils/deduplicate.go`)
- **TODO**: 实现优先级队列 (`core/priority_queue.go`)
- **TODO**: 添加限速控制 (`utils/rate_limiter.go`)
- **TODO**: 实现robots.txt解析 (`utils/robots.go`)

### 阶段3：插件系统
- **TODO**: 实现代理支持 (`plugins/proxy_fetcher.go`)
- **TODO**: 添加缓存系统 (`plugins/cache_fetcher.go`)
- **TODO**: 开发自定义UA支持 (`plugins/useragent.go`)
- **TODO**: 创建Cookie管理器 (`plugins/cookie_jar.go`)

### 阶段4：数据持久化
- **TODO**: 实现文件存储 (`plugins/file_storage.go`)
- **TODO**: 添加数据库存储选项 (`plugins/db_storage.go`)
- **TODO**: 开发导出功能 (`plugins/exporter.go`)

### 阶段5：高级功能
- **TODO**: 添加分布式支持 (`core/distributed.go`)
- **TODO**: 实现网站地图生成 (`plugins/sitemap.go`)
- **TODO**: 开发API接口 (`plugins/api.go`)
- **TODO**: 添加监控和统计 (`utils/metrics.go`)

## 演示功能

项目中添加了一个演示模式，可以展示Go语言的各种基础特性：

1. 内存对齐
2. 接口类型断言
3. 闭包
4. defer和panic恢复
5. Channel
6. Goroutine正确退出方式

运行演示：
```bash
go run main.go demo
```

## 如何贡献

1. Fork项目
2. 创建功能分支
3. 提交更改
4. 推送到分支
5. 创建Pull Request

## 许可证

MIT License 