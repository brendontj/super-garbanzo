package records

import (
	"github.com/brendontj/super-ganbanzo/core/domain/entry"
	"github.com/brendontj/super-ganbanzo/core/domain/game"
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
			kill := entry.ParseKill(logLine)
			gr.RegisterKillEntry(kill.Killer, kill.Killed, kill.Reason)
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
