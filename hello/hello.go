package main

import (
	"fmt"
)

const prefixDefault = "hello "
const prefixJapanese = "konnichiwa "
const prefixFrench = "bonjour "
const langJapanese = "japanese"
const langFrench = "french"

// Hello is greeting function
func Hello(name string, lang string) string {
	if name == "" {
		name = "world"
	}

	return getPrefix(lang) + name
}

func getPrefix(lang string) (prefix string) {
	switch lang {
	case langJapanese:
		prefix = prefixJapanese
	case langFrench:
		prefix = prefixFrench
	default:
		prefix = prefixDefault
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}