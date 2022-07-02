package main

import (
	"math/rand"
	"time"
)

func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	t1 := NewTeam(r1)
	t2 := NewTeam(r1)

	m := NewMatch(&t1, &t2, r1)

	// todo : ticker ile 5 sn bir call et , toplam 48 kere call yapılabilir , os.exit(-1)
	err := m.Play()
	if err != nil {
		return
	}
}
