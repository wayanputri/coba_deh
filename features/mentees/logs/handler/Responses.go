package handler

import (
	"alta/immersive-dashboard-api/features/mentees/logs"
	"time"
)

type ResponseUser struct {
	FullName string
	Team     string
}

type ResponseClass struct {
	Name string
}

type ResponseMentee struct {
	FullName  string
	NickName  string
	Mayor     string
	Institusi string
	Phone     string
	Telegram  string
	Discord   string
	Email     string
	Status    string
}

type ResponseLog struct {
	Feedback string
	Status   string
	CreateAd time.Time
	Class ResponseClass
	User ResponseUser
	Mentee ResponseMentee
}

func ResponseUserCore(input logs.CoreUsers)ResponseUser{
	return ResponseUser{
		FullName: input.FullName,
		Team: input.Team,
	}
}

func ResponseMenteeCore(input logs.CoreMentee) ResponseMentee{
	return ResponseMentee{
		FullName: input.FullName,
		NickName: input.NickName,
		Mayor: input.Mayor,
		Institusi: input.Institusi,
		Phone: input.Phone,
		Telegram: input.Telegram,
		Discord: input.Discord,
		Email: input.Email,
		Status: input.Status,

	}
}

func ResponseClassCore(input logs.CoreClasses) ResponseClass{
	return ResponseClass{
		Name: input.Name,
	}
}

func ResponseLogCore(input logs.Core) ResponseLog{
	return ResponseLog{
		Feedback: input.Feedback,
		Status: input.Status,
		CreateAd: input.CreatedAt,
		Class: ResponseClassCore(input.Class),
		User: ResponseUserCore(input.User),
		Mentee: ResponseMenteeCore(input.Mentee),
	}
}