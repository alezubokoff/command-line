package app

import (
	"fmt"
	"log"
	"net"

	"github.com/urfave/cli"
)

func Generate() *cli.App {
	app := cli.NewApp()
	app.Name = "Aplicação de linha de comando"
	app.Usage = "Buscar IPs de paginas da internet"

	flags := []cli.Flag{
		cli.StringFlag{
			Name:  "host",
			Value: "google.com.br",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:   "ip",
			Usage:  "Comando para exibir IPs",
			Flags:  flags,
			Action: findIPs,
		},
		{
			Name:   "server",
			Usage:  "Comando para exibir IPs",
			Flags:  flags,
			Action: findServers,
		},
	}
	return app
}

func findIPs(c *cli.Context) {
	host := c.String("host")
	ips, err := net.LookupIP(host)

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, ip := range ips {
		fmt.Println(ip)
	}
}

func findServers(c *cli.Context) {
	host := c.String("host")
	servers, err := net.LookupNS(host)

	if err != nil {
		log.Fatal(err.Error())
	}

	for _, server := range servers {
		fmt.Println(server.Host)
	}
}
