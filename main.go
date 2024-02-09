package main

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

var connections = make(map[string]*discordgo.VoiceConnection)

func main() {
	// read file called token
	token, err := os.ReadFile("token")
	if err != nil {
		panic(err)
	}

	if token == nil {
		panic("Token not found")
	}

	discord, err := discordgo.New("Bot " + string(token))
	if err != nil {
		panic(err)
	}

	defer discord.Close()

	err = discord.Open()
	if err != nil {
		panic(err)
	}

	guilds, err := discord.UserGuilds(100, "", "")
	if err != nil {
		panic(err)
	}

	for {
		for _, guild := range guilds {
			channels, err := discord.GuildChannels(guild.ID)
			if err != nil {
				fmt.Println(err)
				continue
			}

			for _, channel := range channels {
				if channel.Type == discordgo.ChannelTypeGuildVoice && strings.ToLower(channel.Name) == "afghanistan" {
					connections[guild.Name], err = discord.ChannelVoiceJoin(guild.ID, channel.ID, false, false)
					if err != nil {
						fmt.Println(err)
					}
				}
			}
		}

		if len(connections) == 0 {
			fmt.Println("No channels found")
			return
		}

		occupying := "Occupying: "
		for guildName, connection := range connections {
			if connection.Ready {
				occupying += guildName + ", "
			}
		}
		fmt.Println(occupying)
		time.Sleep(10 * time.Second)
	}
}
