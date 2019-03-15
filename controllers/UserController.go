package controllers

import (
	"net/http"
	"strconv"

	"easy-gin/models"

	"github.com/gin-gonic/gin"
)

// get one
func UserGet(ctx *gin.Context) {
	id, _ := strconv.Atoi(ctx.Param("id"))
	userModel := models.User{}

	data, err := userModel.UserGet(id)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}

// get list
func UserGetList(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	pageSize := ctx.DefaultQuery("page_size", "20")

	pageInt, _ := strconv.Atoi(page)
	pageSizeInt, _ := strconv.Atoi(pageSize)

	userModel := models.User{}

	users, err := userModel.UserGetList(pageInt, pageSizeInt)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":  "success",
		"data": users,
	})
}

// add one
func UserPost(ctx *gin.Context) {
	userModel := models.User{}
	if err := ctx.ShouldBind(&userModel); nil != err {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	id, err := userModel.UserAdd()

	if nil != err {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"msg": "success",
		"uid": id,
	})
}

// update
func UserPut(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)

	if nil != err || 0 == idInt {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "resource identifier not found",
		})
		return
	}

	userModel := models.User{}

	if err := ctx.ShouldBind(&userModel); nil != err {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	_, err = userModel.UserUpdate(idInt)

	if nil != err {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 更新成功返回 204
	ctx.JSON(http.StatusNoContent, gin.H{})
}

// delete
func UserDelete(ctx *gin.Context) {
	id := ctx.Param("id")
	idInt, err := strconv.Atoi(id)

	if nil != err || 0 == idInt {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "resource identifier not found",
		})
		return
	}

	userModel := models.User{}

	_, err = userModel.UserDelete(idInt)

	if nil != err {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
		return
	}

	// 删除成功返回 204
	ctx.JSON(http.StatusNoContent, gin.H{})
}
