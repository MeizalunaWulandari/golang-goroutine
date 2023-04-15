package golang_goroutine

import(
	"fmt"
	"testing"
	"sync"
	"time"
)

func RunAsynchronous(group *sync.WaitGroup){
	defer group.Done()

	group.Add(1)

	fmt.Println("Hello")
	time.Sleep(1 * time.Second)
}

func TestWaitGroup(t *testing.T){
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		go RunAsynchronous(group)
	}

	group.Wait()
	fmt.Println("Selesai")

}

/************************
 * WAITGROUP
 * WaitGroup adalah fitur yang bisa digunakan untuk menunggu sebuah proses selesai dilakukan
 * Hal ini diperlukan, misalnya ketika kita ingin menjalankan beberapa proses menggunakan goroutine,
 * tapi kita ingin semua proses selesai terlebih dahulu sebelum aplikasi selesai
 * Pada kasus seperti ini bisa menggunakan WaitGroup
 * Untuk menandai bahwa ada proses goroutine, kita bisa menggunakan method Add(int), setelah 
 * proses goroutine selesai kita  bisa menggunakan  method Done()
 * Untuk menunggu proses selesai kita bisa menggunakan method Wait()
 ************************/