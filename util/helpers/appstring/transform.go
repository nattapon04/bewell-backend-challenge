package appstring

import (
	"github.com/gobeam/stringy"
)

func ToSnakeCase(str string) string {
	return stringy.New(str).SnakeCase().ToLower()
}
