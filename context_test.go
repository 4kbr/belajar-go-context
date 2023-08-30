package belajar_go_context

import (
	"context"
	"fmt"
	"runtime"
	"testing"
	"time"
)

// /context
func TestContext(t *testing.T) {
	//buat context secara manual
	background := context.Background()
	todo := context.TODO()

	fmt.Println(background)
	fmt.Println(todo)
}

func TestContextWithValue(t *testing.T) {
	contextA := context.Background()

	// child contextA
	contextB := context.WithValue(contextA, "b", "B")
	contextC := context.WithValue(contextA, "c", "C")

	// child contextB
	contextD := context.WithValue(contextB, "d", "D")
	contextE := context.WithValue(contextB, "e", "E")

	// child contextC
	contextF := context.WithValue(contextC, "f", "F")
	contextG := context.WithValue(contextC, "g", "G")

	fmt.Println("contextA", contextA)
	fmt.Println("contextB", contextB)
	fmt.Println("contextC", contextC)
	fmt.Println("contextD", contextD)
	fmt.Println("contextE", contextE)
	fmt.Println("contextF", contextF)
	fmt.Println("contextG", contextG)

	fmt.Println(contextF.Value("f")) //ada
	fmt.Println(contextF.Value("c")) //ada
	fmt.Println(contextF.Value("d")) //beda parent (nil)
	fmt.Println(contextA.Value("b")) //tidak bisa mengambil value childnya sendiri (nil)

}

func CreateCounter(ctx context.Context) chan int {
	destination := make(chan int)

	go func() {
		defer close(destination)
		counter := 1
		for {
			select {
			case <-ctx.Done():
				return
			default:
				destination <- counter
				counter++
				time.Sleep(1 * time.Second)
			}
		}
	}()

	return destination
}

// Cancel
func TestContextWithCancel(t *testing.T) {
	fmt.Println("Before, total goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithCancel(parent)

	destination := CreateCounter(ctx)
	for n := range destination {
		fmt.Println("Counter", n)
		if n == 10 {
			break
		}
	}

	cancel() //mengirim sinyal cancel ke context

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}

func TestContextWithTimeout(t *testing.T) {
	fmt.Println("Before, total goroutine", runtime.NumGoroutine())
	parent := context.Background()
	ctx, cancel := context.WithTimeout(parent, 5*time.Second)
	defer cancel()

	destination := CreateCounter(ctx)

	for n := range destination {
		fmt.Println("Counter", n)
	}

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
func TestContextWithDeadline(t *testing.T) {
	//dengan deadline
	fmt.Println("Before, total goroutine", runtime.NumGoroutine())

	parent := context.Background()
	ctx, cancel := context.WithDeadline(parent, time.Now().Add(10*time.Second))
	defer cancel()

	destination := CreateCounter(ctx)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	for n := range destination {
		fmt.Println("Counter", n)
	}

	fmt.Println("Total Goroutine", runtime.NumGoroutine())
	time.Sleep(2 * time.Second)
	fmt.Println("Total Goroutine", runtime.NumGoroutine())
}
