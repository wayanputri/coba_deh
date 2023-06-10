package data

import (
	"alta/immersive-dashboard-api/features/mentees/logs"

	"gorm.io/gorm"
)

type MenteeLogs struct {
    gorm.Model
    Status    string `gorm:"type:varchar(50)"`
    Feedback  string `gorm:"type:text"`
    MenteeID  uint
    UserID    uint
    User      Users  `gorm:"foreignKey:UserID"`
    Mentee    Mentees `gorm:"foreignKey:MenteeID"`
	Class Classes `gorm:"foreignKey:UserID"`
}

type Users struct {
	gorm.Model
	FullName			string					`gorm:"type:varchar(100)"` 
	Email				string					`gorm:"type:varchar(50);unique;notNull" validate:"required,email"`
	Password			string					`gorm:"type:varchar(500);notNull"`
	Team 				string 					`gorm:"type:varchar(50)"`
	Role				string 					`gorm:"type:enum('user','admin');default:'user'"`
	Status				string 					`gorm:"type:enum('active','not-active','deleted');default:'active'"`
}

type Classes struct {
	gorm.Model
	Name			string 		`gorm:"type:varchar(50);notNull;unique"`
	Tag				string		`gorm:"type:varchar(10);notNull;unique"`
	UserID			uint
}

type Mentees struct {
	gorm.Model
	FullName 				string 	`gorm:"type:varchar(100)"`
	NickName				string	`gorm:"type:varchar(10)"`
	ClassID					uint
	Status					string	`gorm:"type:varchar(50)"`
	Category				string 	`gorm:"type:varchar(50)"`
	Gender					string	`gorm:"type:varchar(50)"`
	Graduate				string	`gorm:"type:varchar(50)"`
	Mayor					string	`gorm:"type:varchar(50)"`
	Phone					string 	`gorm:"type:varchar(50)"`
	Telegram				string	`gorm:"type:varchar(50)"`
	Discord					string	`gorm:"type:varchar(50)"`
	Institusi 				string	`gorm:"type:varchar(50)"`
	Email					string	`gorm:"type:varchar(50)"`
	EmergencyName			string	`gorm:"type:varchar(50)"`
	EmergencyPhone			string	`gorm:"type:varchar(50)"`
	EmergencyStatus			string	`gorm:"type:varchar(50)"`
}

func CoreToModelLogs(input logs.Core) MenteeLogs{
	return MenteeLogs{
		Status: input.Status,
		Feedback: input.Feedback,
		MenteeID: input.MenteeID,
		UserID: input.UserID,
	}
}

func LogsModelToCore(input MenteeLogs) logs.Core{
	return logs.Core{
		Id: input.ID,
		Status:  input.Status,
		Feedback: input.Feedback,
		MenteeID: input.MenteeID,
		UserID:   input.UserID,
	}
}

func ModelToCoreGetAll(input MenteeLogs)logs.Core{
	return logs.Core{
		Id:        input.ID,
		Status:    input.Status,
		Feedback:  input.Feedback,
		Mentee:    ModelMeente(input.Mentee),
		User:      UserCore(input.User),
		Class: ClassCore(input.Class),
		CreatedAt: input.CreatedAt,

	}
}

func ModelMeente(input Mentees)logs.CoreMentee{
	return logs.CoreMentee{
		FullName: input.FullName,
		NickName: input.NickName,
		Mayor: input.Mayor,
		Institusi: input.Institusi,
		Phone: input.Phone,
		Telegram: input.Telegram,
		Discord: input.Discord,
		Email: input.Email,
		Status: input.Status,
	}
}

func UserCore(input Users) logs.CoreUsers{
	return logs.CoreUsers{
		FullName: input.FullName,
		Team: input.Team,
	}
}

func ClassCore(input Classes) logs.CoreClasses{
	return logs.CoreClasses{
		Name: input.Name,
	}
}

