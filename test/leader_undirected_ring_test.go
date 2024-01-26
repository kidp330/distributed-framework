package test

import (
	"io/ioutil"
	"log"
	"testing"

	"github.com/krzysztof-turowski/distributed-framework/leader/undirected_ring/sync_higham_przytycka"
)

func TestUndirectedRingHighamPrzytycka(t *testing.T) {
	log.SetOutput(ioutil.Discard)
	for n := 2; n <= 100; n++ {
		sync_higham_przytycka.Run(n)
	}
}

func BenchmarkUndirectedRingHighamPrzytycka(b *testing.B) {
	log.SetOutput(ioutil.Discard)
	for i := 0; i < b.N; i++ {
		sync_higham_przytycka.Run(100)
	}
}
