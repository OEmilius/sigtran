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
		}
	case 4:
		s += " H0: Transfer-prohibited-allowed-restricted messages\n"
		switch mtn.H1 {
		case 1:
			s += " H1: TFP\n"
		case 5:
			s += " H1: TFA\n"
		}
	default:
		s += "Not implemented\n"
	}
	return s

}
