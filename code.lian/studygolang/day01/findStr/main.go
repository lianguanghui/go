package main

import (
	"fmt"
	"strings"
)

func main() {
	//查找字符串中是否存在he
	fmt.Println(strings.Contains("hello", "he"))

	//查找字符串中是否存在1或2
	fmt.Println(strings.ContainsAny("nihao", "12"))

	//查找字符在字符串中的数量
	fmt.Println(strings.Count("hello word", "l"))

	//返回字符在字符串中的第一次出现的index，不存在的时候返回-1,如果str为空则返回字符串s的长度。
	fmt.Println(strings.Index("hello", "l"))

	//返回字符在字符串中最后一次出现的index，不存在的时候返回-1,如果str为空则返回字符串s的长度。
	fmt.Println(strings.LastIndex("hello,world", ""))

	//返回任意字符在字符串中第一次出现的index，如果找不到或str为空则返回-1。
	fmt.Println(strings.IndexAny("hello", "12asfl"))

	//返回任意字符在字符串中最后一次出现的index，如果找不到或str为空则返回-1。
	fmt.Println(strings.LastIndexAny("hello", "12asfl"))

	//以字符为切割符号将字符串切割成n份，切割后不包含字符
	fmt.Println(strings.SplitN("helppplo,world", "l", 1))
	fmt.Println(strings.SplitN("helppplo,world", "l", 2))
	fmt.Println(strings.SplitN("helppplo,world", "l", 3))
	fmt.Println(strings.SplitN("helppplo,world", "l", 4))
	fmt.Println(strings.SplitN("helppplo,world", "l", 5))

	str := "apple,orange,banana,grape"
	fmt.Println(strings.SplitN(str, ",", 1))
	fmt.Println(strings.SplitN(str, ",", 2))
	fmt.Println(strings.SplitN(str, ",", 3))
	fmt.Println(strings.SplitN(str, ",", 4))
	fmt.Println(strings.SplitN(str, ",", 5))

	fmt.Println("------>", strings.Split(str, ","))

	//查看字符串的开头和结尾
	fmt.Println("是否以字符he开头===>", strings.HasPrefix("hello,world", "he"))
	fmt.Println("是否以字符ld结尾===>", strings.HasSuffix("hello,world", "ld"))

	//重复字符拼接
	fmt.Println(strings.Repeat("hello", 2))

	//删除字符串前后的字符  TrimLeft删除字符串左边某字符，TrimRight删除字符串右边某字符
	fmt.Println(strings.Trim("\"hello world\"", "\""))

	//删除字符开头的某字符，只会删除一次，TrimSuffix删除字符串尾部的某字符，只会删除一次
	fmt.Println(strings.TrimPrefix("///hello world/", "/"))

	//比较两个字符串是否相同不区分大小写
	fmt.Println("不区分大小写是否相同====>", strings.EqualFold("hello", "HELLo"))
	//比较两个字符是否相同，区分大小写.相等为0，s1>s2为-1，s1<s2为1
	fmt.Println("区分大小写比较是否相同---->", strings.Compare("hello world", "Hello world"))
}
