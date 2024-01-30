package main

import "flag"

func main() {
	var folder string
	flag.StringVar(&folder, "add", "", "novo diretório para análise")
	flag.Parse()
	print(folder)
}
