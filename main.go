package main

import (
	"github.com/urfave/cli"
	"os"
	"strconv"
	"fmt"
)

const APP_NAME = "GDG Perm currencies converter"
const APP_USAGE = "Convert one currency to another"

func main() {
	app := cli.NewApp()
	app.Name = APP_NAME
	app.Usage = APP_USAGE

	app.Flags = []cli.Flag {
		cli.StringFlag{
			Name: "currency, c",
			Usage: "Base currency",
		},
		cli.StringFlag{
			Name: "amount, a",
			Usage: "Amount of base currency",
		},
	}

	app.Action = func(c *cli.Context) error {
		var amount float64

		if c.String("currency") == "" {
			fmt.Println("Currency should be specified")
			os.Exit(1)
		}

		if c.String("amount") == "" {
			fmt.Println("Amount should be specified")
			os.Exit(1)
		}

		currency := Currency{c.String("currency")}

		amount, err := strconv.ParseFloat(c.String("amount"),64)
		if err != nil {
			fmt.Println("Amount should be parsable float")
			os.Exit(1)
		}

		result, err := Convert(currency, amount)
		if err != nil {
			panic(err)
		}

		for currency, amount := range result {
			fmt.Println(currency.Name(), amount)
		}

		return nil
	}

	app.Run(os.Args)
}