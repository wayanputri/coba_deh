package handler

import (
	"alta/immersive-dashboard-api/app/helper"
	"alta/immersive-dashboard-api/features/mentees/mentee"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type MenteeHandler struct {
	menteeService mentee.MenteeServiceInterface
}

func New(service mentee.MenteeServiceInterface) *MenteeHandler {
	return &MenteeHandler{
		menteeService: service,
	}
}

func (handler *MenteeHandler) PostMenteeHandler(c echo.Context) error {
	var payload mentee.RequestCore
	if errBind := c.Bind(&payload); errBind != nil {
		return helper.StatusBadRequestResponse(c, "error bind payload" + errBind.Error())
	}

	payloadMap := mentee.RequestToCoreMentee(payload)
	menteeId, err := handler.menteeService.AddMentee(payloadMap)
	if err != nil {
		if strings.Contains(err.Error(), "validation") {
			return helper.StatusBadRequestResponse(c, "error validation: " + err.Error())
		} else {
			return helper.StatusInternalServerError(c, err.Error())
		}
	} else {
		mentee, err := handler.menteeService.GetMenteeById(menteeId)
		if err != nil {
			return helper.StatusInternalServerError(c, err.Error())
		}
		return helper.StatusCreated(c, "Berhasil menambahkan mentee", map[string]any{
			"mentee": mentee,
		})
	}
}

func (handler *MenteeHandler) GetMenteeByIdHandler(c echo.Context) error {
	menteeId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return helper.StatusBadRequestResponse(c, "error get param: " + err.Error())
	}
	mentee, err := handler.menteeService.GetMenteeById(uint(menteeId))
	if err != nil {
		return helper.StatusInternalServerError(c, err.Error())
	}
	return helper.StatusOKWithData(c, "Berhasil mendapatkan data mentee", map[string]any{
		"mentee": mentee,
	})
}

func (handler *MenteeHandler) GetMenteesHandler(c echo.Context) error {
	querys := map[string]any{}
	classId := c.QueryParam("classId")
	if classId != "" {
		querys["class_id"] = classId
	}
	status := c.QueryParam("status")
	if status != "" {
		querys["status"] = status
	}
	category := c.QueryParam("category")
	if category != "" {
		querys["category"] = category
	}
	allmentees, err := handler.menteeService.GetMentees(querys)
	if err != nil {
		return helper.StatusInternalServerError(c, err.Error())
	} else {
		return helper.StatusOKWithData(c, "Berhasil mendapatkan semua data mentee", map[string]any{
			"mentees": allmentees,
		})
	}
}

func (handler *MenteeHandler) UpdateMenteeHandler(c echo.Context) error {
	menteeId, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil {
		return helper.StatusBadRequestResponse(c, "error get param:" + errParam.Error())
	}
	var payload mentee.RequestCore
	if errBind := c.Bind(&payload); errBind != nil {
		return helper.StatusBadRequestResponse(c, "error bind payload: " + errBind.Error())
	}
	payloadMap := mentee.RequestToCoreMentee(payload)
	mentee, err := handler.menteeService.EditMentee(uint(menteeId), payloadMap) 
	if err != nil {
		return helper.StatusInternalServerError(c, err.Error())
	} else {
		return helper.StatusOKWithData(c, "Berhasil memperbarui data pengguna", map[string]any{
			"mentee": mentee,
		})
	}
}

func (handler *MenteeHandler) DeleteMenteeHandler(c echo.Context) error {
	menteeId, errParam := strconv.Atoi(c.Param("id"))
	if errParam != nil {
		return helper.StatusBadRequestResponse(c, "error get param:" + errParam.Error())
	}
	if err := handler.menteeService.DeleteMentee(uint(menteeId)); err != nil {
		return helper.StatusInternalServerError(c, err.Error())
	}
	return nil
}