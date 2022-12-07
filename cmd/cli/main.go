package main

import (
	"apoor-ssh/pkg/letters"
	"fmt"
)

func main() {
	txt := "hello world"
	fmt.Println(letters.FormatWord(txt))
	fmt.Println(txt)
}
