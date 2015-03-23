package isup

import (
	"strconv"
)

//page 85 T-REC-Q.763-199303-S!!PDF-E.pdf AnnexB General description of component encoding rules
type Parameter struct {
	Code     int
	Length   int
	Contents []byte
}

//page 69 T-REC-Q.763-199303-S!!PDF-E.pdf
type IAM struct {
	//		*Parameter*					*Length(octets) *
	//Nature of connection indicators	1
	//Forward call indicators			2
	//Calling party’s category			1
	//Transmission medium requirement	1
	Called_party //переменная длинна
	//даллее опциональные параметры
	Calling_party //переменная длинна
	//Original_Called_Number
	//Redirecting_Number
}

func (iam *IAM) String() string {
	s := iam.Called_party.String()
	s += iam.Calling_party.String()
	return s
}

type Called_party struct { //длинна переменная указывается в length
	//Pointer               int
	Length                uint8
	Odd_even              uint8 //если 1 значит количество цифр в номере не четное
	Nature_of_address_ind uint8
	INN_ind               uint8
	Numbering_plan        uint8
	Number                string
}

type Calling_party struct { //длинна переменная указывается в length
	Length                          uint8
	Odd_even                        uint8 //если 1 значит количество цифр в номере не четное
	Nature_of_address_ind           uint8
	NI_indicator                    uint8
	Numbering_plan                  uint8
	Address_presentation_restricted uint8
	Screening_indicator             uint8
	Number                          string
}

func DecodeIAM(data []byte) IAM { //в data передаем весь ISDN User Part
	iam := IAM{}
	//пропускаем
	//CIC data[:2]
	//isup.Code = int(data[2])
	//Nature of connection indicators = data[4]
	//Forward call indicators = data[4:6]
	//Calling party’s category = data[6]
	//Transmission medium requirement = data[7]

	pointer_to_parameter := int(data[8])
	pointer_to_start_of_optional_part := int(data[9]) //через сколько byte находится код параметра
	cd := decode_called(data[10:])
	iam.Called_party = cd
	//fmt.Println("here optional para code =", data[9+pointer_to_start_of_optional_part])
	//fmt.Println("here optional para length =", data[10+pointer_to_start_of_optional_part])
	//опциональная часть начинается с data[9+pointer_to_start_of_optional_part:]
	//l := int(data[10+pointer_to_start_of_optional_part])
	if pointer_to_parameter != 2 { //если мандаторный параметр с variable длинной не один
		return iam
	}
	op := decode_optional_part(data[9+pointer_to_start_of_optional_part:])
	for _, p := range op {

		switch p.Code {
		case 10: //Calling Party Number
			cg := decode_calling(p.Contents)
			iam.Calling_party = cg
		case 40: //Original Called Number
		case 11: //Redirecting Number
		case 19: //Redirection information
		}
	}
	//fmt.Printf("что осталось %x \r\n", data[10+l+1+pointer_to_start_of_optional_part:])
	return iam
}
func decode_optional_part(data []byte) []Parameter {
	op_data := data
	params := []Parameter{}
	for {
		if len(op_data) < 4 {
			break
		} else {
			p := Parameter{}
			p.Code = int(op_data[0])
			p.Length = int(op_data[1])
			//p.Contents = op_data[2 : 2+p.Length] //без длинны
			p.Contents = op_data[1 : 2+p.Length] //пох вместе с длинной передам
			//fmt.Printf("param %x \r\n", p.Contents)
			params = append(params, p)
			op_data = op_data[p.Length+2:]
			//fmt.Println("следующий параметр", op_data)
		}
	}
	return params
}

func decode_called(data []byte) Called_party {
	cd := Called_party{}
	cd.Length = (data[0])
	cd.Odd_even = uint8(data[1]) >> 7
	cd.Nature_of_address_ind = uint8(data[1]) & 0x7f
	cd.INN_ind = data[2] >> 7
	cd.Numbering_plan = uint8(data[2]>>4) & 0x7
	//cd.Called_party_number = data[3:cd.Length +1]
	cd.Number = get_number(cd.Odd_even, data[3:cd.Length+1])
	return cd
}

func decode_calling(data []byte) Calling_party {
	cg := Calling_party{}
	cg.Length = (data[0])
	cg.Odd_even = uint8(data[1]) >> 7
	//cg.Nature_of_address_ind
	//cg.NI_indicator
	//cg.Numbering_plan
	//cg.Address_presentation_restricted
	//cg.Screening_indicator
	cg.Number = get_number(cg.Odd_even, data[3:cg.Length+1])
	return cg
}

func get_number(odd uint8, data []byte) string {
	s := ""
	for _, v := range data {
		if (int(v) & 0xf) != 15 {
			s += strconv.Itoa(int(v) & 0xf)
		} else {
			s += "F"
		}
		if int(v)>>4 != 15 {
			s += strconv.Itoa(int(v) >> 4)
		} else {
			s += "F"
		}
	}
	if odd == 0 {
		return s
	} else {
		return s[:len(s)-1]
	}
}

var address_indicator = map[uint8]string{
	1: "subscriber number",
	2: "unknown",
	3: "National(significant)number",
	4: "international number",
}

var num_plan = map[uint8]string{
	0: "spare",
	1: "ISDN(Telephony)numbering plan",
	2: "spare",
	3: "Data numbering plan",
	4: "Telex numbering plan",
	5: "reserved for national use",
	6: "reserved for national use",
	7: "spare",
}

func (cd Called_party) String() string {
	s := "Called Party Number\r\n"
	s += "Nature of address indicator: " + address_indicator[cd.Nature_of_address_ind]
	s += "\r\n"
	s += "Number_plan: " + num_plan[cd.Numbering_plan]
	s += "\r\n"
	s += "Number: " + cd.Number
	s += "\r\n"
	return s
}

var presentation = map[uint8]string{
	0: "allowed",
	1: "restricted",
	2: "address not available",
	3: "spare",
}

func (cg Calling_party) String() string {
	s := "Calling Party Number\r\n"
	s += "Nature of address indicator: " + address_indicator[cg.Nature_of_address_ind]
	s += "\r\n"
	s += "Number_plan: " + num_plan[cg.Numbering_plan]
	s += "\r\nAddress presentation restricted indicator: "
	s += presentation[cg.Address_presentation_restricted]
	s += "\r\nNumber: " + cg.Number
	s += "\r\n"
	return s
}
