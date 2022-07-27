package main

import (
    "fmt"
    "net/http"
	"time"

    "goke/config"
    "html/template"
)

func hello(w http.ResponseWriter, req *http.Request) {

    fmt.Fprintf(w, "hello\n")
}

func headers(w http.ResponseWriter, req *http.Request) {

    for name, headers := range req.Header {
        for _, h := range headers {
            fmt.Fprintf(w, "%v: %v\n", name, h)
        }
    }
}



func main() {

    http.HandleFunc("/hello", hello)
    http.HandleFunc("/headers", headers)
    http.HandleFunc("/templates", templateFile)

    fs := http.FileServer(http.Dir("static/"))
    http.Handle("/static/", http.StripPrefix("/static/", fs))


    http.ListenAndServe(":8090", nil)
}


func hello2(w http.ResponseWriter, req *http.Request) {

    test := config.Config{"test","test123"}
    fmt.Println(test)

    ctx := req.Context()
    fmt.Println("server: hello handler started")
    defer fmt.Println("server: hello handler ended")

    select {
    case <-time.After(10 * time.Second):
        fmt.Fprintf(w, "hello\n")
    case <-ctx.Done():

        err := ctx.Err()
        fmt.Println("server:", err)
        internalError := http.StatusInternalServerError
        http.Error(w, err.Error(), internalError)
    }
}

func templateFile(w http.ResponseWriter, req *http.Request) {
    tmpl := template.Must(template.ParseFiles("templates/index.html.gotpl"))
    cfg := config.Config{"test_tile","0.0.1-dev"}
    tmpl.Execute(w, cfg)
}