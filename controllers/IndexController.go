package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func IndexHome(ctx *gin.Context) {

	//// query string
	//queryVal1 := ctx.Query("val1")
	//queryVal2 := ctx.DefaultQuery("val2", "val2_default")
	//
	//// post form data
	//formVal3 := ctx.PostForm("val3")
	//formVal4 := ctx.DefaultPostForm("val4", "val4_default")
	//
	//// path info
	//pathVal5 := ctx.Param("val5")

	//ctx.String(http.StatusOK, "hello %s %s %s %s %s", queryVal1, queryVal2, formVal3, formVal4, pathVal5)
	ctx.HTML(http.StatusOK, "index/index.html", gin.H{
		"msg": "easy gin",
	})
}
