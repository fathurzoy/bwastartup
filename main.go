package main

import (
	"bwastartup/handler"
	"bwastartup/user"
	"log"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal(err.Error())
	}

	userRepository := user.NewRepository(db)
	userService := user.NewService(userRepository)

	// tes password
	// input := user.LoginInput{
	// 	Email: "fathur@gmail.com",
	// 	Password: "11111111",
	// }
	// user, err := userService.Login(input)
	// if err != nil{
	// 	fmt.Println("Terjadi kesalahan")
	// 	fmt.Println(err.Error())
	// }

	// fmt.Println(user.Email)
	// fmt.Println(user.Name)

	// tes find user
	// userByEmail, err := userRepository.FindByEmail("fathur@gmail.com")
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }

	// if (userByEmail.ID == 0){
	// 	fmt.Println("User tidak ditemukan")
	// }else{
	// 	fmt.Println(userByEmail.Name)
	// }

	userHandler := handler.NewUserHandler(userService)
	
	router := gin.Default()
	api := router.Group("/api/v1")

	api.POST("/users", userHandler.RegisterUser)

	router.Run()


	// --------------
	// user := user.User{
	// 	Name: "Zoyzoy",
	// }
	// userRepository.Save(user)

	// userInput := user.RegisterUserInput{}
	// userInput.Name = "Tes simpan dari service"
	// userInput.Email = "contoh@gmail.com"
	// userInput.Occupation = "bolang"
	// userInput.Password = "11111111"
	// userService.RegisterUser(userInput)


	// input dari user
	// handler, mapping input dari user -> struct input
	// service : melakukan mapping dari struct input ke struct user
	// repository
	// db

	// --------------------
	// fmt.Println("Connection to database is good")

	// var users []user.User
	// db.Find(&users)

	// for _, user := range users {
	// 	fmt.Println(user.Name)
	// 	fmt.Println(user.Email)
	// 	fmt.Println("===============")
	// }

	// router := gin.Default()
	// router.GET("/handler", handler)
	// router.Run()
}

// --------------- test func get users
// func handler(c *gin.Context){
//   dsn := "root:@tcp(127.0.0.1:3306)/bwastartup?charset=utf8mb4&parseTime=True&loc=Local"
//   db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})

// 	if err != nil {
// 		log.Fatal(err.Error())
// 	}

// 	var users []user.User
// 	db.Find(&users)

// 	c.JSON(http.StatusOK, users)

// 	// input
// 	// handler mapping input ke struct
// 	// service mapping ke struct user
// 	// repository save struct ke db
// 	// db
// }