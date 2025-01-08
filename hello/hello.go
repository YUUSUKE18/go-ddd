package main

import "fmt"

func Hello(name string, language string) string {
	if name == "" {
		name = "world"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case "spanish":
		prefix = "Hola, "
	case "french":
		prefix = "Bonjour, "
	default:
		prefix = "Hello, "
	}
	return
}

func main() {
	fmt.Println(Hello("Jhon", "spanish"))
}
