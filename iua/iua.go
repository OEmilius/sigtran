// IUA project IUA.go
package iua

import (
	"encoding/binary"
	"fmt"
	"sigtran"
)

type IUA struct {
	Common_header sigtran.Common_header
	Interface_id  sigtran.Parameter
	DLCI          sigtran.Parameter
	Protocol_data sigtran.Parameter
}

func DecodeIUA(data []byte) IUA {
	iua := IUA{}
	iua.Common_header = sigtran.Decode_Common_header(data[:8])
	p_map := sigtran.Decode_params(data[8:])
	//fmt.Println(p_map)
	iua.Interface_id = p_map[1]
	iua.DLCI = p_map[5]
	if p, ok := p_map[14]; ok { //Protocol data
		iua.Protocol_data = p
	}
	return iua
}

func (iua *IUA) String() string {
	s := "ISDN Q.921-User Adaptaion\n"
	s += iua.Common_header.String()
	s += iua.Interface_id.String()
	s += iua.DLCI.String()
	if iua.Protocol_data.Tag != 0 {
		s += iua.Protocol_data.String()
	}
	//s += m.Protocol_data.String()
	return s
}

func (iua *IUA) Get_Interface_id() string {
	//binary.BigEndian.Uint32(iua.Interface_id.Value)
	return fmt.Sprintf("%d", binary.BigEndian.Uint32(iua.Interface_id.Value))
}
