package main

import (
	"fmt"

	"github.com/urfave/cli"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"path/filepath"
	"time"
)

func downloadHtml(url string, filename string) {

	fmt.Println("Downloading " + url + " ...")
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	f, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	io.Copy(f, resp.Body)

}

func main() {
	app := cli.NewApp()
	app.Name = "random-downloader"
	app.Compiled = time.Now()
	app.Usage = "Download website sources using the Random Page function"
	app.UsageText = "url [-n number] [-d directory] [-r]"
	app.Description = "A crawler to download websites using the random page link"
	app.HideVersion = true
	app.Authors = []cli.Author{
		cli.Author{
			Name:  "Pablo Ovelleiro",
			Email: "pablo1@mailbox.org",
		},
	}

	var numpages int
	var mUrl string
	var directory string
	var retries int

	// Define avaitible flags
	app.Flags = []cli.Flag{
		cli.IntFlag{
			Name:        "n, number",
			Usage:       "Number of pages to download",
			Destination: &numpages,
		},

		cli.StringFlag{
			Name:        "url, u",
			Usage:       "The url to use",
			Destination: &mUrl,
		},

		cli.IntFlag{
			Name:        "retries, r",
			Usage:       "Retries to get new page",
			Destination: &retries,
		},
		cli.StringFlag{
			Name:        "directory, d",
			Usage:       "The download destination directry",
			Value:       "./downloaded-content",
			Destination: &directory,
		},
	}

	app.Action = func(c *cli.Context) error {

		// Check if necessary parameters are provided
		if mUrl != "" {

			// Create dir if not exists
			filename := filepath.Join(directory, path.Base(mUrl))
			os.MkdirAll(directory, os.ModePerm)

			// fmt.Println("Checking if " + filename + " exists ...")
			if _, err := os.Stat(filename); os.IsNotExist(err) {
				downloadHtml(mUrl, filename)
				fmt.Println(filename + " saved!")
			} else {
				fmt.Println(filename + " already exists!")
			}
		} else {
			cli.ShowAppHelp(c)
		}
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
