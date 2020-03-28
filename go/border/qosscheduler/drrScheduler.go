// Copyright 2020 ETH Zurich
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package qosscheduler

import (
	"github.com/scionproto/scion/go/border/qosqueues"
	"github.com/scionproto/scion/go/border/rpkt"
	"github.com/scionproto/scion/go/lib/log"
)

// This is a deficit round robin dequeuer. Queues with higher priority will have more packets dequeued at the same time.

type deficitRoundRobinScheduler struct {
	quantumSum  int
	totalLength int
}

var _ SchedulerInterface = (*deficitRoundRobinScheduler)(nil)

func (sched *deficitRoundRobinScheduler) Init(routerConfig qosqueues.InternalRouterConfig) {

	sched.quantumSum = 0
	sched.totalLength = len(routerConfig.Queues)

	for i := 0; i < sched.totalLength; i++ {
		sched.quantumSum = sched.quantumSum + routerConfig.Queues[i].GetPriority()
	}

}

func (sched *deficitRoundRobinScheduler) dequeue(routerConfig qosqueues.InternalRouterConfig, forwarder func(rp *rpkt.RtrPkt), queueNo int) {

	length := routerConfig.Queues[queueNo].GetLength()
	var nopkts int = 64 * (routerConfig.Queues[queueNo].GetPriority() / sched.quantumSum)
	pktToDequeue := min(1, nopkts)

	log.Debug("The queue has length", "length", length)
	log.Debug("Dequeueing packets", "quantum", pktToDequeue)

	if length > 0 {
		qps := routerConfig.Queues[queueNo].PopMultiple(max(length, pktToDequeue))
		for _, qp := range qps {
			forwarder(qp.Rp)
		}
	}
}

func (sched *deficitRoundRobinScheduler) Dequeuer(routerConfig qosqueues.InternalRouterConfig, forwarder func(rp *rpkt.RtrPkt)) {

	for {
		for i := 0; i < sched.totalLength; i++ {
			sched.dequeue(routerConfig, forwarder, i)
		}
	}
}
