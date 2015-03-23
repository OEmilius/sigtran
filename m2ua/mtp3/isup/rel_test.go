package isup

import (
	"encoding/hex"
	"fmt"
	//	"testing"
)

func Example_DecodeREL() {
	//h, _ := hex.DecodeString("44000c0200028483")
	//h, _ := hex.DecodeString("44000c0200028483")
	//h, _ := hex.DecodeString("44000c0200028483")
	//h, _ := hex.DecodeString("44000c0200028483")
	h, _ := hex.DecodeString("62000c0200028090")
	rel := DecodeREL(h)
	fmt.Println(rel.Cause_indicator)
	fmt.Println(&rel.Cause_indicator)
	//Output:92
	//1
}

func Example_DecodeREL_2() {
	//h, _ := hex.DecodeString("44000c0200028483")
	//h, _ := hex.DecodeString("44000c0200028483")
	//h, _ := hex.DecodeString("44000c0200028483")
	//h, _ := hex.DecodeString("44000c0200028483")
	h, _ := hex.DecodeString("62000c0200028090")
	rel := DecodeISUP(h)
	fmt.Println(&rel)
	//Output:92
	//1
}
