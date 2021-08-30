package model

import (
	"fmt"
	"time"
)

// Certificate - is an interface that describes information about certificates
type Certificate struct {
	ID        uint64    `db:"id"`
	UserID    uint64    `db:"user_id"`
	Created   time.Time `db:"created"`
	Link      string    `db:"link"`
	IsDeleted bool      `db:"is_deleted"`
}

// String - method allows you to display the certificate as a string
func (c Certificate) String() string {
	return fmt.Sprintf("ID: %d\nUserId: %d\nCreated: %s\nLink: %s\n Link: %t\n",
		c.ID, c.UserID, c.Created.Format("02-01-2006 15:04:05"), c.Link, c.IsDeleted)
}
