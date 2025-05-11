package models

import "time"

type CreateLink struct {
	Label string `json:"label"`
	Url   string `json:"url"`
}

type Link struct {
	Fingerprint string    `json:"fingerprint"`
	Label       string    `json:"label"`
	Url         string    `json:"url"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
