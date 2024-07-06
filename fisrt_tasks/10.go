package main

import "fmt"

type Rectangle struct {
	W, H int
}

func (r Rectangle) Square() int {
	return r.H * r.W
}

func task_10() {
	var r Rectangle
	fmt.Println("Введите ширину и высоту прямоугольниика: ")
	fmt.Scan(&r.W, &r.H)
	fmt.Println("Площадь прямоугольника = ", r.Square())
}
