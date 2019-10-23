package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/lujiahaoo/go-study/models"
	"github.com/lujiahaoo/go-study/utils"
)

func getYearAndMonth() string {
	year := time.Now().Format("2006")
	month := time.Now().Format("01")
	day := time.Now().Format("02")

	return year + "-" + month + "-" + day
}

func Register(w http.ResponseWriter, r *http.Request) {
	models.Do()
	var validate registerBody
	//可以在这之前短信验证
	json.NewDecoder(r.Body).Decode(&validate)
	//validate.CreatedAt = time.Now().Format("2006-01-02 15:04:05")
	_, err := govalidator.ValidateStruct(validate)
	if err != nil {
		fmt.Println(err)
		utils.ResponseWithJson(
			w,
			utils.Response{Code: http.StatusBadRequest, Msg: "bad params of register body"},
			http.StatusBadRequest,
		)
		return
	}
	var user models.User
	user.Username = validate.UserName
	user.Password = validate.Password
	user.Phone = validate.Phone
	// var count int
	// models.ReadDb.Model(&models.User{}).Where(&user).Count(&count)
	// if count != 0{
	// 	utils.ResponseWithJson(
	// 		w,
	// 		utils.Response{Code: http.StatusBadRequest, Msg: "has exists user"},
	// 		http.StatusBadRequest,
	// 	)
	// 	return
	// }
	if err := models.WriteDb.Create(&user).Error; err != nil {
		fmt.Println(err)
		utils.ResponseWithJson(
			w,
			utils.Response{Code: http.StatusBadRequest, Msg: "user has exists"},
			http.StatusBadRequest,
		)
		return
	}

	utils.ResponseWithJson(
		w,
		utils.Response{Code: http.StatusOK},
		http.StatusOK,
	)
	return
}

func Login(w http.ResponseWriter, r *http.Request) {
	var validate loginBody
	json.NewDecoder(r.Body).Decode(&validate)
	if _, err := govalidator.ValidateStruct(validate); err != nil {
		fmt.Println(err)
		utils.ResponseWithJson(
			w,
			utils.Response{Code: http.StatusBadRequest, Msg: "bad params of login body"},
			http.StatusBadRequest,
		)
		return
	}

	var user models.User
	user.Phone = validate.Phone
	user.Password = validate.Password
	var count int
	models.ReadDb.Model(&models.User{}).Where(&user).Count(&count)
	if count != 0 {
		fmt.Println(count)
		tokenString, err := utils.GenerateToken(validate.Phone) //先查ID再传入，不要用手机号
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
