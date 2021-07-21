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

func NewCertificate(id uint64, userId uint64, created time.Time, link string) *Certificate {
	return &Certificate{
		Id:      id,
		UserId:  userId,
		Created: created,
		Link:    link,
	}
}

func (c Certificate) String() string {
	return fmt.Sprintf("Id: %d\nUserId: %d\nCreated: %s\nLink: %s\n",
		c.Id, c.UserId, c.Created.Format("02-01-2006 15:04:05"), c.Link)
}
