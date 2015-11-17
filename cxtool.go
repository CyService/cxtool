package main

import (
	"os"
	"github.com/codegangsta/cli"
	"github.com/keiono/cyjs-util-go/converter"
)

// File formats
const(
	csv string = "csv"
	tsv string = "tsv"
	cx string = "cx"
)

func main() {
	app := cli.NewApp()
	app.Name = "cytool"
	app.Usage = "Utility to convert Cytoscape.js JSON into many other formats."
	app.Version = "0.1.0"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: "source, s",
			Value: "source.cx",
			Usage: "Source file to be converted.",
		},

		cli.StringFlag{
			Name: "format, f",
			Value: "cx",
			Usage: "Source file format.  Default input file format is CX.",
		},
	}

	app.Action = func(c *cli.Context) {
		commandLineArgs := c.Args()

		inFileFormat := c.String("format")
		if inFileFormat == "" {
			inFileFormat = cx
		}

		// Two cases: Run from file or piped text stream
		if len(commandLineArgs) == 0 {
			runConversionStream()
		} else {
			source := commandLineArgs[0]
			runConversion(source)
		}
	}

	app.Run(os.Args)
}

func runConversion(source string) {
	var con converter.Converter
	con = converter.Cx2Cyjs{}
	con.Convert(source)
}

func runConversionStream() {
	var con converter.Converter
	con = converter.Cx2Cyjs{}
	con.ConvertFromStdin()
}
