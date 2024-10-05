package main

import "fmt"

func main() {

	//a%b = a - a / b * b
	fmt.Println("--->", -10%3)
	fmt.Println("--->", 10%-3)
	fmt.Println("--->", -10%-3)

	//还有97天放假，问XX个星期零XX天
	i := 97
	n := 97 % 7
	m := 97 / 7
	fmt.Printf("%d天是%d个星期零%d天\n", i, m, n)

	//华氏摄氏度与摄氏温度的转换
	var h float64 = 134.2
	w := 5.0 / 9 * (h - 100)
	fmt.Printf("%v华氏温度对应的摄氏温度是%v\n", h, w)

	//三个数的最大值
	var y, u, t int64 = 101, 220, 3012
	var max int64
	if y > u {
		max = y
	} else {
		max = u
	}
	if t > max {
		max = t
	}
	fmt.Println("最大值是：", max)
}
