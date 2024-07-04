package main

import "fmt"

func task_3() {
	var x int
	fmt.Println("Введите одно целое число: ")
	fmt.Scan(&x)
	if x%2 == 0 {
		fmt.Printf("Число %d чётное", x)
	} else {
		fmt.Printf("Число %d нечётное", x)
	}
	fmt.Println()
}
