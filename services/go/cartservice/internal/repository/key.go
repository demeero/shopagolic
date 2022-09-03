package repository

import "fmt"

func key(prefix, val string) string {
	return fmt.Sprintf("%s:%s", prefix, val)
}
