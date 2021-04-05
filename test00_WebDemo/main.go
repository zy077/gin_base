package main

import (
    "fmt"
    "net/http"
)

// http.ResponseWriter：代表响应，传递到前端的
// *http.Request：表示请求，从前端传递过来的
func SayHello(w http.ResponseWriter, r *http.Request){
    _, _ = fmt.Fprintln(w, "Hello Golang");
}

func main(){
    http.HandleFunc("/hello", SayHello)
    err := http.ListenAndServe(":9080", nil)
    if err != nil {
        fmt.Printf("Server start err: %s\n", err)
        return
    }
}
