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

const (
	TeamA TeamKey = "team a"
	TeamB TeamKey = "team b"
)

type TeamKey string

type Team struct {
	ScoreCount  int
	AttackCount int
	Players     []Player
	Random      *rand.Rand
}

func NewTeam(rnd *rand.Rand) Team {
	return Team{
		Players: []Player{
			{
				Name: "player1",
			},
			{
				Name: "player2",
			},
			{
				Name: "player3",
			},
			{
				Name: "player4",
			},
			{
				Name: "player5",
			},
		},
		Random: rnd,
	}
}

func (t Team) Attack() {
	rnd := t.Random.Float64()

	//Todo : it may be retry for the same players with a method
	assistRnd := t.Random.Intn(len(t.Players))
	shooterRnd := t.Random.Intn(len(t.Players))

	if rnd >= 0.9 {
		t.ScoreCount += 2
		if rnd >= 0.95 {
			t.ScoreCount++
		}
		t.Players[shooterRnd].SuccessCount++
	} else {
		t.Players[shooterRnd].FailureCount++
	}

	t.Players[assistRnd].AssistCount++
	t.AttackCount++
}

type Player struct {
	Name         string
	AssistCount  int
	SuccessCount int
	FailureCount int
}
