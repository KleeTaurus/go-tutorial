# Go 语言

* [Go 官方文档](https://go.dev/doc/)
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
* [等待 goroutines 执行结束](https://yourbasic.org/golang/wait-for-goroutines-waitgroup/)

## Interfaces

## Security

## Testing

## Reflection

## 标准库

[Go 语言的3种排序方式](https://yourbasic.org/golang/how-to-sort-in-go/)

1. 利用标准库函数实现 int, float64 和 string 的排序
2. 自定义 comparator 函数 (一次性排序场景)
3. 自定义类型并实现 sort.Interface 接口 (代码可重复的排序场景)
4. 对 map 进行排序 (前提是将 key 或 value 转化为 slice 先行排序)

# 存储

## MySQL

[github](https://github.com/go-sql-driver/mysql)

## Redis

[github](https://github.com/redis/go-redis)

## Kafka

[github](https://github.com/segmentio/kafka-go)

# 框架

## Gin Web Framework

[github](https://github.com/gin-gonic/gin/)

## GORM

[github](https://github.com/go-gorm/gorm/blob/master/README.md)

# 设计模式

[DESIGN PATTERNS in GO](https://refactoring.guru/design-patterns/go)

## 创建模式

* [单例模式](https://refactoring.guru/design-patterns/singleton/go/example#example-0)

## 行为模式

## 结构模式

# 其他开发资源

* [Git](https://kapeli.com/cheat_sheets/Git.docset/Contents/Resources/Documents/index.html)
* [ASCII Control Codes](https://kapeli.com/cheat_sheets/ASCII_Tables.docset/Contents/Resources/Documents/index)
* [HTTP Header Fields](https://kapeli.com/cheat_sheets/HTTP_Header_Fields.docset/Contents/Resources/Documents/index.html)
* [HTTP Status Codes](https://kapeli.com/cheat_sheets/HTTP_Status_Codes.docset/Contents/Resources/Documents/index.html)
* [regex101](https://regex101.com)

# 代码示例

* [Go by Example](https://gobyexample.com/)
* [yourbasic](https://yourbasic.org/golang/)
* [GeeksforGeeks](https://www.geeksforgeeks.org/golang/?ref=ghm)
* [Learn Web Programming in Go by Examples](https://gowebexamples.com/)
* [Learn Go with tests](https://studygolang.gitbook.io/learn-go-with-tests/)
* [Programming Guide of Go](https://programming.guide/go/)
* [Golang By Example](https://golangbyexample.com/)
