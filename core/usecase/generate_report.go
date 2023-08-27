package usecase

import (
	"github.com/brendontj/super-ganbanzo/core/domain/records"
	"github.com/brendontj/super-ganbanzo/internal/render"
	"github.com/brendontj/super-ganbanzo/pkg/sorter"
	"strconv"
)

type (
	ReportGenerator interface {
		GenerateMatchReport(gameIdx int) error
		GenerateAllMatchesReport()
		GenerateScoreMatchReport(gameIdx int) error
		GenerateScoreReport()
		GenerateDeathCausesMatchReport(gameIdx int) error
		GenerateDeathCausesReport()
	}

	reportGenerator struct {
		gameRecords records.GameRecords
	}

	matchPresenter struct {
		TotalKills int            `json:"total_kills"`
		Players    []string       `json:"players"`
		Kills      map[string]int `json:"kills"`
	}
)

func NewReportGenerator(gameRecords records.GameRecords) ReportGenerator {
	return reportGenerator{gameRecords: gameRecords}
}

func (rg reportGenerator) GenerateMatchReport(gameIdx int) error {
	game, err := rg.gameRecords.GetGame(gameIdx)
	if err != nil {
		return err
	}
	matchReport := make(map[string]matchPresenter)
	matchReport["game_"+strconv.FormatInt(int64(game.ID), 10)] = matchPresenter{
		TotalKills: game.TotalKills,
		Players:    game.Players(),
		Kills:      game.Kills,
	}

	render.JSONRender(matchReport)
	return nil
}

func (rg reportGenerator) GenerateAllMatchesReport() {
	games := rg.gameRecords.GetGames()
	matchReport := make(map[string]matchPresenter)
	for _, g := range games {
		matchReport["game_"+strconv.FormatInt(int64(g.ID), 10)] = matchPresenter{
			TotalKills: g.TotalKills,
			Players:    g.Players(),
			Kills:      g.Kills,
		}
	}

	render.JSONRender(matchReport)
}

func (rg reportGenerator) GenerateScoreMatchReport(gameIdx int) error {
	game, err := rg.gameRecords.GetGame(gameIdx)
	if err != nil {
		return err
	}

	render.ScoreTableRender(
		sorter.SortStringToIntMapIntoASlice(game.Kills),
		true,
		gameIdx)
	return nil
}

func (rg reportGenerator) GenerateScoreReport() {
	render.ScoreTableRender(
		sorter.SortStringToIntMapIntoASlice(rg.gameRecords.GetGeneralScore()),
		false,
		0)
}

func (rg reportGenerator) GenerateDeathCausesMatchReport(gameIdx int) error {
	game, err := rg.gameRecords.GetGame(gameIdx)
	if err != nil {
		return err
	}
	render.DeathCausesTableRender(
		sorter.SortStringToIntMapIntoASlice(game.DeathCauseRecord),
		true,
		gameIdx)
	return nil
}

func (rg reportGenerator) GenerateDeathCausesReport() {
	render.DeathCausesTableRender(
		sorter.SortStringToIntMapIntoASlice(rg.gameRecords.GetDeathCauses()),
		false,
		0)
}
