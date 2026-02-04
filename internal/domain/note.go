package models

import "time"

type Note struct {
	ID       int64     `binding:required`
	Title    string    `binding:required`
	Text     string    `binding:required`
	UserID   int64     `binding:required`
	DateTime time.Time `binding:required`

}