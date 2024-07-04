package main

import "fmt"

func task_8() {
	var S string
	fmt.Println("Введите строку:")
	fmt.Scan(&S)
	s := []rune(S)
	for i, j := len(s)-1, 0; i > j; i, j = i-1, j+1 {
		s[i], s[j] = s[j], s[i]
	}
	fmt.Println("Перевернутая строка:", string(s))
}
