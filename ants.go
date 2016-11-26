package main

import (
	"io/ioutil"
	"log"
	"os"

	"gopkg.in/yaml.v2"
)
var o Words

type Words struct {
	Hashable      OneOfManyStrings `json:"hashable"`
	UserAgents      OneOfManyStrings `json:"user_agents"`
}
type OneOfManyStrings []string


func init() {
	words, err := ioutil.ReadFile("./ants.yml")
	if err != nil {
		log.Fatalln("Unable to load words")
		os.Exit(1)
	}
	log.Println("Setting up Generator")
	err = yaml.Unmarshal((words), &o)
	if err != nil {
		log.Println(err)
		log.Fatalln("There is a syntax error in words.yml")
		os.Exit(1)
	}
}

func main() {

}
