# go-tutorial

A lot of go examples collected at https://www.golinuxcloud.com/getting-started-with-golang/

## Maps

1. 使用 var 声明 map
2. 使用 map literal 初始化 map
3. 使用 make 函数初始化 map
4. 为 map 添加 key/value
5. 访问 map 中的 value
6. 更新 map 中的 value
7. 删除 map 中的元素
8. 检查 map 中的 key 是否存在
9. 遍历 map 中的元素
10. 顺序遍历 map 中的元素

## Interfaces

1. 定义一个空接口
2. 实现简单的接口
3. 通过接口计算工资之和
4. 将接口 Slice 作为参数，使得代码更简单
5. 实现内置 Error 接口

## Testing

1. 编写单元测试用例
2. 编写性能测试代码
3. 通过设置 timer 提升 Benchmark 的精准度

如何在命令行运行性能测试?

```bash
$ go test -bench . -run notest
$ go test -bench . -benchtime 2s -count 2 -benchmem -cpu 4 -run notest
```

## Type Assertion

1. 使用 ok 语法检查断言状态
2. 使用 switch 语句检查接口类型

## init

1. init() 的基本使用方法
2. 多个 init() 的执行顺序
3. 多个 package 中 init() 的执行顺序

## Generics

1. 未采用 generic 的代码(重复)
2. 使用 generic 简化代码
3. 通过 type constraint 简化类型定义

## Packages

1. 调用 package 中公开的方法或变量
2. 调用自定义 package 中的方法
3. encode 或 decode YAML

## Mutex

1. Race condition 暴露的问题
2. 使用 mutex 锁避免并发写
3. 使用 RWMutex 提高性能
4. 在结构体中定义 mutex

## Time

### Timing

1. 使用 time.Since 和 time.Sub 方法计算时长
2. 使用 defer 语法计算时长（倾向此方法，对代码污染比较小）

### Time Format

1. 使用预定义格式输出时间
2. 使用自定义格式输出时间
3. 从字符串中解析时间
4. 从字符串中解析时间（本地 Location）
5. ParseDuration 方法介绍

## Logging

### Log

1. 修改默认 log 格式
2. 默认支持的 flags 参数及使用方法
3. 修改日期和时间格式
4. 通过 SetPrefix 添加日志级别
5. 自定义日志格式并写入文件
6. 自定义日志格式同时写入 console 和文件
