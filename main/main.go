package main


import (
  "fmt"
  "net/http"
  "log"
  "github.com/onyewuenyi/restful_api_V1/handlers"
)

const PORT string = ":8000"






func handleRequest() {
  fmt.Println("Starting Server on localhost" + PORT)
  http.HandleFunc("/", handlers.HomePage)
  http.HandleFunc("/all", handlers.ReturnAllArticles)
  log.Fatal(http.ListenAndServe(PORT, nil))
}



func main() {
  handleRequest()
}
