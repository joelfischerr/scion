package main

import (
	"fmt"
	"testing"
	"time"

	"github.com/scionproto/scion/go/border/rpkt"
	"github.com/scionproto/scion/go/lib/addr"
	"github.com/scionproto/scion/go/lib/common"
)

var canDequeuePacket = make(chan bool, 1)

func TestLoadSampleConfig(t *testing.T) {
	r, _ := setupTestRouter(t)

	r.loadConfigFile("sample-config.yaml")

	fmt.Println("The Queue is: ", r.config.Queues[0])
	fmt.Println("The Rule is: ", r.config.Rules[0])
	// t.Errorf("Output: %v", r.config)

}

func TestLoadSampleConfigQueues(t *testing.T) {
	r, _ := setupTestRouter(t)

	r.loadConfigFile("sample-config.yaml")

	fmt.Println("The Queue is: ", r.config.Queues[0])
	fmt.Println("The Rule is: ", r.config.Rules[0])
	fmt.Println("We have this number of queues: ", len(r.config.Queues))
	t.Errorf("Output: %v", r.config)

	if r.config.Queues[0].ID != 0 {
		t.Errorf("Incorrect Queue ID")
	}
	if r.config.Queues[1].ID != 1 {
		t.Errorf("Incorrect Queue ID")
	}

	if r.config.Queues[0].Name != "General Queue" {
		t.Errorf("Incorrect Queue Name is %v but should be %v", r.config.Queues[0].Name, "General Queue")
	}
	if r.config.Queues[1].Name != "Speedy Queue" {
		t.Errorf("Incorrect Queue Name is %v but should be %v", r.config.Queues[0].Name, "Speedy Queue")
	}
}

func TestQueueSinglePacket(t *testing.T) {
	r, _ := setupTestRouter(t)

	r.loadConfigFile("sample-config.yaml")

	r.initQueueing()
	r.forwarder = r.forwardPacketWithBarrierTest

	fmt.Println("The Queue is: ", r.config.Queues[0])
	fmt.Println("The Queue is: ", r.config.Queues[1])
	fmt.Println("The Queue is: ", r.config.Queues[2])
	fmt.Println("The Rule is: ", r.config.Rules[0])
	fmt.Println("The Rule is: ", r.config.Rules[1])
	fmt.Println("We have this number of queues: ", len(r.config.Queues))

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
		common.L4ProtocolType(1))

	r.queuePacket(pkt)

	fmt.Println("Queue 0 has length ", r.config.Queues[0].getLength())
	fmt.Println("Queue 1 has length ", r.config.Queues[1].getLength())
	fmt.Println("Queue 2 has length ", r.config.Queues[2].getLength())
	if r.config.Queues[2].getLength() != 1 {
		t.Errorf("Queue has incorrect length")
	}

	canDequeuePacket <- true

	<-packets

	fmt.Println("Queue  has length ", r.config.Queues[2].getLength())
	if r.config.Queues[2].getLength() != 0 {
		t.Errorf("Queue has incorrect length")
	}

	// t.Errorf("Incorrect Queue ID")
}

func checkQueueLength(t *testing.T, r *Router, noPacketsPerQueue []int) {

	for i := 0; i < len(noPacketsPerQueue); i++ {
		fmt.Println("Queue", i, "has length", r.config.Queues[i].getLength())
	}
	for i := 0; i < len(noPacketsPerQueue); i++ {
		if r.config.Queues[i].getLength() != noPacketsPerQueue[i] {
			t.Errorf("Queue %d has incorrect length %d instead of %d", i, r.config.Queues[i].getLength(), noPacketsPerQueue[i])
		}
	}
}

func allowDequeuePackets(r *Router, noPackets int) {

	// Unblock forwarder from previous packet
	// Note that initially always one packet needs to be sent before testing is done
	// otherwise the forwarder does not block correctly
	canDequeuePacket <- true
	<-packets

	// Dequeue the first blocking packet
	fmt.Println("Dequeue Queue 1 only")
	for i := 0; i < (noPackets - 1); i++ {
		canDequeuePacket <- true
		<-packets
	}

}

func checkDequeuePackets(t *testing.T, r *Router, precondition []int, queueNo int, noPackets int) {

	postcondition := precondition
	postcondition[queueNo] = precondition[queueNo] - noPackets
	allowDequeuePackets(r, noPackets)
	checkQueueLength(t, r, postcondition)
}

func checkForwarderForLeftovers(t *testing.T) {
	// There will be one last packet stuck in the forwarder
	// This unblocks the forwarder and we will be good
	// If a packet forwarded too early this will time out
	canDequeuePacket <- true
	<-packets

	// This checks whether there are any packets left.
	// At this point there should not be any packets left
	canDequeuePacket <- true
	select {
	case _ = <-packets:
		t.Errorf("There are too many packets on the queue")
	default:
		fmt.Println("Should pass 🥳")
	}
}

func (r *Router) forwardPacketWithBarrierTest(rp *rpkt.RtrPkt) {

	// defer rp.Release()

	fmt.Println("Queued Packet!")

	<-canDequeuePacket
	fmt.Println("We can now dequeue the packet")

	packets <- rp

}

func TestQueueMultPacketsBasic(t *testing.T) {
	r, _ := setupTestRouter(t)

	r.loadConfigFile("sample-config.yaml")
	r.initQueueing()
	r.forwarder = r.forwardPacketWithBarrierTest

	noPacketsPerClass := [3]int{50, 30, 20}

	fmt.Println("The Queue is: ", r.config.Queues[0])
	fmt.Println("The Queue is: ", r.config.Queues[1])
	fmt.Println("The Queue is: ", r.config.Queues[2])
	fmt.Println("The Rule is: ", r.config.Rules[0])
	fmt.Println("The Rule is: ", r.config.Rules[1])
	fmt.Println("We have this number of queues: ", len(r.config.Queues))

	packetSource1, packetDestination1 := "1-ff00:0:110", "1-ff00:0:111"
	packetSource2, packetDestination2 := "1-ff00:0:111", "1-ff00:0:112"
	packetSource3, packetDestination3 := "1-ff00:0:112", "1-ff00:0:110"

	sources := [3]string{packetSource1, packetSource2, packetSource3}
	dests := [3]string{packetDestination1, packetDestination2, packetDestination3}
	L4Types := [3]int{1, 1, 1}

	// This packet will block forwarding so that we can check the queue
	var pkt = rpkt.JFPrepareRtrPacketWithSrings(sources[0], dests[0], L4Types[0])
	r.queuePacket(pkt)
	time.Sleep(time.Millisecond * 50)

	for i := 0; i < len(noPacketsPerClass); i++ {
		for j := 0; j < noPacketsPerClass[i]; j++ {
			pkt := rpkt.JFPrepareRtrPacketWithSrings(sources[i], dests[i], L4Types[i])
			r.queuePacket(pkt)
		}
	}
	// I want to find a way to block the dequeueing
	// But most likely I will need to test enqueue and dequeue seperately and lock the queue inbetween.
	// Enqueue will be tested by just checking the packets coming out of the forwarder
	// Dequeue also needs to test the order in which packets are dequeued, therefore lock and manually enqueue and then remove the locks and check the order in which packets arrive.
	// Does this need some consistency models?

	checkQueueLength(t, r, []int{50, 0, 50})

	allowDequeuePackets(r, 50)

	checkQueueLength(t, r, []int{0, 0, 50})

	allowDequeuePackets(r, 50)

	checkQueueLength(t, r, []int{0, 0, 0})

	checkForwarderForLeftovers(t)

	// t.Errorf("Show Log")
}

func TestQueueMultPacketsEnqueueTwice(t *testing.T) {
	r, _ := setupTestRouter(t)

	r.loadConfigFile("sample-config.yaml")
	r.initQueueing()
	r.forwarder = r.forwardPacketWithBarrierTest

	noPacketsPerClass := [3]int{50, 30, 20}

	fmt.Println("The Queue is: ", r.config.Queues[0])
	fmt.Println("The Queue is: ", r.config.Queues[1])
	fmt.Println("The Queue is: ", r.config.Queues[2])
	fmt.Println("The Rule is: ", r.config.Rules[0])
	fmt.Println("The Rule is: ", r.config.Rules[1])
	fmt.Println("We have this number of queues: ", len(r.config.Queues))

	packetSource1, packetDestination1 := "1-ff00:0:110", "1-ff00:0:111"
	packetSource2, packetDestination2 := "1-ff00:0:111", "1-ff00:0:112"
	packetSource3, packetDestination3 := "1-ff00:0:112", "1-ff00:0:110"

	sources := [3]string{packetSource1, packetSource2, packetSource3}
	dests := [3]string{packetDestination1, packetDestination2, packetDestination3}
	L4Types := [3]int{1, 1, 1}

	// This packet will block forwarding so that we can check the queue
	var pkt = rpkt.JFPrepareRtrPacketWithSrings(sources[0], dests[0], L4Types[0])
	r.queuePacket(pkt)
	time.Sleep(time.Millisecond * 50)

	for i := 0; i < len(noPacketsPerClass); i++ {
		for j := 0; j < noPacketsPerClass[i]; j++ {
			pkt := rpkt.JFPrepareRtrPacketWithSrings(sources[i], dests[i], L4Types[i])
			r.queuePacket(pkt)
		}
	}
	// I want to find a way to block the dequeueing
	// But most likely I will need to test enqueue and dequeue seperately and lock the queue inbetween.
	// Enqueue will be tested by just checking the packets coming out of the forwarder
	// Dequeue also needs to test the order in which packets are dequeued, therefore lock and manually enqueue and then remove the locks and check the order in which packets arrive.
	// Does this need some consistency models?

	fmt.Println("Queuecheck 1")
	checkQueueLength(t, r, []int{50, 0, 50})

	allowDequeuePackets(r, 50)

	fmt.Println("Queuecheck 2")
	checkQueueLength(t, r, []int{0, 0, 50})

	for j := 0; j < 100; j++ {
		pkt := rpkt.JFPrepareRtrPacketWithSrings(sources[2], dests[2], L4Types[2])
		r.queuePacket(pkt)
	}

	fmt.Println("Queuecheck 2")
	checkQueueLength(t, r, []int{100, 0, 50})

	allowDequeuePackets(r, 50)

	checkQueueLength(t, r, []int{100, 0, 0})

	allowDequeuePackets(r, 100)

	checkQueueLength(t, r, []int{0, 0, 0})

	for j := 0; j < 2000; j++ {
		pkt := rpkt.JFPrepareRtrPacketWithSrings(sources[0], dests[0], L4Types[0])
		r.queuePacket(pkt)
	}

	checkQueueLength(t, r, []int{0, 0, 1024})

	allowDequeuePackets(r, 1024)

	checkForwarderForLeftovers(t)

	// t.Errorf("Show Log")
}

func TestOptionsTuut(t *testing.T) {
	testOptions()
	// t.Errorf("Show Log")
}

func TestQueueSingleDequeue(t *testing.T) {
	r, _ := setupTestRouter(t)

	r.loadConfigFile("sample-config.yaml")
	r.initQueueing()
	r.forwarder = r.forwardPacketWithBarrierTest

	noPacketsPerClass := [3]int{50, 30, 20}

	fmt.Println("The Queue is: ", r.config.Queues[0])
	fmt.Println("The Queue is: ", r.config.Queues[1])
	fmt.Println("The Queue is: ", r.config.Queues[2])
	fmt.Println("The Rule is: ", r.config.Rules[0])
	fmt.Println("The Rule is: ", r.config.Rules[1])
	fmt.Println("We have this number of queues: ", len(r.config.Queues))

	packetSource1, packetDestination1 := "1-ff00:0:110", "1-ff00:0:111"
	packetSource2, packetDestination2 := "1-ff00:0:111", "1-ff00:0:112"
	packetSource3, packetDestination3 := "1-ff00:0:112", "1-ff00:0:110"

	sources := [3]string{packetSource1, packetSource2, packetSource3}
	dests := [3]string{packetDestination1, packetDestination2, packetDestination3}
	L4Types := [3]int{1, 1, 1}

	// This packet will block forwarding so that we can check the queue
	var pkt = rpkt.JFPrepareRtrPacketWithSrings(sources[0], dests[0], L4Types[0])
	r.queuePacket(pkt)
	time.Sleep(time.Millisecond * 50)

	for i := 0; i < len(noPacketsPerClass); i++ {
		for j := 0; j < noPacketsPerClass[i]; j++ {
			pkt := rpkt.JFPrepareRtrPacketWithSrings(sources[i], dests[i], L4Types[i])
			r.queuePacket(pkt)
		}
	}

	checkQueueLength(t, r, []int{50, 0, 50})

	checkDequeuePackets(t, r, []int{50, 0, 50}, 0, 1)

	for i := 0; i < 49; i++ {
		checkDequeuePackets(t, r, []int{49 - i, 0, 50}, 0, 1)
	}

	checkDequeuePackets(t, r, []int{50, 0, 0}, 0, 50)

	checkForwarderForLeftovers(t)

	// t.Errorf("Show Log")
}

func BenchmarkQueueSingleDequeue(b *testing.B) {
	r := Router{}

	r.loadConfigFile("sample-config.yaml")
	r.initQueueing()
	r.forwarder = r.forwardPacketWithBarrierTest

	noPacketsPerClass := [3]int{1, 1, 1}
	_ = noPacketsPerClass

	fmt.Println("The Queue is: ", r.config.Queues[0])
	fmt.Println("The Queue is: ", r.config.Queues[1])
	fmt.Println("The Queue is: ", r.config.Queues[2])
	fmt.Println("The Rule is: ", r.config.Rules[0])
	fmt.Println("The Rule is: ", r.config.Rules[1])
	fmt.Println("We have this number of queues: ", len(r.config.Queues))

	packetSource1, packetDestination1 := "1-ff00:0:110", "1-ff00:0:111"
	packetSource2, packetDestination2 := "1-ff00:0:111", "1-ff00:0:112"
	packetSource3, packetDestination3 := "1-ff00:0:112", "1-ff00:0:110"

	sources := [3]string{packetSource1, packetSource2, packetSource3}
	dests := [3]string{packetDestination1, packetDestination2, packetDestination3}
	L4Types := [3]int{1, 1, 1}

	// This packet will block forwarding so that we can check the queue
	var pkt = rpkt.JFPrepareRtrPacketWithSrings(sources[0], dests[0], L4Types[0])
	r.queuePacket(pkt)
	time.Sleep(time.Millisecond * 50)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		r.queuePacket(pkt)
		allowDequeuePackets(&r, 1)
	}

	// t.Errorf("Show Log")
}
