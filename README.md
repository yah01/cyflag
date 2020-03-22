# cyflag

## README.md
- [English](README.md)
- [中文](README.zh_CN.md)

*cyflag* is a lib for parsing flag, which being used frequently in CLI(Command-Line Interface). it's easier to use and has less limitation than the go standard lib *flag*.

## Usage

*cyflag* looks like go standard package [flag](https://golang.org/pkg/flag/), as well the usage.

### Sample
```go

```

### Bind variable with flag
*cyflag* could bind 3 kinds of variable:
- bool
- int
- string

They are all similar, for example:
```go
cyflag.BoolVar(&variable,"-boolflag",false,"it's a bool flag")
```

The arguments of the function above are:
1. the address of the variable
2. the flag name
3. default value
4. usage information

**The flag does not have to start with character '-'**, different with go standard lib *flag* that adding '-' automatically.

### Parser Type

### Parse
**after all bindings having finished**, just call
```go
cyflag.Parse()
```

*cyflag* will parse the CLI arguments and store the value into the binding variables.

**Notice:** The arguments format rule is a little different with go standard lib *flag*:
- *flagname*: only way to parse bool variable
- *flagname value*: only way to parse non-bool variable

the *non-flag* arguments will store in *cyflag.Args*

**There are no limitation in the order of arguments and flags**, which is totally different with go standard lib *flag*, the latter parses until meet first non-flag arguments.
