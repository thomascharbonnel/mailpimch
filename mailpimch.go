package main

import (
//  "gopkg.in/gomail.v2"
  "bufio"
  "os"
  "fmt"
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

// Send the emails
func SendEmails(email_addresses []string) {
}

// Parse CLI args and execute SendEmails
func main() {
}
