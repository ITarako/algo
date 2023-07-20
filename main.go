package main

import "fmt"

func main() {

	//q := structure.NewQueue()
	//q.Push(1)
	//q.Push(2)
	//q.Push(3)
	//fmt.Println(q.Top())
	//q.Pop()
	//fmt.Println(q.Top())
	//q.Pop()
	//q.Pop()

	a := []int{1, 2, 3, 4, 5}
	b := a[:1:1]
	fmt.Println(b)
	b[0] = 1111
	fmt.Println(b)
	//b = b[:5]
	b = append(b, 2222)
	fmt.Println(a)
	fmt.Println(b)
}
