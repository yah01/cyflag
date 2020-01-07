# cyFlag

## Dependences

- [cyDS](https://github.com/yah01/cyDS-GO): some data structures implements

## Usage

cyFlag looks like go standard package [flag](https://golang.org/pkg/flag/), as well the usage.

### Bind variable with flag
cyFlag could bind 3 kinds of variable:
- bool
- int
- string

They are all similar, for example:
> cyFlag.BoolVar(&varible,"-boolflag",false,"it's a bool flag")

**The flag does not have to start with character '-'**

### Parse
after **all binding having finished**, just call
> cyFlag.Parse()

there is a difference between bool variable parsing and the others:
- *flag*: only way to parse bool variable
- *flag arg*: only way to parse non-bool variable

the *non-flag* arguments will store in *cyFlag.Args*

**There are no limitation in the order of arguments and flags**
