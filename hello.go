package main
import (
    "fmt"
    "log"
    "net/http"
    "os"
)
func handler(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Hello World V1 %s!", r.URL.Path[1:])
    //fmt.Fprintf(w, "Hello World1 %s!", r.URL.Path[1:])
    //fmt.Fprintf(w, "Hello World2 %s!", r.URL.Path[1:])
    fmt.Println("RESTfulServ. on:8093, Controller:",r.URL.Path[1:])
}
func main() {
    http.HandleFunc("/", handler)
    fmt.Println("Starting Restful services...")
    fmt.Println("Using port:8093")
    err := http.ListenAndServe(":8093", nil)
    log.Print(err)
    errorHandler(err)
}
func errorHandler(err error){
if err!=nil {
    fmt.Println(err)
    os.Exit(1)
}
}
