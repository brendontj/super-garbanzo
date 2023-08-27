package render

import (
	"encoding/json"
	"fmt"
	"github.com/jedib0t/go-pretty/v6/table"
)

func ScoreTableRender(rows [][]interface{}, scoreFromSpecificMatch bool, gameIdx int) {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Player name", "Player Score"})
	for _, r := range rows {
		t.AppendRow(r)
	}

	if scoreFromSpecificMatch {
		t.SetCaption(fmt.Sprintf("Player ranking - match %d.\n", gameIdx))
	} else {
		t.SetCaption("Player ranking - all matches.\n")
	}

	t.SetAutoIndex(true)
	fmt.Println(t.Render())
}

func DeathCausesTableRender(rows [][]interface{}, deathCausesFromSpecificMatch bool, gameIdx int) {
	t := table.NewWriter()
	t.AppendHeader(table.Row{"Death Cause", "Times"})
	for _, r := range rows {
		t.AppendRow(r)
	}

	if deathCausesFromSpecificMatch {
		t.SetCaption(fmt.Sprintf("Death causes - match %d.\n", gameIdx))
	} else {
		t.SetCaption("Death causes - all matches.\n")
	}

	t.SetAutoIndex(true)
	fmt.Println(t.Render())
}

func JSONRender(data any) {
	b, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(b))
}
