package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/andrewvota/piper"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	cfgFile   string
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
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.piper.yaml)")
	rootCmd.PersistentFlags().StringVar(&token, "token", "", "Discord bot token")
	rootCmd.PersistentFlags().StringVar(&channelID, "channelID", "", "Discord channel ID")

	_ = viper.BindPFlag("token", rootCmd.PersistentFlags().Lookup("token"))
	_ = viper.BindPFlag("channelID", rootCmd.PersistentFlags().Lookup("channelID"))
}

func initConfig() {
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		home, err := os.UserHomeDir()
		cobra.CheckErr(err)
		viper.AddConfigPath(home)
		viper.SetConfigName(".piper")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}

func run(cmd *cobra.Command, args []string) error {
	token := viper.GetString("token")
	channelID := viper.GetString("channelID")

	if token == "" || channelID == "" {
		return fmt.Errorf("token and channelID must be set either as flags or in the config file")
	}

	pipe, err := piper.NewPipe(token, channelID)
	if err != nil {
		return fmt.Errorf("error creating pipe: %w", err)
	}
	defer func() {
		if err := pipe.Stop(); err != nil {
			fmt.Println("Error stopping pipe:", err)
		}
	}()

	return runFromStdin(pipe)
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
