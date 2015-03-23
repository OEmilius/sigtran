package isup

import (
	"encoding/binary"
	"fmt"
)

//T-REC-Q.763-199303-S!!PDF-E.pdf
type ISUP struct {
	CIC               uint16
	Message_type_code int
	Params            []byte
}

func DecodeISUP(data []byte) ISUP {
	isup := ISUP{}
	isup.CIC = binary.LittleEndian.Uint16(data[:2])
	isup.Message_type_code = int(data[2])
	//isup.Params = data[3:]
	isup.Params = data[:]
	return isup
}

func (i *ISUP) Message_type() string {
	return code_type[i.Message_type_code]
}

//page 85 T-REC-Q.763-199303-S!!PDF-E.pdf AnnexB General description of component encoding rules
//type Parameter struct {
//	Code     int
//	Length   int
//	Contents []byte
//}

var code_type = map[int]string{
	1:  "IAM",
	2:  "SAM Subsequent Address Message",
	3:  "INR Information Request",
	4:  "INF Information",
	5:  "COT Continuity",
	6:  "ACM",
	7:  "CON Connect",
	8:  "FOT Forward Transfer",
	9:  "ANM",
	10: "Reserved",
	11: "Reserved",
	12: "REL",
	13: "SUS Suspend",
	14: "RES Resume",
	15: "Reserved",
	16: "RLC",
	17: "CCR Continuity Check Request",
	18: "RSC Reset Circuit",
	19: "BLO Blocking",
	20: "UBL Unblocking",
	21: "BLA Blocking Acknowledgement",
	22: "UBA Unblocking Acknowledgement",
	23: "GRA Circuit Group Reset Acknowledgement",
	24: "CGB Circuit Group Blocking",
	25: "CGU Circuit Group Unblocking",
	26: "CGBA Circuit Group Blocking Acknowledgement",
	27: "CGUA Circuit Group Unblocking Acknowledgement",
	28: "CMR Call Modification Request",
	29: "CMC Call Modification Completed",
	30: "CMRJ Call Modification Reject",
	31: "FRJ Facility Reject",
	32: "FAA Facility Accepted",
	33: "FAR Facility Request",
	34: "Reserved",
	35: "Reserved",
	36: "LPA Loop Back Acknowledgement",
	37: "Reserved",
	38: "Reserved",
	39: "DRS Delayed Release",
	40: "PAM Pass Along",
	41: "GRS Circuit Group Reset",
	42: "CQM Circuit Group Query",
	43: "CQR Circuit Group Query Response",
	44: "CPG Call Progress",
	45: "USR User to User Information",
	46: "UCIC Unequipped Circuit Identification Code",
	47: "CFN Confusion",
	48: "OLM Overload",
	49: "CRG Charge information",
}

func (isup *ISUP) String() string {
	s := "ISDN User Part\n"
	s += "CIC: " + fmt.Sprintf("%d \n", isup.CIC)
	s += "Message type: " + code_type[isup.Message_type_code] + fmt.Sprintf("(%d)\n", isup.Message_type_code)
	//if cd := isup.Called_party_number(); cd != "" {
	//	s += "Called party number: " + cd + "\n"
	//}
	//if cg := isup.Calling_party_number(); cg != "" {
	//	s += "Calling party number: " + cg + "\n"
	//}
	//if isup.Message_type_code == 12 {
	//	rel := DecodeREL(isup.Params)
	//	s += rel.String()
	//}
	return s
}
