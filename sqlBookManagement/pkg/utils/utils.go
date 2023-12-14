package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func ParseBody(req *http.Request, x interface{}) {
	if body, err := io.ReadAll(req.Body); err == nil {
		fmt.Printf("%v\n", x)
		if err := json.Unmarshal([]byte(body), x); err != nil {

			return
		}
		fmt.Printf("%v", x)
	}
}
