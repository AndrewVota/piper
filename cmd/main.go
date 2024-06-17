package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/andrewvota/piper/piper"
	"github.com/spf13/cobra"
)

var (
	token     string
	channelID string
	useStdin  bool
	rootCmd   = &cobra.Command{
		Use:   "piper",
		Short: "Piper is a tool to pipe stdout or stdin to Discord",
		RunE:  run,
	}
)

func init() {
	rootCmd.PersistentFlags().StringVar(&token, "token", "", "Discord bot token")
	rootCmd.PersistentFlags().StringVar(&channelID, "channelID", "", "Discord channel ID")
	rootCmd.PersistentFlags().BoolVar(&useStdin, "stdin", true, "Read from stdin instead of stdout")
	rootCmd.MarkPersistentFlagRequired("token")
	rootCmd.MarkPersistentFlagRequired("channelID")
}

func run(cmd *cobra.Command, args []string) error {
	pipe, err := piper.NewPipe(token, channelID)
	if err != nil {
		return fmt.Errorf("error creating pipe: %w", err)
	}
	defer pipe.Stop()

	if useStdin {
		return runFromStdin(pipe)
	}

	err = pipe.Start()
	if err != nil {
		return fmt.Errorf("error starting pipe: %w", err)
	}

	return nil
}

func runFromStdin(pipe *piper.Pipe) error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		_, err := pipe.Discord.ChannelMessageSend(pipe.ChannelID, line)
		if err != nil {
			fmt.Printf("Error sending message to Discord: %s\n", err)
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading from stdin: %w", err)
	}

	return nil
}

func Execute() error {
	return rootCmd.Execute()
}

func main() {
	if err := Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
