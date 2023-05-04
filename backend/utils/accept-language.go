package utils

import (
	"strings"
)

// Adapted from http://siongui.github.io/2015/02/22/go-parse-accept-language/

func parseLanguageCode(code string) string {
	switch code {
	case "nb", "nn":
		return "no"
	default:
		return code
	}
}

// ParseAcceptLanguage HTTP header
func ParseAcceptLanguage(acceptLanguage string) []string {
	var lqs []string

	langQStrings := strings.Split(acceptLanguage, ",")
	for _, langQStr := range langQStrings {
		trimmedLangQStr := strings.Trim(langQStr, " ")
		langQ := strings.Split(trimmedLangQStr, ";")
		lq := langQ[0]
		lq2 := strings.Split(lq, "-")
		lqs = append(lqs, parseLanguageCode(lq2[0]))
	}

	return lqs
}
