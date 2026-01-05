package main

import (
	"errors"
	"fmt"
)

func login(username string, password string) bool {
	if username == "admin" && password == "123456" {
		return true
	}
	return false
}

func register(username string, password string) (int, error) {
	if username == "" || password == "" {
		return 0, errors.New("username or password cannot be empty")
	}
	return 1, nil
}

func userCenter(userID int) string {
	if userID == 1 {
		return "Welcome to the user center, admin!"
	}
	return "Welcome to the user center, guest!"
}

func main() {
	// 测试登录功能
	fmt.Println("please enter your control")
	fmt.Println("1. Login")
	fmt.Println("2. Register")
	fmt.Println("3. User Center")

	var index int
	fmt.Scan(&index)

	var funMap = map[int]func(){
		1: func() {
			var username, password string
			fmt.Print("Enter username: ")
			fmt.Scan(&username)
			fmt.Print("Enter password: ")
			fmt.Scan(&password)
			if login(username, password) {
				fmt.Println("Login successful!")
			} else {
				fmt.Println("Login failed!")
			}
		},
		2: func() {
			var username, password string
			fmt.Print("Enter username: ")
			fmt.Scan(&username)
			fmt.Print("Enter password: ")
			fmt.Scan(&password)
			userID, err := register(username, password)
			if err != nil {
				fmt.Println("Registration failed:", err)
			} else {
				fmt.Println("Registration successful! User ID:", userID)
			}
		},
		3: func() {
			var userID int
			fmt.Print("Enter user ID: ")
			fmt.Scan(&userID)
			message := userCenter(userID)
			fmt.Println(message)
		},
	}
	fun, ok := funMap[index]
	if ok {
		fun()
	} else {
		fmt.Println("Invalid choice")
	}

	// switch index {
	// case 1:
	// 	var username, password string
	// 	fmt.Print("Enter username: ")
	// 	fmt.Scan(&username)
	// 	fmt.Print("Enter password: ")
	// 	fmt.Scan(&password)
	// 	if login(username, password) {
	// 		fmt.Println("Login successful!")
	// 	} else {
	// 		fmt.Println("Login failed!")
	// 	}
	// case 2:
	// 	var username, password string
	// 	fmt.Print("Enter username: ")
	// 	fmt.Scan(&username)
	// 	fmt.Print("Enter password: ")
	// 	fmt.Scan(&password)
	// 	userID, err := register(username, password)
	// 	if err != nil {
	// 		fmt.Println("Registration failed:", err)
	// 	} else {
	// 		fmt.Println("Registration successful! User ID:", userID)
	// 	}
	// case 3:
	// 	var userID int
	// 	fmt.Print("Enter user ID: ")
	// 	fmt.Scan(&userID)
	// 	message := userCenter(userID)
	// 	fmt.Println(message)
	// default:
	// 	fmt.Println("Invalid choice")
	// }
}
