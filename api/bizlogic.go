package api

import (
	"API-GOLANG/api/dataservice"
	"API-GOLANG/api/model"
	"database/sql"
)

type IBizLogic interface {
	CreateBookLogic(book model.Book) error
}

type BizLogic struct {
	DB *sql.DB
}

func NewBizLogic(db *sql.DB) *BizLogic {
	return &BizLogic{DB: db}
}

func (bl *BizLogic) CreateBookLogic(book model.Book) error {
	if err := dataservice.CreateBook(bl.DB, book); err != nil {
		return err
	}
	return nil
}
