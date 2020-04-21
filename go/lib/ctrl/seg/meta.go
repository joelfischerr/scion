// Copyright 2017 ETH Zurich
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

<<<<<<< HEAD:go/lib/ctrl/seg/meta.go
package seg

import (
	"fmt"

	"github.com/scionproto/scion/go/proto"
)

type Meta struct {
	Type    proto.PathSegType
	Segment *PathSegment `capnp:"pathSeg"`
=======
package scheduler

import (
	"github.com/scionproto/scion/go/border/qos/queues"
	"github.com/scionproto/scion/go/border/rpkt"
)

type SchedulerInterface interface {
	Init(routerConfig queues.InternalRouterConfig)
	Dequeuer(routerConfig queues.InternalRouterConfig, forwarder func(rp *rpkt.RtrPkt))
	dequeue(routerConfig queues.InternalRouterConfig, forwarder func(rp *rpkt.RtrPkt), queueNo int)
	GetMessages() *chan bool
>>>>>>> refactor.:go/border/qos/scheduler/scheduler.go
}

func NewMeta(s *PathSegment, t proto.PathSegType) *Meta {
	return &Meta{
		Type:    t,
		Segment: s,
	}
}

func (m *Meta) String() string {
	return fmt.Sprintf("Type: %v, Segment: %v", m.Type, m.Segment)
}
