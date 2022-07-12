package alog

import (
	"bytes"
	"strings"
	"testing"
	"time"
)

// 01
func TestNewInitializesShutdownChannelsModule3(t *testing.T) {
	alog := New(nil)
	if alog.shutdownCh == nil {
		t.Error("shutdownCh field not initialized")
	}

	if alog.shutdownCompleteCh == nil {
		t.Error("shutdownCompleteCh field not initialized")
	}
}

// 02

func TestShutdownMethodModule3(t *testing.T) {
	alog := New(nil)
	alog.shutdownCompleteCh = make(chan struct{}, 1)
	alog.shutdown()
	time.Sleep(100 * time.Millisecond)
	select {
	case _, ok := <-alog.msgCh:
		if ok {
			t.Error("msgCh not closed by shutdown() method")
		}
	default:
		t.Error("msgCh not closed by shutdown() method")
	}
	select {
	case <-alog.shutdownCompleteCh:
	default:
		t.Error("shutdown() doesn't send message to shutdownCompleteCh")
	}

}

// 03

func TestStartMethodCallsShutdownModule3(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	alog := New(b)
	alog.shutdownCh = make(chan struct{}, 1)
	alog.shutdownCompleteCh = make(chan struct{}, 1)
	go alog.Start()
	alog.shutdownCh <- struct{}{}
	time.Sleep(100 * time.Millisecond)

	select {
	case _, ok := <-alog.msgCh:
		if ok {
			t.Error("Passing message to shutdownCh doesn't call shutdown()")
		}
	default:
		t.Error("Passing message to shutdownCh doesn't call shutdown()")
	}
	select {
	case <-alog.shutdownCompleteCh:
	default:
		t.Error("Passing message to shutdownCh doesn't call shutdown()")
	}
	if b.Len() != 0 {
		t.Error("Passing message to shutdownCh doesn't break out of the Start method's for loop. " +
			"Note that 'break' statements can be used for select and for loops so a label might be " +
			"required to break out the loop.")
	}
}

// 04

func TestStopMethodModule3(t *testing.T) {
	alog := New(nil)
	alog.shutdownCh = make(chan struct{}, 1)
	alog.shutdownCompleteCh = make(chan struct{}, 1)
	alog.shutdownCompleteCh <- struct{}{}
	alog.Stop()
	select {
	case <-alog.shutdownCh:
	default:
		t.Error("Stop() method doesn't send signal to shutdownCh channel")
	}
	select {
	case <-alog.shutdownCompleteCh:
		t.Error("Stop() method doesn't wait for signal from shutdownCompleteCh channel")
	default:
	}
}

// 05

func TestWriteAllBeforeShutdownModule3(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	alog := New(sleepingWriter{b})
	alog.msgCh = make(chan string, 2)
	go alog.Start()
	alog.msgCh <- "first"
	alog.msgCh <- "second"
	time.Sleep(10 * time.Millisecond)
	doneCh := make(chan struct{})
	go func() {
		alog.Stop()
		written := b.String()
		if !strings.Contains(written, "first") || !strings.Contains(written, "second") {
			t.Error("Not all messages written before logger shutdown")
		}
		doneCh <- struct{}{}
	}()
	select {
	case <-time.Tick(1 * time.Second):
		t.Error("Test timed out, please check that the Done method on the wait group is being called in the write method")
	case <-doneCh:
	}
}
