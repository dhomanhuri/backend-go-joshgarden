package data

import (
	"backend-go/features/users"
	"errors"

	"gorm.io/gorm"
)

type MysqlDB struct {
	DBConn *gorm.DB
}

func Repository(db *gorm.DB) users.Data {
	return &MysqlDB{
		DBConn: db,
	}
}

func (db *MysqlDB) InsertData(data users.Core) error {
	userModel := User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		RoleID:   data.Role.ID,
	}
	result := db.DBConn.Create(&userModel)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func (db *MysqlDB) IsDuplicate(data users.Core) error {
	userModel := User{
		Name:     data.Name,
		Email:    data.Email,
		Password: data.Password,
		RoleID:   data.Role.ID,
	}
	result := db.DBConn.Where("email = ?", data.Email).Find(&userModel)
	if result.Error != nil || result.RowsAffected > 0 {
		return errors.New("duplicate email")
	}
	return nil
}

func (repo MysqlDB) FindUser(email string) (userCore users.Core, err error) {
	userModel := User{}

	result := repo.DBConn.Where("email = ?", email).Find(&userModel)
	if result.RowsAffected == 0 {
		return userCore, errors.New("user not found")
	}
	userCore = userModel.ToCore()
	return userCore, nil
}

func (repo MysqlDB) SelectUser(ID int) (userCore users.Core, err error) {
	userModel := User{}

	result := repo.DBConn.Where("id = ?", ID).Find(&userModel)
	if result.RowsAffected == 0 {
		return userCore, errors.New("user not found")
	}

	userCore = userModel.ToCore()
	return userCore, nil
}

func (repo MysqlDB) SelectAll() ([]users.Core, error) {
	var dataa []User
	result := repo.DBConn.Find(&dataa)
	if result.RowsAffected == 0 {
		return nil, errors.New("failed get data")
	}
	return toCoreList(dataa), nil
}
func (repo MysqlDB) DataDelete(id string) (err error) {
	datauser := User{}
	result := repo.DBConn.Delete(&datauser, id)
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("failed to delete data")
	}
	return nil
}
