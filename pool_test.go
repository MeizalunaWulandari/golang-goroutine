package golang_goroutine

import (
	"fmt"
	"testing"
	"sync"
	"time"
)

func TestPool(t *testing.T){
	pool := sync.Pool{
		New : func() interface{} {
			return "New"
		},
	}

	pool.Put("Meizaluna")
	pool.Put("Wulandari")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(11 * time.Second)
	fmt.Println("Selesai")

}

/*******************
 * POOL
 * Pool adalah implementasi design pattern bernama object pool pattern
 * Sederhananya, design pattern Pool ini digunakan untuk menyimpan data, selanjutnya 
 * untuk menggunakan datanya, kita bisa mengambil dari pool, dan setelah selesai 
 * menggunakan datanya kita bisa menyimpannya kembali kedalam pool
 * Implementasi Pool di golang sudah aman dari problem race condition
 * Untuk default value dari Pool kita bisa menambahkan atrubut New pada struct Pool tersebut
 *******************/