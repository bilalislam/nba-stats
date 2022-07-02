package main

const (
	TeamA TeamKey = "team a"
	TeamB TeamKey = "team b"
)

type Interface interface {
	Float64() float64
	Intn(n int) int
}

type TeamKey string

type Team struct {
	ScoreCount  int
	AttackCount int
	Players     []Player
	Random      Interface
}

// accept interface return structs
func NewTeam(rnd Interface) Team {
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

func (t *Team) Attack() {
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
