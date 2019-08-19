package main

import "fmt"

const spanish = "Spanish"
const french = "French"
const helloEnglishPrefix = "Hello, "
const helloSpanishPrefix = "Hola, "
const helloFrenchPrefix = "Bonjour, "

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return getCorrespondantPrefixByLanguage(language) + name
}

func getCorrespondantPrefixByLanguage(language string) (prefix string) {

	switch language {
	case spanish:
		prefix = helloSpanishPrefix
	case french:
		prefix = helloFrenchPrefix
	default:
		prefix = helloEnglishPrefix
	}

	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
