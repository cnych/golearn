package main

import (
	"errors"
	"log"
)

func main() {
	// go func() {
	// 	defer func() {
	// 		if err := recover(); err != nil {
	// 			log.Printf("error: %v", err)
	// 		}
	// 	}()
	// 	log.Panic("some error before work2")
	// 	fmt.Println("do some work2")
	// }()
	// time.Sleep(time.Second)
	// fmt.Println("do some work!")

	err, _ := f3()
	if err != nil {
		log.Print(err)
	}
}

type R struct {
}

func f1() (error, *R) {
	return errors.New("an error"), nil
}
func f2() (error, *R) {
	err, r := f1()
	if err != nil {
		return err, nil
	}
	return nil, r
}
func f3() (error, *R) {
	err, r := f2()
	if err != nil {
		return err, nil
	}
	return nil, r
}
