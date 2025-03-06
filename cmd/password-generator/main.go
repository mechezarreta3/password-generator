package main

import (
	"flag"
	"fmt"
	"github.com/mechezarreta3/password-generator/pkg/password"
)

func main() {
	// Define command-line flags.
	// -l for password length (default: 20)
	// -d to include at least one digit (default: false)
	// -s to include at least one special character (default: false)
	lengthPtr := flag.Int("l", 20, "Password length")
	includeDigits := flag.Bool("d", false, "Include at least one digit")
	includeSpecial := flag.Bool("s", false, "Include at least one special character")
	flag.Parse()

	// Generate the password with the provided flags.
	pwd, err := password.Generate(*lengthPtr, *includeDigits, *includeSpecial)
	if err != nil {
		fmt.Println("Error generating password:", err)
		return
	}
	fmt.Println("Generated Password:", pwd)
}
