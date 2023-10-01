package controller

import (
	"btpn/app"
	"btpn/models"
	"github.com/gin-gonic/gin"
)

func (i Picture) Insert(ctx *gin.Context) {
	req := new(app.Picture)

	token := ctx.GetHeader("Token")

	if token == "" {
		ctx.JSON(400, app.ReturnResponse[*string]{
			Status:  400,
			Message: "Token not found",
			Data:    nil,
		})
		return
	}

	uname, err := i.jwt.GetUsername(token)
	if err != nil {
		ctx.JSON(403, app.ReturnResponse[*string]{
			Status:  403,
			Message: "Token Not Valid",
			Data:    nil,
		})
		return
	}

	if er := ctx.BindJSON(req); er != nil {
		errmsg := er.Error()
		ctx.JSON(400, app.ReturnResponse[*string]{
			Status:  400,
			Message: "JSON Parse Err",
			Data:    &errmsg,
		})
		return
	}

	err = i.db.Insert(models.Picture{
		ID:       req.ID,
		Title:    req.Title,
		Caption:  req.Caption,
		PhotoURL: req.PhotoURL,
		UserID:   uname,
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
	return
}

func (i Picture) Update(ctx *gin.Context) {
	req := new(app.Picture)

	token := ctx.GetHeader("Token")
	id := ctx.Query("id")

	if token == "" || id == "" {
		ctx.JSON(400, app.ReturnResponse[*string]{
			Status:  400,
			Message: "Token atau id not found",
			Data:    nil,
		})
		return
	}

	uname, err := i.jwt.GetUsername(token)
	if err != nil {
		ctx.JSON(403, app.ReturnResponse[*string]{
			Status:  403,
			Message: "Token Not Valid",
			Data:    nil,
		})
		return
	}

	if er := ctx.BindJSON(req); er != nil {
		errmsg := er.Error()
		ctx.JSON(400, app.ReturnResponse[*string]{
			Status:  400,
			Message: "JSON Parse Err",
			Data:    &errmsg,
		})
	}

	data, err := i.db.GetById(uname, id)
	if err != nil {
		errmsg := err.Error()
		ctx.JSON(404, app.ReturnResponse[*string]{
			Status:  404,
			Message: "Data Not Found",
			Data:    &errmsg,
		})
		return
	}

	data.Title = req.Title
	data.Caption = req.Caption
	data.PhotoURL = req.PhotoURL

	err = i.db.Update(data)
	if err != nil {
		errmsg := err.Error()
		ctx.JSON(500, app.ReturnResponse[*string]{
			Status:  500,
			Message: "Update errorr",
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

func (i Picture) Delete(ctx *gin.Context) {

	token := ctx.GetHeader("Token")
	id := ctx.Query("id")
	if token == "" || id == "" {
		ctx.JSON(400, app.ReturnResponse[*string]{
			Status:  400,
			Message: "Token not found",
			Data:    nil,
		})
		return
	}

	uname, err := i.jwt.GetUsername(token)
	if err != nil {
		ctx.JSON(403, app.ReturnResponse[*string]{
			Status:  403,
			Message: "Token Not Valid",
			Data:    nil,
		})
		return
	}

	if er := i.db.Delete(uname, id); er != nil {
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
		Message: "Success Delete",
		Data:    nil,
	})
	return
}

func (i Picture) GetAll(ctx *gin.Context) {

	token := ctx.GetHeader("Token")
	if token == "" {
		ctx.JSON(400, app.ReturnResponse[*string]{
			Status:  400,
			Message: "Token not found",
			Data:    nil,
		})
		return
	}

	uname, err := i.jwt.GetUsername(token)
	if err != nil {
		ctx.JSON(403, app.ReturnResponse[*string]{
			Status:  403,
			Message: "Token Not Valid",
			Data:    nil,
		})
		return
	}

	data, err := i.db.GetAll(uname)
	if err != nil {
		ctx.JSON(404, app.ReturnResponse[*string]{
			Status:  404,
			Message: "data not found",
			Data:    nil,
		})
		return
	}

	ctx.JSON(200, app.ReturnResponse[*[]models.Picture]{
		Status:  200,
		Message: "Success Get All",
		Data:    &data,
	})

	return

}
