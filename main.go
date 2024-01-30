package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func verifierDir(folder string) bool {
	f, err := os.Open(folder)

	if err != nil {
		panic(err)
	}

	for {
		files, err := f.ReadDir(1)
		if err != nil {
			if err == io.EOF {
				break
			}
			continue
		}

		if files[0].Name() == ".git" {
			return true
		}
	}

	return false
}

func main() {
	var folder string
	flag.StringVar(&folder, "add", "", "novo diretório para análise")
	flag.Parse()
	fmt.Println(verifierDir(folder))
}
