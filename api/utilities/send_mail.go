package utilities

import (
	"fmt"
	"log"
	"net/mail"
	"net/smtp"
	"os"

	"github.com/joho/godotenv"
)

func SendMail(name,email,message string) error {
	
	receiver := "karanidancan20@gmail.com"
	from := os.Getenv("EMAIL")
	if from == ""{
		err := godotenv.Load(".env")
		if err != nil {
			panic(err.Error())
		}
	}
	password := os.Getenv("SMTP_PASSWORD")
	smtpHost := "smtp.gmail.com"
	smtpPort := "587"
	// Compose the email
	subject := "Contact us response"
	// Format the times as strings


	log.Println(from+","+password+","+email)
	// Create the body message
	body := fmt.Sprintf(
		message,
	)
	msg := []byte("Subject: " + subject + "\r\n" +
		"To: " + email + "\r\n" +
		"\r\n" +
		body)

	// Create the "from" address
	// fromAddr := mail.Address{Name: "Dancan", Address: from}
	fromAddr := mail.Address{Address: from}
	// Establish a connection to the SMTP server
	auth := smtp.PlainAuth("", from, password, smtpHost)
	err := smtp.SendMail(smtpHost+":"+smtpPort, auth, fromAddr.Address, []string{receiver}, msg)
	if err != nil {
		return err
	}
	return nil
}