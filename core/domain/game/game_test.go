package game

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRegisterKill(t *testing.T) {
	newGame := NewGame(1)

	newGame.RegisterKill("killer", "killed", "MOD_SHOTGUN")

	assert.Equal(t, 1, newGame.TotalKills)
	assert.Equal(t, 1, newGame.Kills["killer"])
	assert.Equal(t, 1, newGame.DeathCauseRecord["MOD_SHOTGUN"])
}

func TestRegisterKill_WhenThereAreMultipleKills(t *testing.T) {
	g := NewGame(1)

	g.RegisterKill("person-1", "person-2", "MOD_SHOTGUN")
	g.RegisterKill("person-0", "person-1", "MOD_TST")
	g.RegisterKill("person-0", "person-2", "MOD_TST2")
	g.RegisterKill("person-2", "person-1", "MOD_SHOTGUN")
	g.RegisterKill(PlayerWorldIdentifier, "person-1", "MOD_FALLING")

	assert.Equal(t, 5, g.TotalKills)
	assert.Equal(t, 2, g.Kills["person-0"])
	assert.Equal(t, 0, g.Kills["person-1"])
	assert.Equal(t, 1, g.Kills["person-2"])
	assert.Equal(t, 2, g.DeathCauseRecord["MOD_SHOTGUN"])
	assert.Equal(t, 1, g.DeathCauseRecord["MOD_TST"])
	assert.Equal(t, 1, g.DeathCauseRecord["MOD_TST2"])
	assert.Equal(t, 1, g.DeathCauseRecord["MOD_FALLING"])
}
