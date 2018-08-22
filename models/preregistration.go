// Package models contains the structures of data which is stored in the database and
// sent and received via HTTP requests.
package models

import "time"

// Preregistration represents the preregistration that is stored in the database.
//
// The difference between the model in the database and the initial POST is that
// we color it with the timestamp.
type Preregistration struct {
	Email     string    `json:"email"`
	Timestamp time.Time `json:"timestamp"`
}
