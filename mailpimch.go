package main

import (
  "gopkg.in/gomail.v2"
  "bufio"
  "os"
  "fmt"
  "flag"
)

// Helper function to take care of errors
func handle_errors(err error) {
  if err != nil {
    panic(err)
  }
}

// Read the email addresses from a file.
// Each line should only contain one email address
func ReadEmailAddresses(filename string) []string {
  email_addresses := make([]string, 0)

  // Open (and close) the file
  file, err := os.Open(filename)
  handle_errors(err)
  defer file.Close()

  // Read the file
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    email_addresses = append(email_addresses, scanner.Text())
  }
  handle_errors(scanner.Err())

  return email_addresses
}

// Generate the message
func GenerateMessage(email_address string) string {
  return fmt.Sprintf("Hello %s!", email_address) +
    "I hope you are well." +
    "J'espère que vous allez bien." +
    "Mi esperas, ke vi estas bone."
}

// Send the emails.
// We assume an already configured SMTP server running on localhost port 587.
func SendEmails(email_addresses []string) {
  dialer := gomail.Dialer{Host: "localhost", Port: 587}
  message := gomail.NewMessage()

  message.SetHeader("From", "admin@surematics.com")
  message.SetHeader("Subject", "Hello")

  for _, email_address := range email_addresses {
    message.SetHeader("To", email_address)
    message.SetBody("text/plain", GenerateMessage(email_address))

    handle_errors(dialer.DialAndSend(message))
  }
}

// Parse CLI args and execute SendEmails
func main() {
  input_filename := flag.String("input", "filename", "The file containing the email addresses")
  flag.Parse()

  SendEmails(ReadEmailAddresses(*input_filename))
}
