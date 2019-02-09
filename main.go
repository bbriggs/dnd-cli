package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"gopkg.in/yaml.v2"
)

type Character struct {
	// Basic stuff
	Name      string   `yaml:"name"`
	Level     int      `yaml:"level"`
	XP        int      `yaml:"experience"`
	Race      string   `yaml:"race"`
	Class     string   `yaml:"class"`
	Alignment string   `yaml:"alignment"`
	Age       int      `yaml:"age"`
	Height    int      `yaml:"height"`
	Size      string   `yaml:"size"`
	Eyes      string   `yaml:"eyes"`
	Skin      string   `yaml:"skin"`
	Weight    string   `yaml:"weight"`
	Hair      string   `yaml:"hair"`
	Traits    []string `yaml:"traits"`
	Ideals    []string `yaml:"ideals"`
	Bonds     []string `yaml:"bonds"`
	Flaws     []string `yaml:"flaws"`
	Features  []string `yaml:"features"`

	// Gear
	Items     []string `yaml:"items"`
	Equipment []string `yaml:"equipment"`
	CP        int      `yaml:"cp"`
	SP        int      `yaml:"sp"`
	EP        int      `yaml:"ep"`
	GP        int      `yaml:"gp"`
	PP        int      `yaml:"pp"`

	// Attributes
	Speed int `yaml:"speed"`
	Str   int `yaml:"str"`
	Dex   int `yaml:"dex"`
	Con   int `yaml:"con"`
	Int   int `yaml:"int"`
	Wis   int `yaml:"wis"`
	Cha   int `yaml:"cha"`

	// Skills
	Acrobatics     int `yaml:"acrobatics"`
	AnimalHandling int `yaml:"animal_handling"`
	Arcana         int `yaml:"arcana"`
	Athletics      int `yaml:"athletics"`
	Deception      int `yaml:"deception"`
	History        int `yaml:"history"`
	Insight        int `yaml:"insight"`
	Intimidation   int `yaml:"intimidation"`
	Investigation  int `yaml:"investigation"`
	Medicine       int `yaml:"medicine"`
	Nature         int `yaml:"nature"`
	Perception     int `yaml:"perception"`
	Persuasion     int `yaml:"persuasion"`
	Religion       int `yaml:"religion"`
	SlightOfHand   int `yaml:sleigh_of_hand"`
	Stealth        int `yaml:"stealth"`
	Survival       int `yaml:"survival"`

	// Other
	PassiveWisdom int `yaml:"passive_wisdom"`

	// Combat
	AC     int `yaml:"ac"`
	HP     int `yaml:"hp"`
	TempHP int `yaml:"temp_hp"`
}

const HELP = "Generic help text"

func (c *Character) LivePrefix() (string, bool) {
	if c.Name == "" {
		return "(none) > ", true
	}
	return c.Name + "> ", true
}

var rootCmds = []prompt.Suggest{
	{Text: "exit", Description: "Exit D&D CLI"},
	{Text: "get", Description: "Display various information about your character"},
	{Text: "set", Description: "Set or update various attributes"},
	{Text: "add", Description: "Increment a value. Use this for tracking things that fluctuate frequently, such as HP."},
	{Text: "sub", Description: "Decrement a value. Use this for tracking things that fluctuate frequently, such as HP."},
}

func (c *Character) executor(in string) {
	in = strings.TrimSpace(in)
	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "help":
		fmt.Println(HELP)
		return
	case "get":
		if len(blocks) > 1 {
			switch blocks[1] {
			case "traits", "ideals", "bonds", "flaws", "features", "items", "equipment":
				c.printStringSliceAttr(blocks[1])
				return
			case "name", "race", "class", "alignment", "size", "eyes", "skin", "weight", "hair":
				c.printStringAttr(blocks[1])
				return
			default:
				attr, err := c.getAttr(blocks[1])
				if err != nil {
					fmt.Println(err.Error())
				} else {
					fmt.Println(attr)
				}
			}
		} else {
			return
		}
	case "exit":
		os.Exit(0)
	default:
		return
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

func (c *Character) printItems() {
	for _, i := range c.Items {
		fmt.Println("- " + i)
	}
}

func (c *Character) getAttr(attr string) (int, error) {
	attrs := map[string]int{
		"hp":           c.HP,
		"ac":           c.AC,
		"speed":        c.Speed,
		"str":          c.Str,
		"strength":     c.Str,
		"dex":          c.Dex,
		"dexterity":    c.Dex,
		"con":          c.Con,
		"constitution": c.Con,
		"int":          c.Int,
		"intelligence": c.Int,
		"wis":          c.Wis,
		"wisdom":       c.Wis,
		"cha":          c.Cha,
		"charisma":     c.Cha,
	}

	for k, v := range attrs {
		if k == attr {
			return v, nil
		}
	}
	return 0, fmt.Errorf("Attribute not found.")
}

func readState(config string) (*Character, error) {
	c := Character{}
	b, err := ioutil.ReadFile(config)
	if err != nil {
		return &c, err
	}

	err = yaml.Unmarshal([]byte(b), &c)
	return &c, err
}

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

func main() {
	fmt.Println("D&D CLI -- Because D&D alone wasn't nerdy enough")
	c, err := readState("config.yaml")
	if err != nil {
		panic(err)
	}
	p := prompt.New(
		c.executor,
		completer,
		prompt.OptionLivePrefix(c.LivePrefix),
		prompt.OptionTitle("dnd-util"),
	)
	p.Run()
}
