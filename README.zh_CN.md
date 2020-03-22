# cyflag

## README.md
- [English](README.md)
- [中文](README.zh_CN.md)

*cyflag* 是一个用于parse命令行参数中flag的库，相比Go标准库中的flag，它更易于使用，并且限制更少。

## 依赖

- [cyDS](https://github.com/yah01/cyDS-GO): 用Go实现的一些数据结构， *cyflag* 使用其中的trie来快速匹配flag。

## 用法

*cyflag* 看上去很像Go标准库中的[flag](https://golang.org/pkg/flag/)，用法也相近。

### 将变量与flag绑定
*cyflag* 可以目前可以绑定三种类型的变量：
- bool
- int
- string

绑定的方法都差不多，例如对于布尔类型变量：
> cyflag.BoolVar(&varible,"-boolflag",false,"it's a bool flag")

函数的四个参数分别是：
1. 变量地址
2. flag名
3. 默认值
4. 用法提示

***flag* 名不必以字符'-'开头**, 这一点与标准库中的 *flag* 不同，后者会自动在flag前添加字符'-'。

### Parse
**在所有绑定完成后**, 调用
> cyflag.Parse()

*cyflag* 将会parse命令行参数，并且把值存储到绑定的变量中。

**注意:** 参数格式的规则与标准库中 *flag* 有一些不同：
- *flagname*: parse bool类型变量的唯一方法
- *flagname value*: parse 非bool类型变量的唯一方法

命令行参数中不属于flag及其值的参数，将会存储到 *cyflag.Args* 中。

***cyflag* 对参数和flag的位置顺序没有要求**，这一点与标准库中 *flag*不同，后者要求所有flag及其值必须在non-flag参数之前。

## 意外情况
*cyflag.Parse()* 会返回一个error类型值，有两类情况：
1. 没有错误，此时返回nil
2. 对于int型flag，提供的值不能转换为int
3. 对于非bool型flag，没有提供值
