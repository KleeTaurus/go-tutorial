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
