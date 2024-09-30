package main

import "fmt"

func main() {
	age := 11
	for i := 0; i <= age; i++ {
		if i >= 18 {
			fmt.Println("成年人你好啊")
		} else if i < 18 {
			fmt.Println("非成年人")
		} else {
			fmt.Printf("错误的")
			fmt.Println("这是怎么回事")
		}
	}
}
