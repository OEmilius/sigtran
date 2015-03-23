package mtp3

import (
	"encoding/hex"
	"fmt"
	//	"testing"
)

//func Test_decode(t *testing.T) {
//	message := "c5214179c05c00010060010a00020907031094958864600801000a070311949538236700"
//	h, _ := hex.DecodeString(message)
//	mtp3, _ := Decode(h)
//	fmt.Printf("%#v \r\n", mtp3)
//}

func Example_Decode_MTP3() {
	message := "85f845702162000c0200028090"
	h, _ := hex.DecodeString(message)
	mtp3 := DecodeMTP3(h)
	fmt.Println(mtp3.DPC)
	//Output: 1528
}

func Example_Decode_MTP3_2() {
	message := "85f845702162000c0200028090"
	h, _ := hex.DecodeString(message)
	mtp3 := DecodeMTP3(h)
	fmt.Println(&mtp3)
	//Output: 1528
}

func Example_Decode_MTP3_SLTM() {
	message := "81f8057901"
	h, _ := hex.DecodeString(message)
	mtp3 := DecodeMTP3(h)
	fmt.Println(&mtp3)
	//Output: 1528
}

func Example_Decode_MTP3_TFA() {
	message := "80f8057901"
	h, _ := hex.DecodeString(message)
	mtp3 := DecodeMTP3(h)
	fmt.Println(&mtp3)
	//Output: 1528
}
