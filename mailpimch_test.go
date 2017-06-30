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
  fmt.Println(GenerateMessage("trevor.reznik@surematics.com"))
  // Output:
  // Hello trevor.reznik@surematics.com!
  // I hope you are well.
  // J'espère que vous allez bien.
  // Mi esperas, ke vi estas bone.
}

// SendEmails should be tested by catching the outgoing SMTP requests.
