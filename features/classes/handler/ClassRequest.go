package handler

import (
	"alta/immersive-dashboard-api/features/classes"
)

type ClassRequest struct {
	Name   string `json:"name" form:"name"`
	Tag    string `json:"initialClass" form:"initialClass"`
	UserID uint    `json:"userId" form:"userId"`
}

func RequestToCore(input ClassRequest) classes.Core{
	return classes.Core{
		Name: input.Name,
		Tag: input.Tag,
		UserID: input.UserID,
	}
}

