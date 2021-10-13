package response

import (
	"encoding/json"
	"net/http"
)

type data struct {
	Data interface{} `json:"data"`
}

func JsonData(w http.ResponseWriter, response interface{}) {
	data := data{
		Data: response,
	}
	json.NewEncoder(w).Encode(data)
}
