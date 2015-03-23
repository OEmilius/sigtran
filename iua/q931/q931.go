package q931

import (
	//	"encoding/binary"
	"fmt"
)

/*
every message shall consist of the following parts:
a) protocol discriminator;
b) call reference;
c) message type;
d) other information elements, as required. */
//page 63 T-REC-Q.931-199805-I!!PDF-E.pdf
type Q931 struct {
	Protocol_discriminator      uint16
	Call_reference_value_length uint8
	Call_reference_flag         uint8
	Call_reference_value        []byte
	Message_type                uint16
	Information_elements        map[uint16]Element
	elements_order              []uint16
}

/* page 66
Two categories of information elements are defined:
a) single-octet information elements [see diagrams a) and b) of Figure 4-7];
b) variable length information elements [see diagram c) of Figure 4-7].
*/
type Element struct {
	single_octet bool
	Id           uint16
	Length       uint16
	Content      []byte
}

func decodeElements(old_data []byte) (map[uint16]Element, []uint16) {
	m := make(map[uint16]Element)
	data := old_data[:]
	e_order := []uint16{}
	for len(data) > 3 {
		if int(data[0]>>7) == 0 { //variable length information element
			e_len := int(data[1])
			e := decode_variable_octet(data[:2+e_len])
			m[e.Id] = e
			e_order = append(e_order, e.Id)
			data = data[e.Length+2:]
		} else { //single length infromation element
			e := decode_single_octet(data[0])
			m[e.Id] = e
			e_order = append(e_order, e.Id)
			data = data[1:]
		}
	}
	//fmt.Println(e_order)
	//fmt.Println(m)
	return m, e_order
}
func decode_variable_octet(data []byte) Element {
	l := uint16(data[1])
	return Element{
		single_octet: false,
		Id:           uint16(data[0]),
		Length:       l,
		Content:      data[2 : 2+l],
	}
}

func decode_single_octet(b byte) Element {
	return Element{
		single_octet: true,
		Id:           uint16(b & 0x7F),
	}
}

func DecodeQ931(data []byte) Q931 {
	q := Q931{
		Protocol_discriminator: uint16(data[0]),
	}
	q.Call_reference_value_length = data[1] & 0xF
	q.Call_reference_flag = data[2] >> 7
	if q.Call_reference_value_length == 2 {
		q.Call_reference_value = []byte{data[2] & 0x7F, data[3]}
		q.Message_type = uint16(data[4])
		q.Information_elements, q.elements_order = decodeElements(data[5:])
	} else {
		q.Call_reference_value = []byte{data[2] & 0x7F}
		q.Message_type = uint16(data[3])
		q.Information_elements, q.elements_order = decodeElements(data[4:])
	}

	return q
}

func (q *Q931) Get_Call_ref() string {
	return fmt.Sprintf("%X", q.Call_reference_value)
}

//page 86
func (q *Q931) Get_Channel_number() string {
	e := q.Information_elements[24]
	l := len(e.Content)
	return fmt.Sprintf("%d", e.Content[l-1]&0x7F)
}

func (q *Q931) Get_Called_party_number() string {
	if e, ok := q.Information_elements[112]; ok {
		return get_number(e.Content[1:])
	}
	return ""
}

func (q *Q931) Get_Calling_party_number() string {
	if e, ok := q.Information_elements[108]; ok {
		//fmt.Printf("%X\n", e.Content[2:])
		return get_number(e.Content[2:])
	}
	return ""
}

func get_number(data []byte) string {
	s := ""
	for _, b := range data {
		s += fmt.Sprintf("%d", b&0xF)
	}
	return s
}

func (q *Q931) Get_Message_type() string {
	return message_name[q.Message_type]
}

/* page 65
Table 4-2/Q.931 – Message types
Bits
87654321
00000000 Escap et onationall yspecifi cmessag etyp e(Note)
000----- Call establishment message:
   0 0 0 0 1  –ALERTING
   0 0 0 1 0 – CALL PROCEEDING
   0 0 1 1 1 – CONNECT
   0 1 1 1 1 – CONNECT ACKNOWLEDGE
   0 0 0 1 1  –PROGRESS
   0 0 1 0 1  –SETUP
   0 1 1 0 1 – SETUP ACKNOWLEDGE
001----- Call information phase message:
   0 0 1 1 0  –RESUME
   0 1 1 1 0 – RESUME ACKNOWLEDGE
   0 0 0 1 0 – RESUME REJECT
   0 0 1 0 1 – SUSPEND
   0 1 1 0 1 – SUSPEND ACKNOWLEDGE
   0 0 0 0 1 – SUSPEND REJECT
   0 0 0 0 0 – USER INFORMATION
010----- Call clearing messages:
   0 0 1 0 1 – DISCONNECT
   0 1 1 0 1 – RELEASE
   1 1 0 1 0 – RELEASE COMPLETE
   0 0 1 1 0  –RESTART
   0 1 1 1 0 – RESTART ACKNOWLEDGE
011----- Miscellaneous messages:
   0 0 0 0 0  –SEGMENT
   1 1 0 0 1 – CONGESTION CONTROL
   1 1 0 1 1 – INFORMATION
   0 1 1 1 0  –NOTIFY
   1 1 1 0 1 – STATUS
   1 0 1 0 1 – STATUS ENQUIRY
*/
func (q *Q931) String() string {
	s := "Q.931\n"
	s += "Protocol discriminator: "
	switch q.Protocol_discriminator {
	case 8:
		s += "Q.931"
	default:
		s += "Not implemented"
	}
	s += fmt.Sprintf("(%d)\n", q.Protocol_discriminator)
	s += "Call reference flag: "
	if q.Call_reference_flag == 0 {
		s += "Message sent from originating side"
	} else {
		s += "Message sent to originating side"
	}
	s += fmt.Sprintf("(%d)\n", q.Call_reference_flag)
	s += fmt.Sprintf("Call reference value: %d (%X)\n", q.Call_reference_value, q.Call_reference_value)
	s += "Message type: " + message_name[q.Message_type]
	s += fmt.Sprintf("(%d)[%X]\n", q.Message_type, q.Message_type)
	if q.Message_type == 5 { //SETUP
		if cd := q.Get_Calling_party_number(); cd != "" {
			s += "Calling_party_number: " + cd + "\n"
		}
		if cg := q.Get_Called_party_number(); cg != "" {
			s += "Called_party_number: " + cg + "\n"
		}
	}
	if q.Message_type == 123 { //information
		if cd := q.Get_Calling_party_number(); cd != "" {
			s += "Calling_party_number: " + cd + "\n"
		}
	}
	return s
}

var message_name = map[uint16]string{
	//000----- Call establishment message:
	1:  "ALERTING",
	2:  "CALL PROCEEDING",
	7:  "CONNECT",
	15: "CONNECT ACKNOWLEDGE",
	3:  "PROGRESS",
	5:  "SETUP",
	13: "SETUP ACKNOWLEDGE",
	//001----- Call information phase message:
	38: "RESUME",
	46: "RESUME ACKNOWLEDGE",
	34: "RESUME REJECT",
	37: "SUSPEND",
	45: "SUSPEND ACKNOWLEDGE",
	33: "SUSPEND REJECT",
	32: "USER INFORMATION",
	//010----- Call clearing messages:
	69: "DISCONNECT",
	77: "RELEASE",
	90: "RELEASE COMPLETE",
	70: "RESTART",
	78: "RESTART ACKNOWLEDGE",
	//011----- Miscellaneous messages:
	96:  "SEGMENT",
	121: "CONGESTION CONTROL",
	123: "INFORMATION",
	110: "NOTIFY",
	125: "STATUS",
	117: "STATUS ENQUIRY",
}
