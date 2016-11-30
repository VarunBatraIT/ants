package main

import "strings"

type SelfMarketing struct {
	Who     OneOfManyStrings `json:"who"`
	What    OneOfManyStrings `json:"what"`
	Contact OneOfManyStrings `json:"contact"`
}

func (l SelfMarketing) generate() string {
	var sentence []string
	sentence = append(sentence, l.Who.sample())
	sentence = append(sentence, l.What.sample())
	sentence = append(sentence, l.Contact.sample())
	return strings.Join(sentence, " ")
}
func GetSelfMarketing() string {
	return o.SelfMarketing.generate()
}

func PostSelfMarketingOnLinkedin() {
	selfMarketing := GetSelfMarketing()
	linkedin := Linkedin{comment: selfMarketing}
	PostOnLinkedin(linkedin)
}
