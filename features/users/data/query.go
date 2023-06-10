package data

import (
	"alta/immersive-dashboard-api/app/helper"
	"alta/immersive-dashboard-api/features/users"
	"errors"

	"gorm.io/gorm"
)

type UserData struct {
	db *gorm.DB
}

// Insert implements users.UserDataInterface
func (repo *UserData) Insert(data users.Core) (uint, error) {
	hashPassword, err := helper.HashPasword(data.Password)
	if err != nil {
		return 0, errors.New("error hashing password: " + err.Error())
	}
	data.Password = hashPassword
	userData := CoreToUsersModel(data)

	tx := repo.db.Create(&userData)
	if tx.Error != nil {
		return 0, tx.Error
	} else if tx.RowsAffected == 0 {
		return 0, errors.New("insert data user failed, rows affected 0 ")
	}
	return userData.ID, nil
}

// Update implements users.UserDataInterface
func (repo *UserData) Update(userId uint, data users.Core) error {
	var user Users
	tx := repo.db.Where("id = ?", userId).First(&user)
	if tx.Error != nil {
		return tx.Error
	}
	hashPassword, err := helper.HashPasword(data.Password)
	if err != nil {
		return errors.New("error hashing password: " + err.Error())
	}

	data.Password = hashPassword
	if tx := repo.db.Model(&user).Updates(CoreToUsersModel(data)); tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Select implements users.UserDataInterface
func (repo *UserData) Select(userId uint) (users.Core, error) {
	var user Users
	tx := repo.db.Where("id = ?", userId).First(&user)
	if tx.Error != nil {
		return users.Core{}, tx.Error
	}

	mapUser := UsersModelToCore(user)

	return mapUser, nil
}

// SelectAll implements users.UserDataInterface
func (repo *UserData) SelectAll() ([]users.Core, error) {
	var _users []Users
	tx := repo.db.Find(&_users).Where("status != ?", "deleted")
	if tx.Error != nil {
		return nil, tx.Error
	}

	var allUsers []users.Core
	for _, user := range _users {
		var data = UsersModelToCore(user)
		allUsers = append(allUsers, data)
	}

	return allUsers, nil
}

// Delete implements users.UserDataInterface
func (repo *UserData) Delete(userId uint) error {
	if errChangeStatusUser := repo.changeStatusUser(userId, "deleted"); errChangeStatusUser != nil {
		return errChangeStatusUser
	}
	tx := repo.db.Delete(&Users{}, userId)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Login implements users.UserDataInterface
func (repo *UserData) Login(email string, password string) (int, error) {
	var user Users
	tx := repo.db.Where("email = ?", email).First(&user)
	if tx.Error != nil {
		return 0, errors.New("email tidak terdaftar")
	}

	match := helper.CheckPaswordHash(password, user.Password)
	if !match {
		return 0, errors.New("kredensial tidak cocok")
	}

	errChangeStatusUser := repo.changeStatusUser(user.ID, "active")
	if errChangeStatusUser != nil {
		return 0, errChangeStatusUser
	}

	return int(user.ID), nil
}

// Logout implements users.UserDataInterface
func (repo *UserData) Logout(userId uint) error {
	if err := repo.changeStatusUser(userId, "not-active"); err != nil {
		return err
	}	
	return nil
}

func (repo *UserData) changeStatusUser(userId uint, status string) error {
	var user Users

	tx := repo.db.Model(&user).Where("id = ?", userId).Update("status", status)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

func New(db *gorm.DB) users.UserDataInterface {
	return &UserData{
		db: db,
	}
}
