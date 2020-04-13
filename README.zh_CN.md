# cyflag

## README.md
- [English](README.md)
- [中文](README.zh_CN.md)

*cyflag* 是一个用于解析命令行参数（你也可以把它用于其他情况下的参数解析）的库，相比Go标准库中的flag，它更易于使用，并且限制更少。

## 快速上手
```go
var (
		parser cyflag.Parser
		args   = `i love "cyber flag" -best -times 95`

		best  bool
		times int
		love  string
	)

	parser.Bind(&best, "-best", false, "whether the best")
	parser.Bind(&times, "-times", 0, "-times [int]")
	parser.Bind(&love, "love", "something", "love [string]")

	parser.ParseString(args)

	fmt.Printf("best: %+v\n"+
		"times: %+v\n"+
		"love: %+v\n",
		best, times, love)
```
绑定函数的4个参数分别是：
- 变量地址
- flag名称
- 默认值
- flag使用方法提示

上面的例子会输出
```yaml
best: true
times: 95
love: cyber flag
```

## 用法

cyflag 看上去很像Go标准库中的flag包，实际上它们的用法也相似。

### 将变量与flag绑定
cyflag 可以绑定5种类型的变量：
- bool
- int （也可绑定 int8，int16，……）
- uint （也可绑定 uint8，uint16，……）
- float64 （也可绑定 float32）
- string
他们的绑定方法几乎都一样，分别是：
```go
parser.BoolVar(&boolVariable,"-bool",false,"it's a bool flag")
parser.IntVar(&intVariable,"-int",0,"it's a int flag")
parser.UintVar(&uintVariable,"-uint",0,"it's a uint flag")
parser.FloatVar(&float64Variable,"-float64",0.0,"it's a float64 flag")
parser.StringVar(&stringVariable,"-string","empty","it's a string flag")
```

为了方便，cyflag有一个"万能"绑定方法，你可以用它来绑定任意类型的变量（前提是cyflag支持），同时，它也是你绑定非默认类型的唯一方法（例如int64）：
```go
parser.Bind(&variable,"-anytype",defaultValue,"the usage")
```
Bind()方法会在提供的变量类型与默认值类型无法转换时panic，注意，不需要类型完全相同，例如变量类型是int，而默认值类型是int64，这种情况下是可行的。但其中一个是有符号整数，而另一个是无符号整数是不可行的。

flag不一定要以字符'-'开头，这一点与Go标准库中的flag包不同，后者会自动的将字符'-'作为flag的前缀。

### 解析
在所有的绑定都完成后，调用
```go
parser.Parse(args)
```
来解析参数。如果 args == nil ，那么parser会去解析parser.LeftArgs。在解析完成之后，无法解析的参数会存储在parser.LeftArgs中，对于调用cyflag.Parse()进行的解析，无法解析的参数会存储在cyflag.Args中。在一开始，cyflag.Args==os.Args[1:]

**注意**：参数格式的规则：
- *flagname*：这是解析bool类型变量的唯一格式
- *flagname value*：这是解析非bool类型变量的唯一格式
对于字符串类型参数，如果参数中含有空格，应该用引号将整个参数括起来，单引号或双引号都是可行的。

cyflag 对参数的顺序并没有限制，这一点也与Go标准库中的flag包不同，后者会在遇到第一个非flag参数时停止解析，而cyflag会解析所有可以解析的参数

如果是为了解析命令行参数（os.Args[1:]），可以不需要声明一个cyflag.Parser。直接使用cyflag.Bind(...)，cyflag.Parse()即可