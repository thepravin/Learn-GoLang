package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"gopkg.in/gomail.v2"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading the .env file %v", err)
	}

	smtpHost := os.Getenv("SMTP_HOST")
	smtpPortStr := os.Getenv("SMTP_PORT")
	senderEmail := os.Getenv("SENDER_EMAIL")
	senderPassword := os.Getenv("SENDER_PASSWORD")
	senderName := os.Getenv("SENDER_NAME")
	receiverEmail := os.Getenv("RECEIVER_EMAIL")
	attrachmentPath := `C:\Users\pravinn\Downloads\Pravin Nalawade_Resume.pdf`

	smtpPort, err := strconv.Atoi(smtpPortStr)
	if err != nil {
		log.Fatalf("Invalid SMTP_PORT")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(senderEmail, senderName))
	m.SetHeader("To", receiverEmail)
	m.SetHeader("Subject", "Email without attachment")

	m.SetBody("text/html", `
		<h2>Hello</h2>
		<p>Hello From go lang </p>
	`)

	if attrachmentPath != "" {
		m.Attach(attrachmentPath)
	}

	d := gomail.NewDialer(smtpHost, smtpPort, senderEmail, senderPassword)
	if err := d.DialAndSend(m); err != nil {
		log.Fatal("Failed to send email")
	}

	fmt.Println("Email sent successfully")
}
