package model

import (
	"fmt"
	"time"
)

// Certificate - is an interface that describes information about certificates
type Certificate struct {
	Id      uint64    `db:"id"`
	UserId  uint64    `db:"user_id"`
	Created time.Time `db:"created"`
	Link    string    `db:"link"`
}

// String - method allows you to display the certificate as a string
func (c Certificate) String() string {
	return fmt.Sprintf("Id: %d\nUserId: %d\nCreated: %s\nLink: %s\n",
		c.Id, c.UserId, c.Created.Format("02-01-2006 15:04:05"), c.Link)
}
