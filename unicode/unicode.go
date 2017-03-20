package unicode

import (
	"fmt"
	"unicode"
)

// Kode for Oppgave 4a
func Translate(expression string, language string) {
	expression := "\x22\x6E\x6F\x72\x64\x20\x6F\x67\x20\x73\xF8\x72\x22"
	//Japansk
	ja := unicode.translate(expression, "ja")

	fmt.Printf("%s", ja)

}
