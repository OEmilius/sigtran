package q931

import (
	"encoding/hex"
	"fmt"
	"testing"
)

//func ExampleDecodeQ931() {
//	//h, err := hex.DecodeString("0802ba8d071e0282821e02828829060d0911111c1b4c0200c3")
//	//h, err := hex.DecodeString("08021ba87b70028133")
//	h, err := hex.DecodeString("0802ba8d071e0282821e02828829060d0911111c1b4c0200c3")
//	if err != nil {
//		fmt.Println(err)
//	}
//	q := DecodeQ931(h)
//	fmt.Println("message \n", q)
//	fmt.Println(&q)
//	//Output: df
//}

func ExampleDecodeQ931_SETUP() {
	h, err := hex.DecodeString("08023a9205a104038090a31803a98393700ba034393539363032343234")
	if err != nil {
		fmt.Println(err)
	}
	q := DecodeQ931(h)
	//fmt.Println("message \n", q)
	fmt.Println(&q)
	//Output: df
}

//func ExampleQ931_Get_Calling_party_number() {
//	h, err := hex.DecodeString("080229e30504038090a31803a9839d6c0c218134393537343134383137700ca13839363437383436393237740700018f32323339")
//	//h, err := hex.DecodeString("080211ef05a104038090a31803a183986c0c2181393236373139383538387008c1373939393039317d029181")
//	if err != nil {
//		fmt.Println(err)
//	}
//	q := DecodeQ931(h)
//	fmt.Println(q.Get_Calling_party_number())
//	fmt.Println(q.Get_Called_party_number())
//	//Output: 4957414817
//}

//func ExampleQ931Get_Channel_number() {
//	h, err := hex.DecodeString("08023a9205a104038090a31803a98393700ba034393539363032343234")
//	if err != nil {
//		fmt.Println(err)
//	}
//	q := DecodeQ931(h)
//	//fmt.Println("message \n", q)
//	fmt.Println(q.Get_Channel_number())
//	//Output: 19
//}

//func ExampleQ931_Get_Called_party_number() {
//	h, err := hex.DecodeString("08023a9205a104038090a31803a98393700ba034393539363032343234")
//	//h, err := hex.DecodeString("08021ba87b70028133") //only one digit
//	if err != nil {
//		fmt.Println(err)
//	}
//	q := DecodeQ931(h)
//	//fmt.Println("message \n", q)
//	//fmt.Println(q.elements_order)
//	fmt.Println(q.Get_Called_party_number())
//	//Output: 4959602424
//}

func Test_decodeElements(t *testing.T) {
	h, err := hex.DecodeString("a104038090a31803a98393700ba034393539363032343234")
	if err != nil {
		fmt.Println(err)
	}
	elements, _ := decodeElements(h)
	//fmt.Println("message \n", elements)
	//fmt.Println(&elements)
	if len(elements) != 4 {
		t.Errorf("количество найденных элементов не правильное", elements)
	}
}

//func Test_DecodeQ931(t *testing.T) {
//	h, err := hex.DecodeString("08021ba87b70028133")
//	if err != nil {
//		fmt.Println(err)
//	}
//	q := DecodeQ931(h)
//	if q.Call_reference_flag != 0 {
//		t.Errorf("Call_reference_flag must be 0", q.Call_reference_flag)
//	}
//}
