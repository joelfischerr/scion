package main

import (
	"fmt"
	"testing"

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

func (r *Router) forwardPacketWithBarrierTest(rp *rpkt.RtrPkt) {

	// defer rp.Release()

	<-canDequeuePacket
	fmt.Println("We can now dequeue the packet")

	packets <- rp

}
