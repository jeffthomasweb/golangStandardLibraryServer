package main

import "log"
import "net/http"

func exampleHandler(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Content-Type", "text/html; charset=utf-8")
        w.Write([]byte("<p>Hi, this is a web server that was created using the Golang Standard Library</p>"))
}

func main() {
        mux := http.NewServeMux()
        exampleHandlerFunc := http.HandlerFunc(exampleHandler)
        mux.Handle("/", exampleHandlerFunc)
        log.Print("Start listening")
        http.ListenAndServe(":8005", mux)
}
