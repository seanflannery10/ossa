package server

import (
	"github.com/seanflannery10/ossa/assert"
	"os"
	"syscall"
	"testing"
	"time"
)

func TestNew(t *testing.T) {
	srv := New("test", nil)

	assert.Equal(t, srv.Addr, "test")
	assert.SameType(t, srv, &Server{})
}

func TestServer_Run(t *testing.T) {
	t.Run("SIGINT", func(t *testing.T) {
		srv := New("localhost:4444", nil)

		go func() {
			time.Sleep(250 * time.Millisecond)

			p, err := os.FindProcess(os.Getpid())
			if err != nil {
				panic(err)
			}
			p.Signal(syscall.SIGINT)
		}()

		err := srv.Run()
		if err != nil {
			t.Fatal(err)
		}
	})

	t.Run("SIGTERM", func(t *testing.T) {
		srv := New("localhost:4444", nil)

		go func() {
			time.Sleep(250 * time.Millisecond)

			p, err := os.FindProcess(os.Getpid())
			if err != nil {
				panic(err)
			}
			p.Signal(syscall.SIGTERM)
		}()

		err := srv.Run()
		if err != nil {
			t.Fatal(err)
		}
	})
}
