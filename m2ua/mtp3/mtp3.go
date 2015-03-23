//T-REC-Q.704-199607-I!!PDF-E.pdf
//The basic signal unit format which is common to all message signal units is described in clause
//2/Q.703. From the point of view of the Message Transfer Part level 3 functions, common
//characteristics of the message signal units are the presence of:
//– the service information octet;
//– the label, contained in the signalling information field, and, in particular, the routing label.

package mtp3

import (
	"encoding/binary"
	"fmt"
)

//T-REC-Q.704-199607-I!!PDF-E.pdf page 80
type MTP3 struct {
	Service_information_octet
	Routing_label
	Payload []byte
}

type Service_information_octet struct {
	Network_indicator uint8
	Service_indicator uint8
}

type Routing_label struct {
	DPC uint32
	OPC uint32
	SLS uint32
}

func DecodeMTP3(data []byte) MTP3 {
	mtp3 := MTP3{}
	//if len(data) < 5 {
	//	return mtp3, false
	//}
	mtp3.Service_indicator = data[0] & 0xf
	mtp3.Network_indicator = data[0] >> 6
	mtp3.DPC = binary.LittleEndian.Uint32(data[1:5]) & 0xfff
	mtp3.OPC = (binary.LittleEndian.Uint32(data[1:5]) >> 14) & 0x3fff
	mtp3.SLS = binary.LittleEndian.Uint32(data[1:5]) >> 28
	mtp3.Payload = data[5:]
	return mtp3
}

func (mtp3 *MTP3) String() string {
	s := "Message Transfer Part Level 3\n"
	s += " Service information octet\n"
	s += "  Network indicator: "
	switch mtp3.Network_indicator {
	case 3:
		s += "Reserved for national use(3)"
	case 2:
		s += "National network(2)"
	case 1:
		s += "Spare(for international use only)"
	case 0:
		s += "International network(0)"
	}
	s += "\n  Service indicator: "
	switch mtp3.Service_indicator {
	case 1:
		s += "MTN"
	case 5:
		s += "ISUP"
	case 0:
		s += "SNM"
	default:
		s += "Not Implemented"
	}
	s += fmt.Sprintf("(%d)", mtp3.Service_indicator)
	s += "\nRouting Label\n"
	s += " DPC: " + fmt.Sprintf("dec%d, hex %X\n", mtp3.DPC, mtp3.DPC)
	s += " OPC: " + fmt.Sprintf("dec%d, hex %X\n", mtp3.OPC, mtp3.OPC)
	s += " SLS: " + fmt.Sprintf("%d\n", mtp3.SLS)
	//s += "\nPayload:" + fmt.Sprintf("%x \n", mtp3.Payload)
	return s
}
