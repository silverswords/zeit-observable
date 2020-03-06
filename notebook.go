package notebook

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

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

	body, _ := ioutil.ReadAll(resp.Body)
	w.Header().Set("content-type", "application/javascript; charset=utf-8")
	w.Write(body)
}
