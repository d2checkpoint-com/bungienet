package shared

import (
	"fmt"
)

func SetUserAgent(options []any) string {
	var userAgent string
	if len(options) > 3 {
		userAgent = options[1].(string)
	} else {
		userAgent = "github.com/D2Checkpoint-com/bungienet"
	}
	return userAgent
}

func CheckResponse(e int32, s, m string) error {
	if e != 1 {
		return fmt.Errorf("%s:%s", s, m)
	}
	return nil
}
