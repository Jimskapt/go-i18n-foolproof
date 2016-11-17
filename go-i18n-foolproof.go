// Package i18nfoolproof is an Internationalization package that can not fail to get text.
package i18nfoolproof

import (
	"encoding/json"
	"errors"
)

// Locales are a map of all registred locales in i18nfoolproof system.
var Locales = map[string]map[string]string{}
var Redirects = map[string]string{}

// RegisterLocale registers a locale in i18n-foolproof system.
func RegisterLocale(data map[string]string, name string) error {

	if _, exists := Locales[name]; !exists {

		Locales[name] = data
		return nil

	}

	return errors.New("The locale " + name + " is already registered.")

}

// Get returns translated text.
// If the translated text was not found, it returns default text, wich is also the key.
// It returns also warning message if there is.
func Get(text, locale string) (string, Warning) {

	if _, exists := Locales[locale]; !exists {
		return text, WarningLocaleNotFound{locale: locale}
	}

	result := Locales[locale][text]

	if result == "" {
		if Redirects[text] != "" {
			return Get(Redirects[text], locale)
		}

		return text, WarningTextNotFoundInLocale{text: text, locale: locale}
	}

	return result, nil

}

// Get returns translated text, and skip any errors/warnings.
func MustGet(text, locale string) string {

	result, _ := Get(text, locale)

	return result

}

// GetLocale returns all data for a locale.
func GetLocale(name string) (map[string]string, error) {

	if locale, exists := Locales[name]; exists {
		return locale, nil
	}

	return nil, errors.New("The locale " + name + " was not registered.")

}

// LocaleParser is the parser which is an helper for convert your data onto
// i18n-foolproof format in order to registered it.
type LocaleParser interface {
	Parse(data []byte) (map[string]string, []Warning, error)
}

// JSONParser is a LocaleParser for JSON-structured locale.
type JSONParser struct{}

func (p JSONParser) Parse(data []byte) (map[string]string, []Warning, error) {

	result := map[string]string{}
	warnings := []Warning{}

	var structData map[string]interface{}
	err := json.Unmarshal(data, &structData)
	if err != nil {
		return result, nil, err
	}

	for dataKey, value := range structData {
		stringValue, successCast := value.(string)
		if successCast {
			result[dataKey] = stringValue
		} else {
			warnings = append(warnings, WarningUncastedValue{keyname: dataKey})
		}
	}

	if len(warnings) <= 0 {
		warnings = nil
	}

	return result, warnings, nil

}

// Warning is like go's errors but they are not critical errors
type Warning interface {
	// Warning returns the warning message as string.
	Warning() string
}

// OtherWarning is inidentified warning.
type OtherWarning struct {
	message string
}

func (w OtherWarning) Warning() string {
	return w.message
}

// WarningUncastedValue says that parser can not parse a value of a tet as string.
type WarningUncastedValue struct {
	keyname string
}

func (w WarningUncastedValue) Warning() string {
	return "The key " + w.keyname + " can not be casted to string value."
}

// WarningLocaleNotFound says that the locale is not registered.
type WarningLocaleNotFound struct {
	locale string
}

func (w WarningLocaleNotFound) Warning() string {
	return "The locale " + w.locale + " is not registred."
}

// WarningTextNotFoundInLocale says that the locale is not registered.
type WarningTextNotFoundInLocale struct {
	text   string
	locale string
}

func (w WarningTextNotFoundInLocale) Warning() string {
	return "The text <" + w.text + "> in the locale " + w.locale + " was not found."
}
