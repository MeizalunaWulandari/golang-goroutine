package golang_goroutine

import(
	"fmt"
	"testing"
	"sync"
	"time"
)

var locker =  sync.Mutex{}
var cond = sync.NewCond(&locker)
var group = sync.WaitGroup{}

func WaitCondition(value int) {
	defer group.Done()
	group.Add(1)

	cond.L.Lock()
	cond.Wait()
	fmt.Println("Done", value)
	cond.L.Unlock()
}

func TestCond(t *testing.T){
	for i := 0; i < 100; i++ {
		go WaitCondition(i)
	}

	go func() {
		for i := 0; i < 100; i++ {
			time.Sleep(1 * time.Second)
			cond.Signal()
		}
	}()

	// go func() {
	// 	time.Sleep(1 * time.Second)
	// 	cond.Broadcast()
	// }()

	group.Wait()
}


/*******************
 * SYNC COND
 * sync.Cond adalah implementasi locking berbasis kodisi
 * Cond membutuhkan locker (bisa menggunakan Mutex atau RWMutex ) untuk implementasi
 * lockingnya, namun berbeda dengan locker biasanya, di cond terdapag function Wait()
 * untuk menunggu apakah perlu menunggu atau tidak
 * Function Signal() bisa digunakan untuk memberi tahu sebuah goroutine agar tidak 
 * perlu menunggu, sedangkan Broadcast() digunakan untuk memberi tahu semua goroutine
 * agar tidak perlu menunggu lagi
 * Untuk membuat Cond, kita bisa menggunakan function sync.NewCond(Locker)
 *******************/