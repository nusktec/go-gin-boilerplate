package controllers

import (
	"goginapi/core"
	"goginapi/model"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (u UserController) Retrieve(c *gin.Context) {

	dbc := new(core.Orm)
	db := dbc.Done()

	user := model.User{Name: "jamal", Age: 18}
	db.Create(&user)

	c.JSON(200, gin.H{
		"message": "hello gogin api v1.0.1",
		"api":     "golang",
		"version": "v1.0.2",
		"author":  "kamal 2121",
	})
	return
}

func (u UserController) GetSingle(c *gin.Context) {
	user := c.Params.ByName("id")
	uid := c.Param("id")
	b := c.Request.Host
	r := c.Query("name")

	headerGetToken := c.GetHeader("token")

	c.JSON(200, gin.H{
		"uid":         user,
		"u":           uid,
		"fullpath":    c.FullPath(),
		"a":           b,
		"query-name":  r,
		"header-name": headerGetToken,
	})
}
