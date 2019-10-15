package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/lujiahaoo/go-study/models"
	"github.com/lujiahaoo/go-study/utils"
)

func Login(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		utils.ResponseWithJson(
			w,
			utils.Response{Code: http.StatusBadRequest, Msg: "bad params1"},
			http.StatusBadRequest,
		)
		return
	}

	exists := true //models.IsExists()
	if exists {
		tokenString, err := utils.GenerateToken(string(user.UserName))
		if err != nil {
			utils.ResponseWithJson(
				w,
				utils.Response{Code: http.StatusBadRequest, Msg: "wrong in generate token"},
				http.StatusBadRequest,
			)
			return
		}
		utils.ResponseWithJson(
			w,
			utils.Response{Code: http.StatusOK, Data: models.JwtToken{Token: tokenString}},
		)
		return
	} else {
		utils.ResponseWithJson(
			w,
			utils.Response{Code: http.StatusNotFound, Msg: "the user not found"},
			http.StatusNotFound,
		)
		return
	}
}
