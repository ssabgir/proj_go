package main

import "fmt"

func task_9() {
	var l int
	fmt.Println("Введите размер массива целых чисел:")
	fmt.Scan(&l)
	fmt.Println("Введите массив целых чисел:")
	v := make([]int, l)
	for i := 0; i < l; i++ {
		fmt.Scan(&v[i])
	}
	sum := 0
	for i := 0; i < l; i++ {
		sum += v[i]
	}
	fmt.Println("Сумма = ", sum)
}
