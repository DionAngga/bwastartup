package main

import (
	"bwastartup/auth"
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
	authService := auth.NewService()

	token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyMn0.n2TLHjR88325WmZwMywzfxBb_tE0_E63kC0vWNTrWOw")
	if err != nil {
		fmt.Println("==================================================")
		fmt.Println("======                                       =====")
		fmt.Println("======             ERROR                     =====")
		fmt.Println("======                                       =====")
		fmt.Println("==================================================")
	}
	if token.Valid {
		fmt.Println("==================================================")
		fmt.Println("======                                       =====")
		fmt.Println("======             VALID                     =====")
		fmt.Println("======                                       =====")
		fmt.Println("==================================================")
	} else {
		fmt.Println("==================================================")
		fmt.Println("======                                       =====")
		fmt.Println("======             INVALID                   =====")
		fmt.Println("======                                       =====")
		fmt.Println("==================================================")
	}

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	api.POST("/avatars", userHandler.UploadAvatar)
	// api.GET("/users/fetch", authMiddleware(authService, usesService), userHandler.FetchUser)

	router.Run()
}
