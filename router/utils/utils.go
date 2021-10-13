package utils

import "net/http"

func GetParameter(r *http.Request, key string) string {
	fields := r.Context().Value("parameters").(map[string]string)
	return fields[key]
}
