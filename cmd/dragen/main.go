package main

import (
	"os"

	"github.com/alee792/runeterra/pkg/dragen"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var (
		app = kingpin.New("dragen", "Generate Go types from Data Dragon")

		pkg = app.Flag("pkg", "package name for the output file").Default("datadragon").String()

		dTyp = app.Arg("type", "Data Dragon type").Required().HintOptions("core", "set").String()
		in   = app.Arg("in", "path to Data Dragon file").Required().String()
		out  = app.Arg("out", "path to write Go types to file").String()
		// raw  = app.Flag("raw", "unpack Data Dragon into flat structs").Short('r').Bool()
	)

	app.HelpFlag.Short('h')
	kingpin.MustParse(app.Parse(os.Args[1:]))

	switch *dTyp {
	case "core":
		if err := dragen.GenerateCore(*pkg, *in, *out); err != nil {
			panic(err)
		}
	case "set":
		if err := dragen.GenerateSet(*pkg, *in, *out); err != nil {
			panic(err)
		}
	default:
	}
}
