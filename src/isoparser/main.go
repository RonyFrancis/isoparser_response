package main

import (
	"fmt"
	"iso8583"
	"math/rand"
	_ "os"
	"reflect"
	"sort"
	"strconv"
	"time"
)

// Constant Values
const (
	Field4   = "0000000000000000"
	Field24  = "200"
	Field32  = "012"
	Field34  = "00000000000000"
	Field49  = "INR"
	Field41  = "MPASSBOOK       "
	Field42  = "MPASSBOOK      "
	Field123 = "MPB"
)

// iso8583 initalisation
type Data struct {
	No                    *iso8583.Numeric      `field:"3" length:"6" encode:"ascii"`  // bcd value encoding
	No2                   *iso8583.Numeric      `field:"4" length:"16" encode:"ascii"` // bcd value encoding
	Ret2                  *iso8583.Alphanumeric `field:"11" length:"12" encode:"ascii`
	DateTime              *iso8583.Numeric      `field:"12" length:"14" encode:"ascii` // bcd value encoding
	Capturedate           *iso8583.Numeric      `field:"17" length:"8" encode:"ascii"` // bcd value encoding
	Functioncode          *iso8583.Numeric      `field:"24" length:"3" encode:"ascii"` // ascii value encoding
	AcqInsIdeCode         *iso8583.Llvar        `field:"32" length:"11" encode:"ascii`
	AccountNumber         *iso8583.Llvar        `field:"34" length:"14" encode:"ascii"` // bcd length encoding, ascii value encoding
	Rrn                   *iso8583.Alphanumeric `field:"37" length:"12" encode:"ascii"`
	ActionCode            *iso8583.Numeric      `field:"39" length:"3" encode:"ascii`
	CaTerminalId          *iso8583.Alphanumeric `field:"41" length:"16" encode:"ascii`
	CaId                  *iso8583.Alphanumeric `field:"42" length:"15" encode:"ascii`
	AdditionalPrivateData *iso8583.Lllvar       `field:"48" length:"999" encode:"ascii`
	CurrencyCode          *iso8583.Alphanumeric `field:"49" length:"3" encode:"ascii`
	AccountId             *iso8583.Llvar        `field:"102" length:"38" encode:"ascii`
	DeliveryChannel       *iso8583.Lllvar       `field:"123" length:"6" encode:"ascii`
	Reserved1             *iso8583.Lllvar       `field:"125" length:"999" encode:"ascii`
	Reserved2             *iso8583.Lllvar       `field:"126" length:"999" encode:"ascii`
	Reserved3             *iso8583.Lllvar       `field:"127" length:"999" encode:"ascii`
}

// default params
func (data *Data) GetDefaultParams() {
	data.No2 = iso8583.NewNumeric(Field4)
	data.Functioncode = iso8583.NewNumeric(Field24)
	data.AcqInsIdeCode = iso8583.NewLlvar([]byte(Field32))
	data.AccountNumber = iso8583.NewLlvar([]byte(Field34))
	data.CurrencyCode = iso8583.NewAlphanumeric(Field49)
	data.DeliveryChannel = iso8583.NewLllvar([]byte(Field123))
	data.CaTerminalId = iso8583.NewAlphanumeric(Field41)
	data.CaId = iso8583.NewAlphanumeric(Field42)
}

// dynamic Values
func (data *Data) GetDynamicValues() {
	random := random(100000000000, 999999999999)
	r := strconv.Itoa(random)
	data.Rrn = iso8583.NewAlphanumeric(r)
	data.Ret2 = iso8583.NewAlphanumeric(r)
	timestamp := time.Now()
	data.DateTime = iso8583.NewNumeric(timestamp.Format("20060102150405"))
	data.Capturedate = iso8583.NewNumeric(timestamp.Format("20060102"))
}

// Get Account Number
func (data *Data) GetAccountNumber(accountNumber string) {
	formattedAccountNo := "012        " + accountNumber[:4] + "    " + accountNumber
	data.AccountId = iso8583.NewLlvar([]byte(formattedAccountNo))
}

// generate random numbers
func random(min, max int) int {
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}

// generate new raw and newdata for Numeric and Aplhanumeric
func LoadValue(raw string, length int) (string, string) {
	var datastring string
	datastring = raw[0:length]
	raw = raw[length:]
	return datastring, raw
}

func LoadValueLl(raw string, oldlength int) (string, string) {
	var datastring string
	oldlength = 2
	length, err := strconv.Atoi(raw[0:oldlength])
	if err != nil {
		fmt.Println("error convert convert string to int")
		fmt.Println(raw[0:oldlength])
	}
	fmt.Println("length of llvar is", length)
	datastring = raw[oldlength:length]
	raw = raw[length:]
	return datastring, raw
}

func main() {
	fmt.Println("isoparser the begining")

	data := &Data{
		No:        iso8583.NewNumeric("930000"),
		Reserved1: iso8583.NewLllvar([]byte("201702120000000000000000099999999999999999                                10AB")),
	}
	data.GetDefaultParams()
	data.GetDynamicValues()
	data.GetAccountNumber("361103670064591")
	msg := iso8583.NewMessage("1200", data)
	msg.MtiEncode = iso8583.ASCII
	b, err := msg.Bytes()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(string(b))
	data2 := &Data{
		No:            iso8583.NewNumeric(""),
		No2:           iso8583.NewNumeric(""),
		Ret2:          iso8583.NewAlphanumeric(""),
		DateTime:      iso8583.NewNumeric(""),
		Capturedate:   iso8583.NewNumeric(""),
		Functioncode:  iso8583.NewNumeric(""),
		AcqInsIdeCode: iso8583.NewLlvar([]byte("")),
		AccountNumber: iso8583.NewLlvar([]byte("")),
		Rrn:           iso8583.NewAlphanumeric(""),
		ActionCode:    iso8583.NewNumeric(""),
		CaTerminalId:  iso8583.NewAlphanumeric(""),
		CaId:          iso8583.NewAlphanumeric(""),
		AdditionalPrivateData: iso8583.NewLllvar([]byte("")),
		CurrencyCode:          iso8583.NewAlphanumeric(""),
		AccountId:             iso8583.NewLlvar([]byte("")),
		DeliveryChannel:       iso8583.NewLllvar([]byte("")),
		Reserved1:             iso8583.NewLllvar([]byte("")),
		Reserved2:             iso8583.NewLllvar([]byte("")),
		Reserved3:             iso8583.NewLllvar([]byte("")),
	}
	parser := iso8583.Parser{}
	parser.MtiEncode = iso8583.ASCII
	//err = parser.Register("1200", data2)
	if err != nil {
		fmt.Println(err)
	}
	//input := []byte{48, 49, 48, 48, 242, 60, 36, 129, 40, 224, 152, 0, 0, 0, 0, 0, 0, 0, 1, 0, 49, 54, 52, 50, 55, 54, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 55, 55, 55, 48, 48, 48, 55, 48, 49, 49, 49, 49, 56, 52, 52, 48, 48, 48, 49, 50, 51, 49, 51, 49, 56, 52, 52, 48, 55, 48, 49, 49, 57, 48, 50, 6, 67, 57, 48, 49, 48, 50, 48, 54, 49, 50, 51, 52, 53, 54, 51, 55, 52, 50, 55, 54, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 53, 61, 49, 50, 51, 52, 53, 54, 55, 56, 57, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 48, 57, 56, 55, 54, 53, 52, 51, 50, 49, 48, 48, 49, 48, 48, 48, 48, 48, 51, 50, 49, 49, 50, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 51, 52, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 84, 101, 115, 116, 32, 116, 101, 120, 116, 100, 48, 1, 2, 3, 4, 5, 6, 7, 8, 49, 50, 51, 52, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 48, 49, 55, 65, 110, 111, 116, 104, 101, 114, 32, 116, 101, 115, 116, 32, 116, 101, 120, 116}
	_, err = parser.Parse(b)
	fmt.Println(err)
	if err != nil {
		fmt.Println("this is a erroer")
		fmt.Println(err)
	}
	fmt.Println(reflect.ValueOf(string(b[4:20])))
	msg2 := iso8583.NewMessage("1200", data2)
	msg2.MtiEncode = iso8583.ASCII
	ret, fields, err := msg2.Convert(b)

	if err != nil {
		fmt.Println("we have an error\n")
		fmt.Println(err)
	}
	fmt.Printf(string(ret.Mti))
	//fmt.Println(reflect.TypeOf(t))

	//fmt.Println(reflect.ValueOf(fields))
	si := make([]int, 0, len(fields))
	for i := range fields {
		si = append(si, i)
	}
	sort.Ints(si)
	//fmt.Println(reflect.ValueOf(si))
	dataString := string(b[20:])
	i := 0
	for _, key := range si {

		//fmt.Println("Key:", key, "Value:", reflect.TypeOf(fields[key].Field))
		var t interface{}
		t = reflect.TypeOf(fields[key].Field)
		//fmt.Println(t)
		// if reflect.TypeOf(fields[key].Field) == Numeric {
		// 	fmt.Println("ssssssssssssssss")
		// } else {
		// 	fmt.Println("nooooooooooooooooooooo")
		// }
		val := reflect.Indirect(reflect.ValueOf(data2))
		val2 := val.Type().Field(i).Name

		var newdata string // data to be stored

		switch t {
		default:
			fmt.Println(t)
			fmt.Printf("unexpected type ") // %T prints whatever type t has
		case reflect.TypeOf(data2.No):
			//fmt.Println("fields[key]")
			//fmt.Println(reflect.ValueOf(fields[key].Index))
			//data, dataString := dataString[:fields[key].Length], dataString[fields[key].Length:]
			//data2.val.Type().Field(i).Name = iso8583.NewNumeric(data)

			//fmt.Println("type is sss")
			//fmt.Println(reflect.TypeOf(val))

			fmt.Println("val2 valius es")
			fmt.Println(reflect.ValueOf(val2))
			switch val2 {
			case "No":
				newdata, dataString = LoadValue(dataString, fields[key].Length)
				data2.No = iso8583.NewNumeric(newdata)
			case "No2":
				newdata, dataString = LoadValue(dataString, fields[key].Length)
				data2.No2 = iso8583.NewNumeric(newdata)
			case "DateTime":
				//timestamp := time.Now()
				newdata, dataString = LoadValue(dataString, fields[key].Length)
				data2.DateTime = iso8583.NewNumeric(newdata)
			case "Capturedate":
				//timestamp := time.Now()
				newdata, dataString = LoadValue(dataString, fields[key].Length)
				fmt.Println("this is the new data for Captures date", newdata)
				data2.Capturedate = iso8583.NewNumeric(newdata)
			case "Functioncode":
				newdata, dataString = LoadValue(dataString, fields[key].Length)
				data2.Functioncode = iso8583.NewNumeric(newdata)
			case "ActionCode":
				newdata, dataString = LoadValue(dataString, fields[key].Length)
				data2.ActionCode = iso8583.NewNumeric(newdata)
			default:
				fmt.Println("fucked")
			}
			//val.Type().val2 = iso8583.NewNumeric(data)
			//fmt.Println(data)
			fmt.Println(dataString)
		case reflect.TypeOf(data2.Ret2): // alphanumeric
			switch val2 {
			case "Ret2":
				newdata, dataString = LoadValue(dataString, fields[key].Length)
				data2.Ret2 = iso8583.NewAlphanumeric(newdata)
			case "Rrn":
				newdata, dataString = LoadValue(dataString, fields[key].Length)
				data2.Rrn = iso8583.NewAlphanumeric(newdata)
			case "CaTerminalId":
				newdata, dataString = LoadValue(dataString, fields[key].Length)
				data2.CaTerminalId = iso8583.NewAlphanumeric(newdata)
			case "CaId":
				newdata, dataString = LoadValue(dataString, fields[key].Length)
				data2.CaId = iso8583.NewAlphanumeric(newdata)
			case "CurrencyCode":
				newdata, dataString = LoadValue(dataString, fields[key].Length)
				data2.CurrencyCode = iso8583.NewAlphanumeric(newdata)
			}

		case reflect.TypeOf(data2.AcqInsIdeCode): //aLlvar
			//	fmt.Println(reflect.TypeOf(data2.AcqInsIdeCode))
			switch val2 {
			case "AcqInsIdeCode":
				fmt.Println("datastring at AcqInsIdeCode is ", dataString)
				newdata, dataString = LoadValueLl(dataString, 2)
				data2.AcqInsIdeCode = iso8583.NewLlvar([]byte(newdata))
			case "AccountNumber":
				newdata, dataString = LoadValueLl(dataString, 2)
				data2.AccountNumber = iso8583.NewLlvar([]byte(newdata))
			case "AccountId":
				newdata, dataString = LoadValueLl(dataString, 2)
				data2.GetAccountNumber(newdata)
			}
		case reflect.TypeOf(data2.AdditionalPrivateData): //Lllvar
			//	fmt.Println(reflect.TypeOf(data2.AdditionalPrivateData))
			switch val2 {
			case "AdditionalPrivateData":
				//data2.AdditionalPrivateData = iso8583.NewLllvar([]byte(Field34))
			case "Reserved1":
				newdata, dataString = LoadValueLl(dataString, 3)
				data2.Reserved1 = iso8583.NewLllvar([]byte(newdata))
				//data2.Reserved1 = iso8583.NewLllvar([]byte("201702120000000000000000099999999999999999                                10AB"))
			}
		}
		i = i + 1
	}
	//fmt.Println(string(b[20:26]))
	//data2.No = iso8583.NewNumeric(string(b[20:26]))

	//this will print entire values in struct
	ret, fields, err = msg2.Convert(b)

	if err != nil {
		fmt.Println("we have an error\n")
		fmt.Println(err)
	}
	for _, key := range si {
		fmt.Println("Field:", reflect.ValueOf(fields[key].Index), reflect.ValueOf(fields[key].Length), "Value:", reflect.ValueOf(fields[key].Field))
	}

	// fmt.Println("data2.N0", reflect.ValueOf(data2.No))
	// fmt.Println("data2.N02", reflect.ValueOf(data2.No2))
	// fmt.Println(data2.Ret2)
	// fmt.Println(data2.Rrn)
	// fmt.Println(data2.DateTime)
	fmt.Println(data2.Capturedate)
	// fmt.Println(data2.ActionCode)
	// fmt.Println(data2.CaTerminalId)
	// fmt.Println(data2.CaId)
	// fmt.Println(data2.CurrencyCode)
	// fmt.Printf("%s\n", data2.AcqInsIdeCode)
	// fmt.Printf("%s\n", data2.AccountNumber)
	// fmt.Printf("%s\n", data2.AccountId)
	// fmt.Printf("%s\n", data2.Reserved1)

}
