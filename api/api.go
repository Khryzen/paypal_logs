package api

import (
	"net/http"
	"strings"
)

func APIhandler(w http.ResponseWriter, r *http.Request) {
	r.URL.Path = strings.TrimPrefix(r.URL.Path, "/api/")
	api := strings.TrimSuffix(r.URL.Path, "/")

	if strings.HasPrefix(api, "request") {
		PaypalAPIHandler(w, r)
		return
	}
}
