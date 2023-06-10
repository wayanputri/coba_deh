package data

import (
	"alta/immersive-dashboard-api/features/classes"
	"errors"

	"gorm.io/gorm"
)

type classQuery struct {
	db *gorm.DB
}

// SelectById implements classes.ClassDataInterface
func (repo *classQuery) SelectById(id int) (classes.Core, error) {
	var classData Classes
	tx := repo.db.First(&classData, id)
	if tx.Error != nil {
		return classes.Core{}, tx.Error
	}

	userCore := ModelToCore(classData)

	return userCore, nil
}

// SelectAll implements classes.ClassDataInterface
func (repo *classQuery) SelectAll() ([]classes.Core, error) {
	var classDataAll []Classes
	tx := repo.db.Find(&classDataAll)
	if tx.Error != nil {
		return []classes.Core{}, tx.Error
	}

	var classAll []classes.Core
	for _, value := range classDataAll {
		classCore := ModelToCore(value)
		classAll = append(classAll, classCore)
	}
	return classAll, nil
}

// Delete implements classes.ClassDataInterface
func (repo *classQuery) Deleted(id int) error {
	var classData Classes
	errDelete := repo.db.Delete(&classData, "id=?", id)
	if errDelete.Error != nil {
		return errDelete.Error
	}
	return nil
}

// Update implements classes.ClassDataInterface
func (repo *classQuery) Update(id int, input classes.Core) error {
	classInput := CoreToModel(input)
	err := repo.db.Model(&Classes{}).Where("id=?", id).Updates(UpdateClass(classInput))
	if err != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("no rows affected, update failed")
	}
	return nil
}

// Insert implements classes.ClassDataInterface
func (repo *classQuery) Insert(input classes.Core) (error,uint) {
	classInput := CoreToModel(input)
	tx := repo.db.Create(&classInput)
	if tx.Error != nil {
		return tx.Error,0
	}
	if tx.RowsAffected == 0 {
		return errors.New("insert failes, row affected = 0"),0
	}
	output := ModelToCore(classInput)
	id := output.Id
	return nil,id
}

func New(db *gorm.DB) classes.ClassDataInterface {
	return &classQuery{
		db: db,
	}
}
