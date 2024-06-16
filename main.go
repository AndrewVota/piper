package main

import (
	"fmt"
	"log"
	"os"

	"github.com/andrewvota/piper/piper"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading environment variables")
	}

	var (
		token     = os.Getenv("DISCORD_TOKEN")
		channelID = os.Getenv("DISCORD_CHANNEL_ID")
	)

	piper, err := piper.NewPipe(token, channelID)
	if err != nil {
		fmt.Println("Error creating new pipe:", err)
	}

	err = piper.Start()
	if err != nil {
		fmt.Println("Error starting piper:", err)
		return
	}
	defer piper.Stop()

	fmt.Println("This is a test output.")
	fmt.Println("This is another test output.")

	for i := range 10 {
		fmt.Printf("CURRENT INDEX %d\n", i)
	}
}
