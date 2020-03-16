package main

import (
	"fmt"
	"math"
	"sync"
	"testing"
	"time"

	"github.com/scionproto/scion/go/border/rpkt"
	"github.com/scionproto/scion/go/lib/addr"
	"github.com/scionproto/scion/go/lib/common"
)

var packets = make(chan *rpkt.RtrPkt, 100)

/*
Things to do:

1. Set up router with a topology
2. Create a packet

*/

func setupQueue() packetQueue {

	bandwidth := 0
	priority := 1

	bucket := tokenBucket{MaxBandWidth: bandwidth, tokens: bandwidth, lastRefill: time.Now(), mutex: &sync.Mutex{}}
	que := packetQueue{MaxLength: 128, MinBandwidth: priority, MaxBandWidth: priority, mutex: &sync.Mutex{}, tb: bucket}

	return que

}

func setupQueuePaket() qPkt {

	return qPkt{queueNo: 0, rp: nil}
}

func TestBasicEnqueue(t *testing.T) {
	que := setupQueue()
	pkt := setupQueuePaket()
	que.enqueue(&pkt)
	length := que.getLength()
	if length != 1 {
		t.Errorf("Enqueue one packet should give length 1 gave length %d", length)
	}
}

func TestMultiEnqueue(t *testing.T) {
	que := setupQueue()
	pkt := setupQueuePaket()
	j := 15
	i := 0

	for i < j {
		que.enqueue(&pkt)
		i++
	}
	length := que.getLength()

	if length != j {
		t.Errorf("Enqueue one packet should give length %d gave length %d", j, length)
	}
}

func BenchmarkEnqueue(b *testing.B) {
	que := setupQueue()
	pkt := setupQueuePaket()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		que.enqueue(&pkt)
	}
}

func benchmarkPop(popNo int, b *testing.B) {
	que := setupQueue()
	pkt := setupQueuePaket()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		j := 0
		for j < popNo {
			que.enqueue(&pkt)
			j++
		}
		b.StartTimer()
		que.popMultiple(popNo)
	}
}

func BenchmarkPop1(b *testing.B) { benchmarkPop(1, b) }
func BenchmarkPop5(b *testing.B) { benchmarkPop(10, b) }

func TestCallPacketGen(t *testing.T) {
	_ = rpkt.JFPrepareRtrPacketSample(t)
}

func (r *Router) forwardPacketTest(rp *rpkt.RtrPkt) {

	// defer rp.Release()

	packets <- rp

}

func TestHundredPacketQueue(t *testing.T) {

	r, _ := setupTestRouter(t)

	r.initQueueing()
	r.forwarder = r.forwardPacketTest

	ps := make([]*rpkt.RtrPkt, 100)

	packetSource, packetDestination := "1-ff00:0:110", "1-ff00:0:111"

	srcIA, err := addr.IAFromString(packetSource)

	if err != nil {
		fmt.Println(err)
		t.Errorf("")
	}

	dstIA, err := addr.IAFromString(packetDestination)

	if err != nil {
		fmt.Println(err)
		t.Errorf("")
	}

	pkt := rpkt.JFPrepareRtrPacketWith(
		srcIA,
		dstIA,
		common.L4ProtocolType(4))

	for i := 0; i < 100; i++ {
		r.queuePacket(pkt)
		ps[i] = pkt
	}

	time.Sleep(2 * time.Second)

	for i := 0; i < 100; i++ {
		select {
		case returnedPacket := <-packets:
			if returnedPacket != ps[i] {
				t.Errorf("Returned wrong packet!")
			} else {
				fmt.Println("We got the packet back 🥳👯‍♂️👯‍♀️")
			}
		default:
			t.Errorf("Returned no packet!")
		}
	}
}

func checkActionBasicTest(t *testing.T, queueFillLevel int, actionFillLevel int, action int, shouldTakeAction bool) {

	ap1 := actionProfile{actionFillLevel, 100, policeAction(action)}

	que := setupQueue()
	que.Profile = []actionProfile{ap1}
	pkt := setupQueuePaket()
	n := int(float64(que.MaxLength) / float64(100) * float64(queueFillLevel))
	fmt.Println(n)
	for i := 0; i < n; i++ {
		que.enqueue(&pkt)
	}
	fmt.Println(que.getFillLevel())
	fmt.Println(que.getLength())
	if (que.checkAction() == policeAction(action)) != shouldTakeAction {
		t.Errorf("Action is %d but we wanted %d", que.checkAction(), action)
	}

}

func TestCheckActionBasic(t *testing.T) {
	checkActionBasicTest(t, 0, 0, 1, true)
	checkActionBasicTest(t, 0, 50, 1, false)
	checkActionBasicTest(t, 0, 50, 1, false)
	checkActionBasicTest(t, 49, 50, 1, false)
	checkActionBasicTest(t, 50, 50, 1, true)
	checkActionBasicTest(t, 51, 50, 1, true)
	checkActionBasicTest(t, 100, 50, 2, true)
}

func checkActionMultiRuleTest(t *testing.T, queueFillLevel int, act int) {

	ap0 := actionProfile{10, 100, policeAction(0)}
	ap1 := actionProfile{20, 100, policeAction(1)}
	ap2 := actionProfile{30, 100, policeAction(1)}
	ap3 := actionProfile{40, 100, policeAction(1)}
	ap4 := actionProfile{50, 100, policeAction(2)}
	ap5 := actionProfile{60, 100, policeAction(2)}
	ap6 := actionProfile{70, 100, policeAction(2)}
	ap7 := actionProfile{80, 100, policeAction(3)}
	ap8 := actionProfile{90, 100, policeAction(2)}
	ap9 := actionProfile{100, 100, policeAction(3)}

	que := setupQueue()
	que.Profile = []actionProfile{ap0, ap1, ap2, ap3, ap4, ap5, ap6, ap7, ap8, ap9}
	pkt := setupQueuePaket()
	n := int(math.Round(float64(que.MaxLength) / float64(100) * float64(queueFillLevel)))
	fmt.Println(n)
	for i := 0; i < n; i++ {
		que.enqueue(&pkt)
	}
	fmt.Println(que.getFillLevel())
	fmt.Println(que.getLength())
	if que.checkAction() != policeAction(act) {
		t.Errorf("Action is %d but we wanted %d", que.checkAction(), act)
	}

}

func TestCheckActionMultiRule(t *testing.T) {
	checkActionMultiRuleTest(t, 0, 0)
	checkActionMultiRuleTest(t, 10, 0)

	checkActionMultiRuleTest(t, 19, 0)
	checkActionMultiRuleTest(t, 20, 1)
	checkActionMultiRuleTest(t, 21, 1)

	checkActionMultiRuleTest(t, 49, 1)
	checkActionMultiRuleTest(t, 50, 2)
	checkActionMultiRuleTest(t, 51, 2)

	checkActionMultiRuleTest(t, 81, 3)
	checkActionMultiRuleTest(t, 85, 3)
	checkActionMultiRuleTest(t, 99, 2)
	checkActionMultiRuleTest(t, 100, 3)
}
