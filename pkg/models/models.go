package models

import (
	"encoding/json"
	"errors"
	"time"
)

var (
	ErrNoRecord           = errors.New("models: no matching record found")
	ErrInvalidCredentials = errors.New("models: invalid credentials")
	ErrDuplicateEmail     = errors.New("models: duplicate email")
)

type Bookmarks struct {
	ID         int
	Name       string
	RawMessage json.RawMessage
	Created    time.Time
}
