package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	dsn := "root:mysql123@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)
	// userDelete, err := userRepository.DeleteId(6)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// fmt.Println(userDelete, "deleted")

	userByEmail, err := userRepository.FindByEmail("dio@gmail.com")
	if err != nil {
		fmt.Println(err.Error())
	}
	if userByEmail.ID == 0 {
		fmt.Println("User tidak ditemukan")
	} else {
		fmt.Println(userByEmail.Name)
	}
	userHandler := handler.NewUserHandler(userService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()
}
