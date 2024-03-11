package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "
const frenchHelloPrefix = "Bonjour, "
const Spanish = "spanish"
const French = "french"

func Hello(name string, language string) string {

	if name == "" {
		name = "World"
	}

  

	return languagePrefix(language) + name
}

func languagePrefix(language string) (prefix string) {

  switch language{
    case Spanish:
      prefix = spanishHelloPrefix
    case French:
      prefix = frenchHelloPrefix
    default:
      prefix = englishHelloPrefix
    }

    return 
  }

func main() {
	fmt.Println(Hello("world",""))
}

