package main

import "fmt"

/*
Лабораторная работа №2, задание №3: Вычислить сумму первых N элементов ряда: 1/4 + 2/9 + 3/16 + 4/25 + ...
*/

func main() {
	sum := 0.0
	var N, i int

	fmt.Print("Input your N = ")
	_, err := fmt.Scanf("%v", &N)
	if err != nil {
		return
	}

	for i = 1; i < N; i++ {
		sum = sum + float64(i)/float64((i*i+i*2.0+1.0))
	}

	fmt.Printf("The sum is %.5g", sum)
}
