package main

import "fmt"

func main() {
	var day, month int32
	fmt.Scan(&day)
	fmt.Scan(&month)

	if betweenDays(day, 22, 31) && month == 12 {
		fmt.Print("capricornio")
	}
	if betweenDays(day, 1, 20) && month == 1 {
		fmt.Print("capricornio")
	}
	if betweenDays(day, 21, 31) && month == 1 {
		fmt.Print("acuario")
	}
	if betweenDays(day, 1, 19) && month == 2 {
		fmt.Print("acuario")
	}
	if betweenDays(day, 20, 29) && month == 2 {
		fmt.Print("piscis")
	}
	if betweenDays(day, 1, 20) && month == 3 {
		fmt.Print("piscis")
	}
	if betweenDays(day, 21, 31) && month == 3 {
		fmt.Print("aries")
	}
	if betweenDays(day, 1, 20) && month == 4 {
		fmt.Print("aries")
	}
	if betweenDays(day, 21, 30) && month == 4 {
		fmt.Print("tauro")
	}
	if betweenDays(day, 1, 21) && month == 5 {
		fmt.Print("tauro")
	}
	if betweenDays(day, 22, 31) && month == 5 {
		fmt.Print("geminis")
	}
	if betweenDays(day, 1, 21) && month == 6 {
		fmt.Print("geminis")
	}
	if betweenDays(day, 22, 30) && month == 6 {
		fmt.Print("cancer")
	}
	if betweenDays(day, 1, 23) && month == 7 {
		fmt.Print("cancer")
	}
	if betweenDays(day, 24, 31) && month == 7 {
		fmt.Print("leo")
	}
	if betweenDays(day, 1, 23) && month == 8 {
		fmt.Print("leo")
	}
	if betweenDays(day, 24, 31) && month == 8 {
		fmt.Print("virgo")
	}
	if betweenDays(day, 1, 23) && month == 9 {
		fmt.Print("virgo")
	}
	if betweenDays(day, 24, 30) && month == 9 {
		fmt.Print("libra")
	}
	if betweenDays(day, 1, 23) && month == 10 {
		fmt.Print("libra")
	}
	if betweenDays(day, 24, 31) && month == 10 {
		fmt.Print("escorpio")
	}
	if betweenDays(day, 1, 22) && month == 11 {
		fmt.Print("escorpio")
	}
	if betweenDays(day, 23, 30) && month == 11 {
		fmt.Print("sagitario")
	}
	if betweenDays(day, 1, 21) && month == 12 {
		fmt.Print("sagitario")
	}
}

func betweenDays(day, initialDay, limitDay int32) bool {
	return day >= initialDay && day <= limitDay
}
