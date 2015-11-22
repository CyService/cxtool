/*

	cxtool: A commandline tool to convert CX and related files

	by Keiichiro Ono (kono at ucsd edu)

	(c) 2015 The Cytoscape Consortium

	MIT License

*/
package main

import (
	builder "github.com/idekerlab/cxtool/appbuilder"
	"os"
)

func main() {

	// Create new cxtool app instance
	app := builder.BuildApp()

	// Run with command line options...
	app.Run(os.Args)
}
