package main

import "fmt"

func main() {

	//控制台输入信息

	var name string
	var age int32
	var sal float32
	var isPass bool

	//方式1:Scanln
	fmt.Println("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Println("请输入年龄")
	fmt.Scanln(&age)
	fmt.Println("请输入薪水")
	fmt.Scanln(&sal)
	fmt.Println("请输入是否通过")
	fmt.Scanln(&isPass)

	fmt.Printf("姓名：%v，年龄%v，薪水：%v，是否通过：%v\n", name, age, sal, isPass)

	//方式2:Scanf
	fmt.Println("请输入姓名，年龄，薪水，是否通过(空格隔开)")
	fmt.Scanf("%v,%v,%v,%v", name, age, sal, isPass)
	fmt.Printf("姓名：%v，年龄%v，薪水：%v，是否通过：%v\n", name, age, sal, isPass)
}
