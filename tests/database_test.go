package tests

import (
	"sync"
	"testing"

	"github.com/latext/rice"
)

func TestNewMaria(t *testing.T) {
	db, err := rice.NewMaria("root:root@tcp(localhost:3306)/test?parseTime=True&loc=Local&charset=utf8mb4")
	if err != nil {
		t.Error(err)
	}
	defer db.Close()
}

func TestNewPostgres(t *testing.T) {

	var wg sync.WaitGroup

	for i := 0; i < 10000; i++ {
		wg.Add(2)
		go func() {
			db, err := rice.NewMaria("root:root@tcp(localhost:3306)/test?parseTime=True&loc=Local&charset=utf8mb4")
			if err != nil {
				t.Error(err)
			}
			defer db.Close()
			wg.Done()
		}()

		go func() {
			db, err := rice.NewPostgres("postgres://postgres:123456@localhost:5432/test?sslmode=disable")
			if err != nil {
				t.Error(err)
			}
			defer db.Close()
			wg.Done()
		}()
	}

	wg.Wait()
}
