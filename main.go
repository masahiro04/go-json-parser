package main

import (
	"encoding/json"
	"fmt"
	"go-json-parser/lexer"
	"go-json-parser/parser"
	"io/ioutil"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) < 2 {
		panic("no valid file name or path provided for file!")
	}

	path := os.Args[1]
	absPath, _ := filepath.Abs(path)
	data, err := ioutil.ReadFile(absPath)
	if err != nil {
		panic(err.Error())
	}

	l := lexer.NewLexer(data)
	p := parser.NewParser(l)

	ast := p.Parse()

	js, _ := json.MarshalIndent(ast, "", "    ")
	fmt.Println(string(js))
}

// ./sample.jsonを読み込ませる
