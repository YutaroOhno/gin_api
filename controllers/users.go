package controllers

import (
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-xorm/xorm"

	"gin_api/models"
)

func Index(c *gin.Context) {
	engine, err := xorm.NewEngine("mysql", "root:@/bookshelf")
	if err != nil {
		log.Fatal(err)
	}

	users := []models.User{}
	engine.Find(&users)
	c.JSON(200, users)
}

func Show(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)

	if err != nil {
		log.Fatal(err)
	}

	if id <= 0 {
		c.JSON(400, gin.H{"message": "idは1以上にしてね"})
		return
	}

	user := models.User{
		ID: id,
	}

	engine, err := xorm.NewEngine("mysql", "root:@/bookshelf")
	if err != nil {
		log.Fatal(err)
	}
	engine.Get(&user)

	c.JSON(200, user)
}

func Create(c *gin.Context) {
	newUser := models.User{}
	id, err := strconv.Atoi(c.PostForm("id"))
  if err != nil {
    log.Fatal(err)
  }

  newUser.ID = id
 	newUser.NickName = c.PostForm("nickname")

	engine, err := xorm.NewEngine("mysql", "root:@/bookshelf")
  if err != nil {
    log.Fatal(err)
  }
	engine.Insert(newUser)

	c.JSON(200, newUser)
}

func Delete(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		log.Fatal(err)
	}

	engine, err := xorm.NewEngine("mysql", "root:@/bookshelf")
	if err != nil {
		log.Fatal(err)
	}
	user := models.User{ID: id}
	_, err = engine.Id(id).Delete(&user)
	if err != nil {
		log.Fatal(err)
	}

	c.JSON(200, "")
}
