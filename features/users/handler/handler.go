package handler

import (
	"alta/immersive-dashboard-api/app/helper"
	"alta/immersive-dashboard-api/app/middlewares"
	"alta/immersive-dashboard-api/features/users"
	"fmt"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	userService users.UserServiceInterface
}

func New(service users.UserServiceInterface) *UserHandler {
	return &UserHandler{
		userService: service,
	}
}

func (handler *UserHandler) PostUserHandler(c echo.Context) error {
	payload := users.RequestUser{}
	fmt.Println(payload)
	if err := c.Bind(&payload); err != nil {
		if err == echo.ErrBadRequest {
			return helper.StatusBadRequestResponse(c, "error bind payload " + err.Error())
		} 
	}

	userId := middlewares.ExtracTokenUserId(c)
	userLoggedIn, err := handler.userService.GetUser(uint(userId))
	if err != nil {
		return err
	}
	
	if userLoggedIn.Role == "admin" {
		payloadMap := users.RequestToCoreUser(payload)
		userId, err := handler.userService.AddUser(payloadMap); if err != nil {
			if strings.Contains(err.Error(), "validation") {
				return helper.StatusBadRequestResponse(c, "error validate payload: " + err.Error())
			} else if strings.Contains(err.Error(), "Duplicate entry") {
				return helper.StatusBadRequestResponse(c, "email tidak tersedia")
			}                   
		}
		user, errGetUser := handler.userService.GetUser(userId);
		if errGetUser != nil {
			return helper.StatusInternalServerError(c, err.Error())
		}
		responseDataUser := users.CoreToResponseUser(user)
		return helper.StatusCreated(c, "User berhasil ditambahkan", map[string]any{
			"user": responseDataUser,
		})
	} else {
		return helper.StatusForbiddenResponse(c, "Anda haruslah seorang admin agar bisa menambahkan user")
	}
}

func (handler *UserHandler) PutUserHandler(c echo.Context) error {
	userId, errParam := strconv.Atoi(c.Param("id"))
	if userId == 0 || errParam != nil {
		return helper.StatusNotFoundResponse(c, errParam.Error())
	}
	newData := users.RequestUser{}
	if errBind := c.Bind(&newData); errBind != nil {
		if errBind == echo.ErrBadRequest {
			return helper.StatusBadRequestResponse(c, "error bind payload " + errBind.Error())
		}
	}

	userSession := middlewares.ExtracTokenUserId(c)
	userLoggedIn, err := handler.userService.GetUser(uint(userSession))
	if err != nil {
		return err
	}
	
	if userLoggedIn.Role == "admin" {
		newDataMap := users.RequestToCoreUser(newData)
		if err := handler.userService.EditUser(uint(userId), newDataMap); err != nil {
			if strings.Contains(err.Error(), "validation") {
				return helper.StatusBadRequestResponse(c, "error validate payload: " + err.Error())
			} else {
				return helper.StatusInternalServerError(c, err.Error())
			}
		}
		user, errGetUser := handler.userService.GetUser(uint(userId));
		if errGetUser != nil {
			return helper.StatusInternalServerError(c, err.Error())
		}
		return helper.StatusOKWithData(c, "Berhasil memperbarui data pengguna", user)
	} else {
		return helper.StatusForbiddenResponse(c, "Anda harus admin jika ingin mengedit resource ini")
	}
}

func (handler *UserHandler) GetUserHandler(c echo.Context) error {
	userId, err := strconv.Atoi(c.Param("id"))
	if userId == 0 || err != nil {
		return helper.StatusNotFoundResponse(c, "Data user tidak ditemukan")
	}
	user, err := handler.userService.GetUser(uint(userId))
	if err != nil {
		return helper.StatusBadRequestResponse(c, err.Error())
	}
	mapUser := users.CoreToResponseUser(user)
	return helper.StatusOKWithData(c, "Berhasil mendapatkan data user", map[string]any{
		"user": mapUser,
	})
}

func (handler *UserHandler) GetAllUsersHandler(c echo.Context) error {
	allUsers, errGetAll := handler.userService.GetAllUser()
	if errGetAll != nil {
		return helper.StatusInternalServerError(c, errGetAll.Error())
	}
	var mapAllUsers []users.ResponseUser
	for _, user := range allUsers {
		mapUser := users.CoreToResponseUser(user)
		mapAllUsers = append(mapAllUsers, mapUser)
	}

	return helper.StatusOKWithData(c, "Berhasil mendapatkan semua data pengguna terdaftar", map[string]any{
		"users": mapAllUsers,
	})
}

func (handler *UserHandler) DeleteUserHandler(c echo.Context) error {
	userId, errParam := strconv.Atoi(c.Param("id"))
	if userId == 0 || errParam != nil {
		return helper.StatusNotFoundResponse(c, errParam.Error())
	}
	user := middlewares.ExtracTokenUserId(c)
	userLoggedIn, err := handler.userService.GetUser(uint(user))
	if err != nil {
		return err
	}

	if userLoggedIn.Role == "admin" {
		if errDelete := handler.userService.DeleteUser(uint(userId)); errDelete != nil {
			if strings.Contains(err.Error(), "validation") {
				return helper.StatusBadRequestResponse(c, "error validate payload: " + err.Error())
			} else {
				return helper.StatusInternalServerError(c, err.Error())
			}
		}
		return helper.StatusOK(c, "Berhasil menghapus data pengguna")
	} else {
		return helper.StatusForbiddenResponse(c, "Anda harus admin jika ingin menghapus resource ini")
	}
}

func (handler *UserHandler) PostLoginUserHandler(c echo.Context) error {
	var payload users.LoginUser
	if errBind := c.Bind(&payload); errBind != nil {
		if errBind == echo.ErrBadRequest {
			return helper.StatusBadRequestResponse(c, "error bind payload " + errBind.Error())
		}
	}

	userId, err := handler.userService.LoginUser(payload.Email, payload.Password)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return helper.StatusAuthorizationErrorResponse(c, "Input tidak valid, harap isi email dan password sesuai ketentuan")
		}
		if strings.Contains(err.Error(), "email tidak terdaftar") {
			return helper.StatusBadRequestResponse(c, "Email yang anda berikan tidak terdaftar")
		}
		if strings.Contains(err.Error(), "kredensial tidak cocok") {
			return helper.StatusAuthorizationErrorResponse(c, "Kredensial yang anda berikan tidak valid") 
		}
	}

	accessToken, err := middlewares.CreateToken(userId)
	if err != nil {
		return err
	}

	userData, err := handler.userService.GetUser(uint(userId))
	if err != nil {
		return err
	}

		return helper.StatusCreated(c, "Login Berhasil", map[string]any{
		"accessToken": accessToken, 
		"user": userData,
	})
}

func (handler *UserHandler) PutLogoutHandler(c echo.Context) error {
	userId := middlewares.ExtracTokenUserId(c)
	if err := handler.userService.LogoutUser(uint(userId)); err != nil {
		return helper.StatusBadRequestResponse(c, err.Error())
	}

	return helper.StatusOK(c, "Berhasil logout")
}
