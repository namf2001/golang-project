package cache

import (
	"net/http"
	"time"
)

// CacheObject is a struct that contains the response, response body, created time
type CacheObject struct {
	Response     *http.Response
	ResponseBody []byte
	Created      time.Time
}
