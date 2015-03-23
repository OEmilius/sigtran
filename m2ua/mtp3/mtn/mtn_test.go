package mtn

import (
	"encoding/hex"
	"fmt"
)

func Example_Decode() {
	h, _ := hex.DecodeString("11a0aaaaaaaaaaaaaaaaaaaa")
	m := Decode(h)
	fmt.Println(m)
	//Output: MTP3 Managment
	//  H0: Changeover and changeback messages
	//  H1: SLTM
}

func Example_Decode_SLTA() {
	h, _ := hex.DecodeString("216055a111000164")
	m := Decode(h)
	fmt.Println(m)
	//Output: MTP3 Managment
	//  H0: Changeover and changeback messages
	//  H1: SLTA
}
