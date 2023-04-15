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

type BankAccount struct{
    RWMutex sync.RWMutex
    Balance int
}

func (account *BankAccount) AddBalance(amount int) {
    account.RWMutex.Lock()
    account.Balance = account.Balance + amount
    account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance()int{
    account.RWMutex.RLock()
    balance := account.Balance
    account.RWMutex.RUnlock()
    return balance
}

func TestRWMutex(t *testing.T){
    account := BankAccount{}

    for i := 0; i < 100; i++{
        go func() {
            for j := 0; j < 100; j++{
                account.AddBalance(1)
                fmt.Println(account.GetBalance())
            }
        }()
    }

    time.Sleep(5 * time.Second)
    fmt.Println("Total Balance", account.GetBalance())

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
 * 
 * SYNC.RWMUTEX
 * RWMutex(Write Write Mutex)
 * Digunakan ketika kita tidak hanya ingin melakukan locking pada proses menngubah data,
 * tetapi juga merubah data
 * Sebenarnya bisa menggunakan Mutex saja, namun masalahkan proses antara proses membaca dan mengubah data
 * akan rebutan
 * Di golang telah disediakan struct RWMutex untuk menangani hal ini, dimana mutex jenis ini memiliki dua 
 * dua lock, lock untuk Read dan Lock untuk Write
 * RWMutex hanya diperlukan ketika membuat data yang akan diakses oleh beberapa goroutine
 */