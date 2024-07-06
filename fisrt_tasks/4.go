package main

import "fmt"

func task_4() {
	var x, y, z int
	fmt.Println("Введите три целых числа: ")
	fmt.Scan(&x, &y, &z)
	fmt.Println("Максимум из трех чисел: ", max(x, y, z))
}
