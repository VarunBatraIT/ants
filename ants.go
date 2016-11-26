package main

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)

var c Credentials
var o Ants

type Credentials struct {
	Linkedin OneOfOneMap
	Facebook OneOfOneMap
	Twitter  OneOfOneMap
}

type Ants struct {
	Hashable   OneOfManyStrings
	UserAgents OneOfManyStrings `yaml:"user_agents"`
	Blogs      OneOfManyMaps
}
type OneOfManyStrings []string
type OneOfManyMaps map[string][]string
type OneOfOneMap map[string]string

func readSettings(variable interface{}, fileName string) {
	words, err := ioutil.ReadFile(fileName)
	log.Println("Loading Settings!")
	if err != nil {
		log.Fatalln("Unable to load settings")
		os.Exit(1)
	}
	err = yaml.Unmarshal((words), variable)
	if err != nil {
		log.Println(err)
		log.Fatalln("There is a syntax error in words.yml")
		os.Exit(1)
	}
}

func init() {
	readSettings(&o, "ants.yml")
	readSettings(&c, "credentials.yml")
}

func main() {

}
