package main

import (
    "bytes"
    "io/ioutil"
    "fmt"
    "log"
    "net/http"
    "strings"
    "encoding/json"

)

type Article struct {
    Id string `json:"Id"`
    Title string `json:"Title"`
    SubTitle string `json:"SubTitle"`
    Content string `json:"Content"`
    TimeStamp string `json:"TimeStamp"`
}
var Articles []Article

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    // json.NewEncoder(w).Encode(Articles)
    for _, article := range Articles {
        json.NewEncoder(w).Encode(article)
    }
}

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func createNewArticle(w http.ResponseWriter, r *http.Request) {
  jsonData := map[string]string{
      "Id": "5",
     "Title": "Crater forms under Hyd Metro station amid rains, officials say structure safe",
     "SubTitle": "A crater formed under the Moosapet metro station in Hyderabad amid heavy rainfall in the city",
     "Content": "Hyderabad Metro Rail Limited MD, NVS Reddy said. The area around the crater has been barricaded to avoid accidents as the road is still open for traffic",
     "TimeStamp": "987643638",
  }
  jsonValue, _ := json.Marshal(jsonData)
  response, err := http.Post("https://httpbin.org/post", "application/json", bytes.NewBuffer(jsonValue))
  if err != nil {
      fmt.Printf("The HTTP request failed with error %s\n", err)
  } else {
      data, _ := ioutil.ReadAll(response.Body)
      var article Article
      json.Unmarshal(jsonValue, &article)
      Articles = append(Articles, article)
      json.NewEncoder(w).Encode(article)
      fmt.Println(string(data))
  }
  defer response.Body.Close()
}




func handleRequests() {
    http.HandleFunc("/", homePage)
    http.HandleFunc("/create", createNewArticle)
    http.HandleFunc("/articles", returnAllArticles)
    http.HandleFunc("/article/", returnSingleArticle)
    http.HandleFunc("/article/search", searchArticle)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func searchArticle(w http.ResponseWriter, r *http.Request){

    key:=r.URL.Query().Get("q")
    fmt.Println("GET params were:", r.URL.Query())
    for _, article := range Articles {
      if (strings.Contains(article.Title,key) || strings.Contains(article.SubTitle,key) || strings.Contains(article.Content,key) ){
        json.NewEncoder(w).Encode(article)
      }
    }
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request){

    key:=r.URL.Query().Get("id")
      fmt.Println("GET params were:", r.URL.Query())
    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}

func main() {
  Articles = []Article{
      Article{Id: "1",Title: "Russia approves second COVID-19 vaccine after preliminary trials", SubTitle: "Russia approves second COVID-19 vaccine after preliminary trials Russia has granted regulatory approval to a second COVID-19 vaccine, Russian President Vladimir Putin announced on Wednesday at a government meeting", Content: "The jab has been developed by the Vector Institute in Siberia and completed early-stage human trials last month. Results of the trials have not been published yet and a large-scale trial, known as Phase III, has not yet begu", TimeStamp:"123456789"},
      Article{Id: "2",Title: "Prithvi Shaw is like baby Sehwag: Swann", SubTitle: "Former England spinner Graeme Swann has said that Delhi Capitals opener Prithvi Shaw is like a baby Sehwag", Content: "He is like a miniature version of Virender Sehwag, who was one of my favourite Indian players of all time.Shaw has scored 202 runs in IPL 2020 so far",TimeStamp:"123456789"},
      Article{Id: "3",Title: "Captains should be allowed to review wide balls or waist-high full-tosses", SubTitle: "RCB captain Virat Kohli has said that captains should be allowed to have ability to review wide balls or waist-high full-tosses", Content: "He added We ve seen...in high-profile tournaments like IPL, these small things can be a big factor. He further said, If you lose a game by a run and aren't able to review that one decision...it makes a massive difference",TimeStamp:"123456789"},
      Article{Id: "4",Title: "Uttarakhand schools to reopen for class 10 and 12 from November 1", SubTitle: "Schools will reopen in Uttarakhand for class 10 and 12 on November 1 after a gap of nearly seven months", Content: "he decision was taken on district magistrates' feedback after they consulted with district education officers and parents, Cabinet Minister Madan Kaushik said. Kaushik added that decisions regarding the resumption of other classes will be taken in due course of time",TimeStamp:"123456789"},
  }
  handleRequests()
}
