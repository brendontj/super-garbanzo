package records

import (
	"fmt"
	"github.com/brendontj/super-ganbanzo/core/domain/entry"
	"github.com/brendontj/super-ganbanzo/core/domain/game"
	"strings"
)

type GameRecords struct {
	games             map[int]*game.Game
	currentGameNumber int
}

func NewGameRecords() GameRecords {
	return GameRecords{
		make(map[int]*game.Game, 0),
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
	g := game.NewGame(gr.currentGameNumber)
	gr.games[gr.currentGameNumber] = &g
}

func (gr *GameRecords) RegisterKillEntry(killer, death, reason string) {
	g := gr.games[gr.currentGameNumber]
	g.RegisterKill(killer, death, reason)
}

func (gr *GameRecords) GetGame(gameId int) (game.Game, error) {
	g, exist := gr.games[gameId]
	if !exist {
		return game.Game{}, fmt.Errorf("game not found")
	}
	return *g, nil
}

func (gr *GameRecords) GetGames() []game.Game {
	var games []game.Game
	for _, v := range gr.games {
		games = append(games, *v)
	}
	return games
}

func (gr *GameRecords) GetDeathCauses() map[string]int {
	dc := make(map[string]int, 0)
	for _, v := range gr.games {
		for k, vv := range v.DeathCauseRecord {
			_, exist := dc[k]
			if !exist {
				dc[k] = vv
				continue
			}
			dc[k] += vv
		}
	}
	return dc
}

func (gr *GameRecords) GetGeneralScore() map[string]int {
	gs := make(map[string]int, 0)
	for _, v := range gr.games {
		for k, vv := range v.Kills {
			_, exist := gs[k]
			if !exist {
				gs[k] = vv
				continue
			}
			gs[k] += vv
		}
	}
	return gs
}
