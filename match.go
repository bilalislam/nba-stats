package main

import (
	"errors"
	"time"
)

type Match struct {
	Teams  map[TeamKey]*Team
	Time   time.Duration
	Random Interface
}

func NewMatch(teamA *Team, teamB *Team, rnd Interface) Match {
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
// hangi takıma denk geldiyse o takım 5sn içinde 60sn max duration olmak sartıyla attack'a kalkar,5sn içinde kac attack yaparsa yapar
// tüm 60sn lik süre içinde sonuna kadar beklemez, 5sn dolarsa biter.Takım attack'a kalkamaz.
// todo: unit test
func (m *Match) Play() (error, *Match) {

	if m.Time > time.Minute*47 {
		return errors.New("match has been finished"), nil
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
	return nil, m
}

func (m *Match) attack(key TeamKey) {
	m.Teams[key].Attack()
}
