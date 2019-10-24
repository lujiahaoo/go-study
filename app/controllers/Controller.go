package controllers

import (
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type registerBody struct {
	UserName string `name:"username" valid:"stringlength(15|20),required"`
	Password string `name:"password" valid:"stringlength(15|20),required"`
	Phone    string `name:"phone" valid:"numeric,length(11|11),required"`
}

type loginBody struct {
	Phone    string `name:"phone" valid:"numeric,length(11|11),required"`
	Password string `name:"password" valid:"stringlength(15|20),required"`
}
