package game

const PlayerWorldIdentifier = "<world>"

type Game struct {
	ID               int
	TotalKills       int
	Kills            map[string]int
	DeathCauseRecord map[string]int
	WasClosed        bool
}

func NewGame(gameIdentifier int) Game {
	return Game{
		ID:               gameIdentifier,
		TotalKills:       0,
		Kills:            make(map[string]int, 0),
		DeathCauseRecord: make(map[string]int, 0),
	}
}

func (g *Game) RegisterKill(killer, death, reason string) {
	g.TotalKills += 1
	_, exist := g.DeathCauseRecord[reason]
	if !exist {
		g.DeathCauseRecord[reason] = 1
		return
	}
	g.DeathCauseRecord[reason] += 1

	if killer == PlayerWorldIdentifier {
		g.removeKillFromPlayer(death)
		return
	}

	g.addKillToPlayer(killer)
}

func (g *Game) Players() []string {
	playerNames := make([]string, 0, len(g.Kills))
	for k := range g.Kills {
		playerNames = append(playerNames, k)
	}
	return playerNames
}

func (g *Game) addKillToPlayer(playerName string) {
	_, exist := g.Kills[playerName]
	if !exist {
		g.Kills[playerName] = 1
		return
	}

	g.Kills[playerName] += 1
}

func (g *Game) removeKillFromPlayer(playerName string) {
	_, exist := g.Kills[playerName]
	if !exist {
		g.Kills[playerName] = -1
		return
	}

	g.Kills[playerName] -= 1
}
