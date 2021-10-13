package router

import (
	"context"
	"encoding/json"
	"net/http"
	"regexp"
	"strings"
)

type route struct {
	method  string
	path    string
	handler http.HandlerFunc
}

func Serve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

routes:
	for _, route := range routes {
		if route.method != r.Method {
			continue
		}

		requestUrl := r.URL.Path
		requestUrl = strings.TrimRight(requestUrl, "/")
		requestUrlParts := strings.Split(requestUrl, "/")
		routeParts := strings.Split(route.path, "/")
		requestUrlParts = requestUrlParts[1:]
		routeParts = routeParts[1:]

		if routeParts[0] == "" && len(requestUrlParts) == 0 {
			route.handler(w, r)
			return
		}

		if len(routeParts) != len(requestUrlParts) {
			continue
		}

		parameters := make(map[string]string)
		for i := 0; i < len(routeParts); i++ {
			routePart := routeParts[i]
			matched, _ := regexp.MatchString(`^:`, routePart)
			if matched {
				routePart = strings.TrimLeft(routePart, ":")
				parameters[routePart] = requestUrlParts[i]
			} else if routePart != requestUrlParts[i] {
				continue routes
			}
		}

		ctx := context.WithValue(r.Context(), "parameters", parameters)
		route.handler(w, r.WithContext(ctx))
		return
	}

	w.WriteHeader(http.StatusNotFound)
	resp := make(map[string]string)
	resp["message"] = "Not Found"
	jsonResp, _ := json.Marshal(resp)
	w.Write(jsonResp)
	return
}

func get(path string, handler http.HandlerFunc) route {
	return route{"GET", path, handler}
}

func post(path string, handler http.HandlerFunc) route {
	return route{"POST", path, handler}
}

func put(path string, handler http.HandlerFunc) route {
	return route{"PUT", path, handler}
}

func delete(path string, handler http.HandlerFunc) route {
	return route{"DELETE", path, handler}
}
