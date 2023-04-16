package golang_goroutine

import(
	"testing"
	"fmt"
	"sync"
	"sync/atomic"
)

func TestAtomic(t *testing.T){
   var x int64 = 0
   group := sync.WaitGroup{}

    for i := 1; i <= 100; i++{
        group.Add(1)
        go func() {
            for j := 1; j <= 100; j++{
                atomic.AddInt64(&x, 1)
            }
            group.Done()
        }()
    }

    group.Wait()
    fmt.Println("Counter = ", x)
}

/*******************
 * ATOMIC
 * Golang memiliki package bernama sync/atomic
 * Atomic merupakan package yang digunakan untuk menggunakan data primitive
 * secara aman pada proses concurrent
 *******************/