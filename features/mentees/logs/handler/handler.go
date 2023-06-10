package handler

import (
	"alta/immersive-dashboard-api/app/helper"
	"alta/immersive-dashboard-api/app/middlewares"
	"alta/immersive-dashboard-api/features/mentees/logs"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type LogsHandler struct {
	logsService logs.LogsServiceInterface
}

func New(service logs.LogsServiceInterface) *LogsHandler {
	return &LogsHandler{
		logsService: service,
	}
}

func (handler *LogsHandler) CreateLogs(c echo.Context) error{
	userId := middlewares.ExtracTokenUserId(c)
	logsInput := LogsRequest{}
	errBind := c.Bind(&logsInput)
	if errBind != nil{
		return helper.StatusBadRequestResponse(c, "error bind data")
	}
	logsCore := RequestToCoreLogs(logsInput)

	id,err := handler.logsService.Add(logsCore,uint(userId) )
	if err != nil{
		if strings.Contains(err.Error(),"validation"){
			return helper.StatusBadRequestResponse(c, err.Error())
		} else {
			return helper.StatusInternalServerError(c, err.Error())
		}
	}

	errGetUser := handler.logsService.GetById(id);
	if errGetUser != nil {
		return helper.StatusInternalServerError(c, errGetUser.Error())
	}

	return helper.StatusOK(c,"Data feedback berhasil ditambahkan")
}

func (handler *LogsHandler) EditLogs(c echo.Context) error{

	logsInput := LogsRequest{}

	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil{
		return helper.StatusBadRequestResponse(c, "id error")
	}

	errBind := c.Bind(&logsInput)
	if errBind != nil{
		return helper.StatusBadRequestResponse(c, "bind error, update failed")
	}

	logsCore :=RequestToCoreLogs(logsInput)

	errGetUser := handler.logsService.GetById(uint(idConv));
	if errGetUser != nil {
		return helper.StatusBadRequestResponse(c, "id tidak valid error, update failed")
	}

	if logsCore.Feedback == "" && logsCore.Status == "" && logsCore.MenteeID == 0 {
		return helper.StatusBadRequestResponse(c, "data tidak ditemukan")
	}else{
		errUpdate := handler.logsService.Edit(logsCore,uint(idConv))
		if errUpdate != nil{
			if strings.Contains(errUpdate.Error(),"validation"){
				return helper.StatusBadRequestResponse(c, errUpdate.Error())
			} else {
				return helper.StatusInternalServerError(c, errUpdate.Error())
			}
		}

	}

	return helper.StatusOK(c,"Data feedback berhasil diupdate")

	}

	func (handler *LogsHandler) DeleteLogs(c echo.Context) error{
		id := c.Param("id")
		idConv, errConv := strconv.Atoi(id)
		if errConv != nil{
			return helper.StatusBadRequestResponse(c, "Delete error")
		}
	
		err := handler.logsService.GetById(uint(idConv))
		if err != nil {
			return helper.StatusNotFoundResponse(c, "id not found")
		}
	
		if err :=handler.logsService.Deleted(uint(idConv));err != nil {
			if strings.Contains(err.Error(), "validation") {
				return helper.StatusBadRequestResponse(c, "error validate payload: " + err.Error())
			} else {
				return helper.StatusInternalServerError(c, err.Error())
			}
		}
		
		return helper.StatusOK(c, "Success delete class")
	
	}

	func (handler *LogsHandler) GetAllLogs(c echo.Context) error{

		data,err := handler.logsService.GetAll()

		var Logs []ResponseLog
		for _,value := range data{
			dataAll := ResponseLogCore(value)
			Logs = append(Logs, dataAll)
		}
		
		if err != nil {
			return helper.StatusInternalServerError(c, err.Error())
		} else {
			return helper.StatusOKWithData(c, "Berhasil mendapatkan semua data mentee", map[string]any{
				"mentees": Logs,
			})
		}
	
	}