package users

import "time"

type Core struct {
	Id 					uint 		
	FullName		string   
	Email				string `validate:"required,email"`
	Password		string `validate:"required"`
	Team				string
	Role				string 
	Status			string
	CreatedAt		time.Time
	UpdatedAt		time.Time
}

type LoginUser struct {
	Email			string `validate:"required,email"`
	Password	string `validate:"required"` 
}

type RequestUser struct{
	FullName		string		`json:"fullname" form:"fullname"`			
	Email			string		`json:"email" form:"email"`				
	Password		string		`json:"password" form:"password"`				
	Team 			string 		`json:"team" form:"team"`				
	Role			string 		`json:"role" form:"role"`		
	Status			string 		`json:"status" form:"status"`				
}

type ResponseUser struct{
	Id 					uint		`json:"id"`
	FullName		string		`json:"fullname"`			
	Email			string		`json:"email"`							
	Team 			string 		`json:"team"`				
	Role			string 		`json:"role"`		
	Status			string 		`json:"status"`
	CreatedAt		time.Time `json:"createdAt"`
	UpdatedAt		time.Time 	`json:"updatedAt"`				
}

func RequestToCoreUser(input RequestUser) Core{
	return Core{
		FullName: input.FullName,
		Email: input.Email,
		Password: input.Password,
		Team: input.Team,
		Role: input.Role,
		Status: input.Status,
	}
}

func CoreToResponseUser(data Core) ResponseUser {
	return ResponseUser{
		Id:	data.Id,
		FullName: data.FullName,
		Email: data.Email,
		Team: data.Team,
		Role: data.Role,
		Status: data.Status,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
	}
}
type UserDataInterface interface {
	Insert(data Core) (uint, error)
	Update(userId uint, data Core) error
	Select(userId uint) (Core, error)
	SelectAll() ([]Core, error)
	Delete(userId uint) error
	Login(email string, password string) (int, error)
	Logout(userId uint) error
}

type UserServiceInterface interface {
	AddUser(data Core) (uint, error)
	EditUser(userId uint, data Core) error
	GetUser(userId uint) (Core, error)
	GetAllUser() ([]Core, error)
	DeleteUser(userId uint) error
	LoginUser(email string, password string) (int, error)
	LogoutUser(userId uint) error
}