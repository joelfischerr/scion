package main

import (
	// "fmt"
	"github.com/scionproto/scion/go/border/rpkt"
	// "github.com/scionproto/scion/go/lib/log"
	// "github.com/scionproto/scion/go/lib/ringbuf"
)

var (
	henry int
)

type RouterQueue struct {
	priority int
	length int
	maxLength int
	queue []rpkt.RtrPkt
}

func (rq *RouterQueue) enqueue(rp *rpkt.RtrPkt) () {

	if rq.length < rq.maxLength {
		rq.queue = append(rq.queue, *rp)
		rq.length++
	}

}

func (rq *RouterQueue) dequeue() (rp *rpkt.RtrPkt) {

	return nil
}

// func peek() () {

// }