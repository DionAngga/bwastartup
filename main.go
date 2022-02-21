package main

import (
	"bwastartup/auth"
	"bwastartup/handler"
	"bwastartup/helper"
	"bwastartup/user"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
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
	//fmt.Println(authService.GenerateToken(100))

	// token, err := authService.ValidateToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyMH0.Hod7cMfrJ80_ZtkLb-z0SXgPHhL2-lah-hSgPe66Bf4")
	// if err != nil {
	// 	fmt.Println("==================================================")
	// 	fmt.Println("======                                       =====")
	// 	fmt.Println("======             ERROR                     =====")
	// 	fmt.Println("======                                       =====")
	// 	fmt.Println("==================================================")
	// }
	// if token.Valid {
	// 	fmt.Println("==================================================")
	// 	fmt.Println("======                                       =====")
	// 	fmt.Println("======             VALID                     =====")
	// 	fmt.Println("======                                       =====")
	// 	fmt.Println("==================================================")
	// } else {
	// 	fmt.Println("==================================================")
	// 	fmt.Println("======                                       =====")
	// 	fmt.Println("======             INVALID                   =====")
	// 	fmt.Println("======                                       =====")
	// 	fmt.Println("==================================================")
	// }

	userHandler := handler.NewUserHandler(userService, authService)

	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)
	api.POST("/sessions", userHandler.Login)
	api.POST("/email_checkers", userHandler.CheckEmailAvailability)
	//api.POST("/avatars", userHandler.UploadAvatar)
	api.POST("/avatars", authMiddleware(authService, userService), userHandler.UploadAvatar)
	// api.GET("/users/fetch", authMiddleware(authService, usesService), userHandler.FetchUser)

	router.Run()
}

func authMiddleware(authService auth.Service, userService user.Service) gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")

		if !strings.Contains(authHeader, "Bearer") {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		tokenString := ""
		arrayToken := strings.Split(authHeader, " ")
		if len(arrayToken) == 2 {
			tokenString = arrayToken[1]
		}

		token, err := authService.ValidateToken(tokenString)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}
		fmt.Println(token)
		claim, ok := token.Claims.(jwt.MapClaims)
		fmt.Println(claim)
		if !ok || !token.Valid {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		userID := int(claim["user_id"].(float64))

		user, err := userService.GetUserByID(userID)
		if err != nil {
			response := helper.APIResponse("Unauthorized", http.StatusUnauthorized, "error", nil)
			c.AbortWithStatusJSON(http.StatusUnauthorized, response)
			return
		}

		c.Set("currentUser", user)
	}
}
