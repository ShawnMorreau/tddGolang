//Package hello is a basic hello world program
package hello

const englishHelloPrefix = "Hello, "
const frenchHelloPrefix = "Bonjour, "
const spanishHelloPrefix = "Hola, "
const spanish = "Spanish"
const french = "French"

//Greeting takes a name and language and outputs the corresponding Greeting
func Greeting(name string, language string) string {
	if name == "" {
		name = "World"
	}

	return greetingPrefix(language) + name
}

func greetingPrefix(language string) string {
	var prefix string

	switch language {
	case spanish:
		prefix = spanishHelloPrefix
	case french:
		prefix = frenchHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return prefix
}
