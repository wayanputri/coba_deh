package service

import (
	"alta/immersive-dashboard-api/features/mentees/mentee"

	"github.com/go-playground/validator/v10"
)

type MenteeService struct {
	menteeData mentee.MenteeDataInterface
	validate   *validator.Validate
}

// AddMentee implements mentee.MenteeServiceInterface
func (service *MenteeService) AddMentee(data mentee.Core) (menteeId uint, err error) {
	if errValidate := service.validate.Struct(data); errValidate != nil {
		return 0, errValidate
	}
	menteeId, errData := service.menteeData.Insert(data)
	if err != nil {
		return 0, errData
	}
	return menteeId, nil
}

// GetMenteeById implements mentee.MenteeServiceInterface
func (service *MenteeService) GetMenteeById(menteeId uint) (mentee *mentee.Core, err error) {
	mentee, errData := service.menteeData.Select(menteeId)
	if errData != nil {
		return nil, errData
	} 
	return mentee, nil
}

// GetMentees implements mentee.MenteeServiceInterface
func (service *MenteeService) GetMentees(query map[string]any) (mentees []mentee.Core, err error) {
	mentees, errData := service.menteeData.SelectAll(query)
	if errData != nil {
		return nil, errData
	}
	return mentees, nil
}

// EditMentee implements mentee.MenteeServiceInterface
func (service *MenteeService) EditMentee(menteeId uint, data mentee.Core) (mentee *mentee.Core, err error) {
	if errValidate := service.validate.Struct(data); errValidate != nil {
		return nil, errValidate
	}
	mentee, errData := service.menteeData.Update(menteeId, data)
	if errData != nil {
		return nil, errData
	}
	return mentee, nil
}


// DeleteMentee implements mentee.MenteeServiceInterface
func (service *MenteeService) DeleteMentee(menteeId uint) error {
	if err := service.menteeData.Delete(menteeId); err != nil {
		return err
	}
	return nil
}

func New(data mentee.MenteeDataInterface) mentee.MenteeServiceInterface {
	return &MenteeService{
		menteeData: data,
		validate:   validator.New(),
	}
}
