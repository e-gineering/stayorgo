package main

import (
	"fmt"
	"html"
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
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		// use the greeting function to get the greeting message
		message := greeting()
		_, err := fmt.Fprintf(w, message)
		if err != nil {
			return
		}
	})

	http.HandleFunc("/", serveFrontend)
	http.HandleFunc("/list", handleList)
	http.HandleFunc("/add", handleAdd)
	//http.HandleFunc("/delete", handleDelete)

	// Specify the port you want your server to listen to
	port := "8080"
	fmt.Printf("Server is starting on port %s\n", port)

	// Start the server and log if there is an error
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %+v", err)
	}
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
		htmlBuilder.WriteString(fmt.Sprintf("<li>%s <button hx-post='/delete' hx-target='#items-list' hx-swap='outerHTML' hx-vals='json({\"item\":\"%s\"})'>Delete</button></li>", safeItem, safeItem))
	}
	htmlBuilder.WriteString("</ul>")
	return htmlBuilder.String()

}

func handleAdd(writer http.ResponseWriter, request *http.Request) {
	// assuming application/x-www-form-urlencoded data
	if err := request.ParseForm(); err != nil {
		http.Error(writer, "Invalid Form data", 400)
		return
	}
	item := request.FormValue("item")

	// use Redis client to add the item
	client := newRedisClient()
	if err := client.RPush(ctx, "itemsList", item).Err(); err != nil {
		http.Error(writer, "Failed to add item", http.StatusInternalServerError)
		return
	}

	writer.WriteHeader(http.StatusCreated)
	items, err := client.LRange(ctx, "itemList", 0, -1).Result()
	if err != nil {
		http.Error(writer, "Failed to fetch items", http.StatusInternalServerError)
		log.Fatalf("Failed to start server: %+v", err)
	}
	htmlResponse := generateHTMLforItems(items)
	fmt.Fprintf(writer, htmlResponse)
}

func serveFrontend(writer http.ResponseWriter, request *http.Request) {
	http.ServeFile(writer, request, "static/index.html")
}
