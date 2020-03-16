package main

import (
	"fmt"
	"testing"

	"github.com/scionproto/scion/go/border/rpkt"
	"github.com/scionproto/scion/go/lib/addr"
	"github.com/scionproto/scion/go/lib/common"
)

func TestContains(t *testing.T) {
	if contains([]int{0, 1, 2, 3, 4}, 4) == false {
		t.Errorf("Error contains 4 but did not return true")
	}
	if contains([]int{0, 1, 2, 3, 4, 5, 6}, 4) == false {
		t.Errorf("Error contains 4 but did not return true")
	}
	if contains([]int{4, 1, 2, 3, 4}, 4) == false {
		t.Errorf("Error contains 4 but did not return true")
	}
	if contains([]int{4}, 4) == false {
		t.Errorf("Error contains 4 but did not return true")
	}
	if contains([]int{}, 4) == true {
		t.Errorf("Error: Does not contain 4 but did not return false")
	}
	if contains([]int{-1}, 4) == true {
		t.Errorf("Error: Does not contain 4 but did not return false")
	}
}

func matchRule(t *testing.T, ruleSource string, ruleDestination string, packetSource string, packetDestination string, ruleL4type int, packetL4Type int, shouldMatch bool) {
	cr := classRule{
		Name:          "Test",
		SourceAs:      ruleSource,
		DestinationAs: ruleDestination,
		L4Type:        []int{ruleL4type},
		QueueNumber:   0}

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
		common.L4ProtocolType(packetL4Type))

	if cr.matchRule(pkt) != shouldMatch {
		t.Errorf("Rule should match %t but did not", shouldMatch)
	}

}

func TestMatchRule(t *testing.T) {

	matchRule(t, "1-ff00:0:110", "1-ff00:0:111",
		"1-ff00:0:110", "1-ff00:0:111", 1, 1, true)

	matchRule(t, "1-ff00:0:110", "1-ff00:0:111",
		"1-ff00:0:110", "1-ff00:0:111", 0, 1, false)

	matchRule(t, "1-ff00:0:110", "1-ff00:0:111",
		"1-ff00:0:110", "1-ff00:0:111", 1, 0, false)

	matchRule(t, "1-ff00:0:111", "1-ff00:0:111",
		"1-ff00:0:110", "1-ff00:0:111", 1, 1, false)
}
