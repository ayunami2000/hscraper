package trollcord

import (
	"math/rand"
	"strings"

	"github.com/bwmarrin/discordgo"
)

var bots []*discordgo.Session

func LoadTokens(tokens []string) (fail int) {
	for _, v := range tokens {
		dg, err := discordgo.New("Bot " + v)

		if err != nil {
			fail++
			continue
		}

		dg.Identify.Intents = discordgo.IntentsGuildMessages

		err = dg.Open()

		if err != nil {
			fail++
			continue
		}

		bots = append(bots, dg)
	}

	return fail
}

func Send(ch string, content string) error {
	bot := bots[rand.Intn(len(bots))]
	_, err := bot.ChannelMessageSend(ch, content)
	return err
}

func Tags(guild, category string) (map[string]string, error) {
	bot := bots[0]

	chs, err := bot.GuildChannels(guild)

	if err != nil {
		return nil, err
	}

	var res = make(map[string]string)

	for _, ch := range chs {
		if ch.ParentID != category {
			continue
		}

		tpc := strings.ReplaceAll(ch.Topic, "\n", "")

		if !ch.NSFW {
			tpc = "rating:s " + tpc
		} else {
			tpc = "rating:e " + tpc
		}

		res[ch.ID] = tpc
	}

	return res, nil
}
