//rfc3331
package m2ua

import (
	"encoding/binary"
	"fmt"
	"sigtran"
)

type M2UA struct {
	Common_header sigtran.Common_header
	Interface_id  sigtran.Parameter
	Protocol_data sigtran.Parameter
	//Correlation_id Parameter //optional
}

func DecodeM2UA(data []byte) M2UA {
	m2 := M2UA{}
	m2.Common_header = sigtran.Decode_Common_header(data[:8])
	p_map := sigtran.Decode_params(data[8:])
	if m2.Common_header.MsgType == 1 {
		m2.Interface_id = p_map[1]
		m2.Protocol_data = p_map[768]
	}
	return m2
}

func (m *M2UA) String() string {
	s := "MTP2 User Adaptation\n"
	s += m.Common_header.String()
	s += m.Interface_id.String()
	s += m.Protocol_data.String()
	return s
}

func (m *M2UA) Get_interfaceID() string {
	if m.Interface_id.Tag == 1 {
		id := binary.BigEndian.Uint32(m.Interface_id.Value)
		return fmt.Sprintf("%d", id)
	}
	return string(m.Interface_id.Value)

}
