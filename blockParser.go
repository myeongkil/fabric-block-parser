package main

import (
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/golang/protobuf/proto"
	cb "github.com/hyperledger/fabric-protos-go/common"
)

func usage() {
	fmt.Println(os.Args[0] + " cmd file")
	fmt.Println("Commands:")
	fmt.Println("      info print block information")
	fmt.Println("      hash calculate current hash")
	fmt.Println("")
	fmt.Println("ex: " + os.Args[0] + " info block/test.block")
}

func info(b *cb.Block) {
	fmt.Printf("CurrentHash : %s\n", base64.StdEncoding.EncodeToString(ComputeHeaderSHA2(b.Header)))
	fmt.Printf("Header: \n")
	fmt.Printf("    Number: %d\n", b.Header.Number)
	fmt.Printf("    PreviousHash: %s\n", base64.StdEncoding.EncodeToString(b.Header.PreviousHash))
	fmt.Printf("    DataHash: %s\n", base64.StdEncoding.EncodeToString(b.Header.DataHash))
	fmt.Printf("Data: \n")
	for i, d := range b.Data.Data {
		fmt.Printf("    Transaction %d \n", i)
		env := &cb.Envelope{}
		proto.Unmarshal(d, env)
		payload := &cb.Payload{}
		_ = proto.Unmarshal(env.Payload, payload)
		shdr := &cb.SignatureHeader{}
		_ = proto.Unmarshal(payload.Header.SignatureHeader, shdr)
		chdr := &cb.ChannelHeader{}
		_ = proto.Unmarshal(payload.Header.ChannelHeader, chdr)
		fmt.Printf("        Signature: %s\n", base64.StdEncoding.EncodeToString(env.Signature))
		fmt.Printf("        Version: %d\n", chdr.Version)
		fmt.Printf("        Timestamp: %s\n", chdr.Timestamp.String())
		fmt.Printf("        TxId: %s\n", chdr.TxId)
		fmt.Printf("        Channel ID: %s\n", chdr.ChannelId)
		fmt.Printf("        Creator: %s\n", string(shdr.Creator))
	}
}

func hash(b *cb.Block) {
	fmt.Printf("CurrentHash: %s\n", base64.StdEncoding.EncodeToString(ComputeHeaderSHA2(b.Header)))
	fmt.Printf("PreviousHash: %s\n", base64.StdEncoding.EncodeToString(b.Header.PreviousHash))
	fmt.Printf("Computed Data Hash: %s\n", base64.StdEncoding.EncodeToString(ComputeDataSHA256(b.Data)))
}

func main() {
	if len(os.Args) < 3 {
		usage()
		return
	}
	cmd := os.Args[1]
	blk := os.Args[2]

	data, err := ioutil.ReadFile(blk)
	if err != nil {
		fmt.Errorf("%v", err)
		return
	}

	b := &cb.Block{}
	proto.Unmarshal(data, b)

	fMap := map[string]interface{}{
		"info": info,
		"hash": hash,
	}

	if fMap[cmd] == nil {
		usage()
		return
	}

	fMap[cmd].(func(*cb.Block))(b)
}
