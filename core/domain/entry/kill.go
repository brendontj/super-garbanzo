package entry

import (
	"regexp"
	"strings"
)

type Kill struct {
	Killer string
	Killed string
	Reason string
}

func ParseKill(killLogLine string) Kill {
	var k Kill
	killerPattern := `:([^:]+)\s+killed`
	re := regexp.MustCompile(killerPattern)
	match := re.FindStringSubmatch(killLogLine)
	k.Killer = strings.TrimSpace(match[1])

	killedPattern := `killed\s(.*?)\s+by`
	re = regexp.MustCompile(killedPattern)
	match = re.FindStringSubmatch(killLogLine)
	k.Killed = strings.TrimSpace(match[1])

	reasonPattern := `\s([^\s]+)$`
	re = regexp.MustCompile(reasonPattern)
	match = re.FindStringSubmatch(killLogLine)
	k.Reason = strings.TrimSpace(match[1])
	return k
}
