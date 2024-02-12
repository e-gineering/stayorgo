package main

import (
	"fmt"
	"github.com/rs/cors"
	"html"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

func greeting() string {
	person := os.Getenv("PERSON")
	if person == "" {
		person = "you"
	}
	return fmt.Sprintf("hello %s", person)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// use the greeting function to get the greeting message
		message := greeting()
		_, err := fmt.Fprintf(w, message)
		if err != nil {
			return
		}
	})

	mux.HandleFunc("/", serveFrontend)
	mux.HandleFunc("/list", handleList)
	mux.HandleFunc("/add", handleAdd)
	mux.HandleFunc("/delete", handleDelete)

	// Specify the port you want your server to listen to
	port := "8080"
	fmt.Printf("Server is starting on port %s\n", port)

	// CORS handler
	handler := cors.New(cors.Options{
		//TODO fix cors for any development but local
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
		Debug:            true,
	}).Handler(mux)

	// Start the server and log if there is an error
	if err := http.ListenAndServe(":"+port, handler); err != nil {
		log.Fatalf("Failed to start server: %+v", err)
	}
}

func handleDelete(writer http.ResponseWriter, request *http.Request) {

	client := newRedisClient()
	// For simplicity, we're reading the item to be deleted from the form data
	if err := request.ParseForm(); err != nil {
		http.Error(writer, "Failed to parse request", http.StatusBadRequest)
		return
	}

	item := request.PostFormValue("item")
	if item == "" {
		http.Error(writer, "Item is required", http.StatusBadRequest)
		return
	}

	// Assume 'itemsList' is the Redis list key
	key := "itemsList"

	// Remove the item from the list stored in Redis
	_, err := client.LRem(ctx, key, 0, item).Result()
	if err != nil {
		http.Error(writer, "Failed to delete item", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusAccepted)
	items, err := client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		http.Error(writer, "Failed to fetch items", http.StatusInternalServerError)
		log.Fatalf("Failed to start server: %+v", err)
	}
	htmlResponse := generateHTMLforItems(items)
	fmt.Fprintf(writer, htmlResponse)
}

func handleList(writer http.ResponseWriter, request *http.Request) {

	key := "itemsList"
	client := newRedisClient()

	items, err := client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		http.Error(writer, "Failed to fetch items", http.StatusInternalServerError)
		log.Fatalf("Failed to start server: %+v", err)
	}

	htmlResponse := generateHTMLforItems(items)
	fmt.Fprintf(writer, htmlResponse)
}

func generateHTMLforItems(items []string) string {
	var htmlBuilder strings.Builder
	htmlBuilder.WriteString("<ul id='items-list'>")
	for _, item := range items {
		safeItem := html.EscapeString(item)
		htmlBuilder.WriteString(fmt.Sprintf("<li>%s <button hx-post='/delete' hx-target='#items-list' hx-swap='outerHTML' hx-vals='{\"item\":\"%s\"}'>Delete</button></li>", safeItem, safeItem))
	}
	htmlBuilder.WriteString("</ul>")
	return htmlBuilder.String()

}

func handleAdd(writer http.ResponseWriter, request *http.Request) {

	key := "itemsList"
	writer.Header().Set("Origin", "localhost")

	// assuming application/x-www-form-urlencoded data
	if err := request.ParseForm(); err != nil {
		http.Error(writer, "Invalid Form data", 400)
		return
	}
	item := request.FormValue("item")

	// use Redis client to add the item
	client := newRedisClient()
	if err := client.RPush(ctx, key, item).Err(); err != nil {
		http.Error(writer, "Failed to add item", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	items, err := client.LRange(ctx, key, 0, -1).Result()
	if err != nil {
		http.Error(writer, "Failed to fetch items", http.StatusInternalServerError)
		log.Fatalf("Failed to start server: %+v", err)
	}
	htmlResponse := generateHTMLforItems(items)
	fmt.Fprintf(writer, htmlResponse)
}

type PageData struct {
	Name string
}

func serveFrontend(writer http.ResponseWriter, request *http.Request) {
	t, _ := template.ParseFiles("static/index.html")

	data := PageData{
		Name: os.Getenv("PERSON"),
	}
	err := t.Execute(writer, data)
	if err != nil {
		return
	}

}
