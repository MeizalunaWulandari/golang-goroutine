package golang_goroutine

import(
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld(){
	fmt.Println("Hello World")
}

func TestCreateGoRoutine(t *testing.T){
	go RunHelloWorld()
	fmt.Println("Ups")
	time.Sleep(1 * time.Second)
}

func DisplayNumber(number int){
	fmt.Println("Display ke", number)
}

func TestManyGoroutine(t *testing.T){
	for i := 0; i < 100000; i++{
		go DisplayNumber(i)
	}
	time.Sleep(5 * time.Second)
}

/**
 * Untuk membuat go routine cukup dengan menambahkan keyword go sebelum menjalankan function
 * Namun jika menggunakan go routine jika function tersebut belum selesai dan program telah selesai
 * maka function tersebut akan dimatikan secara otamatis, maka dari itu pastikan untuk membuat jeda 
 * agar function tersebut dapat diselasaikan sebelum program berhenti
 * Goroutine tidak adakan berguna pada function yang mengembalikan return value
 */