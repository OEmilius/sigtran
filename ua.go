package sigtran

import (
	"encoding/binary"
	"fmt"
)

/*3.1  Common Message Header
  The protocol messages for MTP2-User Adaptation require a message
  structure that contains a version, message class, message type,
  message length, and message contents.  This message header is common
  among all signalling protocol adaptation layers:

   0                   1                   2                   3
   0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
  +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
  |    Version    |     Spare     | Message Class | Message Type  |
  +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
  |                        Message Length                         |
  +-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

                 Figure 2  Common Message Header

  All fields in an M2UA message MUST be transmitted in the network byte
  order, unless otherwise stated.
*/
type Common_header struct {
	Version   uint8
	MsgClass  uint8 //3.1.3  Message Class
	MsgType   uint8
	MsgLength uint32
}

func Decode_Common_header(data []byte) Common_header {
	return Common_header{
		Version:   uint8(data[0]),
		MsgClass:  uint8(data[2]),
		MsgType:   uint8(data[3]),
		MsgLength: binary.BigEndian.Uint32(data[4:8]),
	}
}

func (ch *Common_header) String() string {
	//до того как выводить ниже . выводим выше IUA or MTP2
	s := fmt.Sprintf("Version: %d\n Message Class: ", ch.Version)
	switch ch.MsgClass {
	case 6:
		s += "MTP2 User Adaptation (MAUP) Messages"
	case 5:
		s += "Q.921/Q.931 Boundary Primitives Transport (QPTM) Messages"
	default:
		s += "Print Not Implemented"
	}
	s += fmt.Sprintf("(%d)\n", ch.MsgClass)
	s += " Message Type: "
	switch ch.MsgType {
	case 1:
		s += "Data"
	case 2:
		s += "Data Indication"
	case 5:
		s += "Establish Request"
	default:
		s += "Not implemented"
	}
	s += fmt.Sprintf("(%d)\n", ch.MsgType)
	s += fmt.Sprintf(" Message Length: %d\n", ch.MsgLength)
	return s
}

/*3.1.6  Variable-Length Parameter Format
0                   1                   2                   3
 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1 2 3 4 5 6 7 8 9 0 1
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
|          Parameter Tag        |       Parameter Length        |
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+
\                                                               \
/                       Parameter Value                         /
\                                                               \
+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+

Mandatory parameters MUST be placed before optional parameters in a
message.
*/
//Parameter Tag: 16 bits (unsigned integer)
type Parameter struct {
	Tag    uint16
	Length uint16
	Value  []byte
}

func Decode_params(data []byte) map[uint16]Parameter {
	data_len := len(data)
	total_len := 0
	p_map := make(map[uint16]Parameter)
	for total_len < data_len-3 {
		l := binary.BigEndian.Uint16(data[2:4])
		p := decode_para(data[:l])
		p_map[p.Tag] = p
		total_len += int(l)
		data = data[l:]
	}
	return p_map
}

func decode_para(data []byte) Parameter {
	l := binary.BigEndian.Uint16(data[2:4])
	return Parameter{
		Tag:    binary.BigEndian.Uint16(data[0:2]),
		Length: l,
		Value:  data[4:l],
	}
}

/*
Parameter Value     Parameter Name
---------------     --------------
      0 (0x00)       Reserved
      1 (0x01)       Interface Identifier (Integer)
      2 (0x02)       Unused
      3 (0x03)       Interface Identifier (Text)
      4 (0x04)       Info String
      5 (0x05)       DLCI
      6 (0x06)       Unused
      7 (0x07)       Diagnostic Information
      8 (0x08)       Interface Identifier (Integer Range)
      9 (0x09)       Heartbeat Data
     10 (0x0a)       Unused
     11 (0x0b)       Traffic Mode Type
     12 (0x0c)       Error Code
     13 (0x0d)       Status
     14 (0x0e)       Protocol Data
     15 (0x0f)       Release Reason
     16 (0x10)       TEI Status
     17 (0x11)       ASP Identifier
     18 (0x12)       Unused
     19 (0x13)       Correlation Id
 768 (0x0300)      Protocol Data 1
 769 (0x0301)      Protocol Data 2 (TTC)
 770 (0x0302)      State Request
 771 (0x0303)      State Event
 772 (0x0304)      Congestion Status
 773 (0x0305)      Discard Status
 774 (0x0306)      Action
 775 (0x0307)      Sequence Number
 776 (0x0308)      Retrieval Result
 777 (0x0309)      Link Key
 778 (0x030a)      Local-LK-Identifier
 779 (0x030b)      Signalling Data Terminal (SDT) Identifier
 780 (0x030c)      Signalling Data Link (SDL) Identifier
 781 (0x030d)      Registration Result
 782 (0x030e)      Registration Status
 783 (0x030f)      De-Registration Result
 784 (0x0310)      De-Registration Status
*/

var tag_name = map[uint16]string{
	1:   "Interface Identifier (Integer)",
	3:   "Interface Identifier (Text)",
	5:   "DLCI",
	14:  "Protocol Data",
	768: "Protocol Data 1",
	769: "Protocol Data 2 (TTC)",
}

func (p *Parameter) String() string {
	s := "Parameter Tag: "
	s += tag_name[p.Tag] + fmt.Sprintf("(%d)\n", p.Tag)
	s += fmt.Sprintf("Parameter Length: %d\n", p.Length)
	switch p.Tag {
	case 1:
		s += fmt.Sprintf("ID (integer): %d\n", binary.BigEndian.Uint32(p.Value))
	case 3:
		s += fmt.Sprint(p.Value)
	}
	return s
}
