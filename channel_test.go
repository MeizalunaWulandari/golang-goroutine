package golang_goroutine

import(
    "fmt"
    "testing"
    "time"
)

func TestCreateChannel(t *testing.T){
    channel := make(chan string)
    defer close(channel)

    go func() {
        time.Sleep(2 * time.Second)
        channel <- "Meizaluna Wulandari"
        fmt.Println("Selesai Mengirim data ke Channel")
    }()

    data := <- channel
    fmt.Println(data)
    time.Sleep(5 * time.Second)
}

/** MEMBUAT CHANNEL
 * Channel di golang direpresentasikan dengan tipe data chan
 * Untuk membuat channel kita bisa menggunakan function make()
 * Saat pembuatan channel kita harus menentukan type data yang bisa dimasukkan kedalam channel tersebut
 * make(chan TypeData)
 *  
 * untuk mengirim data kita bisa menggunakan kode: channel <- data
 * sedangkan untuk menerima data bisa gunakan kode data <- channel
 * Jika selesai jangan lupa untuk menutup channel menggunakan function close()
 * Jika membuat channel pastikan ada yang mengirim dan menerima, karena jika salah satunya tidak ada 
 * maka akan terjadi error
 */