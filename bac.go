package main

import (
	"encoding/json"
	"io/ioutil"
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

func Answer(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var counter int
	err := r.ParseForm()
	if err != nil {
		log.Fatal(err)
	}	
	var answer string
	for _, values := range r.Form {   
		for _, value := range values {
			answer = value
		}
	}
	fmt.Println("Answer ", answer)

	movieApiUrl := "https://api.themoviedb.org/3/search/movie?include_adult=false&api_key="
	apiKey :=  Credentials["movieApiKey"]

	sanitisedAnswer := strings.Replace(answer, " ", "%2B", -1)
	
	url := movieApiUrl + apiKey + "&language=en&" + "query=" + sanitisedAnswer

	fmt.Println("URL", url)

	payload := strings.NewReader("{}")

	req, _ := http.NewRequest("GET", url, payload)
	res, _ := http.DefaultClient.Do(req)

	body, _ := ioutil.ReadAll(res.Body)

	var response = new(MovieReponse)
	err = json.Unmarshal(body, &response)
	if err != nil {
		fmt.Println("Could not unmarshal the response from the API", err)
	}

	var movieList []string

	for _, movie := range response.Results {
		movieList = append(movieList, strings.ToLower(movie.Title))
	}

	fmt.Println("Movie List", movieList)

	if Contains(movieList, answer) {
		fmt.Println("Correct answer")
		counter ++
	} else {
		fmt.Println("Bad answer")
	}

	redirect(w, r, "/game")
	defer res.Body.Close()
}

func redirect(w http.ResponseWriter, r *http.Request, url string) {
	http.Redirect(w, r, url, 301)
}

func main() {
	router := httprouter.New()
	router.GET("/", Index)
	router.GET("/game", Game)
	router.POST("/game", Answer)
	
	http.Handle("/css/", http.StripPrefix("/css/", http.FileServer(http.Dir("css"))))

	log.Fatal(http.ListenAndServe(":8080", router))

}

type MovieReponse struct {
	Page         int `json:"page"`
	TotalResults int `json:"total_results"`
	TotalPages   int `json:"total_pages"`
	Results      []struct {
		VoteCount        int     `json:"vote_count"`
		ID               int     `json:"id"`
		Video            bool    `json:"video"`
		VoteAverage      float64 `json:"vote_average"`
		Title            string  `json:"title"`
		Popularity       float64 `json:"popularity"`
		PosterPath       string  `json:"poster_path"`
		OriginalLanguage string  `json:"original_language"`
		OriginalTitle    string  `json:"original_title"`
		GenreIds         []int   `json:"genre_ids"`
		BackdropPath     string  `json:"backdrop_path"`
		Adult            bool    `json:"adult"`
		Overview         string  `json:"overview"`
		ReleaseDate      string  `json:"release_date"`
	} `json:"results"`
}

func Contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

