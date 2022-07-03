package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

type mockRandomForMatch struct {
	floatValue float64
	intValue   int
}

func (m *mockRandomForMatch) Float64() float64 {
	return m.floatValue
}

func (m *mockRandomForMatch) Intn(n int) int {
	return m.intValue
}

func TestPlay_Should_ThrowError_When_TimeHasFinished(t *testing.T) {
	rnd := mockRandomForMatch{
		floatValue: 0,
		intValue:   0,
	}

	teamA := NewTeam(&rnd)
	teamB := NewTeam(&rnd)
	m := NewMatch(&teamA, &teamB, &rnd)
	m.Time = time.Minute * 48

	err := m.Play()
	assert.EqualError(t, err, "match has been finished")
}

func TestPlay_Should_AttackForTeamA_WhenAttackCountIsThreeAndScoreIsNine(t *testing.T) {

	rndForMatch := mockRandomForMatch{
		floatValue: 0.6,
		intValue:   20,
	}

	rndForTeam := mockRandom{
		floatValue: 1,
		intValue:   []int{0, 1},
		index:      -1,
	}

	teamA := NewTeam(&rndForTeam)
	teamB := NewTeam(&rndForTeam)
	m := NewMatch(&teamA, &teamB, &rndForMatch)

	err := m.Play()
	assert.EqualValues(t, nil, err)
	assert.EqualValues(t, time.Minute, m.Time)
	assert.EqualValues(t, 3, teamA.AttackCount)
	assert.EqualValues(t, 9, teamA.ScoreCount)
}

func TestPlay_Should_AttackForTeamB_WhenAttackCountIsSixAndScoreCountIsTwelve(t *testing.T) {
	rndForMatch := mockRandomForMatch{
		floatValue: 0.4,
		intValue:   10,
	}

	rndForTeam := mockRandom{
		floatValue: 0.90,
		intValue:   []int{0, 1},
		index:      -1,
	}

	teamA := NewTeam(&rndForTeam)
	teamB := NewTeam(&rndForTeam)
	m := NewMatch(&teamA, &teamB, &rndForMatch)

	err := m.Play()
	assert.EqualValues(t, nil, err)
	assert.EqualValues(t, time.Minute, m.Time)
	assert.EqualValues(t, 6, teamB.AttackCount)
	assert.EqualValues(t, 12, teamB.ScoreCount)
}
