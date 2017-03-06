# learn golang

## IDE
推荐使用`Gogland`或者`IntelliJ IDEA`安装`golang`插件，其他的`sublime text`或者`vs code`都可以，选择自己熟悉的开发工具就OK~

## 代码结构
1. `scorpio`工程是核心代码，与业务相关的所有代码都放置在该工程下面
2. `stardust-go`工程是项目用到到所有公共代码，包括第三方包(包括用`go get`命令安装)都放置在该工程下面

## 安装
1. 进入`scorpio`目录，执行`setup-gopath.sh`脚本，该脚本会把当前目录和`stardustx`项目设置为`gopath`
2. 提交代码之前先手动执行`format-src.sh`脚本，该脚本会用`golang`的标准方式格式化代码
3. 也可以直接在`git`本地的`pre-commit`的`hook`文件中加上上面格式化的脚本，避免每次手动执行

## 运行
1. 可以在`scorpio`根目录中新建一个`bin`文件夹
2. 进入`bin`文件夹，执行命令`go build scorpio/mathapp`
3. 在`bin`目录下面会生成一个`mathapp`的可执行文件，执行命令`./mathapp`即可得到运行结果
4. 也可以直接在`scorpio`根目录下面直接执行命令`go run src/scorpio/mathapp/main.go`得到运行结果


### 指针

代码路径：[pointers.go](https://github.com/cnych/golearn/blob/master/scorpio/src/scorpio/complextype/pointers.go)

`Go` 具有指针。 指针保存了变量的内存地址。

类型 `*T` 是指向类型 `T` 的值的指针。其零值是 `nil`。
```go
    var p *int
```
`&` 符号会生成一个指向其作用对象的指针。
```go
    i := 42
    p = &i
```
`*` 符号表示指针指向的底层的值。
```go
    fmt.Println(*p) // 通过指针 p 读取 i
    *p = 21         // 通过指针 p 设置 i
```
这也就是通常所说的“间接引用”或“非直接引用”。

与 `C` 不同，`Go` 没有指针运算。


### 结构体

代码路径：[structs.go](https://github.com/cnych/golearn/blob/master/scorpio/src/scorpio/complextype/structs.go)

一个结构体（`struct`）就是一个字段的集合。

（而 `type` 的含义跟其字面意思相符。）

结构体字段使用点号来访问。

结构体字段可以通过结构体指针来访问，通过指针间接的访问是透明的。
