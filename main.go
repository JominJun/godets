package main

import "fmt"

func main() {
	age := 20

	if manAge := age-1; manAge >= 19 {
		fmt.Println("You can buy an alchol")
	} else {
		fmt.Println("You cannot buy an alchol")
	}

	switch manAge := age-1;{
		case manAge <= 13:
			fmt.Println("어린이 요금")
		case manAge < 65:
			fmt.Println("일반 요금")
		case manAge >= 65:
			fmt.Println("노약자 요금")
	}

	a := 10
	b := &a
	a = 11

	fmt.Println(a, *b) // 11 11

	arr := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(arr)

	arr = append(arr, 11)
	fmt.Println(arr)
}