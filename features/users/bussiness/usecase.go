package bussiness

import (
	"backend-go/features/users"
	"backend-go/middlewares"
	"errors"
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

type UserRepo struct {
	userData users.Data
}

// type authController struct {
// 	userData users.Data
// }

func UserBussines(data users.Data) users.Bussines {
	return &UserRepo{
		userData: data,
	}
}

func (repo *UserRepo) Login(userCore users.Core) (id int, token string, err error) {
	result, errLogin := repo.userData.FindUser(userCore.Email)
	if errLogin != nil {
		return 0, "", errLogin
	}
	passCompare := bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(userCore.Password))
	if passCompare != nil {
		return 0, "", errors.New("wrong password")
	}

	token, _ = middlewares.GenerateToken(result.UserID, result.Email)
	fmt.Println(result)
	return result.UserID, token, nil
}

func (repo *UserRepo) Register(data users.Core) error {
	dupErr := repo.userData.IsDuplicate(data)
	if dupErr != nil {
		return dupErr
	}

	passByte := []byte(data.Password)
	hashByte, _ := bcrypt.GenerateFromPassword(passByte, bcrypt.DefaultCost)
	data.Password = string(hashByte)
	result := repo.userData.InsertData(data)
	if result != nil {
		return result
	}
	return nil
}
