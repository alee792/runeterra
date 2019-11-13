package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"regexp"
	"strings"

	pb "github.com/alee792/runeterra/proto"
	"github.com/dave/jennifer/jen"
	"github.com/pkg/errors"
	kingpin "gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	var (
		app  = kingpin.New("dragen", "Generate Go types from Data Dragon")
		dTyp = app.Arg("type", "Data Dragon type").Required().HintOptions("core", "set").String()
		in   = app.Arg("in", "path to Data Dragon file").Required().String()
		out  = app.Arg("out", "path to write Go types to file").String()
	)

	app.HelpFlag.Short('h')
	kingpin.MustParse(app.Parse(os.Args[1:]))

	switch *dTyp {
	case "core":
		if err := generateCore(*in, *out); err != nil {
			panic(err)
		}
	case "set":
		if err := generateSet(*in, *out); err != nil {
			panic(err)
		}
	default:
	}

}

func generateCore(in, out string) error {
	bb, err := ioutil.ReadFile(in)
	if err != nil {
		return err
	}

	var core pb.Core
	if err := json.Unmarshal(bb, &core); err != nil {
		return err
	}

	f := jen.NewFile("datadragon")

	f.Comment("Auto-generated from Data Dragon en_us.")

	var cc []jen.Code
	pkg := "github.com/alee792/runeterra/proto"

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return err
	}

	cc = append(cc, jen.Comment("Keywords"))
	for _, v := range core.GetKeywords() {
		cleanName := reg.ReplaceAllString(v.GetName(), "")
		cc = append(cc,
			jen.Id(fmt.Sprintf("Keyword%s", cleanName)).
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

	if out == "" {
		f.Println()
		return nil
	}

	if err := f.Save(out); err != nil {
		return errors.Wrapf(err, "failed to save to %s", out)
	}

	return nil
}

func generateSet(in, out string) error {
	bb, err := ioutil.ReadFile(in)
	if err != nil {
		return err
	}

	var set []pb.Card
	if err := json.Unmarshal(bb, &set); err != nil {
		return err
	}

	f := jen.NewFile("datadragon")

	f.Comment("Auto-generated from Data Dragon en_us.")

	var cc []jen.Code
	pkg := "github.com/alee792/runeterra/proto"

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		return err
	}

	cc = append(cc, jen.Comment("Cards"))
	for _, v := range set {
		cleanName := reg.ReplaceAllString(v.GetName(), "")
		if v.GetRarity() == "None" {
			code := v.GetCardCode()
			cleanName = fmt.Sprintf("%s%s", cleanName, code[len(code)-2:])
		}
		cc = append(cc,
			jen.Id(fmt.Sprintf("Card%s", cleanName)).
				Op("=").
				Qual(pkg, "Card").
				Values(jen.Dict{
					jen.Id("CardCode"):           jen.Lit(v.GetCardCode()),
					jen.Id("Name"):               jen.Lit(v.GetName()),
					jen.Id("Region"):             jen.Lit(v.GetRegion()),
					jen.Id("RegionRef"):          jen.Lit(v.GetRegionRef()),
					jen.Id("Attack"):             jen.Lit(int(v.GetAttack())),
					jen.Id("Cost"):               jen.Lit(int(v.GetCost())),
					jen.Id("Health"):             jen.Lit(int(v.GetHealth())),
					jen.Id("Description"):        jen.Lit(v.GetDescription()),
					jen.Id("DescriptionRaw"):     jen.Lit(v.GetDescriptionRaw()),
					jen.Id("FlavorText"):         jen.Lit(v.GetFlavorText()),
					jen.Id("ArtistName"):         jen.Lit(v.GetArtistName()),
					jen.Id("SpellSpeed"):         jen.Lit(v.GetSpellSpeed()),
					jen.Id("Rarity"):             jen.Lit(v.GetRarity()),
					jen.Id("RarityRef"):          jen.Lit(v.GetRegionRef()),
					jen.Id("Subtype"):            jen.Lit(v.GetSubtype()),
					jen.Id("Supertype"):          jen.Lit(v.GetSupertype()),
					jen.Id("Collectible"):        jen.Lit(v.GetCollectible()),
					jen.Id("AssociatedCards"):    jen.Index().String().Values(wrapLit(v.GetAssociatedCards())...),
					jen.Id("AssociatedCardRefs"): jen.Index().String().Values(wrapLit(v.GetAssociatedCardRefs())...),
					jen.Id("Keywords"):           jen.Index().String().Values(wrapLit(v.GetKeywords())...),
					jen.Id("KeywordRefs"):        jen.Index().String().Values(wrapLit(v.GetKeywordRefs())...),
					// jen.Id("Assets"):             jen.List(v.GetAssets()),

				}).Op("\n"),
		)
	}

	f.Var().Defs(cc...)

	if out == "" {
		f.Println()
		return nil
	}

	if err := f.Save(out); err != nil {
		return errors.Wrapf(err, "failed to save to %s", out)
	}

	return nil
}

func embeddedID(pkg, nameRef, name string) *jen.Statement {
	return jen.Op("&").Qual(pkg, "ID").Values(jen.Dict{
		jen.Id("NameRef"): jen.Lit(nameRef),
		jen.Id("Name"):    jen.Lit(name),
	})
}

func wrapLit(ss []string) []jen.Code {
	var ll []jen.Code
	for _, s := range ss {
		ll = append(ll, jen.Lit(s))
	}

	return ll
}
