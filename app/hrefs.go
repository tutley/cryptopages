// Code generated by goagen v1.3.1, DO NOT EDIT.
//
// API "cryptopages": Application Resource Href Factories
//
// Command:
// $ goagen
// --design=github.com/tutley/cryptopages/design
// --out=$(GOPATH)/src/github.com/tutley/cryptopages
// --regen=true
// --version=v1.3.1

package app

import (
	"fmt"
	"strings"
)

// UserHref returns the resource href.
func UserHref(username interface{}) string {
	paramusername := strings.TrimLeftFunc(fmt.Sprintf("%v", username), func(r rune) bool { return r == '/' })
	return fmt.Sprintf("/user/%v", paramusername)
}
