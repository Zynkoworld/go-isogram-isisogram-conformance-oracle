package main

import (
	__json "encoding/json"
	__fmt "fmt"
)



import "strings"

func IsIsogram(word string) bool {
	lowerWord := strings.ToLower(word)
	for i, r := range lowerWord {
		if r != ' ' && r != '-' && i < strings.LastIndex(lowerWord, string(r)) {
			return false
		}
	}
	return true
}

type R struct {
	Ok bool        `json:"ok"`
	V  interface{} `json:"v"`
}

func main() {
	inputs := []string{"", "isogram", "eleven", "zzyzx", "subdermatoglyphic", "Alphabet", "alphAbet", "thumbscrew-japingly", "thumbscrew-jappingly", "six-year-old", "Emily Jung Schwartzkopf", "accentor", "angola", "up-to-date"}
	out := []R{}
	for _, x := range inputs {
		func() {
			defer func() { if r := recover(); r != nil { out = append(out, R{false, __fmt.Sprint(r)}) } }()
			out = append(out, R{true, IsIsogram(x)})
		}()
	}
	b, _ := __json.Marshal(map[string]interface{}{"out": out})
	__fmt.Println(string(b))
}
