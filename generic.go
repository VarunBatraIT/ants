package main

import (
	"log"
	"math/rand"
	"reflect"
	"regexp"
	"strings"
)

func in_array(v interface{}, in interface{}) (ok bool, i int) {
	val := reflect.Indirect(reflect.ValueOf(in))
	switch val.Kind() {
	case reflect.Slice, reflect.Array:
		for ; i < val.Len(); i++ {
			if ok = v == val.Index(i).Interface(); ok {
				return
			}
		}
	}
	return
}
func alphanumericsmall(str string) string {
	reg, err := regexp.Compile("[^A-Za-z0-9]+")
	if err != nil {
		log.Fatal(err)
	}
	safe := reg.ReplaceAllString(str, "")
	safe = strings.ToLower(safe)
	return safe
}

type sample interface {
	sample()
}

func (s OneOfManyStrings) sample() string {
	return SampleString(s)
}

func collectKeys(s interface{}) []string {
	log.Println(reflect.TypeOf(s))
	rKeys := reflect.ValueOf(s).MapKeys()
	var keys []string
	for _,eachKey := range rKeys {
		keys = append(keys,eachKey.String())
	}
	return keys;
}

func (s OneOfManyMaps) sample() string {
	keys := collectKeys(s)
	randomKey := SampleString(keys)
	return SampleString(s[randomKey])
}

// Gives a sample string
func SampleString(s []string) string {
	return s[rand.Intn(len(s))]
}


func Truncate(s string, i int) string {
	runes := []rune(s)
	if len(runes) > i {
		return string(runes[:i])
	}
	return s
}

