package main

import (
	"fmt"
	"math/rand"
	"time"
)

/*
Лабораторная работа №4, В массиве из 10 целых чисел обнулить все отрицательные элементы,
если сумма их модулей меньше суммы положительных элементов.
*/

const (
	N = 10
)

func main() {
	var arr [N]int
	negnum := 0
	posnum := 0

	rand.Seed(time.Now().Unix())
	for i := 0; i < N; i++ {
		arr[i] = rand.Intn(100) - 50
	}

	fmt.Println(arr)

	for i := 0; i < N; i++ {
		if arr[i] < 0 {
			negnum += arr[i]
		} else {
			posnum += arr[i]
		}
	}
	negnum = -negnum

	if negnum < posnum {
		for i := 0; i < N; i++ {
			if arr[i] < 0 {
				arr[i] = 0
			}
		}
	}

	fmt.Printf("The negnum = %v\n", negnum)
	fmt.Printf("The posnum = %v\n", posnum)
	fmt.Println(arr)
}
