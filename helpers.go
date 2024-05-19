package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

func RespondWithError(w http.ResponseWriter, code int, msg string) {
	log.Printf("%s\n", msg)
	type errResp struct {
		Error string `json:"error"`
	}
	RespondWithJson(w, code, errResp{
		Error: msg,
	})
}

func RespondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	data, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error: failed marshalling to json: %s", err)
		w.WriteHeader(500)
		return
	}
	w.Write(data)
}

func DecodeRequestParams(r *http.Request, param any) error {
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&param); err != nil {
		return err
	}
	return nil
}

func GetHeaderValue(header, prefix string) (value string, err error) {
	value, found := strings.CutPrefix(header, prefix)
	if !found {
		return "", errors.New("failed to retrieve header value")
	}
	return value, nil
}
