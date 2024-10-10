package app

import (

	"github.com/urfave/cli"
)

//Gerador vai retornar aplicação de linha de comando
func Gerador() *cli.App {
	app:= cli.NewApp()
	app.Name = "Mini aplicação"
	app.Usage = "Busca IP e nome de servidor"

	return app
}