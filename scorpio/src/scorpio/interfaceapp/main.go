package main

import "fmt"

// Shaper 接口
type Shaper interface {
	Area() float32 // 计算面积
}

// 正方形结构体
type Square struct {
	side float32
}

// 实现接口Shaper中的Area的方法
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

type Rectangle struct {
	length, width float32
}

func (r *Rectangle) Area() float32 {
	return r.length * r.width
}

type stockPosition struct {
	ticker     string
	sharePrice float32
	count      float32
}

func (s stockPosition) getValue() float32 {
	return s.sharePrice * s.count
}

type car struct {
	make  string
	model string
	price float32
}

func (c car) getValue() float32 {
	return c.price
}

type valuable interface {
	getValue() float32
}

func showValue(asset valuable) {
	fmt.Printf("Value of the asset is %f\n", asset.getValue())
}

func main() {
	sq1 := new(Square)
	sq1.side = 5

	areaIntf := sq1
	fmt.Printf("The square has area: %f\n", areaIntf.Area())

	r := Rectangle{5, 3}
	q := Square{5}
	shapes := []Shaper{&r, &q}
	fmt.Println("Looping through shapes for area...")
	for _, shape := range shapes {
		fmt.Println("Shape details: ", shape)
		fmt.Println("Area of this shape is:", shape.Area())
	}

	var o valuable = stockPosition{"GOOG", 577.20, 4}
	showValue(o)
	o = car{"BMW", "M3", 66500}
	showValue(o)

}
