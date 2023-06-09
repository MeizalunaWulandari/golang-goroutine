package golang_goroutine

import(
    "fmt"
    "testing"
    "time"
)

func TestRaceCondition(t *testing.T){
    x := 0

    for i := 1; i <= 100; i++{
        go func() {
            for j := 1; j <= 100; j++{
                x = x + 1
            }
        }()
    }

    time.Sleep(2 * time.Second)
    fmt.Println("Counter = ", x)
}

/** RACE CONDITION
 * Saat kita menggunakan goroutine, dia tidak hanya berjalan secara concurrent, tapi juga bisa pararel juga
 * karena bisa ada beberapa thread yang berjalan secara pareler
 * hal ini sangat berbahaya ketika kita melakukan manipulasi data variable yang sama oleh beberapa goroutine 
 * secara bersamaan
 * Hal ini bisa menyebabkan masalah yang namanya Race Condition
 * SYNC.MUTEX
 * Mutex(Mutual Exclusion) 
 * Untuk mengatasi masalah race condition, di golang terdapat sebuah struct bernama sync.Mutex
 * Mutex bisa digunakan untuk melakukan locking dan unlocking , dimana  ketika  melakukan locking terhadap
 * mutex, maka tidak ada yang bisa melakukan locking lagi sampai kita melakukan unlock
 * Dengan Demikan, jika ada beberapa goroutine melakukan lock terhadap Mutex, maka hanya 1 goroutine yang 
 * diperbolehkan, setelah goroutine tersebut melakukan unlock, maka baru goroutine selanjutnya 
 * diperbolehkan melakukan lock lagi
 * Ini sangat cocok sebagai solusi ketika kita ada masalah race condition
 */