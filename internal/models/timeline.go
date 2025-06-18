package main

import "time"

// Base structs (database entities)
type TimelineEntry struct {
	ID        int
	CompanyID int
	Note      *string
	Reminder  *time.Time
	CreatedAt time.Time
}
