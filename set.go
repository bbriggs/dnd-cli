package main

import (
	"fmt"
)

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
