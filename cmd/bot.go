package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/riceandbeas/nekoya/internal/discord"
	"github.com/spf13/cobra"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Start the bot",
	RunE: func(cmd *cobra.Command, args []string) error {
		if profile != "production" {
			err := godotenv.Load()
			if err != nil {
				return err
			}
		}

		log.Println("Starting bot...")
		bot, err := discord.NewBot(os.Getenv("DISCORD_TOKEN"))
		if err != nil {
			return fmt.Errorf("Error creating bot: %w", err)
		}

		return bot.Run()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
