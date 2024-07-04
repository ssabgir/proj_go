package main

import "fmt"

func Sum(x, y int) int {
	return x + y
}

func task_2() {
	var x, y int
	fmt.Println("Введите два целых числа: ")
	fmt.Scan(&x, &y)
	fmt.Println("Сумма: ", Sum(x, y))
}
