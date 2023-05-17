package getpass

import (
	"fmt"
	"golang.org/x/crypto/ssh/terminal"
	"os"
)

func GetPassword(prompt string) (string, error) {
	fmt.Print(prompt)

	// Disable input echoing
	bytePassword, err := terminal.ReadPassword(int(os.Stdin.Fd()))
	if err != nil {
		return "", err
	}

	// Convert the password to a string
	password := string(bytePassword)

	// Add a new line after the input
	fmt.Println()

	return password, nil
}

// func main() {
// 	password, err := getPassword("Enter your password: ")
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return
// 	}

// 	fmt.Println("Password entered:", password)
// }
