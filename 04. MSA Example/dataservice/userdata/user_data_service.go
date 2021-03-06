package userdata

import (
	"MSA.example.com/1/model"
	"errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

type userDAO struct {
	db *gorm.DB
}

func GetUserDAO(db *gorm.DB) *userDAO {
	db.AutoMigrate(
		&model.Users{},
	)
	if !db.HasTable(&model.Users{}) {
		db.CreateTable(&model.Users{})
	}

	return &userDAO{db: db}
}

func (u *userDAO) Find(id uint32) (*model.Users, error) {
	user := model.Users{}
	if u.db.Where("id = ?", id).Find(&user).RowsAffected != 0 {
		return nil, errors.New("A user with that ID already exists")
	}
	return &user, nil
}

func (u *userDAO) FindByUserId(userId string) (*model.Users, error) {
	user := model.Users{}
	if u.db.Where("user_id = ?", userId).Find(&user).RowsAffected != 0 {
		return nil, errors.New("A user with that ID already exists")
	}
	return &user, nil
}

func (u *userDAO) Insert(user *model.Users) (*model.Users, error) {
	var r *model.Users
	err := u.db.Transaction(func(tx *gorm.DB) error {
		if tx = tx.Create(user); tx.Error != nil {
			return tx.Error
		}
		r = tx.Value.(*model.Users)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return r, nil
}


func (u *userDAO) Remove(id uint32) (int64, error) {
	db := u.db.Where("id = ?", id).Delete(model.Users{})
	if db.Error != nil {
		fmt.Println(db.Error)
		return 0, db.Error
	}
	return db.RowsAffected, nil
}