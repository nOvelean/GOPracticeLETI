package main

import (
	"GOPracticeLETI/Structs"
	"fmt"
)

func main() {
	fmt.Println("Starting")
	EmptyList := *Structs.New(10)

	EmptyList.Print()
	//0 <-> 0 <-> 0 <-> 0 <-> 0 <-> 0 <-> 0 <-> 0 <-> 0 <-> 0 <-> 0
	array := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	List := *Structs.NewFormList(array)
	List.Print()
	//1 <-> 2 <-> 3 <-> 4 <-> 5 <-> 6 <-> 7 <-> 8 <-> 9
	List.Add(10)
	fmt.Println(List.Len())
	//10
	List.Pop()

	List.DeleteFrom(0)

	List.UpdateAt(1, 100)

	List.Print()
	//2 <-> 100 <-> 4 <-> 5 <-> 6 <-> 7 <-> 8 <-> 9
	fmt.Println(List.At(7))
	//9
	List.Insert(6, 1)
	List.Print()
	//2 <-> 100 <-> 4 <-> 5 <-> 6 <-> 7 <-> 1 <-> 8 <-> 9
	Structs.Join(&List, &EmptyList).Print()
	//2 <-> 100 <-> 4 <-> 5 <-> 6 <-> 7 <-> 1 <-> 8 <-> 9 <-> 0 <-> 0 <-> 0 <-> 0 <-> 0 <-> 0 <-> 0 <-> 0 <-> 0 <-> 0 <-> 0 <-> 0
}
