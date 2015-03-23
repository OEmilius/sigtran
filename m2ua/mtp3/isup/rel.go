package isup

import (
	"fmt"
)

//page 70
type REL struct {
	//code 12
	Cause_indicator
}

//-> Q.850
type Cause_indicator struct {
	Location        uint8
	Coding_standard uint8
	Cause_value     uint16
}

//page 70
func DecodeREL(data []byte) REL {
	if data[2] != 12 { //this is not release
		return REL{}
	}
	fmt.Printf("%x\n", data[6])
	ci := Cause_indicator{
		Location:        data[6] & 0xF,
		Coding_standard: data[6] >> 5 & 3,
		Cause_value:     uint16(data[7]) & 0x7F,
	}
	return REL{Cause_indicator: ci}
}

func (rel *REL) String() string {
	return rel.Cause_indicator.String()
}

func (ci *Cause_indicator) String() string {
	s := "Cause indicators\n"
	s += "Location: "
	switch ci.Location {
	case 0:
		s += "user(U)"
	case 1:
		s += "private network serving the local user (LPN)"
	case 2:
		s += "public network serving the local user (LN)"
	case 3:
		s += "transit network (TN)"
	case 4:
		s += "public network serving the remote user (RLN)"
	case 5:
		s += "private network serving the remote user (RPN)"
	case 7:
		s += "international network (INTL)"
	case 10:
		s += "network beyond interworking point (BI)"
	default:
		s += "reserved for national use"
	}
	s += fmt.Sprintf("(%d)\n", ci.Location)
	s += "Coding standard: "
	switch ci.Coding_standard {
	case 0:
		s += "ITU-T standardized coding"
	case 1:
		s += "ISO/IEC standard"
	case 2:
		s += "national standard"
	case 3:
		s += "standard specific to identified location"
	}
	s += fmt.Sprintf("(%d)\n", ci.Coding_standard)
	s += "Cause value: " + cause_codes[ci.Cause_value]
	s += fmt.Sprintf("(%d)\n", ci.Cause_value)
	return s
}

//https://support.sonus.net/display/uxapidoc/q.850+cause+codes+-+reference
var cause_codes = map[uint16]string{
	1:   "Unallocated Number",
	2:   "No Route to Transit Network",
	3:   "No Route to Destination",
	4:   "Send Special Information tone",
	5:   "Misdialed Trunk Prefix",
	6:   "Channel Unacceptable",
	7:   "Call Awarded in Established Channel",
	8:   "Preemption",
	9:   "Preemption - Circuit Reserved for Reuse",
	16:  "Normal Call Clearing",
	17:  "User Busy",
	18:  "No User Responding",
	19:  "No Answer from User (user alerted)",
	20:  "Subscriber Absent",
	21:  "Call Rejected",
	22:  "Number Changed",
	23:  "Redirection to New Destination",
	25:  "Exchange Routing Error",
	26:  "Non-selected User Clearing",
	27:  "Destination Out of Order",
	28:  "Invalid Number Format (addr incomplete)",
	29:  "Facility Rejected",
	30:  "Response to STATUS ENQUIRY",
	31:  "Normal, Unspecified",
	34:  "No Circuit//Channel Available",
	38:  "Network Out of Order",
	39:  "Permanent Frame Mode Connection OoS",
	40:  "Permanent Frame Mode Connection Oper",
	41:  "Temporary Failure",
	42:  "Switching Equipment Congestion",
	43:  "Access Information Discarded",
	44:  "Requested Circuit//Channel N//A",
	46:  "Precedence Call Blocked",
	47:  "Resource Unavailable, Unspecified",
	49:  "Quality of Service Not Available",
	50:  "Requested Facility Not Subscribed",
	53:  "Outgoing Calls Barred Within CUG",
	55:  "Incoming Calls Barred Within CUG",
	57:  "Bearer Capability Not Authorized",
	58:  "Bearer Capability Not Available",
	62:  "Inconsistency in Outgoing IE",
	63:  "Service or Option N/A, unspecified",
	65:  "Bearer Capability Not Implemented",
	66:  "Channel Type Not Implemented",
	69:  "Requested Facility Not Implemented",
	70:  "Only Restricted Digital Bearer Cap supported",
	79:  "Service or Option Not Implemented, Unspecified",
	81:  "Invalid Call Reference Value",
	82:  "Identified Channel Does Not Exist",
	83:  "Call Exists, but Call Identity Does Not",
	84:  "Call Identity in User",
	85:  "No Call Suspended",
	86:  "Call with Requested Call Identity has Cleared",
	87:  "User Not Member of CUG",
	88:  "Incompatible Destination",
	90:  "Non-existent CUG",
	91:  "Invalid Transit Network Selection",
	95:  "Invalid Message, Unspecified",
	96:  "Mandatory Information Element is Missing",
	97:  "Message Type Non-existent // Not Implemented",
	98:  "Message Incompatible With Call State or Message Type",
	99:  "IE/Parameter Non-existent or Not Implemented",
	100: "Invalid Information Element Contents",
	101: "Message Not Compatible With Call State",
	102: "Recovery on Timer Expiry",
	103: "Parameter Non-existent // Not Implemented, Passed On",
	110: "Message With Unrecognized Parameter, Discarded",
	111: "Protocol Error, Unspecified",
	127: "Interworking, Unspecified",
}
