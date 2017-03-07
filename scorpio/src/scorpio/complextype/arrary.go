package main

import "fmt"

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	p := []int{2, 3, 5, 7, 11, 13}
	fmt.Println("p ==", p)

	for i := 0; i < len(p); i++ {
		fmt.Printf("p[%d] == %d\n", i, p[i])
	}

	fmt.Println("p[1:4] == ", p[1:4])

	// 省略下标代表从0开始
	fmt.Println("p[:3] == ", p[:3])

	// 省略上标代表到len(s)结束
	fmt.Println("p[4:] == ", p[4:])

	a1 := make([]int, 5)
	printSlice("a", a1)
	b := make([]int, 0, 5)
	printSlice("b", b)
	c := b[:2]
	printSlice("c", c)
	d := c[2:5]
	printSlice("d", d)

	var z []int
	fmt.Println(z, len(z), cap(z))
	if z == nil {
		fmt.Println("nil!")
	}

	var a2 []int
	printSlice("a2", a2)

	// append works on nil slices
	a2 = append(a2, 0)
	printSlice("a2", a2)

	// the slice grows as needed
	a2 = append(a2, 1)
	printSlice("a2", a2)

	// we can add more than one element at a time.
	a2 = append(a2, 2, 3, 4)
	printSlice("a2", a2)

}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
