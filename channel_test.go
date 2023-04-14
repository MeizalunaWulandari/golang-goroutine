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

func GiveMeResponse(channel chan string){
    time.Sleep(2 * time.Second)
    channel <- "Meizaluna Wulandari"
}

func TestChannelAsParameter(t *testing.T){
    channel := make(chan string)
    defer close(channel)

    go GiveMeResponse(channel)

    data := <- channel
    fmt.Println(data)
}

func OnlyIn(channel chan <- string ){
    time.Sleep(2 * time.Second)
    channel <- "Meizaluna Wulandari"   
}

func OnlyOut(channel <- chan string ){
    data := <- channel
    fmt.Println(data)
}

func TestInOutChannel(t *testing.T){
    channel := make(chan string)
    defer close(channel)

    go OnlyIn(channel)
    go OnlyOut(channel)
    time.Sleep(2 * time.Second)
}

func TestBufferedChannel(t *testing.T){
    channel := make(chan string, 3)
    defer close(channel)    

    go func() {
        channel <- "Luna"
        channel <- "Andini"
        channel <- "Rizka"
    }()

    go func() {
        fmt.Println(<- channel)
        fmt.Println(<- channel)
        fmt.Println(<- channel)
    }()
    time.Sleep(2 * time.Second)
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
 * 
 * CHANNEL SEBAGAI PARAMETER
 * Channel secara otomatis menggunakan pass by refence sehingga tidak perlu menggunakan pointer untuk
 * melakukan hal ini
 * 
 * CHANNEL IN DAN OUT
 * Sebuah channel dapat digunakan sebagai pengirim maupun penerima data
 * namun jika menginginkan sebuah channel hanya bisa digunakan sebagai sebuah penerima
 * atau pengim saja bisa dengan memberi penanda pada channel tersebut ( chan <-)
 * penannda in sebagai pengirim data dan penanda out sebagai penerima ( <-chan )
 * 
 * CHANNEL BUFFERED
 * Pada golang channel hanya bisa mengirim 1 data jika ingin menanmbahkan data lagi maka 
 * data tersebut akan dimasukan dalam antrian sampai data pertama atau sebelumnya diambil
 * Buffered berfungsi sebagai penampung data antrian
 ** Buffer capacity
 * Kita bebas memasukkan berapa jumlah kapasitas antrian didalam buffer
 * jika kita set 5 maka buffer tersebut dapat menerima data sebanyak 5 data saja
 * dan jika ingin mengirim data ke 6 maka kita harus menunggu sampai ada buffer yang kosong
 * Untuk membuat buffer cukup menambahkannya capacitynya sebagai argumen pada function make
 * make(chan string, 5)
 */