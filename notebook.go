package notebook

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

var headerKeys = []string{"content-type", "Cache-Control", "Access-Control-Allow-Origin"}

// Notebook -
func Notebook(w http.ResponseWriter, r *http.Request) {
	url := r.URL.RequestURI()
	pramas := r.URL.Query()
	var finalURL string
	if pramas["v"] == nil {
		finalURL = fmt.Sprintf("https://api.observablehq.com%v?v=%v", url, "3")
	} else {
		finalURL = fmt.Sprintf("https://api.observablehq.com%v", url)
	}

	resp, err := http.Get(finalURL)
	if err != nil {
		return
	}
	defer resp.Body.Close()

	header := resp.Header
	body, _ := ioutil.ReadAll(resp.Body)

	for _, key := range headerKeys {
		w.Header().Set(key, header.Get(key))
	}

	w.Write(body)
}
