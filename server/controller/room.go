package controller

import (
	"log"

	"github.com/jinzhu/gorm"
	"server/model"
)

func GetRoom(db *gorm.DB, id string) (*model.Room, error) {
	var err error
	post := new(model.Room)

	if err := db.Where("id = ? ", id).First(&post).Error; err != nil {
		log.Println(err)

		return nil, err
	}

	return post, err
}

func GetRooms(db *gorm.DB, args model.Args) ([]model.Room, int64, int64, error) {
	var rooms []model.Room
	var filteredData, totalData int64

	table := "rooms"
	query := db.Select(table + ".*")
	query = query.Offset(0)
	query = query.Limit(Limit(args.Limit))
	query = query.Order(SortOrder(table, args.Sort, args.Order))
	query = query.Scopes(Search(args.Search))

	if err := query.Find(&rooms).Error; err != nil {
		log.Println(err)
		return rooms, filteredData, totalData, err
	}

	query = query.Offset(0)
	query.Table(table).Count(&filteredData)
	db.Table(table).Count(&totalData)

	return rooms, filteredData, totalData, nil
}

// Both creates and updates post according to if ID field is empty or not
func SaveRoom(db *gorm.DB, room *model.Room) (*model.Room, error) {
	if err := db.Save(&room).Error; err != nil {
		return room, err
	}

	return room, nil
}

// Soft deletes room and all messages
func DeleteRoom(db *gorm.DB, id string) error {
	room := new(model.Room)
	if err := db.Where("id = ? ", id).Delete(&room).Error; err != nil {
		log.Println(err)
		return err
	}

	message := new(model.Message)
	if err := db.Where("room_id = ? ", id).Delete(&message).Error; err != nil {
		log.Println(err)
	}

	return nil
}
