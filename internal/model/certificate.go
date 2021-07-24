package model

import (
	"fmt"
	"time"
)

type Certificate struct {
	Id      uint64
	UserId  uint64
	Created time.Time
	Link    string
}

func (c Certificate) String() string {
	return fmt.Sprintf("Id: %d\nUserId: %d\nCreated: %s\nLink: %s\n",
		c.Id, c.UserId, c.Created.Format("02-01-2006 15:04:05"), c.Link)
}
