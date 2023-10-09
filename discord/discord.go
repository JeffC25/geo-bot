package discord

import (
	"main/config"
	"os"
	"os/signal"

	"github.com/bwmarrin/discordgo"
	"github.com/rs/zerolog"
)

func Run(c config.Config, log zerolog.Logger) {
	session, err := discordgo.New("Bot " + c.DiscordToken)
	if err != nil {
		log.Panic().Err(err).Msg("failed to create disgo client: %v")
	}
	session.Open()
	defer session.Close()

	log.Info().Msg("Starting new bot session")

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)
	<-ch
}
