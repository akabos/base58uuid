package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/btcsuite/btcutil/base58"
	uuid "github.com/satori/go.uuid"
)

var (
	encode bool
	decode bool
)

func init() {
	flag.BoolVar(&encode, "e", false, "encode")
	flag.BoolVar(&decode, "d", false, "decode")
	flag.Parse()
}

func main() {
	if encode == decode {
		flag.Usage()
		os.Exit(1)
	}
	var (
		buf *bytes.Buffer
		arg = flag.Arg(0)
	)
	if arg == "" {
		buf = &bytes.Buffer{}
		_, _ = buf.ReadFrom(os.Stdin)
	} else {
		buf = bytes.NewBufferString(arg)
	}
	arg = strings.TrimSpace(buf.String())

	switch  {
	case encode:
		fmt.Println(base58.Encode(uuid.FromStringOrNil(arg).Bytes()))
	case decode:
		fmt.Println(uuid.FromBytesOrNil(base58.Decode(arg)).String())
	default:
		panic("should never happen")
	}
}
