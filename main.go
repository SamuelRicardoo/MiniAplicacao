package main

import (
	"aplicacao/app"
	"fmt"
	"log"
	"os"
)

func main()  {
	
	fmt.Print("Ponto de Partida")

	aplicacao:= app.Gerador()
	if erro := aplicacao.Run(os.Args); erro != nil{
		log.Fatal(erro)
	}
}