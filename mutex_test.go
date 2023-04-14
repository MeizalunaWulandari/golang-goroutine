package golang_goroutine

import (
    "fmt"
    "testing"
    "time"
    "sync"
)

func TestMutex(t *testing.T){
    x := 0
    var mutex sync.Mutex

    for i := 1; i <= 100; i++{
        go func() {
            for j := 1; j <= 100; j++{
                mutex.Lock()
                x = x + 1
                mutex.Unlock()
            }
        }()
    }

    time.Sleep(2 * time.Second)
    fmt.Println("Counter = ", x)
}

/** SYNC.MUTEX
 * Mutex(Mutual Exclusion) 
 * Untuk mengatasi masalah race condition, di golang terdapat sebuah struct bernama sync.Mutex
 * Mutex bisa digunakan untuk melakukan locking dan unlocking , dimana  ketika  melakukan locking terhadap
 * mutex, maka tidak ada yang bisa melakukan locking lagi sampai kita melakukan unlock
 * Dengan Demikan, jika ada beberapa goroutine melakukan lock terhadap Mutex, maka hanya 1 goroutine yang 
 * diperbolehkan, setelah goroutine tersebut melakukan unlock, maka baru goroutine selanjutnya 
 * diperbolehkan melakukan lock lagi
 * Ini sangat cocok sebagai solusi ketika kita ada masalah race condition
 */