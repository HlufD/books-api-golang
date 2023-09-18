package utils

import (
	"encoding/json"
	"io"
	"net/http"
)

func ParseJsonBody(req *http.Request, bodyStruct interface{}) {
	if body, err := io.ReadAll(req.Body); err == nil {
		if err := json.Unmarshal([]byte(body), bodyStruct); err != nil {
			return
		}
	}
}
