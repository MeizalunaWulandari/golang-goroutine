package golang_goroutine

import (
	"fmt"
	"testing"
	// "sync"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(1 * time.Nanosecond)

	go func() {
		time.Sleep(5 * time.Second)
		ticker.Stop()
	}()

	for time := range ticker.C {
		fmt.Println(time)
	}
}

func TestTick(t *testing.T){
	channel := time.Tick(1 * time.Second)

	for time := range channel {
		fmt.Println(time)
	}
}

/*******************
 * TIME TICKER
 * time.Ticker adalah representasi kejadian berulang
 * Ketika waktu ticker sudah expired, maka event akan dikirim ke dalam channel
 * Untuk membuat ticker kita bisa menggunakan function time.NewTicker(duration)
 * Untuk menghentikan ticker, kita bisa menggunakan Ticker.Stop
 * *time.Tick
 * Kadang kita tidak butuh data Ticker nya, kita hanya butuh channelnya saja
 * Jika demikian, kita bisa menggunakan function time.Tick(duration), funcion ini 
 * tidak akan mengembalikan Ticker, namun hanya mengembalikan channel timernya saja
 *******************/