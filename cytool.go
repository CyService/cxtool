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
			Name: "target, t",
			Value: "out.cyjs",
			Usage: "Target file name (i.e., destination file name).",
		},

		cli.StringFlag{
			Name: "format, f",
			Value: "cx",
			Usage: "Source file format.  Default is CX.",
		},
	}

	app.Action = func(c *cli.Context) {
		source := c.Args()[0]
		target := c.Args()[1]


		inFileFormat := c.String("format")

		if inFileFormat != "" {
//    		println("Format:", c.String("format"))
  		} else {
//    		println("CX format")
			inFileFormat = cx
  		}

		runConversion(source, target)
	}

	app.Run(os.Args)
}

func runConversion(source, target string) {
	var con converter.Converter

//	println("Source File: " + source)
//	println("Destination File: " + target)


	con = converter.Cx2Cyjs{}

	con.Convert(source, target)
}