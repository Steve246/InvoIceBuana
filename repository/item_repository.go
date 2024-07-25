package repository

import (
	"fmt"
	"invoiceBuana/model"
	"invoiceBuana/model/dto"

	"gorm.io/gorm"
)

type ItemRepository interface {
	GetAll(limit, offset string) ([]dto.DisplayItem, error)
	GetDuplicateByName(name string) (bool, error)
	Create(item *model.Item) error
}

type itemRepository struct {
	db *gorm.DB
}

// fungsi untuk tetep display price dengan  2 decimal places
func formatPrice(price float64) string {
	return fmt.Sprintf("%.2f", price)
}

func (i *itemRepository) GetAll(limit, offset string) ([]dto.DisplayItem, error) {
	var items []dto.DisplayItem

	query := `SELECT item_id, item_name, item_type, item_price FROM Item LIMIT ? OFFSET ?`

	rows, err := i.db.Raw(query, limit, offset).Rows()
	if err != nil {
		return nil, fmt.Errorf("failed execute querry ==> %w", err)
	}

	defer rows.Close()

	for rows.Next() {
		var item dto.DisplayItem
		if err := rows.Scan(&item.ItemId, &item.ItemName, &item.ItemType, &item.ItemPrice); err != nil {
			return nil, fmt.Errorf("failed row scann ==> %w", err)
		}
		// Format the item price with 2 decimal places
		item.ItemPriceFormatted = formatPrice(item.ItemPrice)
		items = append(items, item)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("failed row iteration ==> %w", err)
	}

	return items, nil
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
