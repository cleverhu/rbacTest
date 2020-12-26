package handlers

import (
	"github.com/gin-gonic/gin"
	"rbacTest/src/data/Getter"
)

func Get(ctx *gin.Context) {
	ctx.JSON(200, Getter.UserRoleGetter.UserRoleList())
}
