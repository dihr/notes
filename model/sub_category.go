package model

import "database/sql"

type SubCategory struct {
	ID         int            `json:"id"`
	CategoryID int            `json:"category_id"`
	Name       string         `json:"name"`
	Text       string         `json:"text"`
	FlagImg    bool           `json:"flag_img"`
	ImgName    sql.NullString `json:"img_name"`
}
