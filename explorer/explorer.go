package explorer

import (
	"crypto/sha256"
	"fmt"
	fiber "github.com/gofiber/fiber/v2"
	"log"
	"main/core/sendtx"
)

func Explorer() {
	app := fiber.New()
	app.Post("/chainservice/uploadchain", uploadchain)
	log.Fatal(app.Listen(":3005"))
}

type UserInfo struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Name        string `json:"name"`
	Phone       string `json:"phone"`
	Description string `json:"description"`
}
type ErrorResponse struct {
	Error   string `json:"error"`
	Success bool   `json:"success"`
	Data    string `json:"data"`
}

func sha356func(s string) string {
	var sha = sha256.New()
	sha.Write([]byte(s))
	var encrypted = sha.Sum(nil)
	var encryptedString = fmt.Sprintf("%x", encrypted)
	return encryptedString
}

func uploadchain(c *fiber.Ctx) error {
	fmt.Println(string(c.Body()))
	payload := &UserInfo{}
	if err := c.BodyParser(payload); err != nil {
		fmt.Println("Parse error")
		return c.Status(400).JSON(ErrorResponse{
			Error:   err.Error(),
			Success: false,
			Data:    "",
		})
	}
	hash := sendtx.AddOrUpdateUserData(payload.Username, sha356func(payload.Password), payload.Name, payload.Phone, payload.Description)
	fmt.Println(hash)
	return c.Status(200).JSON(ErrorResponse{Error: "", Success: true, Data: hash})
}
