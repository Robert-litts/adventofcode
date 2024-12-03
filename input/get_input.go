package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func downloadFile(day int, cookie string) error {
	// Construct the URL for the given day
	url := fmt.Sprintf("https://adventofcode.com/2024/day/%d/input", day)

	// Create a new HTTP client with the session cookie
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return err
	}

	// Set the session cookie (assuming you have it)
	req.Header.Set("Cookie", "session="+cookie)

	// Perform the GET request
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Check if the status is OK (HTTP 200)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to download file: %s", resp.Status)
	}

	// Create the file to save the response content
	fileName := fmt.Sprintf("day%d.input", day)
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	// Copy the content of the response body into the file
	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	fmt.Printf("Day %d input downloaded successfully as %s\n", day, fileName)
	return nil
}

func main() {

	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		return
	}

	cookie := os.Getenv("SESSION_COOKIE")
	if cookie == "" {
		fmt.Println("SESSION_COOKIE not set in .env file")
		return
	}

	dayFlag := flag.Int("day", 1, "The day of the Advent of Code input to download")

	flag.Parse()

	err = downloadFile(*dayFlag, cookie)
	if err != nil {
		fmt.Println("Error:", err)
	}
}
