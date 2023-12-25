package main

import "fmt"

/*
Лабораторная работа 3 - Виртуальные функции

Вариант 3: Создать абстрактный класс «Кривые» для вычисления координаты y для некоторой x.
Создать производные классы «Прямая», «Эллипс», «Гипербола» со своими функциями вычисления y в зависимости от входного параметра x.
*/

func main() {
	var straight Curve = Straight{2.0, 3.0, 4.0}
	var ellipse Curve = Ellipse{2, 3, 4}
	var hyperbola Curve = Hyperbola{2, 3, 4}
	YStraight := straight.FindY()
	YEllipse := ellipse.FindY()
	YHyperbola := hyperbola.FindY()

	fmt.Printf("The straight = %.3g, the ellipse = %.3g, the hyperbola = %.3g\n", YStraight, YEllipse, YHyperbola)
}
