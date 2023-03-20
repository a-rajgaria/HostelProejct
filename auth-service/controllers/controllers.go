package controllers

import (
	"log"
	"net/http"

	"github.com/a-rajgaria/HostelProject/service"
	"github.com/gin-gonic/gin"
)

func Register(c *gin.Context)  {
		
	var data map[string]string

	if err := c.BindJSON(&data); err!=nil{
		log.Fatalln("Error: ", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":true,
			"message":"Invalid request body",
		})
	}

	customer := service.RegisterCustomer(data)

	c.JSON(http.StatusCreated,customer)
	 	
}

func Login(c *gin.Context) {
	var data map[string]string

	if err := c.BindJSON(&data); err!=nil{
		log.Fatalln("Error: ", err)
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"error":true,
			"message":"Invalid request body",
		})
	}

	token , err2 := service.LoginCustomer(data)
	if err2 != nil {
		log.Println("Error", err2)
		c.JSON(http.StatusBadRequest, err2.Error())
	}
	c.JSON(http.StatusOK,token)
}