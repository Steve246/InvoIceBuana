package repository

import (
	"fmt"
	"invoiceBuana/model"

	"gorm.io/gorm"
)

type ItemRepository interface {
	GetDuplicateByName(name string) (bool, error)

	Create(item *model.Item) error
}

type itemRepository struct {
	db *gorm.DB
}

func (i *itemRepository) GetDuplicateByName(name string) (bool, error) {
	var count int64
	query := "SELECT COUNT(*) FROM Item WHERE item_name = ?"
	result := i.db.Raw(query, name).Scan(&count)
	if result.Error != nil {
		return false, result.Error
	}

	if count > 0 == true {
		return true, nil
	} else {
		return false, nil
	}
}

func (i *itemRepository) Create(item *model.Item) error {
	query := `INSERT INTO Item (item_id, item_name, item_type, item_price)
	          VALUES (?, ?, ?, ?)`
	result := i.db.Exec(query, item.ItemId, item.ItemName, item.ItemType, item.ItemPrice)
	if result.Error != nil {
		return result.Error
	}
	item.ID = uint(result.RowsAffected)
	fmt.Println(item.ID)
	return nil
}

func NewItemRepository(db *gorm.DB) ItemRepository {
	repo := new(itemRepository)
	repo.db = db
	return repo
}
