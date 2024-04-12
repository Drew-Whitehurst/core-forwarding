package main

import (
    "net/http"
    "net/http/httputil"
    "net/url"
)

func handleRequestAndRedirect(res http.ResponseWriter, req *http.Request) {
    url, _ := url.Parse("http://20.20.20.10") // Target host
    proxy := httputil.NewSingleHostReverseProxy(url)
    proxy.ServeHTTP(res, req)
}

func main() {
    http.HandleFunc("/", handleRequestAndRedirect)
    if err := http.ListenAndServe(":9050", nil); err != nil {
        panic(err)
    }
}

