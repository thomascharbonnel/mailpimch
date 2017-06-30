package main

import (
  "testing"
  "fmt"
)

func TestReadEmailAddresses(t *testing.T) {
  // The format for the email file is one email address per line, and only that.

  fmt.Println(ReadEmailAddresses("./test/email_addresses.txt"))
  // Output:
  // patrick.bateman@surematics.com
  // bruce.wayne@surematics.com
}

func TestGenerateMessage(t *testing.T) {
}

// SendEmails should be tested by catching the outgoing SMTP requests.
