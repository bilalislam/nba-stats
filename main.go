package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

//todo: amac 5snde bir cagrılmasını kontrol etmek
func main() {
	s1 := rand.NewSource(time.Now().UnixNano())
	r1 := rand.New(s1)

	numberOfMatch := 2
	var wg sync.WaitGroup
	wg.Add(numberOfMatch)
	for i := 0; i < numberOfMatch; i++ {
		go func() {
			defer wg.Done()
			t1 := NewTeam(r1) // listeye ekle print için
			t2 := NewTeam(r1) // listeye ekle print için
			m := NewMatch(&t1, &t2, r1)
			ticker := time.NewTicker(5 * time.Second)
			for i := 0; i < 48; i++ {
				<-ticker.C // 5sn boyunca bekleyecek
				err, match := m.Play()
				if err != nil {
					return
				}
				for key, team := range match.Teams {
					fmt.Println(key)
					fmt.Println(team.AttackCount)
					fmt.Println(team.ScoreCount)
					for _, player := range team.Players {
						fmt.Println(player.Name)
						fmt.Println(player.SuccessCount)
						fmt.Println(player.AssistCount)
						fmt.Println(player.FailureCount)
					}
				}
			}
			ticker.Stop()
		}()
	}
	wg.Wait()

	// t1 ve t2 için print
}
