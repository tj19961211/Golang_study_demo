package main

const (
	spanish            = "Spanish"
	french             = "French"
	frenchHelloPrefix  = "Bonjour, "
	englisHelloPrefix  = "Hello, "
	spanishHelloPrefix = "Hola, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func main() {
	//fmt.Println(Hello("name"))
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englisHelloPrefix
	}
	return
}
