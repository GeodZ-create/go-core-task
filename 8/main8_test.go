package main

import (
	"testing"
	"time"
)

func TestSemaphoreWait(t *testing.T) {
	wg := NewSemaphore(3)
	counter := 0

	go func() {
		wg.Done()
		counter++
	}()
	go func() {
		wg.Done()
		counter++
	}()
	go func() {
		wg.Done()
		counter++
	}()
	wg.Wait(3)

	if counter != 3 {
		t.Errorf("TestSemaphoreWait is error, got = %d", counter)
	}
}

func TestSemaphoreBlocksUntilDone(t *testing.T) {
	wg := NewSemaphore(1)
	done := make(chan struct{})

	go func() {
		wg.Wait(1)
		close(done)
	}()

	select {
	case <-done:
		t.Fatal("Wait returned before Done")
	case <-time.After(50 * time.Millisecond):

	}

	wg.Done()

	select {
	case <-done:
	case <-time.After(50 * time.Millisecond):
		t.Fatal("Wait not return after Done")
	}

}
