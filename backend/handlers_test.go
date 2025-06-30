package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/google/uuid"
	_ "modernc.org/sqlite"
)

func TestMain(m *testing.M) {
	db = InitDB("file::memory:?cache=shared")
	code := m.Run()
	os.Exit(code)
}

func TestPostMethod(t *testing.T) {
	entry := MoodEntry{
		UserID: "test-id",
		Mood:   "test-mood",
		Note:   "test-note",
	}

	body, _ := json.Marshal(entry)
	req := httptest.NewRequest(http.MethodPost, "/moods", bytes.NewReader(body))
	w := httptest.NewRecorder()

	handleUserMoods(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusCreated {
		t.Fatalf("Expected status 201 Created, got %d", resp.StatusCode)
	}

	var returned MoodEntry
	if err := json.NewDecoder(resp.Body).Decode(&returned); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if returned.Mood != entry.Mood {
		t.Fatalf("Expected Mood %s, got %s", entry.Mood, returned.Mood)
	}
}

func TestGetMoods(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/moods?userId=test-id", nil)
	w := httptest.NewRecorder()

	handleUserMoods(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	var returned []MoodEntry
	if err := json.NewDecoder(resp.Body).Decode(&returned); err != nil {
		t.Fatalf("Failed to decode response: %v", err)
	}

	if len(returned) != 1 {
		t.Fatalf("Expected 1 entry, got %d", len(returned))
	}
}

func seedMood(userID, mood, note string) {
	entry := MoodEntry{
		ID:     uuid.New().String(),
		UserID: userID,
		Mood:   mood,
		Note:   note,
		Date:   "2025-06-25T12:00:00Z",
	}
	_, err := db.Exec(`INSERT INTO mood_entries (id, user_id, mood, note, date) VALUES (?, ?, ?, ?, ?)`,
		entry.ID, entry.UserID, entry.Mood, entry.Note, entry.Date)
	if err != nil {
		panic(err)
	}
}

func TestAdminMoods(t *testing.T) {
	// Seed some test data
	seedMood("admin-user", "😀", "Sunny day")
	seedMood("admin-user", "😐", "")
	seedMood("other-user", "😢", "Bad sleep")

	row := db.QueryRow("SELECT COUNT(*) FROM mood_entries")
	var count int
	row.Scan(&count)
	t.Logf("Number of moods in DB: %d", count)

	t.Run("return all moods", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/admin/moods", nil)
		w := httptest.NewRecorder()

		handleAdminMoods(w, req)
		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status 200 OK, got %d", resp.StatusCode)
		}

		var moods []MoodEntry
		if err := json.NewDecoder(resp.Body).Decode(&moods); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}
		if len(moods) != 4 {
			t.Fatalf("Expected 4 moods, got %d", len(moods))
		}
	})

	t.Run("filters by userId", func(t *testing.T) {
		req := httptest.NewRequest(http.MethodGet, "/admin/moods?userId=admin-user", nil)
		w := httptest.NewRecorder()

		handleAdminMoods(w, req)

		resp := w.Result()
		if resp.StatusCode != http.StatusOK {
			t.Fatalf("Expected status 200 OK, got %d", resp.StatusCode)
		}
		var moods []MoodEntry
		if err := json.NewDecoder(resp.Body).Decode(&moods); err != nil {
			t.Fatalf("Failed to decode response: %v", err)
		}
		for _, m := range moods {
			if m.UserID != "admin-user" {
				t.Errorf("Unexpected userID %s", m.UserID)
			}
		}
	})
}
