package main

import (
	"aplicacao/app"
	"fmt"
	"log"
	"os"
)

func main()  {

	fmt.Println("Resultados")

	aplicacao:= app.Gerador()
	if erro := aplicacao.Run(os.Args); erro != nil{
		log.Fatal(erro)
	}
}