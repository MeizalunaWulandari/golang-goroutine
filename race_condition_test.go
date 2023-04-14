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
 */