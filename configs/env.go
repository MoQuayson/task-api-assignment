package configs

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func LoadEnvFile() {
	file, err := os.Open(".env")
	if err != nil {
		fmt.Println("Error opening .env file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if equal := strings.Index(line, "="); equal >= 0 {
			key := strings.TrimSpace(line[:equal])
			value := strings.TrimSpace(line[equal+1:])
			err := os.Setenv(key, value)
			if err != nil {
				fmt.Println("Error setting environment variable:", err)
				return
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading .env file:", err)
		return
	}

}
