package controller

import (
	"btpn/app"
	"btpn/models"
	"fmt"
	valid "github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
	"time"
)

func (u *Users) Insert(ctx *gin.Context) {
	req := new(app.User)
	if err := ctx.BindJSON(req); err != nil {
		errmsg := err.Error()
		ctx.JSON(400, app.ReturnResponse[*string]{
			Status:  400,
			Message: "JSON Parse Err",
			Data:    &errmsg,
		})
		return
	}

	if val, err := valid.ValidateStruct(req); !val {
		errmsg := err.Error()
		ctx.JSON(400, app.ReturnResponse[*string]{
			Status:  400,
			Message: "Validation Error",
			Data:    &errmsg,
		})
		return
	}

	err := u.db.Insert(models.User{
		ID:        req.ID,
		Usernames: req.Usernames,
		Email:     req.Email,
		Password:  req.Password,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	})

	if err != nil {
		errmsg := err.Error()
		ctx.JSON(500, app.ReturnResponse[*string]{
			Status:  500,
			Message: "Insert errorr",
			Data:    &errmsg,
		})
		return
	}

	ctx.JSON(200, app.ReturnResponse[*string]{
		Status:  200,
		Message: "Success Insert",
		Data:    nil,
	})
	return
}

func (u *Users) Update(ctx *gin.Context) {
	req := new(app.User)
	id := ctx.Param("userId")

	if err := ctx.BindJSON(req); err != nil {
		errmsg := err.Error()
		ctx.JSON(400, app.ReturnResponse[*string]{
			Status:  400,
			Message: "JSON Parse Err",
			Data:    &errmsg,
		})
		return
	}

	if val, err := valid.ValidateStruct(req); !val {
		errmsg := err.Error()
		ctx.JSON(400, app.ReturnResponse[*string]{
			Status:  400,
			Message: "Validation Error",
			Data:    &errmsg,
		})
		return
	}

	data, err := u.db.GetById(id)
	if err != nil {
		ctx.JSON(404, app.ReturnResponse[*string]{
			Status:  404,
			Message: "Data Not Found",
			Data:    nil,
		})
		return
	}

	data.Usernames = req.Usernames
	data.Email = req.Email
	data.Password = req.Password
	data.UpdateAt = time.Now()

	err = u.db.Update(data)
	if err != nil {
		errmsg := err.Error()
		ctx.JSON(500, app.ReturnResponse[*string]{
			Status:  500,
			Message: "Failed Update",
			Data:    &errmsg,
		})
		return
	}

	ctx.JSON(200, app.ReturnResponse[*string]{
		Status:  200,
		Message: "Success Update",
		Data:    nil,
	})
	return
}

func (u *Users) Login(ctx *gin.Context) {
	email := ctx.Query("email")
	password := ctx.Query("password")

	data, err := u.db.GetByUserPassword(email, password)
	if err != nil {
		ctx.JSON(404, app.ReturnResponse[*string]{
			Status:  404,
			Message: "Data Not Found",
			Data:    nil,
		})
		return
	}

	token, err := u.jwt.GenerateToken(fmt.Sprintf("%d", data.ID))
	if err != nil {
		ctx.JSON(500, app.ReturnResponse[*string]{
			Status:  500,
			Message: "error when creating token",
			Data:    nil,
		})
		return
	}

	ctx.JSON(200, app.ReturnResponse[string]{
		Status:  200,
		Message: "Login Success",
		Data:    token,
	})

	return
}

func (u *Users) Delete(ctx *gin.Context) {
	id := ctx.Param("userId")

	data, err := u.db.GetById(id)
	if err != nil {
		ctx.JSON(404, app.ReturnResponse[*string]{
			Status:  404,
			Message: "Data Not Found",
			Data:    nil,
		})
		return
	}

	if er := u.db.Delete(data); er != nil {
		errmsg := er.Error()
		ctx.JSON(500, app.ReturnResponse[*string]{
			Status:  500,
			Message: "error when deleting data",
			Data:    &errmsg,
		})
		return
	}

	ctx.JSON(200, app.ReturnResponse[*string]{
		Status:  200,
		Message: "Login Success",
		Data:    nil,
	})

	return
}
