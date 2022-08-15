package main

import (
	"fmt"
	"os"

	"github.com/gduysens/dbml/parser"
	"github.com/gduysens/dbml/scanner"
)

func main() {
	f, _ := os.Open("test.dbml")
	s := scanner.NewScanner(f)
	parser := parser.NewParser(s)
	dbml, err := parser.Parse()
	fmt.Printf("%#v, %v", dbml, err)
}
