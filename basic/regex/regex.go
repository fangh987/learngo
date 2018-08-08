package main

import (
	"regexp"
	"fmt"
)

const text = `My email is fanghua@email.com
cc abc@qq.com
email is dd@abc.com.cn`
func main() {
	re := regexp.MustCompile(`([a-zA-Z0-9]+)@([a-zA-Z0-9]+)(\.[a-zA-Z0-9.]+)`)
	match := re.FindAllStringSubmatch(text,-1)
	for _,m := range match {
		fmt.Println(m)
	}

}
