package main

import (
	"fmt"
	"os"
	"tree/tree"
)

func main(){
	if len(os.Args) != 3 {
		fmt.Println("not enough args")
	}
	maintext := os.Args[1]
	aftertext := os.Args[2]

	file, _ := os.ReadFile(maintext)
	content := tree.Transformer(string(file))

	os.WriteFile(aftertext,[]byte(content), 0644)

	
}