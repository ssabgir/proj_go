package main

import (
	"fmt"
	"math"
)

func task_7() {
	var x int
	fmt.Println("Введите целое число:")
	fmt.Scan(&x)
	fmt.Println("Простые числа до ", x, ":")
	for i := 2; i <= x; i++ {
		fl := false
		for j := 2; j <= int(math.Sqrt(float64(i))); j++ {
			if i%j == 0 {
				fl = true
				break
			}
		}
		if !fl {
			fmt.Print(i, " ")
		}
	}
}
