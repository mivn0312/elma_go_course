package main

import (
	"strings"
	"strconv"
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d\n", &n)
	fmt.Println(GeneratorSquaresNumbers(n))
	fmt.Println(Solution(n))
	fmt.Println(NumberUniques([]int{1, 2, 3, 4, 5}))
	fmt.Println(Carousel([]int{1, 2, 3, 4, 5}, 1))
}

// генератор квадратов натуральных чисел
func GeneratorSquaresNumbers(n int) (arr []int) {
	for i := 1; i <= n; i++ {
		arr = append(arr, i*i)
	}
	return
}

// максимальная длина последовательности нулей в двоичном представлении целого числа
func Solution(n int) (l int) {
	var arr = strings.Split(strconv.FormatInt(int64(n), 2), "1")
	for _, s := range arr {
		var ls = len(s)
		if ls > l {
			l = ls
		}
	}
	return
}

// количество уникальный чисел в слайсе
func NumberUniques(arr []int) int {
	m := make(map[int]bool)
	for _, v := range arr {
		m[v] = false
	}
	return len(m)
}

// карусель
func Carousel(arr []int, shift int) ([]int) {
	l := len(arr)
	if l <= 1 {
		return arr
	}

	shift = shift % l
	if shift == 0 {
		return arr
	}

	return append(arr[l-shift:], arr[:l-shift]...)
}
