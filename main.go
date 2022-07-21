package main

import (
	"log"

	"github.com/testadenevill/myniceprogram/helpers"
)

func main() {
	log.Println("Hello")

	var myVar helpers.SomeType
	myVar.TypeName = "Testa"
	log.Println(myVar.TypeName)

}
