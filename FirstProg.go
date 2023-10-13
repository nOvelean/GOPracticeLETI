package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	fmt.Println("Hello, world!")
	f, err := ioutil.ReadFile("information.txt")
	if err != nil {
		panic(err)
	} else {
		f_string := string(f)
		fmt.Println(f_string)
	}
	sum := 1
	for i := 0; i < 63; i++ {
		sum *= 2
		fmt.Printf("2^%d is %d\n", i+1, sum)
	}
}
