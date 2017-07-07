package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("text.txt")
	if err != nil {
		return
	}
	defer file.Close()
	a := "aaa"
	b := "bbb"
	//file.WriteString("test\nhello")
	file.WriteString(fmt.Sprintf("%s   %s", a, b))
}
