package main

import (
	"errors"
	"math/rand"
	"time"
)

type Match struct {
	Teams  map[TeamKey]*Team
	Time   time.Duration
	Random *rand.Rand
}

func NewMatch(teamA *Team, teamB *Team, rnd *rand.Rand) Match {
	return Match{
		Random: rnd,
		Teams: map[TeamKey]*Team{
			TeamA: teamA,
			TeamB: teamB,
		},
	}
}

// içerde 1 dk olarak calısacak, toplam 48 dk olacak
// main'de 1 dk = 5sn , 48 dk = 240 sn (bu hesabın play ile ilgisi yok)
// yani play 1 dk olarak herşeyi hesaplacak
// yani 24 sn lik attack'ın 5 snlik mac ile ilgisi yok, Gercek dünyada mac'ın suresinden düşülecek .
// todo: unit test
func (m Match) Play() error {

	if m.Time > time.Minute*47 {
		return errors.New("match has been finished")
	}

	rnd := m.Random.Float64()
	total := 0

	for total < 60 {
		if rnd > 0.5 {
			m.attack(TeamA)
		} else {
			m.attack(TeamB)
		}
		attackRand := m.Random.Intn(24) + 1 // 0sn lik attack olmaz , en az 1 sn lik attack gelebilir
		total += attackRand
	}

	m.Time += time.Minute
	return nil
}

func (m Match) attack(key TeamKey) {
	m.Teams[key].Attack()
}
