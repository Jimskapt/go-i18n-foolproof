package i18nfoolproof

import "encoding/json"

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
