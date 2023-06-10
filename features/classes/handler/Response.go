package handler

import (
	"alta/immersive-dashboard-api/features/classes"
	"time"
)

type Response struct {
	Id        uint
	Name      string
	Tag       string
	UserID    uint
	CreatedAt time.Time
}

func CoreToResponse(input classes.Core) Response{
return Response{
	Id:   input.Id,
	Name:    input.Name,
	Tag:     input.Tag,
	UserID:  input.UserID,
	CreatedAt: input.CreatedAt,
}
}
