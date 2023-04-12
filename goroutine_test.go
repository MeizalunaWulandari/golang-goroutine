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
/**
 * Untuk membuat go routine cukup dengan menambahkan keyword go sebelum menjalankan function
 * Namun jika menggunakan go routine jika function tersebut belum selesai dan program telah selesai
 * maka function tersebut akan dimatikan secara otamatis, maka dari itu pastikan untuk membuat jeda 
 * agar function tersebut dapat diselasaikan sebelum program berhenti
 */