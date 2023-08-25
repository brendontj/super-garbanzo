package entry

import "errors"

const (
	TypeGameInit     = Type("game-start")
	TypeKill         = Type("kill")
	TypeGameFinished = Type("game-finished")
)

type Type string

func FromString(et string) (Type, error) {
	switch et {
	case "InitGame":
		return TypeGameInit, nil
	case "Kill":
		return TypeKill, nil
	case "ShutdownGame":
		return TypeGameFinished, nil
	default:
		return "", errors.New("")
	}
}
