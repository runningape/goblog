package controllers

import (
	"fmt"
	"net/http"

	"github.com/runningape/goblog/app/models/user"

	"github.com/runningape/goblog/pkg/view"
)

type AuthController struct {
}

func (*AuthController) Register(w http.ResponseWriter, r *http.Request) {
	view.RenderSimple(w, view.D{}, "auth.register")
}

func (*AuthController) DoRegister(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")

	_user := user.User{
		Name:     name,
		Email:    email,
		Password: password,
	}

	_user.Create()

	if _user.ID > 0 {
		fmt.Fprint(w, "Create user successful. ID:"+_user.GetStringID())
	} else {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Create user failed, Please contact the administrator")
	}
}
