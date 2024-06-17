package piper

import (
	"log"
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var (
	token     string
	channelID string
)

func TestMain(m *testing.M) {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatalf("Error loading environment variables: %s", err)
	}

	token = os.Getenv("DISCORD_TOKEN")
	channelID = os.Getenv("DISCORD_CHANNEL_ID")

	// Run the tests
	code := m.Run()

	// Exit with the code returned by m.Run
	os.Exit(code)
}

func TestNewPipe(t *testing.T) {
	_, err := NewPipe(token, channelID)
	if err != nil {
		t.Errorf("%v", err)
	}
}

func TestStartStop(t *testing.T) {
	pipe, err := NewPipe(token, channelID)
	if err != nil {
		t.Fatalf("%v", err)
	}

	err = pipe.Start()
	if err != nil {
		t.Fatalf("%v", err)
	}

	err = pipe.Stop()
	if err != nil {
		t.Fatalf("%v", err)
	}
}
