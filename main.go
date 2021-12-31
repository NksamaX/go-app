package main

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"

	"app/types"

	"github.com/urfave/cli"
)

func routes() {
	app := cli.NewApp()
	app.Name = "Nksama"
	app.Version = "0.0.1"
	app.Usage = "Nksama is a CLI tool for memes"

	app.Commands = []cli.Command{
		{
			Name:    "meme",
			Aliases: []string{"m"},
			Usage:   "Generate memes",
			Action: func(c *cli.Context) error {
				res, _ := http.Get("https://nksamamemeapi.pythonanywhere.com/")
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(string(body))
				return nil
			},
		},

		{
			Name:    "dl",
			Aliases: []string{"d"},
			Usage:   "Download memes",
			Action: func(c *cli.Context) error {
				res, _ := http.Get("https://nksamamemeapi.pythonanywhere.com/")
				body, err := ioutil.ReadAll(res.Body)
				if err != nil {
					fmt.Println(err)
				}

				var mem types.Meme

				json.Unmarshal(body, &mem)

				resp, _ := http.Get(mem.Image)
				file, _ := os.Create("meme" + ".jpg")
				defer file.Close()
				_, err = io.Copy(file, resp.Body)
				if err != nil {
					fmt.Println(err)
				}

				return nil
			},
		},

		{
			Name:    "quote",
			Aliases: []string{"q"},
			Usage:   "Get a quote",
			Action: func(c *cli.Context) error {
				res, _ := http.Get("https://animechan.vercel.app/api/random")
				body, _ := ioutil.ReadAll(res.Body)
				var qt types.Quote
				json.Unmarshal(body, &qt)

				fmt.Printf("%s \n   - %s", qt.Quote, qt.Character)
				return nil

			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
	}

}

func main() {
	routes()
}
