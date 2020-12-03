package main

import (
	"encoding/asn1"
	"fmt"
	"math"

	cb "github.com/hyperledger/fabric-protos-go/common"
	"github.com/hyperledger/fabric/common/util"
)

type asn1Header struct {
	Number       int64
	PreviousHash []byte
	DataHash     []byte
}

// ComputeHeaderSHA2 returns SHA2-256 on BlockHeader
func ComputeHeaderSHA2(bh *cb.BlockHeader) []byte {
	return util.ComputeSHA256(headerBytes(bh))
}

// ComputeDataSHA256 returns SHA2-256 on BlockData
func ComputeDataSHA256(bd *cb.BlockData) []byte {
	return util.ComputeSHA256(dataBytes(bd))
}

func headerBytes(b *cb.BlockHeader) []byte {
	asn1Header := asn1Header{
		PreviousHash: b.PreviousHash,
		DataHash:     b.DataHash,
	}
	if b.Number > uint64(math.MaxInt64) {
		panic(fmt.Errorf("Golang does not currently support encoding uint64 to asn1"))
	} else {
		asn1Header.Number = int64(b.Number)
	}
	result, err := asn1.Marshal(asn1Header)
	if err != nil {
		panic(err)
	}
	return result
}

func dataBytes(bd *cb.BlockData) []byte {
	return util.ConcatenateBytes(bd.Data...)
}
