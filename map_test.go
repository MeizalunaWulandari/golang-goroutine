package golang_goroutine

import (
	"fmt"
	"testing"
	"sync"
)

func AddToMap(data *sync.Map, value int, group *sync.WaitGroup){
	defer group.Done()

	group.Add(1)
	data.Store(value, value)
}

func TestMap(t *testing.T){
	data := &sync.Map{}
	group := &sync.WaitGroup{}

	for i := 0; i < 100; i++{
		go AddToMap(data, i, group)
	}
	group.Wait()
	data.Range(func (key, value interface{}) bool{
		fmt.Println(key ,":", value)
		return true
	})
}

/*******************
 * SYNC MAP
 * Golang memiliki sebuah struct bernama sync.Map, Map ini mirip dengan
 * golang map, namun yang membedakan adalah map ini aman untuk menggunakan 
 * concurrent mengunakan goroutine (Aman dari race condition)
 * Ada beberapa function yang bisa digunakan di Map
 * 	* Store(key, value) untuk menyimpan data ke Map
 * 	* Load(key) untuk mengambil data dari Map menggunakan key
 * 	* Delete(key) untuk menghapus data dari Map menggunakan key
 * 	* Range(function(key, value)) untuk melakukan iterasi seluruh data
 * Jadi jika menggunakan goroutine jangan menggunakan map bawaan golang melainkan
 * menggunakan sync.Map karena aman dari race condition
********************/