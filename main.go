package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

var tpl = template.Must(template.ParseFiles("index.html"))

func main() {

	godoErr := godotenv.Load()

	if godoErr != nil {
		log.Println("Not able to load ENVs")
	}

	port := getEnv("PORT", "8080")

	fs := http.FileServer(http.Dir("assets"))

	router := http.NewServeMux()
	router.Handle("/assets/", http.StripPrefix("/assets/", fs))

	router.HandleFunc("/", indexHandler)
	router.HandleFunc("/articles", articlesHandler)
	router.HandleFunc("/search", searchHandler)

	err := http.ListenAndServe(":"+port, router)

	if err != nil {
		panic(err)
	}

}

func getEnv(key, fallback string) string {
	if port, ok := os.LookupEnv(key); ok {
		return port
	}
	return fallback
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tpl.Execute(w, "what is this")

	// w.Write([]byte(r.Method))
	// w.Write([]byte("<h1>Hello Somebody! I am in home page </h1>"))
}

func articlesHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>I am in article page</h1>"))
}

func searchHandler(w http.ResponseWriter, r *http.Request) {
	u, err := r.URL.Parse(r.URL.String())
	if err != nil {
		fmt.Println("URL parse error")
	}

	params := u.Query()
	searchQuery := params.Get("q")

	page := params.Get("page")
	if page == "" {
		x := "1"
		page = x
	}
	fmt.Println("Your search query is: ", searchQuery)
	fmt.Println("Your page is: ", page)

}
