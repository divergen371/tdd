package main

import (
	"fmt"
)

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frechHelloPrefix = "Bonjour, "
const french = "French"
const spanish = "Spanish"

func Hello(name string, language string) string {
	if name == "" {
		return englishHelloPrefix + "World"
	}

	if language == spanish {
		return spanishHelloPrefix + name
	}

	if language == french {
		return frechHelloPrefix + name
	}
	return englishHelloPrefix + name
}

func main() {
	fmt.Println(Hello("Atsushi", "French"))
}
