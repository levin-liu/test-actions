package test_latency

import (
	"encoding/json"
	"fmt"
	"html"
	"net/http"
)

// HelloLatency is an HTTP Cloud Function with a request parameter.
func HelloLatency(w http.ResponseWriter, r *http.Request) {
	var d struct {
		Name string `json:"name"`
	}
	if err := json.NewDecoder(r.Body).Decode(&d); err != nil {
		fmt.Fprint(w, "Hello, Latency!")
		return
	}
	if d.Name == "" {
		fmt.Fprint(w, "Hello, Latency!")
		return
	}
	fmt.Fprintf(w, "Hello, %s!\n", html.EscapeString(d.Name))
}
