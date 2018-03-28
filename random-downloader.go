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

func downloadHtml(url string, directory string, retries int) {

	if retries == 0 {
		return
	}
	var filename string

	// Create dir if not exists
	os.MkdirAll(directory, os.ModePerm)

	// fmt.Println("Checking if " + filename + " exists ...")
	if _, err := os.Stat(filename); os.IsNotExist(err) {

		resp, err := http.Get(url)
		if err != nil {
			panic(err)
		}
		defer resp.Body.Close()

		// Determine final filename
		finalURL := resp.Request.URL.Path
		filename := filepath.Join(directory, path.Base(finalURL))

		f, err := os.Create(filename)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		io.Copy(f, resp.Body)

		fmt.Println(filename + " saved!")
	} else {
		fmt.Println(filename + " already exists! Retries left: " + string(retries-1))
		downloadHtml(url, directory, retries-1)
	}

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
			Value:       1,
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
			Value:       3,
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
			for i := 0; i < numpages; i++ {
				downloadHtml(mUrl, directory, retries)
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
