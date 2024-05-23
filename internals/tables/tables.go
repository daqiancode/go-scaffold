package tables

import (
	"time"

	"github.com/daqiancode/scache"
	"gopkg.in/guregu/null.v4"
)

type Base struct {
	Id        uint64    `gorm:"type:bigint not null;autoIncrement;primarykey" json:"id"`
	CreatedAt time.Time `gorm:"type:datetime not null;" json:"createdAt"`
	UpdatedAt time.Time `gorm:"type:datetime not null;" json:"updatedAt"`
}

func (s Base) GetID() uint64 {
	return s.Id
}
func (s Base) ListIndexes() scache.Indexes {
	return scache.Indexes{}
}

type User struct {
	Base
	Email null.String `gorm:"type:varchar(255);uniqueIndex" json:"email"`
}

func (s User) ListIndexes() scache.Indexes {
	return scache.Indexes{}.Add(scache.NewIndex("email", s.Email.String))
}
