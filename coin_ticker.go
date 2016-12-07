package ct

import (
	"fmt"
	"net/http"
	"os"

	"github.com/urfave/cli"
)

func main() {

	var coinToken string

	httpClient := &http.Client{}
	ctClient := NewClient(httpClient)

	app := cli.NewApp()
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "coin",
			Value:       "bitcoin",
			Usage:       "coin used for testing",
			Destination: &coinToken,
		},
	}

	app.Name = "bitcoin checker"
	app.Usage = "wrapper testing current cryptcurr price"

	app.Action = func(c *cli.Context) error {
		if coinToken == "b" {
			fmt.Println(ctClient.GetBitcoinPrice())
		} else if coinToken == "e" {
			fmt.Println(ctClient.GetEtherPrice())
		} else {
			fmt.Println("unkown option")
		}

		return nil
	}

	app.Run(os.Args)
}
