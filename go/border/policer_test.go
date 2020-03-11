package main

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestBasicRefill(t *testing.T) {
	bucket := tokenBucket{MaxBandWidth: 5 * 1024, tokens: 0, lastRefill: time.Now(), mutex: &sync.Mutex{}}

	fmt.Println(bucket)
	bucket.refill(false)
	fmt.Println(bucket)
	if bucket.tokens != 0 {
		t.Errorf("Got %d, want %d", bucket.tokens, 0)
	}
}

func timedRefill(t *testing.T, maxBandwidth int, waitTime float32) {
	bucket := tokenBucket{MaxBandWidth: maxBandwidth, tokens: 0, lastRefill: time.Now(), mutex: &sync.Mutex{}}

	fmt.Println(bucket)
	time.Sleep(time.Millisecond * time.Duration(waitTime*1000))
	bucket.refill(true)
	fmt.Println(bucket)
	res := int(float32(maxBandwidth) * waitTime)

	if waitTime < 0.1 {
		res = 0
	}

	fmt.Println("Result should be")
	fmt.Println(res)
	if bucket.tokens != min(maxBandwidth, res) {
		t.Errorf("Got %d, want %d", bucket.tokens, min(maxBandwidth, res))
	}
}

func TestTimedRefill(t *testing.T) {

	timedRefill(t, 5*1024, 0.05)
	timedRefill(t, 5*1024, 0.5)
	timedRefill(t, 5*1024, 1)
	timedRefill(t, 5*1024, 2)

	timedRefill(t, 500*1024, 0.05)
	timedRefill(t, 500*1024, 0.5)
	timedRefill(t, 500*1024, 1)
	timedRefill(t, 500*1024, 2)
}

func BenchmarkBasicRefill(b *testing.B) {

	bucket := tokenBucket{MaxBandWidth: 5 * 1024, tokens: 0, lastRefill: time.Now(), mutex: &sync.Mutex{}}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		for j := 0; j < 10; i++ {
			bucket.refill(false)
			b.StopTimer()
			time.Sleep(time.Millisecond * 100)
			b.StartTimer()
		}
	}
}
