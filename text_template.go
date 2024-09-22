package main

import (
	"fmt"
	"os"
	"text/template"
)

func main() {
	
	// fmt.Println(b,c)
	t1 := template.New("t1")
	fmt.Println(t1)
	t1, err := t1.Parse("value is {{.}} \n")
	fmt.Println(t1)

	if err != nil {
		panic(err)
	}
	t1 = template.Must(t1.Parse("Value:{{.}}\n"))
	t1.Execute(os.Stdout, "some text")
	t1.Execute(os.Stdout, 5)
	t1.Execute(os.Stdout,[] string{
		"go",
		"Rust",
	})

	
} 
