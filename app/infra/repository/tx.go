package repository

import (
	"log"

	"gorm.io/gorm"
)

type IGormTx interface {
	Begin() (*gorm.DB, error)
	Commit(*gorm.DB)
	Rollback(*gorm.DB)
}

type psqlTransactor struct {
	connection *gorm.DB
}

func NewGormTx(db *gorm.DB) IGormTx {
	return &psqlTransactor{
		connection: db,
	}
}

func (u *psqlTransactor) Begin() (*gorm.DB, error) {
	tx := u.connection.Begin()

	log.Println("Tx gorm begin")
	return tx, nil
}

func (u *psqlTransactor) Commit(tx *gorm.DB) {
	log.Println("Tx gorm commit")
	tx.Commit()
}

func (u *psqlTransactor) Rollback(tx *gorm.DB) {
	log.Println("Tx gorm rollback")
	tx.Rollback()
}
