// http://golang.site/go/article/107-MySql-%EC%82%AC%EC%9A%A9---%EC%BF%BC%EB%A6%AC

// https://blog.youngbin.xyz/2019-09-09-migrating-skhus-backend-from-nodejs-to-golang/
package main

import (
    "net/http"
)
 
func main() {
    http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
        w.Write([]byte("Hello World"))
    })
 
    http.ListenAndServe(":5000", nil)
}