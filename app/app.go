package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

//Gerador vai retornar aplicação de linha de comando
func Gerador() *cli.App {
	app:= cli.NewApp()
	app.Name = "Mini aplicação"
	app.Usage = "Busca IP e nome de servidor"

	app.Commands = []cli.Command{

		{
			Name: "ip",
			Aliases: []string{"busca-ip"},
			Flags: []cli.Flag{
				cli.StringFlag{
					Name: "host",
					Value: "Mini Aplicacao",
				},
			},
			Action: BuscaDeIp,

		},
	}


	return app
}


func BuscaDeIp(c * cli.Context)  {

	host:= c.String("host")
	
	ips, erro := net.LookupHost(host)

	if erro != nil{
		log.Fatal(erro)
	}

	for _, ip:= range ips{
		fmt.Println(ip)
	}
}