package main

import (
	"errors"
	"fmt"
)

const (
	x = iota // 0
	y = iota // 1
	z = iota // 2
	w        // 常量声明省略值时，默认和之前一个值的字面相同。这里隐式地说w = iota
)

const v = iota // 每遇到一个const关键字，iota就会重置，此时v = 0

const (
	h, i, j = iota, iota, iota // h=i=j=0,iota在同一行的时候值相同
)

const (
	a       = iota // a=0
	b       = "B"
	c       = iota             // c= 2
	d, e, f = iota, iota, iota // d=e=f=3
	g       = iota             // g=4
)

func main() {
	fmt.Printf("%s or 你好，世界 or καλημ ́ρα κóσμ or こんにちはせかい\n", "Hello, world")

	s := "hello"
	c1 := []byte(s)
	c1[0] = 'c'
	s2 := string(c1)
	fmt.Printf("%s\n", s2)

	m := `hello
		world`
	fmt.Println(m)

	err := errors.New("emit macho dwarf: elf header corrupted")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(a, b, c, d, e, f, g, h, i, j, x, y, z, w, v)

	var arr [10]int
	arr[0] = 42
	arr[1] = 13
	fmt.Printf("The first element is %d\n", arr[0])
	fmt.Printf("The last element is %d\n", arr[9])

	// 声明一个含有10个元素元素类型为byte的数组
	var ar = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}
	// 声明两个含有byte的slice
	var a, b []byte
	// a指向数组的第3个元素开始，并到第五个元素结束
	a = ar[2:5]
	b = ar[3:5]
	fmt.Printf("slice_a=%v, slice_b=%v", a, b)
	// 注意slice和数组在声明时的区别：声明数组时，方括号内写明了数组的长度或使用...自动计算长度，而声明slice时，方括号内没有任何字符

	// 从概念上来说slice像一个结构体，包含三个元素：
	// 1. 一个指针，指向数组中slice指定的开始位置
	// 2. 长度，即slice的长度
	// 3. 最大长度，也就是slice开始位置到数组的最后位置的长度

	// 声明一个key是字符串，值为int的字典，这种方式在使用的时候使用make初始化
	//var numbers map[string]int
	numbers := make(map[string]int)
	numbers["one"] = 1
	numbers["ten"] = 10
	numbers["three"] = 3

	fmt.Println("第三个数字是： ", numbers["three"])
	// 1. map是无序的，每次打印出来的map都会不一样，不能通过index获取，只能通过key获取
	// 2. map的长度不固定，也就是和slice一样，也是一种引用类型
	// 3. 内置的len函数同样适用于map，返回map拥有的key的数量
	// 4. map的值可以很方便的修改，通过numbers["one"]=11可以很容易的把key为one的字典改为11
	// 5. map和其他类型不同，它不是thread-safe的，在多个go-routine存取时，必须使用mutex lock机制
	rating := map[string]float32{"C": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	// map有两个返回值，第二个返回值，如果不存在key，那么ok为false，如果存在ok为true
	if csharpRating, ok := rating["C#"]; ok {
		fmt.Println("C# is in the map and its rating is ", csharpRating)
	} else {
		fmt.Println("We have no rating associated with C# in the map")
	}
	delete(rating, "C") // 删除key为C的元素
	// map也是一种引用类型，如果两个map同时指向一个底层，那么一个改变了，另一个也相应的改变了：
	//m := make(map[string]string)
	//m["Hello"] = "Bonjour"
	//m1 := m
	//m1["Hello"] = "Salut"  // m["Hello"]的值已经改变了~

	// make 用于内建类型（map、slice、channel）的内存分配。new用于各种类型的内存分配。
	// new(T)分配了零值填充的T类型的内存空间，并且返回其地址，即一个*T类型的值。
	// 用Go的术语说，它返回了一个指针，指向新分配的类型T的零值。new返回指针。
	// 内建函数make(T, args)与new(T)有着不同的功能，make只能创建slice、map和channel，并且返回一个有初始值的T类型，而不是*T

	sum := 0
	for index := 0; index < 10; index++ {
		sum += index
	}
	fmt.Println("sum is equal to ", sum)

	// go里面的默认switch相当于每个case最后带有break，匹配成功后不会自动向下执行其他case，而是跳出整个switch
	// 可以使用fallthrough强制执行后面的case代码。
	integer := 6
	switch integer {
	case 4:
		fmt.Println("The integer was <= 4")
		fallthrough
	case 5:
		fmt.Println("The integer was <= 5")
		fallthrough
	case 6:
		fmt.Println("The integer was <= 6")
		fallthrough
	case 7:
		fmt.Println("The integer was <= 7")
		fallthrough
	case 8:
		fmt.Println("The integer was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

	x := 3
	fmt.Println("x = ", x)
	x1 := add1(x)
	fmt.Println("x+1=", x1)
	fmt.Println("x=", x)

	x2 := add2(&x)
	fmt.Println("x+1=", x2)
	fmt.Println("x=", x)

	slice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("slice= ", slice)
	odd := filter(slice, isOdd) // 函数当做值来传递
	fmt.Println("Odd elements of slice are: ", odd)
	even := filter(slice, isEven)
	fmt.Println("Even elements of slice are: ", even)
}

func add1(a int) int { // a是x的copy
	a = a + 1
	return a
}

// 1. 传指针使得多个函数能操作同一个对象
// 2. 传指针比较轻量级(8bytes)，只是传内存地址，我们可以用指针传递体积大的结构体。
// 如果用参数传递的话，在每次copy上面就会花费相对较多的系统开销（内存和时间）。
// 3. Go语言中channel、slice、map这三种类型的实现机制类似指针，所以可以直接传递，而不用取地址后传递指针。
func add2(a *int) int {
	*a = *a + 1
	return *a
}

// goto 标签名是大小写敏感
func myFunc() {
	i := 0
Here: // 以冒号结束作为标签
	println(i)
	i++
	goto Here // 跳转到Here去
}

// Go函数也是一种变量，我们可以通过type来定义它，它的类型就是所有拥有相同的参数，相同的返回值的一种类型
// type typeName func(input1 inputType1, input2 inputType2 [, ...]) (result1 resultType1 [, ...])

type testInt func(int) bool

func isOdd(integer int) bool {
	if integer%2 == 0 {
		return false
	}
	return true
}

func isEven(integer int) bool {
	if integer%2 == 0 {
		return true
	}
	return false
}

func filter(slice []int, f testInt) []int {
	var result []int
	for _, value := range slice {
		if f(value) {
			result = append(result, value)
		}
	}
	return result
}
