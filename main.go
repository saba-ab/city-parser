package main

import "github.com/saba-ab/city-parser/parser"

func main() {
	err := parser.ParseFile("data/batumi.html")
	if err != nil {
		panic(err)
	}
}
