package sigtran

import (
	"encoding/hex"
	"fmt"
	"testing"
)

func Example_Decode_Common_header() {
	h, err := hex.DecodeString("0100060100000038000100080000000d03000028c5214179c05c00010060010a00020907031094958864600801000a070311949538236700")
	if err != nil {
		fmt.Println(err)
	}
	ch := Decode_Common_header(h)
	fmt.Println(&ch)
	//Output: User Adaptation
	//  Version: 1
	//  Message Class: MTP2 User Adaptation (MAUP) Messages(6)
	//  Message Type: Data(1)
	//  Message Length: 56

}

func Example_Decode_Common_header_IUA() {
	h, err := hex.DecodeString("010005020000002800010008000000b80005000800030000000e000d08021ca84508028090000000")
	if err != nil {
		fmt.Println(err)
	}
	ch := Decode_Common_header(h)
	fmt.Println(&ch)
	//Output: User Adaptation
	//  Version: 1
	//  Message Class: Q.921/Q.931 Boundary Primitives Transport (QPTM) Messages(5)
	//  Message Type: Data Indication(2)
	//  Message Length: 40

}

func Test_decode_params(t *testing.T) {
	//h, err := hex.DecodeString("010005020000002800010008000000b80005000800030000000e000d08021ba87b70028133000000")
	h, err := hex.DecodeString("0100060100000038000100080000000d03000028c5214179c05c00010060010a00020907031094958864600801000a070311949538236700")
	if err != nil {
		fmt.Println(err)
	}
	ch := Decode_Common_header(h)
	fmt.Println(&ch)
	p_map := Decode_params(h[8:])
	for _, p := range p_map {
		fmt.Println(&p)
	}
}
