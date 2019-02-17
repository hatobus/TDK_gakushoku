package models

import (
	"time"
)

type Student struct {
	No         int       `xorm:"not null pk autoincr INT(11)"`
	Name       string    `xorm:"not null VARCHAR(64)"`
	Sumofcoin  int       `xorm:"not null INT(11)"`
	Lastworked time.Time `xorm:"not null default 'current_timestamp()' DATETIME"`
}
