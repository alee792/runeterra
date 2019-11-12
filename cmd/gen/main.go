package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	pb "github.com/alee792/runeterra/proto"
	"github.com/dave/jennifer/jen"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var (
		path = kingpin.Flag("path", "path to the Data Dragon's data file").Short('p').String()
	)

	kingpin.Parse()

	bb, err := ioutil.ReadFile(*path)
	if err != nil {
		panic(err)
	}

	var core pb.Core
	if err := json.Unmarshal(bb, &core); err != nil {
		panic(err)
	}

	f := jen.NewFile("datadragon")

	f.Comment("Auto-generated from Data Dragon en_us.")

	var cc []jen.Code
	pkg := "github.com/alee792/runeterra/proto"

	cc = append(cc, jen.Comment("Keywords"))
	for _, v := range core.GetKeywords() {
		cc = append(cc,
			jen.Id(fmt.Sprintf("Keyword%s", strings.Replace(v.GetNameRef(), " ", "", -1))).
				Op("=").
				Qual(pkg, "Keyword").
				Values(jen.Dict{
					jen.Id("ID"):          embeddedID(pkg, v.GetNameRef(), v.GetName()),
					jen.Id("Description"): jen.Lit(v.GetDescription()),
				}).Op("\n"),
		)
	}

	cc = append(cc, jen.Comment("Regions"))
	for _, v := range core.GetRegions() {
		cc = append(cc,
			jen.Id(fmt.Sprintf("Region%s", strings.Replace(v.GetNameRef(), " ", "", -1))).
				Op("=").
				Qual(pkg, "Region").
				Values(jen.Dict{
					jen.Id("ID"):               embeddedID(pkg, v.GetNameRef(), v.GetName()),
					jen.Id("Abbreviation"):     jen.Lit(v.GetAbbreviation()),
					jen.Id("IconAbsolutePath"): jen.Lit(v.GetIconAbsolutePath()),
				}).Op("\n"),
		)
	}

	cc = append(cc, jen.Comment("Spell Speeds"))
	for _, v := range core.GetSpellSpeeds() {
		cc = append(cc,
			jen.Id(fmt.Sprintf("SpellSpeed%s", strings.Replace(v.GetNameRef(), " ", "", -1))).
				Op("=").
				Qual(pkg, "SpellSpeed").
				Values(jen.Dict{
					jen.Id("\nID"): embeddedID(pkg, v.GetNameRef(), v.GetName()).Op(",\n"),
				}).Op("\n"),
		)
	}

	cc = append(cc, jen.Comment("Rarities"))
	for _, v := range core.GetRarities() {
		cc = append(cc,
			jen.Id(fmt.Sprintf("Rarity%s", strings.Replace(v.GetNameRef(), " ", "", -1))).
				Op("=").
				Qual(pkg, "Rarity").
				Values(jen.Dict{
					jen.Id("\nID"): embeddedID(pkg, v.GetNameRef(), v.GetName()).Op(",\n"),
				}).Op("\n"),
		)
	}

	f.Var().Defs(cc...)

	if err := f.Save("pkg/datadragon/datadragon.go"); err != nil {
		panic(err)
	}
}

func embeddedID(pkg, nameRef, name string) *jen.Statement {
	return jen.Op("&").Qual(pkg, "ID").Values(jen.Dict{
		jen.Id("NameRef"): jen.Lit(nameRef),
		jen.Id("Name"):    jen.Lit(name),
	})
}
