package main

import (
	"fmt"
)

const (
	englishPrefix = "Hello, "
	spanishPrefix = "Hola, "
	spanish       = "Spanish"
	french        = "French"
	frenchPrefix  = "Bonjour, "
)

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

	return languagePrefix(language) + name
}

func languagePrefix(language string) (prefix string) {
	switch language {
	case spanish:
		return spanishPrefix
	case french:
		return frenchPrefix
	default:
		return englishPrefix
	}
}

func main() {
	fmt.Println(Hello("e", "Spanish"))
}
