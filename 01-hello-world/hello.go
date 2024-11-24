package main

import "fmt"

const spanish = "Spanish"
const portuguese = "Portuguese"
const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const portugueseHelloPrefix = "Olá, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case portuguese:
		prefix = portugueseHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("João", ""))
}
