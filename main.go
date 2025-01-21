package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/joho/godotenv"
)

func performTask() {
	fmt.Println("Task performed at:", time.Now())

	file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	if _, err := file.WriteString("New text to append\n"); err != nil {
		log.Fatal(err)
	}

	cmd := exec.Command("git", "add", ".")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("git", "commit", "-m", "hi")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}

	cmd = exec.Command("git", "push", "origin", "main")
	if err := cmd.Run(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	entries, err := os.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	gitExists := false
	for _, e := range entries {
		if e.Name() == ".git" {
			gitExists = true
			break
		}
	}

	if !gitExists {
		cmd := exec.Command("git", "init")
		if err := cmd.Run(); err != nil {
			fmt.Println("Error initializing git repository")
			log.Fatal(err)
		}

		pat := os.Getenv("GITHUB_PERSONAL_ACCESS_TOKEN")
		if pat == "" {
			log.Fatal("GitHub Personal Access Token not found in environment variables")
		}

		remoteURL := "https://github.com/adityakakarla/auto-commit.git"
		cmd = exec.Command("git", "remote", "add", "origin", remoteURL)
		if err := cmd.Run(); err != nil {
			fmt.Println("Error adding remote origin")
			log.Fatal(err)
		}

		cmd = exec.Command("git", "branch", "-M", "main")
		if err := cmd.Run(); err != nil {
			fmt.Println("Error renaming branch to main")
			log.Fatal(err)
		}

		cmd = exec.Command("git", "add", ".")
		if err := cmd.Run(); err != nil {
			fmt.Println("add")
			log.Fatal(err)
		}

		cmd = exec.Command("git", "commit", "-m", "'hi'")
		if err := cmd.Run(); err != nil {
			fmt.Println("commit")
			log.Fatal(err)
		}

		cmd = exec.Command("git", "push", "-u", "origin", "main", "--force")
		if err := cmd.Run(); err != nil {
			fmt.Println("push")
			log.Fatal(err)
		}
	}

	for {
		now := time.Now()
		nextRun := time.Date(now.Year(), now.Month(), now.Day(), now.Hour(), now.Minute(), now.Second()+5, now.Nanosecond(), now.Location())
		duration := time.Until(nextRun)
		fmt.Println("Next run in:", duration)
		time.Sleep(duration)
		performTask()
	}
}
