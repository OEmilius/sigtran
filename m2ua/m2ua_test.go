package m2ua

import (
	"encoding/hex"
	"fmt"
	//	"testing"
)

func Example_DecodeM2UA() {
	h, err := hex.DecodeString("0100060100000038000100080000000d03000028c5214179c05c00010060010a00020907031094958864600801000a070311949538236700")
	if err != nil {
		fmt.Println(err)
	}
	mtp2 := DecodeM2UA(h)
	//mtp2 := Decode_Common_header(h)
	fmt.Println("message \n", mtp2)
	fmt.Println(&mtp2)
	//Output: df
}

//func Test_Decode_mtp2(t *testing.T) {
//	h, err := hex.DecodeString("010006010000002800010008000000010300001581e4057e0121a0aaaaaaaaaaaaaaaaaaaa000000")
//	if err != nil {
//		fmt.Println(err)
//	}
//	mtp2 := Decode_MTP2(h)
//	fmt.Println("message \n", mtp2)
//	fmt.Println(&mtp2)

//}

//func Example_Get_INT_interface_id() {
//	//h, err := hex.DecodeString("010006010000002800010008000000010300001581e4057e0121a0aaaaaaaaaaaaaaaaaaaa000000")
//	h, err := hex.DecodeString("0100060100000038000100080000000d03000028c5214179c05c00010060010a00020907031094958864600801000a070311949538236700")
//	if err != nil {
//		fmt.Println(err)
//	}
//	mtp2 := Decode_MTP2(h)
//	fmt.Println(mtp2.Get_INT_interfaceID())
//	//Output: 13

//}
