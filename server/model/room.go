package model

import (
	"github.com/jinzhu/gorm"
)

type Response struct {
	TotalData    int64
	FilteredData int64
	Data         []Room
}

type Args struct {
	Sort   string
	Order  string
	Limit  string
	Search string
}

type Room struct {
	gorm.Model  `swaggerignore:"true"`
	Name        string    `json:"Name" gorm:"type:varchar(255)" example:"SampleName"`
	Description string    `json:"Description" gorm:"type:text" example:"Sample description"`
	IsPrivate   bool      `json:"IsPrivate" gorm:"type:bool" example:"false"`
	Messages    []Message `json:",omitempty" swaggerignore:"true"`
}

type Message struct {
	gorm.Model `swaggerignore:"true"`
	Sender     string `json:"Sender" gorm:"type:varchar(255)"`
	Body       string `json:"Body" gorm:"type:text"`
	RoomId     uint   `json:"RoomId"`
}
