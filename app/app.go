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

	flagGenerica := []cli.Flag{
		cli.StringFlag{
			Name: "host",
			Value: "Mini Aplicacao",
		},
	}


	app.Commands = []cli.Command{

		{
			Name: "ip",
			Aliases: []string{"busca-ip"},
			Flags: flagGenerica,
			Action: buscaDeIp,

		},
		{
			Name: "servidor",
			Aliases: []string{"busca-servidor"},
			Flags: flagGenerica,
			Action: buscaServidor,
		},
	}


	return app
}

func buscaServidor(c * cli.Context)  {
	host:= c.String("host")

	servidores, erro := net.LookupNS(host)

	if erro != nil{
		log.Fatal(erro)
	}

	for _, servidor:= range servidores{
		fmt.Println(servidor.Host)
	}
}



func buscaDeIp(c * cli.Context)  {

	host:= c.String("host")
	
	ips, erro := net.LookupHost(host)

	if erro != nil{
		log.Fatal(erro)
	}

	for _, ip:= range ips{
		fmt.Println(ip)
	}
}