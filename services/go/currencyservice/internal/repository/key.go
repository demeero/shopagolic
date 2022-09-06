package repository

import "fmt"

const internalPrefix = ":currency_conversion:"

func key(prefix, val string) string {
	return fmt.Sprintf("%s%s%s", prefix, internalPrefix, val)
}
