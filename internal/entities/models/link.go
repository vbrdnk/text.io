package models

import "time"

type Link struct {
	Fingerprint string    `json:"fingerprint"`
	Label       string    `json:"label"`
	Url         string    `json:"url"`
	CreatedBy   string    `json:"created_by"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
}
