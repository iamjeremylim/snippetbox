package models

import {
	"error"
	"time"
}

var ErrNoRecord = errors.New("models: no matching record found")

type Snipper struct {
	ID int
	Title string
	Content string
	Create time.Time
	Expires time.Time
}
