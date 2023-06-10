package logs

import "time"

type Core struct {
	Id        uint
	Status    string
	Feedback  string
	MenteeID  uint
	UserID    uint
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt time.Time
	Mentee CoreMentee
	User CoreUsers
	Class CoreClasses
}

type CoreMentee struct {
	Id        uint
	CreatedAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
	FullName        string  	
	NickName        string 	
	ClassID         uint		
	Status          string 		
	Category        string 	
	Gender          string 
	Graduate        string 
	Mayor           string 
	Phone           string 
	Telegram        string 
	Discord         string 
	Institusi		string
	Email           string 
	EmergencyName   string 
	EmergencyPhone  string 
	EmergencyStatus string 
	
}

type CoreUsers struct {
	Id 				uint 		
	FullName		string   
	Email			string 
	Password		string 
	Team			string
	Role			string 
	Status			string
	CreatedAt		time.Time
	UpdatedAt		time.Time
}

type CoreClasses struct {
	Id        uint
	CreatedAt time.Time
	UpdateAt time.Time
	DeleteAt time.Time
	Name    string   
	Tag     string	 
	UserID uint
}

type LogsDataInterface interface {
	Insert(input Core, userId uint) (uint,error)
	Update(input Core,id uint) error
	SelectById(id uint)error
	Deleted(id uint) error
	SelectAll()([]Core,error)
}

type LogsServiceInterface interface {
	Add(input Core, userId uint) (uint,error)
	Edit(input Core, id uint) error
	GetById(id uint)error
	Deleted(id uint) error
	GetAll()([]Core,error)
}
