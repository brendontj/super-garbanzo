package usecase

import (
	"github.com/brendontj/super-ganbanzo/core/domain/records"
	"github.com/brendontj/super-ganbanzo/internal/render"
	"github.com/brendontj/super-ganbanzo/pkg/sorter"
	"strconv"
)

type ReportGenerator interface {
	GenerateMatchReport(gameIdx int)
	GenerateAllMatchesReport()
	GenerateScoreMatchReport(gameIdx int)
	GenerateScoreReport()
	GenerateDeathCausesMatchReport(gameIdx int)
	GenerateDeathCausesReport()
}

type reportGenerator struct {
	gameRecords records.GameRecords
}

type MatchPresenter struct {
	TotalKills int            `json:"total_kills"`
	Players    []string       `json:"players"`
	Kills      map[string]int `json:"kills"`
}

func NewReportGenerator(gameRecords records.GameRecords) ReportGenerator {
	return reportGenerator{gameRecords: gameRecords}
}

func (rg reportGenerator) GenerateMatchReport(gameIdx int) {
	game, err := rg.gameRecords.GetGame(gameIdx)
	if err != nil {
		//Todo: Log error
		return
	}
	matchReport := make(map[string]MatchPresenter)
	matchReport["game_"+strconv.FormatInt(int64(game.ID), 10)] = MatchPresenter{
		TotalKills: game.TotalKills,
		Players:    game.Players(),
		Kills:      game.Kills,
	}

	render.JSONRender(matchReport)
}

func (rg reportGenerator) GenerateAllMatchesReport() {
	games := rg.gameRecords.GetGames()
	matchReport := make(map[string]MatchPresenter)
	for _, g := range games {
		matchReport["game_"+strconv.FormatInt(int64(g.ID), 10)] = MatchPresenter{
			TotalKills: g.TotalKills,
			Players:    g.Players(),
			Kills:      g.Kills,
		}
	}

	render.JSONRender(matchReport)
}

func (rg reportGenerator) GenerateScoreMatchReport(gameIdx int) {
	game, err := rg.gameRecords.GetGame(gameIdx)
	if err != nil {
		//Todo: Log error
		return
	}

	render.ScoreTableRender(
		sorter.SortStringToIntMapIntoASlice(game.Kills),
		true,
		gameIdx)
}

func (rg reportGenerator) GenerateScoreReport() {
	render.ScoreTableRender(
		sorter.SortStringToIntMapIntoASlice(rg.gameRecords.GetGeneralScore()),
		false,
		0)
}

func (rg reportGenerator) GenerateDeathCausesMatchReport(gameIdx int) {
	game, err := rg.gameRecords.GetGame(gameIdx)
	if err != nil {
		//Todo: Log error
		return
	}
	render.DeathCausesTableRender(
		sorter.SortStringToIntMapIntoASlice(game.DeathCauseRecord),
		true,
		gameIdx)
}

func (rg reportGenerator) GenerateDeathCausesReport() {
	render.DeathCausesTableRender(
		sorter.SortStringToIntMapIntoASlice(rg.gameRecords.GetDeathCauses()),
		false,
		0)
}
