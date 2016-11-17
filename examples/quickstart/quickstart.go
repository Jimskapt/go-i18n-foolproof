package main

import (
	"fmt"

	i18n "github.com/Jimskapt/go-i18n-foolproof"
)

func main() {

	// DEFINING locales :

	en_EN := `{
		"This is an example of translated text.":
			"This is an example of translated text.",
		"This is another text ...":
			"This is another text ..."
	}`

	fr_FR := `{
		"This is an example of translated text.":
			"Ceci est un example de texte traduit."
	}`

	es_ES := `{
		"This is an example of translated text.":
			"Este es un ejemplo de texto traducido.",
		"This is another text ...":
			"Este es otro texto ..."
	}`

	// REGISTERING locales :

	parsed, _, _ := i18n.JSONParser{}.Parse([]byte(en_EN))
	i18n.RegisterLocale(parsed, "en_EN")

	parsed, _, _ = i18n.JSONParser{}.Parse([]byte(fr_FR))
	i18n.RegisterLocale(parsed, "fr_FR")

	parsed, _, _ = i18n.JSONParser{}.Parse([]byte(es_ES))
	i18n.RegisterLocale(parsed, "es_ES")

	// DISPLAY translations :

	// Normal display :
	fmt.Println(i18n.MustGet("This is an example of translated text.", "fr_FR"))
	// prints `Ceci est un example de texte traduit.`

	fmt.Println(i18n.MustGet("This is another text ...", "es_ES"))
	// prints `Este es otro texto ...`

	// Unknown text in (every) locale :
	fmt.Println(i18n.MustGet("Text not registered.", "fr_FR"))
	// prints `Text not registered.`

	// Redirected text :
	i18n.Redirects["Text not registered."] = "Redirected text."
	fmt.Println(i18n.MustGet("Text not registered.", "fr_FR"))
	// prints `Redirected text.`

	// Unknown locale :
	fmt.Println(i18n.MustGet("This is an example of translated text.", "zh_CN"))
	// prints `This is an example of translated text.`

	// Volontary forgot in the following locale :
	fmt.Println(i18n.MustGet("This is another text ...", "fr_FR"))
	// prints `This is another text ...`

}
