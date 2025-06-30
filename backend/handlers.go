package main

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/google/uuid"
)

func handleUserMoods(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		var entry MoodEntry
		if err := json.NewDecoder(r.Body).Decode(&entry); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if entry.UserID == "" || entry.Mood == "" {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		entry.ID = uuid.New().String()
		entry.Date = time.Now().Format(time.RFC3339)

		_, err := db.Exec(`INSERT INTO mood_entries (id, user_id, mood, note, date) VALUES (?, ?, ?, ?, ?)`,
			entry.ID, entry.UserID, entry.Mood, entry.Note, entry.Date)
		if err != nil {
			http.Error(w, "Insert failed", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(entry)
	case http.MethodGet:
		userID := r.URL.Query().Get("userId")
		if userID == "" {
			http.Error(w, "Missing userId", http.StatusBadRequest)
			return
		}
		rows, err := db.Query(`SELECT id, user_id, mood, note, date FROM mood_entries WHERE user_id = ? ORDER BY date DESC LIMIT 50`, userID)
		if err != nil {
			http.Error(w, "Query failed", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var results []MoodEntry
		for rows.Next() {
			var m MoodEntry
			rows.Scan(&m.ID, &m.UserID, &m.Mood, &m.Note, &m.Date)
			results = append(results, m)
		}
		json.NewEncoder(w).Encode(results)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func handleAdminMoods(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	var (
		rows *sql.Rows
		err  error
	)

	//if userId == "" {
	//	rows, err = db.Query(`SELECT * FROM mood_entries WHERE user_id = ?`, userId)
	//} else {
	//	rows, err = db.Query(`SELECT id, user_id, mood, note, date FROM mood_entries WHERE user_id = ? ORDER BY date DESC LIMIT 50`, userId)
	//}

	if userId == "" {
		// No filter – get all entries
		rows, err = db.Query(`SELECT id, user_id, mood, note, date FROM mood_entries ORDER BY date DESC LIMIT 50`)
	} else {
		// Filter by userId
		rows, err = db.Query(`SELECT id, user_id, mood, note, date FROM mood_entries WHERE user_id = ? ORDER BY date DESC LIMIT 50`, userId)
	}

	if err != nil {
		http.Error(w, "Query failed", http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var results []MoodEntry
	for rows.Next() {
		var m MoodEntry
		rows.Scan(&m.ID, &m.UserID, &m.Mood, &m.Note, &m.Date)
		results = append(results, m)
	}
	json.NewEncoder(w).Encode(results)
}
