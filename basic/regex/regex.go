package main

import (
	"fmt"
	"regexp"
)

const text = `My email is ccmouse@gmail.com@dwd
email1 is abc@def.org
email2 is    kkk@qq.com
email3 is ddd@abc.com.cn
`

func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text, -1) // -1はすべてを見つける
	//fmt.Println(match)     // [[ccmouse@gmail.com ccmouse gmail com] [abc@def.org abc def org] [kkk@qq.com kkk qq com] [ddd@abc.com.cn ddd abc.com cn]]
	for _, m := range match {
		fmt.Println(m)

	}
}
