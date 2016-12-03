package main

import (
	"log"
	"math/rand"
	"reflect"
	"regexp"
	"strings"

	"github.com/parnurzeal/gorequest"
)

func expandUrl(uri string) (string, []error) {
	request := gorequest.New()
	res, _, err := request.Get(uri).
		Set("User-Agent", o.UserAgents.sample()).
		End()

	return res.Request.URL.String(), err
}

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
	i = -1
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
	return SampleString(s).(string)
}

func collectKeys(s interface{}) []string {
	rKeys := reflect.ValueOf(s).MapKeys()
	var keys []string
	for _, eachKey := range rKeys {
		keys = append(keys, eachKey.String())
	}
	return keys
}

func (s OneOfManyMaps) sample() string {
	keys := collectKeys(s)
	randomKey := SampleString(keys).(string)
	return SampleString(s[randomKey]).(string)
}

// Gives a sample string
func SampleString(s interface{}) interface{} {
	vs := reflect.ValueOf(s)
	index := rand.Intn(vs.Len())
	val := vs.Index(index)
	return val.Interface()
}

func Truncate(s string, i int) string {
	runes := []rune(s)
	if len(runes) > i {
		return string(runes[:i])
	}
	return s
}
