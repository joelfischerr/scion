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

package main

import (
	"sync"
	"time"

	"github.com/scionproto/scion/go/lib/log"
)

type tokenBucket struct {
	MaxBandWidth     int // In bps
	tokens           int // One token is 1 b
	tokenSpent       int
	timerGranularity int
	lastRefill       time.Time
	mutex            *sync.Mutex
}

func (tb *tokenBucket) refill(shouldLog bool) {

	now := time.Now()

	timeSinceLastUpdate := now.Sub(tb.lastRefill).Milliseconds()

	if shouldLog {
		log.Debug("Last update was", "ms ago", timeSinceLastUpdate)
	}

	if timeSinceLastUpdate > 100 {

		newTokens := ((tb.MaxBandWidth) * int(timeSinceLastUpdate)) / (1000)
		tb.lastRefill = now

		if shouldLog {
			log.Debug("Add new tokens", "#tokens", newTokens)
			log.Debug("On Update: Spent token in last period", "#tokens", tb.tokenSpent)
		}

		tb.tokenSpent = 0

		if tb.tokens+newTokens > tb.MaxBandWidth {
			tb.tokens = tb.MaxBandWidth
		} else {
			tb.tokens = tb.tokens + newTokens
		}
	}

}

func (pq *packetQueue) police(qp *qPkt, shouldLog bool) policeAction {
	pq.tb.mutex.Lock()
	defer pq.tb.mutex.Unlock()

	packetSize := (qp.rp.Bytes().Len()) // In byte

	tokenForPacket := packetSize * 8 // In bit

	if shouldLog {
		log.Debug("Overall available bandwidth per second", "MaxBandWidth", pq.tb.MaxBandWidth)
		log.Debug("Spent token in last period", "#tokens", pq.tb.tokenSpent)
		log.Debug("Available bandwidth before refill", "bandwidth", pq.tb.tokens)
	}

	pq.tb.refill(shouldLog)

	if shouldLog {
		log.Debug("Available bandwidth after refill", "bandwidth", pq.tb.tokens)
		log.Debug("Tokens necessary for packet", "tokens", tokenForPacket)
		log.Debug("Tokens necessary for packet", "bytes", qp.rp.Bytes().Len())
	}

	if pq.tb.tokens-tokenForPacket > 0 {
		pq.tb.tokens = pq.tb.tokens - tokenForPacket
		pq.tb.tokenSpent += tokenForPacket
		qp.act.action = PASS
		qp.act.reason = None
	} else {
		qp.act.action = DROP
		qp.act.reason = BandWidthExceeded
	}

	if shouldLog {
		log.Debug("Available bandwidth after update", "bandwidth", pq.tb.tokens)
	}

	return qp.act.action
}

func (qp *qPkt) sendNotification() {
	select {
	case r.notifications <- qp:
	default:
	}
}