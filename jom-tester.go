package main

import (
	"fmt"
	"os"

	"github.com/SimonEverts/joml"
)

func main() {

	file, err := os.Open("d:/devenv/go-yoml-env/test.joml")
	defer file.Close()

	if err != nil {
		panic(err)
	}

	parser := joml.NewParser(file)
	mymap := parser.ParseRootObject()

	fmt.Printf("%s", mymap)

}
