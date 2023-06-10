package classes

import (
	"alta/immersive-dashboard-api/features/mentees/mentee"
	"time"
)

type Core struct {
	Id        uint
	CreatedAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
	Name    string   `validation:"required,unique"`
	Tag     string	 `validation:"required,unique"`
	UserID uint
	Mentees []mentee.Core
}

type ClassDataInterface interface {
	Insert(input Core) (error,uint)
	Update(id int, input Core) error
	Deleted(id int) error
	SelectAll()([]Core,error)
	SelectById(id int)(Core,error)

}

type ClassServiceInterface interface {
	Create(input Core) (error,uint)
	Edit(id int, input Core) error
	Deleted(id int) error
	GetAll()([]Core,error)
	GetById(id int)(Core,error)
}