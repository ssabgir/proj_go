package main

import "fmt"

func task_5() {
	var x uint64
	fmt.Println("Введите целое число: ")
	fmt.Scan(&x)
	fact := uint64(1)
	for x > 1 {
		fact *= x
		x--
	}
	fmt.Println("Факториал числа равен ", fact)
}
