package main

import "fmt"

func main() {
	factor := 996
	factor = 10551396
	hitcount := 0
	for i := 1; i <= factor; i++ {
		if factor % i == 0 {
			hitcount += i
			fmt.Println(hitcount)
		}
		/*
		for j := 1; j <= factor; j++ {
			if i * j == 996 {
				hitcount += i
				fmt.Println(hitcount)
			}
		}
		*/
	}
	fmt.Println(hitcount)
}
