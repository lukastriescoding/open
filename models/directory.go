package models

import "database/sql"

type Directory struct {
	Name    string
	Path    string
	MainApp sql.NullString
}
