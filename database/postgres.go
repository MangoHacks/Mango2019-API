// Package database deals with the initialization
// of our database connection.
package database

import (
	"strings"
	"time"

	"github.com/MangoHacks/Mango2019-API/models"

	// To avoid using pq directly.
	_ "github.com/lib/pq"
)

// PostgreSQL Queries
var (
	PostgresInsertPreregistrationQuery = `
		INSERT 
		INTO preregistrations(email, timestamp) 
		VALUES($1, $2) 
		RETURNING email
	`
	PostgresSelectPreregistrationsQuery = `
		SELECT * 
		FROM preregistrations
		ORDER BY timestamp ASC
	`
	PostgresDeletePreregistrationQuery = `
		DELETE 
		FROM preregistrations
		WHERE email = $1
	`
)

func (db *DB) postgresInsertPreregistrations(email string) error {
	if _, err := db.postgres.Exec(PostgresInsertPreregistrationQuery, email, time.Now()); err != nil {
		if strings.Contains(err.Error(), "duplicate key value violates unique constraint") {
			return ErrDuplicate
		}
		return err
	}
	return nil
}

func (db *DB) postgresSelectPreregistrations() ([]models.Preregistration, error) {
	rws, err := db.postgres.Query(PostgresSelectPreregistrationsQuery)
	if err != nil {
		return nil, err
	}

	var prrs []models.Preregistration
	for rws.Next() {
		var eml string
		var t time.Time
		rws.Scan(&eml, &t)
		prrs = append(prrs, models.Preregistration{
			Email:     eml,
			Timestamp: t,
		})
	}
	return prrs, nil
}

func (db *DB) postgresDeletePreregistration(email string) error {
	_, err := db.postgres.Exec(PostgresDeletePreregistrationQuery, email)
	return err
}
