// Package i18nfoolproof is an Internationalization package that can not fail to get text.
package i18nfoolproof

import "errors"

// Locales are a map of all registred locales in i18nfoolproof system.
var Locales = map[string]map[string]string{}
var Redirects = map[string]string{}

// Deprecated : you should use "Locales" variable instead.
// RegisterLocale registers a locale in i18n-foolproof system.
func RegisterLocale(data map[string]string, name string) error {

	if _, exists := Locales[name]; !exists {

		Locales[name] = data
		return nil

	}

	return errors.New("The locale " + name + " is already registered.")

}

// Get returns translated text.
// If the translated text was not found, it returns default text, which is also the key.
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

// Deprecated : you should use "Locales" variable instead.
// GetLocale returns all data for a locale.
func GetLocale(name string) (map[string]string, error) {

	if locale, exists := Locales[name]; exists {
		return locale, nil
	}

	return nil, errors.New("The locale " + name + " was not registered.")

}
