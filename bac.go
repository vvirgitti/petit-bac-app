package main

import (
	"strings"
	"fmt"
	"math/rand"
	"time"
	"net/http"
	"html/template"
)

//PageData is data used on the html
type PageData struct {
	PageTitle string
	Letter string
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

func main() {
	tmpl, err := template.ParseFiles("index.html")
	if err != nil {
		fmt.Println("Error while parsing the html", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		data := PageData{
			PageTitle: "Petit Bac App",
			Letter: strings.ToUpper(generateLetter()),
		}
		tmpl.Execute(w, data)
	})
	http.ListenAndServe(":8080", nil)
}

