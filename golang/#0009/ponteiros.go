package main

import "fmt"

func main() {

	var variavelone int = 10
	var variaveltwo int = variavelone

	fmt.Println(variavelone, variaveltwo)

	variavelone ++
	fmt.Println(variavelone, variaveltwo)

	var variavelthree int
	var ponteiro *int

	fmt.Println(variavelthree, ponteiro)

	variavelthree = 100
	ponteiro = &variavelthree

	fmt.Println(variavelthree, *ponteiro)
}