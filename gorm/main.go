package main

import (
	"fmt"
	"gorm/database"
	"gorm/repository"
	"strings"
)

func main() {
	db := database.StartDB()
	userRepo := repository.NewUserRepo(db)

	// user := models.User{
	// 	Email: "arsyad@mail.com",
	// }

	// err := userRepo.CreateUser(&user)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }
	// fmt.Println("Created Success!")

	employees, err := userRepo.GetAllUsers()
	if err != nil {
		fmt.Println("error :", err.Error())
		return
	}

	for i, emp := range *employees {
		fmt.Println("User :", i+1)
		emp.Print()
		fmt.Println(strings.Repeat("=", 10))
	}

	// Get Employee By ID
	// emp, err := userRepo.GetUserByID(2)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }
	// emp.Print()

	// Update User By ID
	// newUser := models.User{
	// 	Email: "testupdate@mail.com",
	// }

	// err = userRepo.UpdateUser(4, &newUser)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }
	// fmt.Println("Update user success!")

	// Delete User
	// err = userRepo.DeleteUser(4)
	// if err != nil {
	// 	fmt.Println("error :", err.Error())
	// 	return
	// }
	// fmt.Println("Delete user success!")
}
