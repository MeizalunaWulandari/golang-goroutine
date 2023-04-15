package golang_goroutine

import (
	"fmt"
	"sync"
	"testing"
)

var counter = 0

func OnlyOnce(){
	counter++
}

func TestOnce(t *testing.T){
	once := sync.Once{}
	group := sync.WaitGroup{}

	for i := 0; i < 100; i++{
		go func() {
			group.Add(1)
			once.Do(OnlyOnce)
			group.Done()
		}()
	}
	group.Wait()
	fmt.Println("Counter", counter)
}

/*******************
 * SYNC.ONCE
 * sync.Once adalah fitur di golang yang bisa digunakan untuk memastikan bahwa sebuah 
 * function dieksekusi hanya sekali
 * Jadi berapa banyak pun goroutine yang mengakses, bisa dipastikan bahwa goroutine 
 * yang pertamalah yang bisa mengeksekusi function tersebut 
 * Goroutine lainnya akan dihiraukan, artinya function tidak akan dieksekusi lagi
 * Untuk membuat once gunakan struct Do 
 * Do(nameFunction)
 * Pada nama function cukup menuliskan nama function tanpa ()
 *******************/