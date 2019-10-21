package main

import (
	"c1freitas/captable/cmd"
	"c1freitas/captable/data"
	"errors"
	"log"
	"os"
	"time"

	"github.com/shopspring/decimal"
	"github.com/urfave/cli"
)

func main() {

	var filterDateStr string
	var filePath string

	// overide the default marshalling format
	decimal.MarshalJSONWithoutQuotes = true

	app := cli.NewApp()
	app.Name = "captable - A program for generating Capitalization Tables"
	app.Version = "0.0.1"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "date,d",
			Usage:       "The date `YYYY-MM-DD` used when generating the captable. Optional, if not supplied todays date is used.",
			Destination: &filterDateStr,
			Required:    false,
		},
		cli.StringFlag{
			Name:        "file,f",
			Usage:       "The `FILE` used to generate the Capitalization Table, required",
			Destination: &filePath,
			Required:    true,
		},
	}

	app.Action = func(c *cli.Context) error {
		filterDate := time.Now()
		var err error
		if len(filterDateStr) > 0 {
			filterDate, err = time.Parse(data.DateFormat, filterDateStr)
		}
		if err != nil {
			return errors.New("Date is in an incorrect format, must be YYYY-MM-DD")
		}
		capTable, err := cmd.ProcessFile(filePath, filterDate)
		if err != nil {
			return err
		}
		// Run the parse cmd
		jsonBytes, err := cmd.RenderData(capTable)

		// write out the capTable
		_, err = log.Writer().Write(jsonBytes)
		return err
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
