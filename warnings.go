package i18nfoolproof

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
