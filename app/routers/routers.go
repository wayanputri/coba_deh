package routers

import (
	"alta/immersive-dashboard-api/app/middlewares"
	userData "alta/immersive-dashboard-api/features/users/data"
	userHandler "alta/immersive-dashboard-api/features/users/handler"
	userService "alta/immersive-dashboard-api/features/users/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"

	_classData "alta/immersive-dashboard-api/features/classes/data"
	_classHandler "alta/immersive-dashboard-api/features/classes/handler"
	_classService "alta/immersive-dashboard-api/features/classes/service"

	_menteeData "alta/immersive-dashboard-api/features/mentees/mentee/data"
	_menteeHandler "alta/immersive-dashboard-api/features/mentees/mentee/handler"
	_menteeService "alta/immersive-dashboard-api/features/mentees/mentee/service"

	_logsData "alta/immersive-dashboard-api/features/mentees/logs/data"
	_logsHandler "alta/immersive-dashboard-api/features/mentees/logs/handler"
	_logsService "alta/immersive-dashboard-api/features/mentees/logs/service"
)

func InitRouters(db *gorm.DB, e *echo.Echo) {
	UserData := userData.New(db)
	UserService := userService.New(UserData)
	UserHandler := userHandler.New(UserService)
  
  classData := _classData.New(db)
	classService := _classService.New(classData)
	classHandlerAPI := _classHandler.New(classService)

	menteeData := _menteeData.New(db)
	menteeService := _menteeService.New(menteeData)
	menteeHandler := _menteeHandler.New(menteeService)

	logsData := _logsData.New(db)
	logsService := _logsService.New(logsData)
	logsHandlerAPI := _logsHandler.New(logsService)

	e.POST("/users", UserHandler.PostUserHandler, middlewares.JWTMiddleware())
	e.PUT("/users/:id", UserHandler.PutUserHandler, middlewares.JWTMiddleware())
	e.GET("/users/:id", UserHandler.GetUserHandler, middlewares.JWTMiddleware())
	e.GET("/users", UserHandler.GetAllUsersHandler, middlewares.JWTMiddleware())
	e.DELETE("/users/:id", UserHandler.DeleteUserHandler, middlewares.JWTMiddleware())
	e.POST("/login", UserHandler.PostLoginUserHandler)
	e.GET("/logout", UserHandler.PutLogoutHandler, middlewares.JWTMiddleware())

	e.POST("/classes",classHandlerAPI.CreateClass,middlewares.JWTMiddleware())
	e.PUT("/classes/:id",classHandlerAPI.UpdateClass,middlewares.JWTMiddleware())
	e.DELETE("/classes/:id",classHandlerAPI.DeleteClass,middlewares.JWTMiddleware())
	e.GET("/classes",classHandlerAPI.GetAll,middlewares.JWTMiddleware())
	e.GET("/classes/:id",classHandlerAPI.GetByIdClass,middlewares.JWTMiddleware())

	e.POST("/mentees", menteeHandler.PostMenteeHandler, middlewares.JWTMiddleware())
	e.GET("/mentees", menteeHandler.GetMenteesHandler, middlewares.JWTMiddleware())
	e.GET("/mentees/:id", menteeHandler.GetMenteeByIdHandler, middlewares.JWTMiddleware())
	e.PUT("/mentees/:id", menteeHandler.UpdateMenteeHandler, middlewares.JWTMiddleware())
	e.DELETE("/mentees/:id", menteeHandler.DeleteMenteeHandler, middlewares.JWTMiddleware())

	e.POST("/feedbacks",logsHandlerAPI.CreateLogs,middlewares.JWTMiddleware())
	e.PUT("/feedbacks/:id",logsHandlerAPI.EditLogs,middlewares.JWTMiddleware())
	e.DELETE("/feedbacks/:id", logsHandlerAPI.DeleteLogs,middlewares.JWTMiddleware())
	e.GET("/feedbacks", logsHandlerAPI.GetAllLogs, middlewares.JWTMiddleware())
}
