package main

import (
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
				err := m.Play()
				if err != nil {
					return
				}
			}
			ticker.Stop()
		}()
	}
	wg.Wait()

	// t1 ve t2 için print
}
