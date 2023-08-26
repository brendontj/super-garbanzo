package entry

import (
	"errors"
)

const (
	TypeGameInit = Type("game-start")
	TypeKill     = Type("kill")
)

type Type string

func FromString(et string) (Type, error) {
	switch et {
	case "InitGame":
		return TypeGameInit, nil
	case "Kill":
		return TypeKill, nil
	default:
		return "", errors.New("unable to detect entry type")
	}
}
