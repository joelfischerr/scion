package main

import (
	"github.com/scionproto/scion/go/lib/log"
	// "fmt"
)

var (
	mTicker = 0
)

type Manager struct {
	// confDir is the directory containing the configuration file.
	confDir string
	rules []string
	queues []RouterQueue
}

func initManager() () {

}

func (manag *Manager) managerProcess() () {
	// Process an incoming packet and enqueue it to a queue
	log.Debug("Manager: Called manager")
	
	for i, rule := range manag.rules {
		_ = i
		_ = rule
	}
}

func (manag *Manager) preRoute() () {

	// for i, queue := range manag.queues {
	// 	_ = i
	// 	_ = queue
	// 	if i % queue.priority == 0 {
	// 		rp = &queue.dequeue()
	// 	}

		
	// }
}

// func route (rp *rpkt.RtrPkt) () {

// }