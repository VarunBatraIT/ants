package main

import "strings"

//Hashes in twitter
var TWITTER_HASH = 3
//Hashes in linked
var LINKEDIN_HASH = 2

func RemoveHashes(str string) string {
	words := strings.Split(str, " ")
	for i := 0; i < len(words); i++ {
		if strings.Index(words[i], "#") == 0 {
			words[i] = strings.Replace(words[i], "#", "", -1)
		}
	}
	str = strings.Join(words, " ")
	return str
}

// If total = 0 no limit on hash
// If total = -1 then no hashes will be added.
func AddHashes(str string, total int) string {
	//hashes must be removed once! (O2)
	str = RemoveHashes(str)
	words := strings.Split(str, " ")
	if total == 0 {
		total = len(words)
	}
	for i := 0; i < len(words); i++ {
		ok, _ := in_array(alphanumericsmall(words[i]), o.Hashable)
		if ok && total > -1 {
			words[i] = "#" + words[i]
			total--
		}
	}
	str = strings.Join(words, " ")
	return str
}
