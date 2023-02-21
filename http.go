package main

import (
	"io"
	"net/http"
	"strings"
	"time"
)

var address string = ":80"

func httpHandleFunc(urlPath string, filepath string, contentType string, output string) {
	urlPath = strings.ToLower(urlPath)
	logger(address + "/" + urlPath + " <--> " + filepath + " <" + contentType + ">")
	http.HandleFunc("/"+urlPath, func(w http.ResponseWriter, r *http.Request) {
		logger(r.Method + " request -> " + address + "/" + urlPath + " <" + contentType + ">")

		w.Header().Add("Content-Type", contentType)

		io.WriteString(w, output)
	})
}

func httpGetBody(url string) []byte {
	c := &http.Client{Timeout: 10 * time.Second}
	r, err := c.Get(url)
	errorPanic(err)
	defer r.Body.Close()

	var body []byte
	body, err = io.ReadAll(r.Body)
	errorPanic(err)

	return body
}
