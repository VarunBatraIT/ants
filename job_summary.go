package main

import "strings"

type JobSummary struct {
	LookingFor    OneOfManyStrings `yaml:"looking_for"`
	MarketingText OneOfManyStrings `yaml:"marketing_text"`
	Growth        OneOfManyStrings `yaml:"growth"`
	JoinUs        OneOfManyStrings `yaml:"join_us"`
	Contact       OneOfManyStrings `yaml:"contact"`
}

func (j JobSummary) generate() string {
	var sentence []string
	sentence = append(sentence, j.LookingFor.sample())
	sentence = append(sentence, j.MarketingText.sample())
	sentence = append(sentence, j.Growth.sample())
	sentence = append(sentence, j.JoinUs.sample())
	sentence = append(sentence, j.Contact.sample())
	return strings.Join(sentence, " ")
}

func GetJobSummary() string {
	return o.JobSummary.generate()
}
