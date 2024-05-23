package tables

import (
	"time"

	"gopkg.in/guregu/null.v4"
)

type Base struct {
	Id        uint64    `gorm:"type:bigint not null;autoIncrement;primarykey" json:"id"`
	CreatedAt time.Time `gorm:"type:datetime not null;" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:datetime not null;" json:"updatedAt"`
}

type User struct {
	Base
	Email null.String `gorm:"type:varchar(255);uniqueIndex" json:"email"`
}
