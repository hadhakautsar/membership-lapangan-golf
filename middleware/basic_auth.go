package middleware

import (
	config "membership-lapangan-golf/configs"
	"membership-lapangan-golf/models"

	"github.com/labstack/echo/v4"
)

func BasicAuthDB(email, password string, c echo.Context) (bool, error) {
	var member models.Member
	db, err := config.InitDB()
	if err != nil {
		return false, err
	}
	err = db.Where("email = ? AND password = ?", email, password).First(&member).Error
	if err != nil {
		return false, err
	}
	return true, nil
}
