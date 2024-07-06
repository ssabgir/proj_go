package main

import "fmt"

func task_6() {
	var ch string
	fmt.Println("Введите символ: ")
	fmt.Scan(&ch)
	if ch == "а" || ch == "е" || ch == "ё" ||
		ch == "и" || ch == "о" || ch == "у" ||
		ch == "э" || ch == "ю" || ch == "я" {
		fmt.Println("Буква является гласной")
	} else {
		fmt.Println("Буква является coгласной")
	}
}
