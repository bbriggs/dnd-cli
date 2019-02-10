package main

func (c *Character) setStringAttr(attr string, val string) {
	attrs := map[string]*string{
		"name":      &c.Name,
		"race":      &c.Race,
		"class":     &c.Class,
		"alignment": &c.Alignment,
		"size":      &c.Size,
		"eyes":      &c.Eyes,
		"skin":      &c.Skin,
		"weight":    &c.Weight,
		"hair":      &c.Hair,
	}
	for k, v := range attrs {
		if attr == k {
			*v = val
		}
	}
}

func (c *Character) setIntAttr(attr string, val int) {
	attrs := map[string]*int{
		"ac":             &c.AC,
		"cha":            &c.Cha,
		"charisma":       &c.Cha,
		"con":            &c.Con,
		"constitution":   &c.Con,
		"dex":            &c.Dex,
		"dexterity":      &c.Dex,
		"hp":             &c.HP,
		"int":            &c.Int,
		"intelligence":   &c.Int,
		"level":          &c.Level,
		"speed":          &c.Speed,
		"str":            &c.Str,
		"strength":       &c.Str,
		"wis":            &c.Wis,
		"wisdom":         &c.Wis,
		"xp":             &c.XP,
		"exp":            &c.XP,
		"age":            &c.Age,
		"height":         &c.Height,
		"cp":             &c.CP,
		"sp":             &c.SP,
		"ep":             &c.EP,
		"gp":             &c.GP,
		"pp":             &c.PP,
		"acrobatics":     &c.Acrobatics,
		"animalHandling": &c.AnimalHandling,
		"arcana":         &c.Arcana,
		"athletics":      &c.Athletics,
		"deception":      &c.Deception,
		"history":        &c.History,
		"insight":        &c.Insight,
		"intimidation":   &c.Intimidation,
		"investigation":  &c.Investigation,
		"medicine":       &c.Medicine,
		"nature":         &c.Nature,
		"perception":     &c.Perception,
		"persuasion":     &c.Persuasion,
		"religion":       &c.Religion,
		"sleightOfHand":  &c.SleightOfHand,
		"stealth":        &c.Stealth,
		"survival":       &c.Survival,
		"passiveWisdom":  &c.PassiveWisdom,
		"tempHP":         &c.TempHP,
	}

	for k, v := range attrs {
		if k == attr {
			*v = val
		}
	}
}
