# Go 语言

* [Go 官方文档](https://go.dev/doc/), [标准库](https://pkg.go.dev/std), [Effective Go](https://go.dev/doc/effective_go)
* [Go 开发者 Roadmaps](https://roadmap.sh/golang)

这是一个标准 Go 项目的[目录结构](https://yourbasic.org/golang/library-package-example-template/)

## 开发工具

* [Go 开发工具](https://www.alexedwards.net/blog/an-overview-of-go-tooling)
* [如何使用第三方工具 Delve debug Go 程序](https://golang.cafe/blog/golang-debugging-with-delve.html)

## Basic

* [iota 的使用方法](https://yourbasic.org/golang/iota/)
* [获取对象类型](https://yourbasic.org/golang/find-type-of-object/)
* [读取环境变量](https://yourbasic.org/golang/environment-variables/)

## Arrays, Slices 和 Maps

* [Arrays 的基本用法](https://gobyexample.com/arrays)
* [Slices 的基本用法](https://gobyexample.com/slices), 此处为 Go 官方的 [Slices 详解](https://go.dev/blog/slices-intro)
* [Maps 的基本用法](https://gobyexample.com/maps)
* [两种从 slice 中删除元素的方法](https://yourbasic.org/golang/delete-element-slice/)
* [获取或删除 slice 中的最后一个元素](https://yourbasic.org/golang/last-item-in-slice/)
* [从 slice 中删除重复的 item](https://www.geeksforgeeks.org/how-to-remove-duplicate-values-from-slice-in-golang/)
* [如何对 map 中的 key 或 value 进行排序](https://yourbasic.org/golang/sort-map-keys-values/)

## Strings

* [float 与 string 之间的互转](https://yourbasic.org/golang/convert-string-to-float/)
* [int, int64 与 string 之间的互转](https://yourbasic.org/golang/convert-int-to-string/)
* [将 interface 转换为 string](https://yourbasic.org/golang/interface-to-string/)
* [string 常用操作](https://yourbasic.org/golang/string-functions-reference-cheat-sheet/)

## Files

* [逐行读取文件内容](https://yourbasic.org/golang/read-file-line-by-line/)
* [往文件中追加内容](https://yourbasic.org/golang/append-to-file/)
* [读取指定目录下的文件列表](https://yourbasic.org/golang/list-files-in-directory/)

## Time 和 Date

* [获取系统当前时间](https://yourbasic.org/golang/current-time/)
* [计算代码执行时长](https://yourbasic.org/golang/measure-execution-time/)

## Random numbers

* [生成随机数或字母](https://yourbasic.org/golang/generate-number-random-range/)
* [随机打散 slice 或 array](https://yourbasic.org/golang/shuffle-slice-array/)

## Goroutine 和 Channels

* [什么是 Data races](https://yourbasic.org/golang/data-races-explained/), 如何通过 channel 避免 Data races 的发生
* [如何通过 --race 参数检查代码是否存在 Data races 问题](https://yourbasic.org/golang/detect-data-races/)
* [等待 goroutines 执行结束](https://yourbasic.org/golang/wait-for-goroutines-waitgroup/)
* [mutex 互斥锁的使用方法](https://yourbasic.org/golang/mutex-explained/)
* [What are channels used for?](https://stackoverflow.com/questions/39826692/what-are-channels-used-for)
* [Concurrency in Go with Goroutines and Channel](https://blog.knoldus.com/achieving-concurrency-in-go/)

## Interfaces

## Security

## Testing

* [Testify: easy assertions and mocking](https://github.com/stretchr/testify/)

## Reflection

## 标准库

[Go 语言的3种排序方式](https://yourbasic.org/golang/how-to-sort-in-go/)

1. 利用标准库函数实现 int, float64 和 string 的排序
2. 自定义 comparator 函数 (一次性排序场景)
3. 自定义类型并实现 sort.Interface 接口 (代码可重复的排序场景)
4. 对 map 进行排序 (前提是将 key 或 value 转化为 slice 先行排序)

# 存储

## MySQL

* [go-sql-driver](https://github.com/go-sql-driver/mysql)
* [如何访问关系型数据库](https://go.dev/doc/tutorial/database-access)
* [MySQL 教程](https://www.w3schools.com/MySQL/)

## Redis

* [go-redis](https://github.com/redis/go-redis)
* [Redis 官网文档](https://redis.io/docs/)

## Kafka

* [kafka-go](https://github.com/segmentio/kafka-go)

# 框架

## Gin Web Framework

* [gin-gonic](https://github.com/gin-gonic/gin/) Gin is a web framework written in Go
* [API 文档及示例](https://pkg.go.dev/github.com/gin-gonic/gin)

## GORM

* [GORM](https://github.com/go-gorm/gorm/) The fantastic ORM library for Golang, aims to be developer friendly
* [API 文档及示例](https://gorm.io/docs/)

# 系统设计

[System Design Roadmap](https://roadmap.sh/system-design)

# 设计模式

[DESIGN PATTERNS in GO](https://refactoring.guru/design-patterns/go)

## 创建模式

* [单例模式](https://refactoring.guru/design-patterns/singleton/go/example#example-0)

## 行为模式

## 结构模式

# 其他开发资源

* [Git](https://kapeli.com/cheat_sheets/Git.docset/Contents/Resources/Documents/index.html)
* [Gitflow Workflow](https://www.atlassian.com/git/tutorials/comparing-workflows/gitflow-workflow)
* [ASCII Control Codes](https://kapeli.com/cheat_sheets/ASCII_Tables.docset/Contents/Resources/Documents/index)
* [HTTP Header Fields](https://kapeli.com/cheat_sheets/HTTP_Header_Fields.docset/Contents/Resources/Documents/index.html)
* [HTTP Status Codes](https://kapeli.com/cheat_sheets/HTTP_Status_Codes.docset/Contents/Resources/Documents/index.html)

# 在线工具

* [JSON Formatter & Validator](https://jsonformatter.curiousconcept.com/#) JSON 在线格式化/验证工具
* [Convert CSV to JSON](https://www.convertcsv.com/csv-to-json.htm) 在线 CSV 转 JSON
* [regex101](https://regex101.com) 正则表达式验证工具
* [Epoch Converter](https://www.epochconverter.com/) Epoch & Unix Timestamp 转换工具

# 代码示例

* [Go by Example](https://gobyexample.com/)
* [yourbasic](https://yourbasic.org/golang/)
* [GeeksforGeeks](https://www.geeksforgeeks.org/golang/?ref=ghm)
* [Learn Web Programming in Go by Examples](https://gowebexamples.com/)
* [Learn Go with tests](https://studygolang.gitbook.io/learn-go-with-tests/)
* [Programming Guide of Go](https://programming.guide/go/)
* [Golang By Example](https://golangbyexample.com/)
* [Go tips and snippets](https://freshman.tech/snippets/go/)

# 其他好的文章

* [How To Benchmark HTTP Latency with wrk on Ubuntu 14.04](https://www.digitalocean.com/community/tutorials/how-to-benchmark-http-latency-with-wrk-on-ubuntu-14-04#simulate-advanced-http-requests-with-lua-scripts)
* [Benchmarking in Golang: Improving function performance](https://blog.logrocket.com/benchmarking-golang-improve-function-performance/)
