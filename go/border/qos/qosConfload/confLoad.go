package qosconfload

import (
	"github.com/scionproto/scion/go/lib/addr"
	"github.com/scionproto/scion/go/lib/common"
	"github.com/scionproto/scion/go/lib/l4"
)

type RpktInterface interface {
	SrcIA() (addr.IA, error)
	DstIA() (addr.IA, error)
	Bytes() common.RawBytes
	Release()
	L4Hdr(verify bool) (l4.L4Header, error)
}

type QosConfigInterface interface{}
