package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
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
	SleightOfHand  int `yaml:sleight_of_hand"`
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

func (c *Character) executor(in string) {
	in = strings.TrimSpace(in)
	blocks := strings.Split(in, " ")
	switch blocks[0] {
	case "help":
		fmt.Println(HELP)
		return
	case "get":
		if len(blocks) > 1 {
			// the lack of flexibility in case statements is troubling
			// hence we aren't using one. It's a scalability thing, you know?
			if itemInSlice(blocks[1], stringSliceAttrs) {
				c.printStringSliceAttr(blocks[1])
				return
			} else if itemInSlice(blocks[1], stringAttrs) {
				c.printStringAttr(blocks[1])
				return
			} else {
				// lol better hope the attr is in here
				c.printIntAttr(blocks[1])
				return
			}
		} else {
			return
		}
	case "set":
		if len(blocks) > 1 {
			if len(blocks) == 2 {
				fmt.Println("Not enough arguments. Usage: set <attr> <value>")
				return
			} else if len(blocks) > 2 {
				if itemInSlice(blocks[1], stringAttrs) {
					c.setStringAttr(blocks[1], blocks[2])
				} else if itemInSlice(blocks[1], intAttrs) {
					i, err := strconv.Atoi(blocks[2])
					if err != nil {
						fmt.Println("Value must be an engineer")
					}
					c.setIntAttr(blocks[1], i)
				}
				return
			}
		}
	case "exit":
		os.Exit(0)
	default:
		return
	}
}

func itemInSlice(item string, slice []string) bool {
	for _, v := range slice {
		if item == v {
			return true
		}
	}
	return false
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
