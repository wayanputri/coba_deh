package data

import (
	"alta/immersive-dashboard-api/features/classes"
	"alta/immersive-dashboard-api/features/mentees/mentee/data"

	"gorm.io/gorm"
)

type Classes struct {
	gorm.Model
	Name			string 		`gorm:"type:varchar(50);notNull;unique"`
	Tag				string		`gorm:"type:varchar(10);notNull;unique"`
	UserID		uint
	Mentees 	[]data.Mentees `gorm:"foreignKey:ClassID"`
}

func CoreToModel(input classes.Core) Classes{
	return Classes{
		Model:   gorm.Model{},
		Name:    input.Name,
		Tag:     input.Tag,
		UserID:  input.UserID,
		Mentees: []data.Mentees{},
	}
}
func ModelToCore(input Classes) classes.Core{
	return classes.Core{
		Id:   input.ID,
		Name:    input.Name,
		Tag:     input.Tag,
		UserID:  input.UserID,
		CreatedAt: input.CreatedAt,
		
	}
}

func UpdateClass(input Classes)Classes{
	return Classes{
		Name: input.Name,
		Tag: input.Tag,
		UserID: input.UserID,
	}
}