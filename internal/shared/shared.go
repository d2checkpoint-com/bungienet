package shared

import (
	"encoding/json"
	"fmt"
	"github.com/d2checkpoint-com/bungienet/internal/rest"
	"github.com/d2checkpoint-com/bungienet/pkg/Models"
	"strings"
	"time"
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

func SetCacheBreakOpt(options []any) bool {
	var cacheBreaker bool
	if len(options) > 3 {
		cacheBreaker = options[2].(bool)
	} else {
		cacheBreaker = false
	}
	return cacheBreaker
}

func SetCacheBreakUrl(url string, cacheBreaker bool) string {
	if cacheBreaker {
		if strings.Contains(url, "?") {
			// URL already has query strings, append using '&'
			url += fmt.Sprintf("&t=%d", time.Now().Unix())
		} else {
			// URL has no query strings, add using '?'
			url += fmt.Sprintf("?t=%d", time.Now().Unix())
		}
	}
	return url
}

func SendRequest(request rest.Request) (*Models.Response, error) {
	// Send request
	response, err := rest.Send(request)
	if err != nil {
		return nil, err
	}

	return UnmarshalResponse(response)
}

func UnmarshalResponse(response *rest.Response) (*Models.Response, error) {
	var ret *Models.Response
	err := json.Unmarshal([]byte(response.Body), &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
