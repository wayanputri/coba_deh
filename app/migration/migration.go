package migration

import (
	classData "alta/immersive-dashboard-api/features/classes/data"
	logData "alta/immersive-dashboard-api/features/mentees/logs/data"
	menteeData "alta/immersive-dashboard-api/features/mentees/mentee/data"
	userData "alta/immersive-dashboard-api/features/users/data"

	"gorm.io/gorm"
)

func InitMigrate(db *gorm.DB) error {
	// TODO: Tambahkan setiap model fitur didalam parameter pisahkan dengan koma
	err := db.AutoMigrate(&userData.Users{}, &classData.Classes{}, &menteeData.Mentees{}, &logData.MenteeLogs{})
	if err != nil {
		return err
	}

	return nil
}