package golang_goroutine

import(
    "fmt"
    "testing"
    "runtime"
    "time"
    "sync"
)

func TestGetGomaxprocs(t *testing.T){

    group := sync.WaitGroup{}

    for i := 0; i < 100; i++{
        group.Add(1)
        go func() {
            time.Sleep(3 * time.Second)
            group.Done()
        }()
    }

    totalCpu := runtime.NumCPU()
    fmt.Println("Total CPU : ", totalCpu)


    totalThread := runtime.GOMAXPROCS(-1)
    fmt.Println("Total Thread : ", totalThread)

    totalGoroutine := runtime.NumGoroutine()
    fmt.Println("Total Goroutine : ", totalGoroutine)
    group.Wait()
}

func TestChangeThreadNum(t *testing.T){

    group := sync.WaitGroup{}

    for i := 0; i < 100; i++{
        group.Add(1)
        go func() {
            time.Sleep(3 * time.Second)
            group.Done()
        }()
    }

    totalCpu := runtime.NumCPU()
    fmt.Println("Total CPU : ", totalCpu)

    runtime.GOMAXPROCS(20)
    totalThread := runtime.GOMAXPROCS(-1)
    fmt.Println("Total Thread : ", totalThread)

    totalGoroutine := runtime.NumGoroutine()
    fmt.Println("Total Goroutine : ", totalGoroutine)
    group.Wait()
}

/*******************
 * GOMAXPROCS
 * Gomaxprocs yaitu sebuah function di package runtime yang bisa digunakan
 * untuk mengetahui jumlah thread atau mengambil jumlah tread
 * Secara defalut, jumlah thread di golang itu sebanyak jumlah CPU di kompoter kita
 * Kita juga bisa melihat berapa jumlah CPU kita dengan menggunakan function
 * runtime.NumCpu()
 * 
 * * Mengubah jumlah Thread
 * Untuk merubah jumlah thread bisa dengan runtime.GOMAXPROCS(theead)
 * runtime.GOMAXPROCS(10)
 *******************/