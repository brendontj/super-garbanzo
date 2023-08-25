package records

import (
	"github.com/brendontj/super-ganbanzo/core/domain/entry"
	"github.com/brendontj/super-ganbanzo/core/domain/game"
	"regexp"
	"strings"
)

type GameRecords struct {
	games             map[int]game.Game
	currentGameNumber int
}

func NewGameRecords() GameRecords {
	return GameRecords{
		make(map[int]game.Game, 0),
		0,
	}
}

func (gr *GameRecords) ParseRecord(logLine string) {
	dataSlice := strings.Split(strings.TrimSpace(logLine), " ")
	logLineMightBeParsed := len(dataSlice) > 1
	if logLineMightBeParsed {
		entryTypeString := strings.TrimSuffix(dataSlice[1], ":")
		entryType, err := entry.FromString(entryTypeString)
		if err != nil {
			//ignore: entry type not implemented
			return
		}
		switch entryType {
		case entry.TypeGameInit:
			gr.StartNewGame()
		case entry.TypeKill:
			pattern := `:([^:]+)\s+killed`
			re := regexp.MustCompile(pattern)
			match := re.FindStringSubmatch(logLine)
			killer := strings.TrimSpace(match[1])

			pattern = `killed\s(.*?)\s+by`
			re = regexp.MustCompile(pattern)
			match = re.FindStringSubmatch(logLine)
			death := strings.TrimSpace(match[1])

			pattern = `\s([^\s]+)$`
			re = regexp.MustCompile(pattern)
			match = re.FindStringSubmatch(logLine)
			reason := strings.TrimSpace(match[1])

			gr.RegisterKillEntry(killer, death, reason)
		case entry.TypeGameFinished:
			//panic()
		}
	}
}

func (gr *GameRecords) StartNewGame() {
	gr.currentGameNumber += 1
	gr.games[gr.currentGameNumber] = game.NewGame(gr.currentGameNumber)
}

func (gr *GameRecords) RegisterKillEntry(killer, death, reason string) {
	game := gr.games[gr.currentGameNumber]
	game.RegisterKill(killer, death, reason)
}
