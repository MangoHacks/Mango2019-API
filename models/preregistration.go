package models

import "time"

type Preregistration struct {
	Email     string    `json:"email"`
	Timestamp time.Time `json:"timestamp"`
}
