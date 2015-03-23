package iua

import (
	"encoding/hex"
	"fmt"
	//	"testing"
)

func ExampleDecodeIUA() {
	//h, err := hex.DecodeString("010005020000002800010008000000b80005000800030000000e000d08021ba87b70028133000000")
	h, err := hex.DecodeString("010005050000001800010008000000b10005000800010000")
	if err != nil {
		fmt.Println(err)
	}
	iua := DecodeIUA(h)
	fmt.Println("message \n", iua)
	fmt.Println(&iua)
	//Output: df
}

func ExampleQ931_Get_Interface_id() {
	h, err := hex.DecodeString("0100050200000044000100080000028c0005000800030000000e002b080203780504038090a31803a9839e6c0c0081343935373431343530317008803731333334383300")
	if err != nil {
		fmt.Println(err)
	}
	iua := DecodeIUA(h)
	//fmt.Println("message \n", iua)
	fmt.Println(iua.Get_Interface_id())
	//Output: 652

}
