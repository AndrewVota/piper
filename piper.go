package piper

import (
	"fmt"
	"io"
	"os"

	"github.com/bwmarrin/discordgo"
)

type Pipe struct {
	Token          string
	ChannelID      string
	Discord        *discordgo.Session
	originalStdout *os.File
	pipeReader     *os.File
	pipeWriter     *os.File
	outputChannel  chan string
	done           chan bool
}

func NewPipe(token string, channelID string) (*Pipe, error) {
	discord, err := discordgo.New("Bot " + token)
	if err != nil {
		return nil, fmt.Errorf("error creating Discord session: %w", err)
	}

	err = discord.Open()
	if err != nil {
		return nil, fmt.Errorf("error creating Discord connection: %w", err)
	}

	return &Pipe{
		Token:          token,
		ChannelID:      channelID,
		Discord:        discord,
		originalStdout: os.Stdout,
		outputChannel:  make(chan string, 1),
		done:           make(chan bool),
	}, nil
}

func (c *Pipe) Start() error {
	var err error
	c.pipeReader, c.pipeWriter, err = os.Pipe()
	if err != nil {
		return fmt.Errorf("error returning the os pipe: %w", err)
	}

	os.Stdout = c.pipeWriter

	go func() {
		var buf []byte = make([]byte, 1024)
		for {
			n, err := c.pipeReader.Read(buf)
			if err != nil {
				if err != io.EOF {
					c.outputChannel <- ""
				}
				break
			}
			if n > 0 {
				output := buf[:n]
				_, _ = c.originalStdout.Write(output)
				c.outputChannel <- string(output)
				_, err := c.Discord.ChannelMessageSend(c.ChannelID, string(output))
				if err != nil {
					fmt.Printf("error sending message to Discord: %s", err)
				}
			}
		}
		c.done <- true
	}()

	return nil
}

func (c *Pipe) Stop() error {
	err := c.pipeWriter.Close()
	if err != nil {
		return fmt.Errorf("error closing the pipe writer: %w", err)
	}

	<-c.done

	os.Stdout = c.originalStdout
	return nil
}
