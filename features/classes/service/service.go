package service

import (
	//"be17/main/feature/user"

	"alta/immersive-dashboard-api/features/classes"
	"fmt"

	"github.com/go-playground/validator/v10"
)

type classService struct {
	classData classes.ClassDataInterface
	validate  *validator.Validate
}

// GetById implements classes.ClassServiceInterface
func (service *classService) GetById(id int) (classes.Core, error) {
	data, err := service.classData.SelectById(id)
	return data, err
}

// GetAll implements classes.ClassServiceInterface
func (service *classService) GetAll() ([]classes.Core, error) {
	dataClass, err := service.classData.SelectAll()
	return dataClass, err
}

// Delete implements classes.ClassServiceInterface
func (service *classService) Deleted(id int) error {
	err := service.classData.Deleted(id)
	return err
}

// Edit implements classes.ClassServiceInterface
func (service *classService) Edit(id int, input classes.Core) error {
	err := service.classData.Update(id, input)
	if err != nil {
		return fmt.Errorf("failed to update classses with ID %d:%w", id, err)
	}
	return nil
}

// Create implements classes.ClassServiceInterface
func (service *classService) Create(input classes.Core) (error,uint) {
	errValidate := service.validate.Struct(input)
	if errValidate != nil {
		return errValidate,0
	}

	errInsert,nil := service.classData.Insert(input)
	return errInsert,nil
}

func New(repo classes.ClassDataInterface) classes.ClassServiceInterface {
	return &classService{
		classData: repo,
		validate:  validator.New(),
	}
}
