package golang_goroutine

import(
	"fmt"
	"testing"
	"time"
	"sync"
)

func TestTimer(t *testing.T){
	timer := time.NewTimer(5 * time.Second)
	fmt.Println(time.Now())

	time := <- timer.C
	fmt.Println(time)
}

func TestAfter(t *testing.T){
	channel := time.After(5 * time.Second)
	fmt.Println(time.Now())

	// time := <- timer.C // DENGAN TIMER
	time := <- channel // DENGAN AFTER

	fmt.Println(time)
}

func TestTimerFunc(t *testing.T){
	group :=  sync.WaitGroup{}
	group.Add(1)

	time.AfterFunc(5 * time.Second, func (){
		fmt.Println(time.Now())
		group.Done()
	})
	fmt.Println(time.Now())
	group.Wait()

}

/*******************
 * TIME TIMER
 * Timer adalah representasi satu kejadian
 * Ketika waktu timer sudah expire, maka event akan dikirim kedalam channel
 * Untuk membuat channel kita bisa menggunakan time.NewTimer(duration)
 * 
 * * time.After
 * Kadang kita hanya butuh channelnya saja, tidak membutuhkan data Timernya
 * Untuk melakukan hal itu kita bisa menggunakan function time.After(Duration)
 * 
 * * time.AfterFunc
 * Kadang ada kebutuhan kita ingin menjalankan sebuah function dengan delay waktu tertentu
 * Kita bisa memanfaatkan Timer dengan menggunakan funtion time.AfterFunc()
 * Kita tidak perlu lagi menggunakan channelnya, melainkan cukup kirim function 
 * yang akan dipanggil ketika Timer mengirim kejadian
 *******************/