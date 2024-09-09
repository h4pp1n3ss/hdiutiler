package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// saveOutput writes the output to the specified file
func SaveOutput(filePath string, output []byte) {
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Printf("Error creating output file: %s\n", err)
		os.Exit(1)
	}
	defer file.Close()

	_, err = file.Write(output)
	if err != nil {
		fmt.Printf("Error writing to output file: %s\n", err)
		os.Exit(1)
	}

	fmt.Printf("[!] Output saved to %s\n", filePath)

}

func ContainsWord(input string, word string) bool {
	// Convert input to uppercase to perform a case-insensitive comparison
	upperInput := strings.ToUpper(input)
	// Check if the word "VALID" is present in the input
	return strings.Contains(upperInput, word)
}

func SearchStringInFile(searchString string, fileName string) {

	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("[!] Error opening file: ", err)
		return
	}

	defer file.Close()

	// Create a scanner to read the file line by line
	scanner := bufio.NewScanner(file)

	// Iterate over each line in the file
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, searchString) {
			fmt.Println("[!] Check:", line)
			return
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
	} else {
		fmt.Println("String not found")
	}
}
