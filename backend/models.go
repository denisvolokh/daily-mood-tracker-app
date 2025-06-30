package main

type MoodEntry struct {
	ID     string `json:"id"`
	UserID string `json:"userId"`
	Mood   string `json:"mood"`
	Note   string `json:"note,omitempty"`
	Date   string `json:"date"`
}
