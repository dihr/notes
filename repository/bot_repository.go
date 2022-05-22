package repository

import (
	"database/sql"
	"fmt"
	"notes/model"
)

type (
	BotRepository interface {
		GetAllSubCategories() ([]model.SubCategory, error)
		GetAllCategories() ([]model.Category, error)
		GetSubCategory(parameters string, args ...interface{}) (model.SubCategory, error)
		GetSubCategories(parameters string, args ...interface{}) ([]model.SubCategory, error)
	}

	botRepositoryImp struct {
		conn *sql.DB
	}
)

func NewBotRepository(conn *sql.DB) BotRepository {
	return &botRepositoryImp{
		conn: conn,
	}
}

func (b *botRepositoryImp) GetAllSubCategories() ([]model.SubCategory, error) {
	rows, err := b.conn.Query("select * from sub_category")
	if err != nil {
		return nil, err
	}
	subCategories := make([]model.SubCategory, 0)
	for rows.Next() {
		tempSubCategory := model.SubCategory{}
		if err := rows.Scan(&tempSubCategory.ID, &tempSubCategory.CategoryID,
			&tempSubCategory.Name, &tempSubCategory.Text); err != nil {
			return nil, err
		}
		subCategories = append(subCategories, tempSubCategory)
	}
	return subCategories, nil
}

func (b *botRepositoryImp) GetSubCategories(parameters string, args ...interface{}) ([]model.SubCategory, error) {
	rows, err := b.conn.Query(fmt.Sprintf("select * from sub_category where %s",
		parameters), args...)
	if err != nil {
		return nil, err
	}
	subCategories := make([]model.SubCategory, 0)
	for rows.Next() {
		tempSubCategory := model.SubCategory{}
		if err := rows.Scan(&tempSubCategory.ID, &tempSubCategory.CategoryID,
			&tempSubCategory.Name, &tempSubCategory.Text); err != nil {
			return nil, err
		}
		subCategories = append(subCategories, tempSubCategory)
	}
	return subCategories, nil
}

func (b *botRepositoryImp) GetSubCategory(parameters string, args ...interface{}) (model.SubCategory, error) {
	row := b.conn.QueryRow(fmt.Sprintf("select * from sub_category where %s",
		parameters), args...)
	if row.Err() != nil {
		return model.SubCategory{}, row.Err()
	}
	subCategory := model.SubCategory{}
	if err := row.Scan(&subCategory.ID, &subCategory.CategoryID,
		&subCategory.Name, &subCategory.Text); err != nil {
		return model.SubCategory{}, err
	}
	return subCategory, nil
}

func (b *botRepositoryImp) GetAllCategories() ([]model.Category, error) {
	rows, err := b.conn.Query("select * from category")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	categories := make([]model.Category, 0)
	for rows.Next() {
		tempCategory := model.Category{}
		if err := rows.Scan(&tempCategory.ID, &tempCategory.Name); err != nil {
			return nil, err
		}
		categories = append(categories, tempCategory)
	}
	return categories, nil
}
