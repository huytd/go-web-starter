package main

import "time"

type Post struct {
	ID        int
	Message   string `sql:"size:255"`
	CreatedAt time.Time
}
