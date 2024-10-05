package main

import "fmt"

func main() {
	i := 1001
	j := 100
	fmt.Println("i的内存地址是:", &i) //& 用于取变量的内存地址
	fmt.Println("j的内存地址是:", &j)

	var po *int = &i

	*po = 100 //通过内存地址修改值
	fmt.Printf("po内存地址指向的值是：%v\n", *po)
}
