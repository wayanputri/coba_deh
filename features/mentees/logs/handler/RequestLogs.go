package handler

import "alta/immersive-dashboard-api/features/mentees/logs"

type LogsRequest struct {
	Status   string `json:"proof" form:"proof"`	
	Feedback string `json:"notes" form:"notes"`
	MenteeID uint `json:"id_mentee" form:"id_mentee"`
	UserID   uint `json:"id_user" form:"id_user"`
}

type Response struct{
	Status   string 
	Feedback string
	MenteeID uint
	UserID   uint
}


func RequestToCoreLogs(input LogsRequest) logs.Core{
	return logs.Core{
		Status: input.Status,
		Feedback: input.Feedback,
		MenteeID: input.MenteeID,
		UserID: input.UserID,
	}
}

func CoreToResponseLogs(input logs.Core) Response{
	return Response{
		Status: input.Status,
		Feedback: input.Feedback,
		MenteeID: input.MenteeID,
		UserID: input.UserID,
	}
}


