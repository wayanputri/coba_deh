package handler

import (
	"alta/immersive-dashboard-api/app/helper"
	"alta/immersive-dashboard-api/features/classes"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ClassHandler struct {
	classService classes.ClassServiceInterface
}

func New(service classes.ClassServiceInterface) *ClassHandler{
	return &ClassHandler{
		classService: service,
	}
}

func (handler *ClassHandler) CreateClass(c echo.Context) error{

	classInput := ClassRequest{}
	errBind := c.Bind(&classInput)
	if errBind != nil{
		return helper.StatusBadRequestResponse(c, "error bind data")
	}
	classCore := RequestToCore(classInput)

	err, id := handler.classService.Create(classCore)
	if err!= nil {
		if strings.Contains(err.Error(), "validation") {
			return helper.StatusBadRequestResponse(c, "error validate: " + err.Error())
		} else {
			return helper.StatusInternalServerError(c, err.Error())
		}
	}
	class, errGetUser := handler.classService.GetById(int(id));
	if errGetUser != nil {
		return helper.StatusInternalServerError(c, errGetUser.Error())
	}
	data := CoreToResponse(class)
	return helper.StatusOKWithData(c, "Berhasil menambah data class", data)

}

func (handler *ClassHandler) UpdateClass(c echo.Context) error{

	id := c.Param("id")
	classInput := ClassRequest{}
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil{
		return helper.StatusBadRequestResponse(c, "id error")
	}
	errBind := c.Bind(&classInput)
	if errBind != nil{
		return helper.StatusBadRequestResponse(c, "bind error, update failed")
	}

	classCore :=RequestToCore(classInput)
	if classCore.Name == ""&&classCore.UserID==0 && classCore.Tag==""{
		return helper.StatusBadRequestResponse(c, "data tidak ditemukan")
	}else{

	errUpdate := handler.classService.Edit(idConv,classCore) 
	if errUpdate!= nil {
		if strings.Contains(errUpdate.Error(), "validation") {
			return helper.StatusBadRequestResponse(c, "error validate: " + errUpdate.Error())
		} else {
			return helper.StatusInternalServerError(c, errUpdate.Error())
		}
	}
	}
		
	class, errGetUser := handler.classService.GetById(idConv);
	if errGetUser != nil {
		return helper.StatusInternalServerError(c, errGetUser.Error())
	}
	
	data := CoreToResponse(class)
	return helper.StatusOKWithData(c, "Berhasil menambah data class", data)
}

func (handler *ClassHandler) DeleteClass(c echo.Context) error{
	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil{
		return helper.StatusBadRequestResponse(c, "Delete error")
	}

	result, err := handler.classService.GetById(idConv)
	if err != nil {
		return helper.StatusBadRequestResponse(c, "id not found")
	}
	data := CoreToResponse(result)

	if err :=handler.classService.Deleted(idConv);err != nil {
		if strings.Contains(err.Error(), "validation") {
			return helper.StatusBadRequestResponse(c, "error validate payload: " + err.Error())
		} else {
			return helper.StatusInternalServerError(c, err.Error())
		}
	}
	
	return helper.StatusOKWithData(c, "Success delete class", data)

}

func (handler *ClassHandler) GetAll(c echo.Context) error{

	dataClass, err := handler.classService.GetAll()
	if err != nil{
		return helper.StatusBadRequestResponse(c, "error read class")
	}
	var ClassResponAll []Response
	for _,value := range dataClass{
		dataResponse :=CoreToResponse(value)
		ClassResponAll = append(ClassResponAll, dataResponse)
	}
	return helper.StatusOKWithData(c, "Success read data class", ClassResponAll)
}

func (handler *ClassHandler) GetByIdClass(c echo.Context) error{

	id := c.Param("id")
	idConv, errConv := strconv.Atoi(id)
	if errConv != nil{
		return helper.StatusBadRequestResponse(c, "error, id tidak sesuai")
	}

	result, err := handler.classService.GetById(idConv)
	if err != nil {
		return helper.StatusBadRequestResponse(c, "error read data")
	}
	ClassResponse := CoreToResponse(result)

	return helper.StatusOKWithData(c, "Success read data class", ClassResponse)
	}