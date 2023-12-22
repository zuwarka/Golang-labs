package main

import "fmt"

/*
Лабораторная работа №1, задание №3: Бассейн объемом V литров полностью заполнен водой.
По одной трубе вода из бассейна вытекает со скоростью ν1 литров в минуту, по другой трубе вода поступает.
Бассейн стал пустым через время t.
Какова скорость поступления воды ν2?
*/

func main() {
	var V, u1, t float32

	fmt.Print("The capacity of the pool V = ")
	fmt.Scanf("%g", &V)

	fmt.Print("Empty time (min), t = ")
	fmt.Scanf("%g", &t)

	fmt.Print("The speed of leakage (l/min), u1 = ")
	fmt.Scanf("%g", &u1)

	fmt.Printf("The speed of filling the pool, u2 = %g", u1-V/t)
}
