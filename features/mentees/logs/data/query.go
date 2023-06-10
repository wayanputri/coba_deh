package data

import (
	"alta/immersive-dashboard-api/features/mentees/logs"
	"errors"

	"gorm.io/gorm"
)

type LogsData struct {
	db *gorm.DB
}


// SelectAll implements logs.LogsDataInterface
func (repo *LogsData) SelectAll() ([]logs.Core, error) {

	var Logs []MenteeLogs
	
	err := repo.db.Preload("User").Preload("Mentee").Preload("Class").Find(&Logs).Error
	if err != nil {
		return nil, err
	}

	var logsAll []logs.Core
	for _,value := range Logs{
		valueLogs := ModelToCoreGetAll(value)
		logsAll = append(logsAll, valueLogs)	
	}
	return logsAll, nil

}

// Deleted implements logs.LogsDataInterface
func (repo *LogsData) Deleted(id uint) error {
	var LogsData MenteeLogs
	errDelete := repo.db.Delete(&LogsData, "id=?", id)
	if errDelete.Error != nil {
		return errDelete.Error
	}
	return nil
}

// SelectById implements logs.LogsDataInterface
func (repo *LogsData) SelectById(id uint) error {

	var logsData MenteeLogs

	tx := repo.db.Where("id = ?", id).First(&logsData)
	if tx != nil {
		return tx.Error
	}

	return nil
}

// Update implements logs.LogsDataInterface
func (repo *LogsData) Update(input logs.Core, id uint) error {
	var logs MenteeLogs

	tx := repo.db.Model(&logs).Where("id=?", id).Updates(CoreToModelLogs(input))
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Insert implements logs.LogsDataInterface
func (repo *LogsData) Insert(input logs.Core, userId uint) (uint, error) {
	logsInput := CoreToModelLogs(input)
	logsInput.UserID = userId
	tx := repo.db.Create(&logsInput)
	if tx.Error != nil {
		return 0, tx.Error
	}
	if tx.RowsAffected == 0 {
		return 0, errors.New("insert failes, row affected = 0")
	}
	Output := LogsModelToCore(logsInput)
	id := Output.Id
	return id, nil
}

func New(db *gorm.DB) logs.LogsDataInterface {
	return &LogsData{
		db: db,
	}
}
