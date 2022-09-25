package misc

import (
	"strings"
)

func GetFriendlyName(title string) string {
	return strings.ToLower(strings.ReplaceAll(title, " ", "-"))
}
