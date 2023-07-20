package bitwise

import "fmt"

func Bitwise() {
	var n uint = 17
	var s uint = 1 << 2

	fmt.Printf("bit: %b, dec: %d\n", n, n)
	fmt.Printf("bit: 00%b, dec: %d\n", s, s)

	fmt.Println("----------Устанавливаем бит в 1----------")
	n = n | s
	fmt.Printf("bit: %b, dec: %d\n", n, n)

	fmt.Println("----------Устанавливаем бит в 0----------")
	n = n & ^s
	fmt.Printf("bit: %b, dec: %d\n", ^s, ^s)
	fmt.Printf("bit: %b, dec: %d\n", n, n)

	fmt.Println("----------Меняем значение бита на противоположное----------")
	n = n ^ s
	fmt.Printf("bit: %b, dec: %d\n", n, n)
	n = n ^ s
	fmt.Printf("bit: %b, dec: %d\n", n, n)
	n = n ^ s
	fmt.Printf("bit: %b, dec: %d\n", n, n)
	n = n ^ s
	fmt.Printf("bit: %b, dec: %d\n", n, n)

	fmt.Println("---------Проверяем значение бита-----------")
	if n&s == 0 {
		fmt.Println("zero")
	}
	if n&s != 0 {
		fmt.Println("not zero")
	}
}

func Sets() {
	data := []int{1, 2, 3}
	n := len(data)

	for mask := 0; mask < (1 << n); mask++ {
		fmt.Printf("{")
		first := true
		for i := 0; i < n; i++ {
			if mask&(1<<i) != 0 {
				if !first {
					fmt.Printf(", ")
				}
				fmt.Printf("%d", data[i])
				first = false
			}
		}
		fmt.Printf("}\n")
	}
}
