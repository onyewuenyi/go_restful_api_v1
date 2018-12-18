package handlers


import (
  "fmt"
  "net/http"
  "encoding/json"
  "github.com/onyewuenyi/restful_api_V1/model"
)

type Articles []model.Article


func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
  articles := Articles {
    model.Article{Title: "Hello", Desc: "Article Description", Content: "Article Content"},
    model.Article{Title: "Hello 2", Desc: "BlahBlah Description", Content: "Blah Blah Content"},
  }
  fmt.Println("Endpoint Hit: returnAllArticles")
  json.NewEncoder(w).Encode(articles)
}

func HomePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}
