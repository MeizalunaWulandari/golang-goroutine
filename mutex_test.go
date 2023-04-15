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

type UserBalance struct{
    sync.Mutex
    Name string
    Balance int
}

func (user *UserBalance) Lock(){
    user.Mutex.Lock()
}

func (user *UserBalance) Unlock(){
    user.Mutex.Unlock()
}

func (user *UserBalance) Change(amount int){
    user.Balance = user.Balance + amount
}


func Transfer(user1 *UserBalance, user2 *UserBalance, amount int){
    user1.Lock()
    fmt.Println("Lock user1", user1.Name)
    user1.Change(-amount)

    time.Sleep(1 * time.Second)

    user2.Lock()
    fmt.Println("Lock user2", user2.Name)
    user2.Change(amount)

    time.Sleep(1 * time.Second)

    user1.Unlock()
    user2.Unlock()
}

func TestDeadlock(t *testing.T){
    user1 := UserBalance{
        Name : "Andini",
        Balance : 1000000,
    }
    user2 := UserBalance{
        Name : "Luna",
        Balance : 1000000,
    }

    go Transfer(&user1, &user2, 100000)
    go Transfer(&user2, &user1, 200000)
    
    time.Sleep(1 * time.Second)

    fmt.Println("User ", user1.Name, ", Balance", user1.Balance)
    fmt.Println("User ", user2.Name, ", Balance", user2.Balance)

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
 * 
 * DEADLOCK
 * Deadlock adalah masalah yang biasa terjadi ketika membuat aplikasi yang concurrent dan paraller
 * Deadlock sendiri terjadi karena sebuah goroutine saling menunggu sehingga tidak ada satupu 
 * goroutine yang berjalan, hal ini biasanya disebabkan oleh locking
 */