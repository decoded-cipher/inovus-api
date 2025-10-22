package models

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/knadh/goyesql/v2"
	goyesqlx "github.com/knadh/goyesql/v2/sqlx"
)

// Queries contains all prepared SQL queries.
type Queries struct {
	GetAllUsers   *sqlx.Stmt `query:"get-all-users"`
	GetUsersCount *sqlx.Stmt `query:"get-users-count"`
	GetUser       *sqlx.Stmt `query:"get-user"`
	CreateUser    *sqlx.Stmt `query:"create-user"`
	UpdateUser    *sqlx.Stmt `query:"update-user"`
	DeleteUser    *sqlx.Stmt `query:"delete-user"`
}

// LoadQueries loads and prepares SQL queries from a file using goyesqlx.
// This automatically parses queries.sql and prepares all statements.
func LoadQueries(db *sqlx.DB, filePath string) (*Queries, error) {
	queries := &Queries{}

	// Parse queries from the SQL file
	parsedQueries := goyesql.MustParseFile(filePath)

	// Prepare and bind queries to the struct
	if err := goyesqlx.ScanToStruct(queries, parsedQueries, db); err != nil {
		return nil, fmt.Errorf("failed to load queries: %w", err)
	}

	return queries, nil
}
