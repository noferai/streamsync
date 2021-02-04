package v1

import (
	"log"
	"server/core/models"

	"github.com/jinzhu/gorm"
)

func GetRoom(db *gorm.DB, id string) (*models.Room, error) {
	var err error
	post := new(models.Room)

	if err := db.Where("id = ? ", id).First(&post).Error; err != nil {
		log.Println(err)

		return nil, err
	}

	return post, err
}

func GetRooms(db *gorm.DB, args models.Args) ([]models.Room, int64, int64, error) {
	var rooms []models.Room
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
func SaveRoom(db *gorm.DB, room *models.Room) (*models.Room, error) {
	if err := db.Save(&room).Error; err != nil {
		return room, err
	}

	return room, nil
}

// Soft deletes room and all messages
func DeleteRoom(db *gorm.DB, id string) error {
	room := new(models.Room)
	if err := db.Where("id = ? ", id).Delete(&room).Error; err != nil {
		log.Println(err)
		return err
	}

	message := new(models.Message)
	if err := db.Where("room_id = ? ", id).Delete(&message).Error; err != nil {
		log.Println(err)
	}

	return nil
}
