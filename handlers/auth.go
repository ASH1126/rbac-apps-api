package handlers

import (
	"rbac-api/models"
	"rbac-api/utils"
	"time"
)

type envelope map[string]interface{}

func Login(Username, Password string) utils.Respon {
	var Respon utils.Respon

	// look up the user by email
	user, err := models.GetByEmail(Username)
	if err != nil {
		return utils.ErrorJSON(err.Error())
	}

	// validate the user's password
	validPassword, err := user.PasswordMatches(Password)
	if err != nil || !validPassword {
		return utils.ErrorJSON("invalid username/password")
	}

	// make sure user is active
	if user.Active == 0 {
		return utils.ErrorJSON("user is not active")
	}

	// we have a valid user, so generate a token
	token, err := models.GenerateToken(user.ID, 24*time.Hour)
	if err != nil {
		return utils.ErrorJSON(err.Error())
	}

	// save it to the database
	err = models.Insert(*token, *user)
	if err != nil {
		return utils.ErrorJSON(err.Error())
	}

	Respon.Success = true
	Respon.Data = envelope{"token": token, "user": user}
	Respon.Message = "Berhasil Login"
	return Respon
}
