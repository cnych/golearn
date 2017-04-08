package main

import (
	"encoding/json"
	"fmt"
)

type Animal struct {
	Name    string  `json:"name"`
	Order   string  `json:",omitempty"`
	Age     float64 `json:"age"`
	Married bool    `json:"married"`
	Baby
}

type Baby struct {
	Toy string
}

func main() {
	var jsonBlob = []byte(`[{"name": "Platypus", "order": "Monotremata"}, {"name": "Quoll", "order": "Dasyuromorphia"}]`)

	var animals []Animal

	err := json.Unmarshal(jsonBlob, &animals)
	if err != nil {
		fmt.Println("error:", err)
	}
	fmt.Printf("%+v\n", animals)

	fmt.Println("-------------------")

	var jsonBlob1 = []byte(`{"name": "Platypus", "order":"haha", "age": 343}`)
	var animal1 Animal
	err = json.Unmarshal(jsonBlob1, animal1)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("带数字的--->%+v\n", animal1)

	fmt.Println("-------------------")

	a := Animal{"haha", "", 24, false, Baby{"bear"}}
	res, _ := json.Marshal(a)
	fmt.Println(string(res))

	fmt.Println("=================")
	b := []byte(`{"Name":"Wednesday","Age":6,"Parents":["Gomez","Morticia"]}`)
	var f interface{}
	err = json.Unmarshal(b, &f)
	if err != nil {
		fmt.Println(err)
	}
	if m, ok := f.(map[string]interface{}); ok {
		for k, v := range m {
			switch vv := v.(type) {
			case string:
				fmt.Println(k, "is string =", vv)
			case int:
				fmt.Println(k, "is int =", vv)
			case float64:
				fmt.Println(k, "is float64 =", vv)
			case []interface{}:
				fmt.Println(k, "is an array =")
				for i, u := range vv {
					fmt.Println("    ", i, u)
				}
			default:
				fmt.Println(k, "is of a type I don't know how to handle")
			}
		}
	}
}
