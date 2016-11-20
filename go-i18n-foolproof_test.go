package i18nfoolproof

import (
	"reflect"
	"testing"
)

func TestGet(t *testing.T) {

	// reset registered locales :
	Locales = map[string]map[string]string{}

	// registrer a test locale
	fr_FR := map[string]string{}
	fr_FR["test_text"] = "lorem ipsum"
	Locales["fr_FR"] = fr_FR

	// ____________________________________________________________

	res, warn := Get("test_text", "fr_FR")
	if warn != nil {
		switch warn.(type) {
		case *WarningLocaleNotFound:
			t.Error("Unexpected warning : the locale should exists.")
		case *WarningTextNotFoundInLocale:
			t.Error("Unexpected warning : the text should exists in the locale.")
		default:
			t.Errorf("Unexpected warning : unknown warning : %+v", warn)
		}
	}

	if res != fr_FR["test_text"] {
		t.Errorf("The translated text should be the translated text, <%s>, got <%s>.", fr_FR["test_text"], res)
	}

	// ____________________________________________________________

	expectedText := "this is not registered"
	res, warn = Get(expectedText, "fr_FR")
	if warn == nil {
		t.Error("Getting translation should returns a warning because it doesn't exists in the requested locale.")
	}

	if _, goodType := warn.(WarningTextNotFoundInLocale); !goodType {
		t.Errorf("Wrong warning type : %+v.", reflect.TypeOf(warn))
	}

	if res != expectedText {
		t.Errorf("The translated text should be the requested text, <%s>, got <%s>.", expectedText, res)
	}

	// ____________________________________________________________

	res, warn = Get(expectedText, "unknown")
	if warn == nil {
		t.Error("Getting translation should returns a warning because the requested locale doesn't exists.")
	}

	if _, goodType := warn.(WarningLocaleNotFound); !goodType {
		t.Errorf("Wrong warning type : %+v.", reflect.TypeOf(warn))
	}

	if res != expectedText {
		t.Errorf("The translated text should be the requested text, <%s>, got <%s>.", expectedText, res)
	}

}

func TestMustGet(t *testing.T) {

	// reset registered locales :
	Locales = map[string]map[string]string{}

	// registrer a test locale
	fr_FR := map[string]string{}
	fr_FR["test_text"] = "lorem ipsum"
	Locales["fr_FR"] = fr_FR

	// ____________________________________________________________

	res := MustGet("test_text", "fr_FR")

	if res != fr_FR["test_text"] {
		t.Errorf("The translated text should be the translated text, <%s>, got <%s>.", fr_FR["test_text"], res)
	}

	// ____________________________________________________________

	expectedText := "this is not registered"
	res = MustGet(expectedText, "fr_FR")

	if res != expectedText {
		t.Errorf("The translated text should be the requested text, <%s>, got <%s>.", expectedText, res)
	}

	// ____________________________________________________________

	res = MustGet(expectedText, "unknown")

	if res != expectedText {
		t.Errorf("The translated text should be the requested text, <%s>, got <%s>.", expectedText, res)
	}

}
