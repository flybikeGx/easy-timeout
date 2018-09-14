package timeout

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestRun(t *testing.T) {
	if !Run(1*time.Second, func() {
		time.Sleep(2 * time.Second)
	}) {
		t.Log("time out: OK")
	} else {
		t.Fatal("time out: FAILED")
	}

	rand.Seed(time.Now().Unix())

	for i := 0; i < 100; i++ {
		timelimit := time.Millisecond * 500
		runtime := time.Millisecond * time.Duration(rand.Float64()*1000)
		fmt.Println(timelimit, runtime)
		if Run(timelimit, func() {
			time.Sleep(runtime)
		}) != (timelimit > runtime) {
			t.Fatal("random error: ", timelimit, runtime)
		}
	}
}

func TestPanic(t *testing.T) {
	timelimit := time.Millisecond * 500
	ok := Run(timelimit, func() {
		defer func() {
			if err := recover(); err != nil {
				fmt.Println(err)
			}
		}()
		panic("noooo")

	})
	fmt.Println(ok)
}
