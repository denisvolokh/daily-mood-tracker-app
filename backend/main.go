package main

import (
	"log"
	"net/http"
)

func main() {
	InitDB("moods.db")

	http.HandleFunc("/moods", handleUserMoods)
	http.HandleFunc("/admin/moods", handleAdminMoods)

	log.Println("Server running on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
