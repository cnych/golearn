# learn golang

## IDE
推荐使用`Gogland`或者`IntelliJ IDEA`安装`golang`插件，其他的`sublime text`或者`vs code`都可以，选择自己熟悉的开发工具就OK~

## 代码结构
1. `scorpio`工程是核心代码，与业务相关的所有代码都放置在该工程下面
2. `stardustx`工程是项目用到到所有公共代码，包括第三方包(包括用`go get`命令安装)都放置在该工程下面

## 安装
1. 进入`scorpio`目录，执行`setup-gopath.sh`脚本，该脚本会把当前目录和`stardustx`项目设置为`gopath`
2. 提交代码之前先手动执行`format-src.sh`脚本，该脚本会用`golang`的标准方式格式化代码
3. 也可以直接在`git`本地的`pre-commit`的`hook`文件中加上上面格式化的脚本，避免每次手动执行

## 运行
1. 可以在`scorpio`根目录中新建一个`bin`文件夹
2. 进入`bin`文件夹，执行命令`go build scorpio/mathapp`
3. 在`bin`目录下面会生成一个`mathapp`的可执行文件，执行命令`./mathapp`即可得到运行结果
4. 也可以直接在`scorpio`根目录下面直接执行命令`go run src/scorpio/mathapp/main.go`得到运行结果

## 学习

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


### 数组

代码路径：[array.go](https://github.com/cnych/golearn/blob/master/scorpio/src/scorpio/complextype/array.go)

类型`[n]T`是一个`n`个类型为`T`的值的数组。

#### 表达式

```go
    var a [10]int
```
定义变量`a`是一个有10个整数的数组。

数组的长度是其类型的一部分，因此数组不能改变大小。

#### slice

一个`slice`会指向一个序列的值，并且包含了长度信息。

`[]T`是一个元素类型为`T`的`slice`。

#### 对slice切片

`slice`可以重新切片，创建一个新的`slice`值指向相同的数组。

`s[lo:hi]`：表示从`lo`到`hi-1`的`slice`元素，含两端。

因此`s[lo:lo]`是空的，而`s[lo:lo+1]`有一个元素。


#### 构造slice
`slice`由函数`make`创建。这会分配一个零长度的数组并且返回一个`slice`指向这个数组：
```go
    a := make([]int, 5) // len(a)=5
```
为了指定容量，可以传递第三个参数到`make`：
```go
    b := make([]int, 0, 5) // len(b)=0, cap(b)=5

    b = b[:cap(b)]  // len(b)=5, cap(b)=5
    b = b[1:] // len(b)=4, cap(b)=4
```

`slice`的零值是`nil`，一个`nil`的slice长度和容量是`0`。

#### 向slice添加元素
向`slice`添加元素是一种常见的操作，`Go`提供了一个内建的`append`函数。

```go
 func append(s []T, vs ...T) []T
```
`append`的第一个参数`s`是一个类型为`T`的数组，其余类型为`T`的值将会添加到`slice`上。
`append`的结果是一个包含原`slice`所有元素加上新添加的元素的`slice`。
如果`s`的底层数组太小，而不能容纳所有的值时，会分配一个更大的数组。返回的`slice`会指向这个新分配的数组。

参阅`slice`文章：[使用和内幕](http://golang.org/doc/articles/slices_usage_and_internals.html)

## interface
代码路径：[interfaces.go](https://github.com/cnych/golearn/blob/master/scorpio/src/scorpio/interfaceapp/main.go)

接口是什么？`Go`语言没有类和几成的概念。但是`Go`语言有非常灵活的接口概念，通过它可以实现很多面向对象的特性。
接口提供了一种方式来说明对象的行为。
接口定义了一组方法集，但是这些方法不包含（实现代码）：它们没有被实现（它们是抽象的）。接口里也不能包含变量。

## grpc
> 参考资料: [grpc](http://www.jianshu.com/p/774b38306c30) [grpc-api](https://xiequan.info/go%E4%BD%BF%E7%94%A8grpc%E4%B8%8Eprotocol-buffers%E6%9E%84%E5%BB%BA%E9%AB%98%E6%80%A7%E8%83%BDapi-%E6%9C%8D%E5%8A%A1/)

`gRPC`是一个高性能、通用的开源RPC框架，其由`Google`主要面向移动应用开发并基于`HTTP/2`协议标准而设计，基于`ProtoBuf(Protocol Buffers)`序列化协议开发，且支持众多开发语言。
`gRPC`提供了一种简单的方法来精确地定义服务和为`iOS`、`Android`和后台支持服务自动生成可靠性很强的客户端功能库。客户端充分利用高级流和链接功能，从而有助于节省带宽、降低的`TCP`链接次数、节省`CPU`使用、和电池寿命。
我们可以从这里[http://www.grpc.io/blog/principles](http://www.grpc.io/blog/principles) `gRPC`的动机和设计原则。

![gRPC](https://o6rr5e4by.qnssl.com/wp-content/uploads/2016/10/1-X7I-IyhPdnNCsYJlm1U0Hw.jpeg)
