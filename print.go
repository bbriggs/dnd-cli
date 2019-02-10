package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

func dumpState(c Character) error {
	d, err := yaml.Marshal(&c)
	if err != nil {
		fmt.Println("Unable to save state!")
		fmt.Println(err.Error())
		return err
	}

	err = ioutil.WriteFile(fmt.Sprintf("%s.yaml", c.Name), d, 0666)
	if err != nil {
		fmt.Println("Unable to save state!")
		fmt.Println(err.Error())
		return err
	}

	return nil
}

func (c *Character) printIntAttr(attr string) {
	attrs := map[string]int{
		"ac":             c.AC,
		"cha":            c.Cha,
		"charisma":       c.Cha,
		"con":            c.Con,
		"constitution":   c.Con,
		"dex":            c.Dex,
		"dexterity":      c.Dex,
		"hp":             c.HP,
		"int":            c.Int,
		"intelligence":   c.Int,
		"level":          c.Level,
		"speed":          c.Speed,
		"str":            c.Str,
		"strength":       c.Str,
		"wis":            c.Wis,
		"wisdom":         c.Wis,
		"xp":             c.XP,
		"exp":            c.XP,
		"age":            c.Age,
		"height":         c.Height,
		"cp":             c.CP,
		"sp":             c.SP,
		"ep":             c.EP,
		"gp":             c.GP,
		"pp":             c.PP,
		"acrobatics":     c.Acrobatics,
		"animalHandling": c.AnimalHandling,
		"arcana":         c.Arcana,
		"athletics":      c.Athletics,
		"deception":      c.Deception,
		"history":        c.History,
		"insight":        c.Insight,
		"intimidation":   c.Intimidation,
		"investigation":  c.Investigation,
		"medicine":       c.Medicine,
		"nature":         c.Nature,
		"perception":     c.Perception,
		"persuasion":     c.Persuasion,
		"religion":       c.Religion,
		"sleightOfHand":  c.SleightOfHand,
		"stealth":        c.Stealth,
		"survival":       c.Survival,
		"passiveWisdom":  c.PassiveWisdom,
		"tempHP":         c.TempHP,
	}

	for k, v := range attrs {
		if k == attr {
			fmt.Println(v)
		}
	}
}

func (c *Character) printStringAttr(attr string) {
	attrs := map[string]string{
		"name":      c.Name,
		"race":      c.Race,
		"class":     c.Class,
		"alignment": c.Alignment,
		"size":      c.Size,
		"eyes":      c.Eyes,
		"skin":      c.Skin,
		"weight":    c.Weight,
		"hair":      c.Hair,
	}
	for k, v := range attrs {
		if attr == k {
			fmt.Println(v)
		}
	}
}

func (c *Character) printStringSliceAttr(attr string) {
	attrs := map[string][]string{
		"traits":    c.Traits,
		"ideals":    c.Ideals,
		"bonds":     c.Bonds,
		"flaws":     c.Flaws,
		"features":  c.Features,
		"items":     c.Items,
		"equipment": c.Equipment,
	}
	for k, v := range attrs {
		if k == attr {
			for _, i := range v {
				fmt.Println("- " + i)
			}
		}
	}
}
