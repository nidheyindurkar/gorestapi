
// main.go
package main

import (
    "encoding/json"
    "fmt"
    "log"
    "io/ioutil"
    "net/http"

    "github.com/gorilla/mux"
)

// User - Our struct for all Users
type User struct {
    // Id      string    `json:"Id"`
    // Title   string `json:"Title"`
    // Desc    string `json:"desc"`
    // Content string `json:"content"`
	id 		int 	`json:id`
	fname	string	`json:fname`
	city	string	`json:city`
	phone	string	`json:phone`
	height	float64	`json:height`
	married	bool	`json:married`
}

var User []Users

func homePage(w http.ResponseWriter, r *http.Request) {
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

func returnAllArticles(w http.ResponseWriter, r *http.Request) {
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Users)
}

func returnSingleArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    key := vars["id"]

    for _, article := range Articles {
        if article.Id == key {
            json.NewEncoder(w).Encode(article)
        }
    }
}


func createNewArticle(w http.ResponseWriter, r *http.Request) {
    // get the body of our POST request
    // unmarshal this into a new Article struct
    // append this to our Articles array.    
    reqBody, _ := ioutil.ReadAll(r.Body)
    var article Article 
    json.Unmarshal(reqBody, &article)
    // update our global Articles array to include
    // our new Article
    Articles = append(Articles, article)

    json.NewEncoder(w).Encode(article)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    id := vars["id"]

    for index, article := range Articles {
        if article.Id == id {
            Articles = append(Articles[:index], Articles[index+1:]...)
        }
    }

}

func handleRequests() {
    myRouter := mux.NewRouter().StrictSlash(true)
    myRouter.HandleFunc("/", homePage)
    myRouter.HandleFunc("/articles", returnAllArticles)
    myRouter.HandleFunc("/article", createNewArticle).Methods("POST")
    myRouter.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
    myRouter.HandleFunc("/article/{id}", returnSingleArticle)
    log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func main() {
    Users = []Users{
        // Article{Id: "1", Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        // Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
		Users{ id: 1, fname: "Nidhey", city: "Nagpur", phone: "9096810774", height: 7.0, married: false },
		Users{ id: 2, fname: "Riwtik", city: "Nagpur", phone: "9096810775", height: 6.7, married: false }
    }
    handleRequests()
}