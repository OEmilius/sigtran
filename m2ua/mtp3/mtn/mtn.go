//The label used for signalling network management messages is also used for
//testing and maintenance messages (see Recommendation Q.707).
package mtn

type MTN struct {
	H0           uint8 //page 82 Heading code
	H1           uint8
	Test_length  uint16
	Test_pattern []byte
}

func Decode(data []byte) *MTN {
	//l := uint16(data[1])
	return &MTN{
		H0: uint8(data[0]) & 0xf,
		H1: uint8(data[0]) >> 4,
		//Test_length: l,
		//Test_pattern: data[2:l],
	}
}

func (mtn *MTN) String() string {
	s := "MTP3 Managment\n"
	switch mtn.H0 {
	case 1:
		s += " H0: Changeover and changeback messages\n"
		switch mtn.H1 {
		case 1:
			s += " H1: SLTM\n"
		case 2:
			s += " H1: SLTA\n"
		default:
			s += "Not implemented\n"
		}
	default:
		s += "Not implemented\n"
	}
	//s += fmt.Sprintf("Test length: %d\n", mtn.Test_length)
	return s
}
