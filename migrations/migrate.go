package migrations

import (
	"fmt"
	"github.com/emerak/golang-bootcamp-2020/domain/model"
	"github.com/jinzhu/gorm"
)

func Migrate(db *gorm.DB) error {
	err := db.AutoMigrate(model.User{}, model.Post{}).Error
	if err != nil {
		fmt.Println("Error Migrate:" + err.Error())
		return err
	}

	return nil
}
