package helper

import (
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func ErrMsgFormat(err string) string {
	t := cases.Title(language.English)
	sliceErr := strings.Split(err, ":")
	firstLeeter := strings.Split(sliceErr[0], " ")
	firstLeeter[0] = t.String(firstLeeter[0])
	sliceErr[0] = strings.Join(firstLeeter, " ")
	if len(sliceErr) > 2 {
		index := len(sliceErr) - 1
		return strings.Join(sliceErr[:index], ":")
	} else {
		return sliceErr[0]
	}
}
