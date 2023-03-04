package handlers

import (
	"rbac-api/models"
	"rbac-api/utils"
)

func Login(Username, Password string) utils.Respon {
	var Respon utils.Respon

	user, err := models.GetByEmail(Username)
	if err != nil {
		return utils.ErrorJSON(err.Error())
	}

	Respon.Success = true
	Respon.Data = user
	Respon.Message = "Berhasil Login"
	return Respon
}
