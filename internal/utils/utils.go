package utils

import (
	"moba-converter-go/internal/config"
)

// #region Value Handling

// applyTemplate applies template data to the session data.
func ApplyTemplate(sessionData, templateData map[string]string) map[string]string {
	for key, value := range templateData {
		sessionData[key] = value
	}
	return sessionData
}

// setDefaultValues sets default values for missing fields in session data.
func SetDefaultValues(sessionData map[string]string, optionsMap config.OptionsMap) map[string]string {
	for key, valueSpec := range optionsMap {
		if _, exists := sessionData[key]; !exists {
			sessionData[key] = valueSpec.Default
		}
	}
	// Custom default for folder
	if _, exists := sessionData["folder"]; !exists {
		sessionData["folder"] = "/"
	}
	return sessionData
}

// applyValueReplacements replaces option values in session data.
func ApplyValueReplacements(sessionData map[string]string, optionsMap config.OptionsMap) map[string]string {
	for key, valueSpec := range optionsMap {
		if valueSpec.Options != nil {

			// TODO: IF something can't be looked up, the original is used. This should be clearly shown in the debug

			if val, exists := sessionData[key]; exists {
				if replacement, found := valueSpec.Options[val]; found {
					sessionData[key] = replacement
				} else {
					print("VALUE1 not found " + key + "\n")
				}
			} else {
				print("VALUE2 not found " + key)
			}
		}
	}
	return sessionData
}

// ReverseValueReplacements replaces -1 => false
func ReverseValueReplacements(sessionData map[string]string, optionsMap config.OptionsMap) map[string]string {

	for key, value := range sessionData {

		if option, ok := optionsMap[key]; ok {
			for k, v := range option.Options {
				if v == value {
					sessionData[key] = k
					break
				}
			}
		} else {
			if key != "SessionName" {
				print("There was no option for " + key + "in the config file")
			}
		}
	}

	return sessionData
}

func ReduceOptions(sessionData map[string]string, optionsMap config.OptionsMap) map[string]string {
	// Remove all default values from sessionData

	for key, value := range sessionData {

		if option, ok := optionsMap[key]; ok {
			if value == option.Default {
				delete(sessionData, key)
			}
		} else {
			if key != "SessionName" && key != "folder" {
				print("There was no option for " + key + "in the config file")
			}
		}
	}

	return sessionData

}

// #region Other

func GroupByFolder(slice []map[string]string) map[string][]map[string]string {
	groupKey := "folder"
	grouped := make(map[string][]map[string]string)
	for _, item := range slice {
		if value, exists := item[groupKey]; exists {
			grouped[value] = append(grouped[value], item)
		}
	}
	return grouped
}

func Check(err error) {
	if err != nil {
		panic(err)
	}
}
