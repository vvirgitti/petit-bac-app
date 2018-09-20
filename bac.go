package main

import (
	"log"
	"fmt"
	"math/rand"
	"time"
	"net/http"
	"github.com/julienschmidt/httprouter"
	"strings"
	"html/template"
)

//PageData is data used on the html
type PageData struct {
	PageTitle string
	Letter string
}

type CategoryData struct {
	Categories []string
}

func generateLetter() string {
	var runes = []rune("abcdefghijklmnopqrstuvwxyz")
	randRune := make([]rune, 1)
	
	for i := range randRune {
		rand.Seed(time.Now().UnixNano())
		randRune[i] = runes[rand.Intn(len(runes))]
	}
	return string(randRune)
}

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	cat, err := template.ParseFiles("category.html", "index.html")
	if err != nil {
		fmt.Println("Error while parsing the html", err)
	}
	category := CategoryData {
		Categories: []string{"TV show", "Movie", "Actor", "Actress"},
	}
	cat.ExecuteTemplate(w, "category", category)
}


func Game(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	tmpl, err := template.ParseFiles("game.html", "index.html")
	if err != nil {
		fmt.Println("Error while parsing the html", err)
	}
	data := PageData{
		PageTitle: "Petit Bac App",
		Letter: strings.ToUpper(generateLetter()),
	}
	tmpl.ExecuteTemplate(w, "game", data)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/game", Game)

	log.Fatal(http.ListenAndServe(":8080", router))

}
