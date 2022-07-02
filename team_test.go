package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockRandom struct {
	floatValue float64
	intValue   []int
	index      int
}

func (m *mockRandom) Float64() float64 {
	return m.floatValue
}

func (m *mockRandom) Intn(n int) int {
	m.index++
	return m.intValue[m.index]
}

func TestAttack_Should_Success_AttackCount_Increment(t *testing.T) {
	rnd := mockRandom{
		floatValue: 1,
		intValue:   []int{0, 1},
		index:      -1,
	}
	team := NewTeam(&rnd)
	team.Attack()
	assert.EqualValues(t, 1, team.AttackCount)
	assert.EqualValues(t, 3, team.ScoreCount)
	assert.EqualValues(t, 1, team.Players[0].AssistCount)
	assert.EqualValues(t, 0, team.Players[0].SuccessCount)
	assert.EqualValues(t, 1, team.Players[1].SuccessCount)
	assert.EqualValues(t, 0, team.Players[1].AssistCount)
}

func TestAttack_Should_Failure_AttackCount_Increment(t *testing.T) {
	rnd := mockRandom{
		floatValue: 0.5,
		intValue:   []int{0, 1},
		index:      -1,
	}
	team := NewTeam(&rnd)
	team.Attack()
	assert.EqualValues(t, 1, team.AttackCount)
	assert.EqualValues(t, 0, team.ScoreCount)
	assert.EqualValues(t, 1, team.Players[0].AssistCount)
	assert.EqualValues(t, 1, team.Players[1].FailureCount)
}
