package entry

import (
	"reflect"
	"testing"
)

func TestParseKill(t *testing.T) {
	tests := []struct {
		name    string
		logLine string
		want    Kill
	}{
		{
			name:    "should parse kill log line when the killer isn't a player",
			logLine: "20:54 Kill: 1022 2 22: <world> killed Isgalamido by MOD_TRIGGER_HURT",
			want:    Kill{Killed: "Isgalamido", Killer: "<world>", Reason: "MOD_TRIGGER_HURT"},
		},
		{
			name:    "should parse kill log line when the killer is a player",
			logLine: "22:06 Kill: 2 3 7: Isgalamido killed Mocinha by MOD_ROCKET_SPLASH",
			want:    Kill{Killed: "Mocinha", Killer: "Isgalamido", Reason: "MOD_ROCKET_SPLASH"},
		},
		{
			name:    "should parse kill log line when the killer has more than a single word in the name",
			logLine: "2:11 Kill: 2 4 6: Dono da Bola killed Zeh by MOD_ROCKET",
			want:    Kill{Killed: "Zeh", Killer: "Dono da Bola", Reason: "MOD_ROCKET"},
		},
		{
			name:    "should parse kill log line when the killed has more than a single word in the name",
			logLine: "2:11 Kill: 2 4 6: Test killed Dono da Bola by MOD_ROCKET",
			want:    Kill{Killed: "Dono da Bola", Killer: "Test", Reason: "MOD_ROCKET"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseKill(tt.logLine); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseKill() = %v, want %v", got, tt.want)
			}
		})
	}
}
