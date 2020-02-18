package dragen

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"github.com/alee792/runeterra/proto"
	"github.com/dave/jennifer/jen"
	"github.com/pkg/errors"
)

// ReadCore a path to a Data Dragon bundle.
func ReadCore(path string) (*proto.Core, error) {
	// Read JSON.
	bb, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var core proto.Core
	if err := json.Unmarshal(bb, &core); err != nil {
		return nil, err
	}

	return &core, nil
}

// ParseCore turns a bundle into generated code.
func ParseCore(pkgName string, core *proto.Core) (*jen.File, error) {
	f := jen.NewFile(pkgName)
	f.Comment("Auto-generated from Data Dragon.")

	reg := regexp.MustCompile("[^a-zA-Z0-9]+")

	// Begin generation.
	cc := []jen.Code{}
	basePkg := "github.com/alee792/runeterra/proto"

	cc = append(cc, jen.Comment("Keywords"))

	for _, v := range core.GetKeywords() {
		cleanName := reg.ReplaceAllString(v.GetName(), "")
		cc = append(cc,
			jen.Id(fmt.Sprintf("Keyword%s", cleanName)).
				Op("=").
				Qual(basePkg, "Keyword").
				Values(jen.Dict{
					jen.Id("ID"):          unpackID(basePkg, v.GetNameRef(), v.GetName()),
					jen.Id("Description"): jen.Lit(v.GetDescription()),
				}).Op("\n"),
		)
	}

	cc = append(cc, jen.Comment("Regions"))
	for _, v := range core.GetRegions() {
		cc = append(cc,
			jen.Id(fmt.Sprintf("Region%s", strings.Replace(v.GetNameRef(), " ", "", -1))).
				Op("=").
				Qual(basePkg, "Region").
				Values(jen.Dict{
					jen.Id("ID"):               unpackID(basePkg, v.GetNameRef(), v.GetName()),
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
				Qual(basePkg, "SpellSpeed").
				Values(jen.Dict{
					jen.Id("\nID"): unpackID(basePkg, v.GetNameRef(), v.GetName()).Op(",\n"),
				}).Op("\n"),
		)
	}

	cc = append(cc, jen.Comment("Rarities"))
	for _, v := range core.GetRarities() {
		cc = append(cc,
			jen.Id(fmt.Sprintf("Rarity%s", strings.Replace(v.GetNameRef(), " ", "", -1))).
				Op("=").
				Qual(basePkg, "Rarity").
				Values(jen.Dict{
					jen.Id("\nID"): unpackID(basePkg, v.GetNameRef(), v.GetName()).Op(",\n"),
				}).Op("\n"),
		)
	}

	f.Var().Defs(cc...)

	return f, nil
}

// GenerateCore using a package name, reading from in and writing to out.
func GenerateCore(pkgName, in, out string) error {
	core, err := ReadCore(in)
	if err != nil {
		return errors.Wrap(err, "failed to read core")
	}

	f, err := ParseCore(pkgName, core)
	if err != nil {
		return errors.Wrap(err, "failed to parse core")
	}

	// Print to stdout if there's no output path.
	if out == "" {
		f.Println()
		return nil
	}

	if err := f.Save(out); err != nil {
		return errors.Wrapf(err, "failed to save to %s", out)
	}

	return nil
}

// ReadSet from a Data Dragon path.
func ReadSet(path string) ([]proto.Card, error) {
	bb, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var set []proto.Card
	if err := json.Unmarshal(bb, &set); err != nil {
		return nil, err
	}

	return set, nil
}

// ParseSet turns a set of cards into generated code.
func ParseSet(pkgName string, set []proto.Card) (*jen.File, error) {
	f := jen.NewFile(pkgName)
	f.Comment("Auto-generated from Data Dragon.")

	reg := regexp.MustCompile("[^a-zA-Z0-9]+")

	cc := []jen.Code{}
	basePkg := "github.com/alee792/runeterra/proto"
	cardDict := jen.Dict{}

	cc = append(cc, jen.Comment("Cards"))

	for _, v := range set {
		cleanName := reg.ReplaceAllString(v.GetName(), "")

		if v.GetRarity() == "None" {
			code := v.GetCardCode()
			cleanName = fmt.Sprintf("%s%s", cleanName, code[len(code)-2:])
		}

		varName := fmt.Sprintf("Card%s", cleanName)
		c := jen.Dict{
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
			jen.Id("Assets"):             jen.Index().Op("*").Qual(basePkg, "Asset").Values(unpackAssets(basePkg, v)...),
		}
		cc = append(cc,
			jen.Id(varName).
				Op("=").
				Qual(basePkg, "Card").
				Values(c).
				Op("\n"),
		)
		cardDict[jen.Lit(v.GetCardCode())] = jen.Id(varName)
	}

	f.Var().Defs(cc...)
	f.Func().Id("Cards").Params().Id("map[string]proto.Card").Block(
		jen.Return(jen.Map(jen.String()).Id("proto.Card").Values(cardDict)),
	)

	return f, nil
}

// GenerateSet using a package name, reading from in and writing to out.
func GenerateSet(pkgName, in, out string) error {
	set, err := ReadSet(in)
	if err != nil {
		return errors.Wrap(err, "failed to read set")
	}

	f, err := ParseSet(pkgName, set)
	if err != nil {
		return errors.Wrap(err, "failed to parse set")
	}

	// Print to stdout if there' no output path.
	if out == "" {
		f.Println()
		return nil
	}

	if err := f.Save(out); err != nil {
		return errors.Wrapf(err, "failed to save to %s", out)
	}

	return nil
}

func unpackAssets(pkg string, v proto.Card) []jen.Code {
	ss := []jen.Code{}

	aa := v.GetAssets()
	if len(aa) == 1 && aa[0].GetFullAbsolutePath() == "" {
		return nil
	}

	for _, a := range aa {
		ss = append(ss,
			jen.Op("\n&").Qual(pkg, "Asset").Values(jen.Dict{
				jen.Id("GameAbsolutePath"): jen.Lit(a.GetGameAbsolutePath()),
				jen.Id("FullAbsolutePath"): jen.Lit(a.GetFullAbsolutePath()),
			}),
		)
	}

	return ss
}

func unpackID(pkg, nameRef, name string) *jen.Statement {
	return jen.Op("&").Qual(pkg, "ID").Values(jen.Dict{
		jen.Id("NameRef"): jen.Lit(nameRef),
		jen.Id("Name"):    jen.Lit(name),
	})
}

func wrapLit(ss []string) []jen.Code {
	ll := []jen.Code{}
	for _, s := range ss {
		ll = append(ll, jen.Lit(s))
	}

	return ll
}
