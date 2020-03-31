// Copyright 2020 ETH Zurich
// Copyright 2018 ETH Zurich, Anapaya Systems
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
	qosconfload "github.com/scionproto/scion/go/border/qos/qosConfload"
	"github.com/scionproto/scion/go/border/qos/qosqueues"
)

type SchedulerInterface interface {
	Init(routerConfig qosqueues.InternalRouterConfig)
	Dequeuer(routerConfig qosqueues.InternalRouterConfig, forwarder func(rp qosconfload.RpktInterface))
	dequeue(routerConfig qosqueues.InternalRouterConfig, forwarder func(rp qosconfload.RpktInterface), queueNo int)
	GetMessages() *chan bool
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
