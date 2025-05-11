package models

import "time"

type CreateCollection struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Collection struct {
	Fingerprint string    `json:"fingerprint"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Published   bool      `json:"published"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
