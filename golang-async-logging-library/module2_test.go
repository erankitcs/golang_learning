package alog

import (
	"bytes"
	"errors"
	"reflect"
	"regexp"
	"sync"
	"testing"
	"time"
)

const messageTimestampPattern = `\[\d{4}-\d{2}-\d{2}\ \d{2}:\d{2}:\d{2}] - `

// 01
func TestMessageChannelModule2(t *testing.T) {
	alog := New(nil)
	if alog.msgCh == nil {
		t.Fatal("msgCh field not initialized. Should have type 'chan string' but it is currently nil")
	}
}

// 02
func TestErrorChannelModule2(t *testing.T) {
	alog := New(nil)
	if alog.errorCh == nil {
		t.Fatal("errorCh field not initialized. Should have type 'chan error' but it is currently nil")
	}
}

// 03
func TestMessageChannelMethodModule2(t *testing.T) {
	alog := New(nil)
	if alog.MessageChannel() != alog.msgCh {
		t.Fatal("MessageChannel method does not return the msgCh field")
	}
	messageChannelDir := reflect.ValueOf(alog.MessageChannel()).Type().ChanDir()
	if messageChannelDir != reflect.SendDir {
		t.Fatal("MessageChannel does not return send-only channel")
	}
}

// 04
func TestErrorChannelMethodModule2(t *testing.T) {
	alog := New(nil)
	if alog.ErrorChannel() != alog.errorCh {
		t.Fatal("ErrorChannel method does not return the errorCh field")
	}
	errorChannelDir := reflect.ValueOf(alog.ErrorChannel()).Type().ChanDir()
	if errorChannelDir != reflect.RecvDir {
		t.Fatal("ErrorChannel does not return receive-only channel")
	}
}

// 05
func TestWritesToWriterModule2(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	alog := New(b)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	alog.write("test", wg)

	written := b.String()
	if written == "" {
		t.Fatal("Nothing written to log")
	}
	if !regexp.MustCompile(messageTimestampPattern + "test\n$").Match([]byte(written)) {
		t.Error("Properly formatted string not written to log. Did you pass the message to 'formatMessage'?")
	}
}

// 06
type errorWriter struct {
	b *bytes.Buffer
}

func (ew errorWriter) Write(data []byte) (int, error) {
	ew.b.Write(data)
	return 0, errors.New("error")
}
func TestWriteSendsErrorsToErrorChannelModule2(t *testing.T) {
	alog := New(&errorWriter{bytes.NewBuffer([]byte{})})
	alog.errorCh = make(chan error, 1)
	wg := &sync.WaitGroup{}
	wg.Add(1)
	alog.write("test", wg)
	go func() {
		if (<-alog.errorCh).Error() != "error" {
			t.Fatal("Did not receive destination writer's error on errorCh")
		}
	}()
	time.Sleep(100 * time.Millisecond)
	alog.errorCh <- errors.New("")
	time.Sleep(100 * time.Millisecond)
}

// 07
type sleepingWriter struct {
	b *bytes.Buffer
}

func (sw sleepingWriter) Write(data []byte) (int, error) {
	sw.b.Write(data)
	time.Sleep(500 * time.Millisecond)
	sw.b.WriteString("write complete")
	return 0, nil
}

func TestStartHandlesMessagesModule2(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	alog := New(sleepingWriter{b})
	alog.msgCh = make(chan string, 2)
	go alog.Start()
	alog.msgCh <- "test message"
	time.Sleep(100 * time.Millisecond)
	written := b.Bytes()
	if !regexp.MustCompile(messageTimestampPattern + "test message\n$").Match(written) {
		t.Error("Message not written to logger's destination")
	}
	shouldRelock := false
	if alog.m != nil {
		mutexState := reflect.ValueOf(*alog.m).FieldByName("state").Int()
		if mutexState != 0 {
			alog.m.Unlock()
			shouldRelock = true
		}
	}
	alog.msgCh <- "second message"
	time.Sleep(100 * time.Millisecond)
	if alog.m != nil {
		if shouldRelock {
			alog.m.Lock()
		}
	}
	written = b.Bytes()
	if !regexp.MustCompile(messageTimestampPattern + "test message\n" + messageTimestampPattern + "second message\n").Match(written) {
		t.Error("write method not called as a goroutine")
	}
}

// 08
type panickingWriter struct {
	b *bytes.Buffer
}

func (pw panickingWriter) Write(data []byte) (int, error) {
	pw.b.Write(data)
	panic("panicking!")
}
func TestWriteSendsWriteRequestsSequentiallyModule2(t *testing.T) {
	b := bytes.NewBuffer([]byte{})
	alog := New(sleepingWriter{b})
	if alog.m == nil {
		t.Fatal("Alog's mutex field 'm' not initialized")
	}
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go alog.write("test message", wg)
	time.Sleep(100 * time.Millisecond)
	go alog.write("second message", wg)
	time.Sleep(1000 * time.Millisecond)
	written := b.Bytes()
	if !regexp.MustCompile(messageTimestampPattern + "test message\nwrite complete" + messageTimestampPattern + "second message\n").Match(written) {
		t.Error("Mutex not protecting Alog.dest#Write from concurrent calls")
	}

	b = bytes.NewBuffer([]byte{})
	alog = New(panickingWriter{b})
	if alog.msgCh == nil {
		t.Fatal("msgCh field is nil")
	}
	if alog.m == nil {
		t.Fatal("mutex field 'm' is nil")
	}
	go func() {
		defer func() {
			recover()
		}()
		alog.write("test message", wg)
	}()
	time.Sleep(100 * time.Millisecond)
	go func() {
		defer func() {
			recover()
		}()
		alog.write("second message", wg)
	}()
	time.Sleep(1000 * time.Millisecond)
	written = b.Bytes()
	if !regexp.MustCompile(messageTimestampPattern + "test message\n" + messageTimestampPattern + "second message\n").Match(written) {
		t.Error("Mutex not unlocked when panicking")
	}
}

// 09
func TestWriteSendsErrorsAsynchronouslyModule2(t *testing.T) {
	TestWriteSendsWriteRequestsSequentiallyModule2(t)
	b := bytes.NewBuffer([]byte{})
	alog := New(&errorWriter{b})
	wg := &sync.WaitGroup{}
	wg.Add(2)
	go alog.write("first", wg)
	time.Sleep(100 * time.Millisecond)
	go alog.write("second", wg)
	time.Sleep(100 * time.Millisecond)
	written := b.Bytes()
	if !regexp.MustCompile(`.*first.*\n.*second.*`).Match(written) {
		t.Fatal("Error messages not sent to error channel asynchronously")
	}
	errorReceived := false
	go func() {
		<-alog.errorCh
		errorReceived = true
	}()
	time.Sleep(100 * time.Millisecond)
	if !errorReceived {
		t.Fatal("Error messages not sent to error channel")
	}
}
