package main

import (
	"adventOfCode2023/day"
	"context"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	args := os.Args
	dayNumber := "1"
	
	if len(args) >= 2 {
		dayNumber = args[1]
	}

	input, err := fetchDayInputWithCaching(dayNumber, len(args) == 3)
	if err != nil {
		fmt.Println("Error fetching input:", err)
		return
	}

	anwser := day.Day2_2(string(input))
	fmt.Println("Answer:", anwser)
}

func fetchDayInputWithCaching(day string, forceRefetch bool) ([]byte, error) {
	if _, err := os.Stat(fmt.Sprintf("input/input_day%s.txt", day)); !os.IsNotExist(err) && !forceRefetch {
		fmt.Println("Input already exists, skipping fetch")
		return os.ReadFile(fmt.Sprintf("input/input_day%s.txt", day))
	}

	fmt.Println("Fetching input for day", day)

	url := fmt.Sprintf("https://adventofcode.com/2023/day/%s/input", day)
	godotenv.Load()

	// Replace "your_session_cookie" with your actual session cookie value
	cookie := os.Getenv("aoc_session_token")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	req, _ := http.NewRequestWithContext(ctx, "GET", url, nil)
	sessionCookie := http.Cookie{
		Name:  "session",
		Value: cookie,
	}
	req.AddCookie(&sessionCookie)
	res, _ := http.DefaultClient.Do(req)

	body, _ := io.ReadAll(res.Body)
	os.WriteFile(fmt.Sprintf("input/input_day%s.txt", day), body, os.FileMode(0644))

	return os.ReadFile(fmt.Sprintf("input/input_day%s.txt", day))
}
