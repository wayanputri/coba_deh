package data

import (
	"alta/immersive-dashboard-api/features/mentees/logs/data"
	"alta/immersive-dashboard-api/features/mentees/mentee"

	"gorm.io/gorm"
)

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
	EmergencyName		string	`gorm:"type:varchar(50)"`
	EmergencyPhone	string	`gorm:"type:varchar(50)"`
	EmergencyStatus	string	`gorm:"type:varchar(50)"`
	Logs						[]data.MenteeLogs `gorm:"foreignKey:MenteeID"`
}

func CoreToMenteeModel(mentee mentee.Core) Mentees {
	return Mentees{
		FullName: mentee.FullName,
		NickName: mentee.NickName,
		ClassID: mentee.ClassID,
		Status: mentee.Status,
		Category: mentee.Category,
		Gender: mentee.Gender,
		Graduate: mentee.Graduate,
		Mayor: mentee.Mayor,
		Phone: mentee.Phone,
		Telegram: mentee.Telegram,
		Discord: mentee.Discord,
		Institusi: mentee.Institusi,
		Email: mentee.Email,
		EmergencyName: mentee.EmergencyName,
		EmergencyPhone: mentee.EmergencyPhone,
		EmergencyStatus: mentee.EmergencyStatus,
		Logs: mentee.Logs,
	}
}

func MenteeModelToCore(menteeModel Mentees) mentee.Core {
	return mentee.Core{
		Id: menteeModel.ID,
		CreatedAt: menteeModel.CreatedAt,
		UpdateAt: menteeModel.UpdatedAt,
		FullName: menteeModel.FullName,
		NickName: menteeModel.NickName,
		ClassID: menteeModel.ClassID,
		Status: menteeModel.Status,
		Category: menteeModel.Category,
		Gender: menteeModel.Gender,
		Graduate: menteeModel.Graduate,
		Mayor: menteeModel.Mayor,
		Phone: menteeModel.Phone,
		Telegram: menteeModel.Telegram,
		Discord: menteeModel.Discord,
		Institusi: menteeModel.Institusi,
		Email: menteeModel.Email,
		EmergencyName: menteeModel.EmergencyName,
		EmergencyPhone: menteeModel.EmergencyPhone,
		EmergencyStatus: menteeModel.EmergencyStatus,
		Logs: menteeModel.Logs,
	}
}