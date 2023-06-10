package data

import (
	classData "alta/immersive-dashboard-api/features/classes/data"
	logData "alta/immersive-dashboard-api/features/mentees/logs/data"
	"alta/immersive-dashboard-api/features/users"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	FullName		string					`gorm:"type:varchar(100)"` 
	Email				string					`gorm:"type:varchar(50);unique;notNull" validate:"required,email"`
	Password		string					`gorm:"type:varchar(500);notNull"`
	Team 				string 					`gorm:"type:varchar(50)"`
	Role				string 					`gorm:"type:enum('user','admin');default:'user'"`
	Status			string 					`gorm:"type:enum('active','not-active','deleted');default:'active'"`
	Classes			[]classData.Classes	`gorm:"foreignKey:UserID"`
	Logs				[]logData.MenteeLogs `gorm:"foreignKey:UserID"`
}

func CoreToUsersModel(data users.Core) Users{
	return Users{
		FullName: data.FullName,
		Email:    data.Email,
		Password: data.Password,
		Team:     data.Team,
		Role:     data.Role,
		Status:   data.Status,	
	}
}

func UsersModelToCore(user Users) users.Core{
	return users.Core{
		Id:        user.ID,
		FullName:  user.FullName,
		Email:     user.Email,
		Team:      user.Team,
		Role:      user.Role,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

