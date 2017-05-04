package main

import (
	"fmt"
	"iso8583"
	"math/rand"
	"net"
	"os"
	"reflect"
	"sort"
	"strconv"
	"time"
)

// Constant Values
const (
	CONN_HOST    = "localhost"
	CONN_PORT    = "33221"
	CONN_TYPE    = "tcp"
	CONN_TIMEOUT = 10 * time.Second
	Field4       = "0000000000000000"
	Field24      = "200"
	Field32      = "012"
	Field34      = "00000000000000"
	Field49      = "INR"
	Field41      = "MPASSBOOK       "
	Field42      = "MPASSBOOK      "
	Field123     = "MPB"
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

// NewData struct
func NewData() *Data {
	return &Data{
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
	rValue := strconv.Itoa(random)
	data.Rrn = iso8583.NewAlphanumeric(rValue)
	data.Ret2 = iso8583.NewAlphanumeric(rValue)
	timestamp := time.Now()
	data.DateTime = iso8583.NewNumeric(timestamp.Format("20060102150405"))
	data.Capturedate = iso8583.NewNumeric(timestamp.Format("20060102"))
}

// Update Numeric Values
func (data *Data) UpdateNumericValues(val2, dataString string, Length int) string {
	var newdata string
	switch val2 {
	case "No":
		newdata, dataString = LoadValue(dataString, Length)
		data.No = iso8583.NewNumeric(newdata)
	case "No2":
		newdata, dataString = LoadValue(dataString, Length)
		data.No2 = iso8583.NewNumeric(newdata)
	case "DateTime":
		newdata, dataString = LoadValue(dataString, Length)
		data.DateTime = iso8583.NewNumeric(newdata)
	case "Capturedate":
		newdata, dataString = LoadValue(dataString, Length)
		data.Capturedate = iso8583.NewNumeric(newdata)
	case "Functioncode":
		newdata, dataString = LoadValue(dataString, Length)
		data.Functioncode = iso8583.NewNumeric(newdata)
	case "ActionCode":
		newdata, dataString = LoadValue(dataString, Length)
		data.ActionCode = iso8583.NewNumeric(newdata)
	default:
		fmt.Println("fucked")
	}
	return dataString
}

// Update Alphanumeric Values
func (data *Data) UpdateAlphaValues(val2, dataString string, Length int) string {
	var newdata string
	switch val2 {
	case "Ret2":
		newdata, dataString = LoadValue(dataString, Length)
		data.Ret2 = iso8583.NewAlphanumeric(newdata)
	case "Rrn":
		newdata, dataString = LoadValue(dataString, Length)
		data.Rrn = iso8583.NewAlphanumeric(newdata)
	case "CaTerminalId":
		newdata, dataString = LoadValue(dataString, Length)
		data.CaTerminalId = iso8583.NewAlphanumeric(newdata)
	case "CaId":
		newdata, dataString = LoadValue(dataString, Length)
		data.CaId = iso8583.NewAlphanumeric(newdata)
	case "CurrencyCode":
		newdata, dataString = LoadValue(dataString, Length)
		data.CurrencyCode = iso8583.NewAlphanumeric(newdata)
	}
	return dataString
}

//Update Llvar Values
func (data *Data) UpdateLlValues(val2, dataString string) string {
	var newdata string
	switch val2 {
	case "AcqInsIdeCode":
		newdata, dataString = LoadValueLl(dataString, 2)
		data.AcqInsIdeCode = iso8583.NewLlvar([]byte(newdata))
	case "AccountNumber":
		newdata, dataString = LoadValueLl(dataString, 2)
		data.AccountNumber = iso8583.NewLlvar([]byte(newdata))
	case "AccountId":
		newdata, dataString = LoadValueLl(dataString, 2)
		data.AccountId = iso8583.NewLlvar([]byte(newdata))
	default:
		fmt.Println("new fucked")
	}
	return dataString
}

//Update Lllvar Values
func (data *Data) UpdateLllValues(val2, dataString string) string {
	var newdata string
	switch val2 {
	case "AdditionalPrivateData":
		newdata, dataString = LoadValueLl(dataString, 3)
		data.AdditionalPrivateData = iso8583.NewLllvar([]byte(newdata))
	case "DeliveryChannel":
		newdata, dataString = LoadValueLl(dataString, 3)
		data.DeliveryChannel = iso8583.NewLllvar([]byte(newdata))
	case "Reserved1":
		newdata, dataString = LoadValueLl(dataString, 3)
		data.Reserved1 = iso8583.NewLllvar([]byte(newdata))
	}
	return dataString
}

// GetAccountNumber ...
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

// generate new raw and newdata for Llvar & Lllvar
func LoadValueLl(raw string, oldlength int) (string, string) {
	var datastring string
	length, err := strconv.Atoi(raw[0:oldlength])
	if err != nil {
		fmt.Println("error convert convert string to int")
		fmt.Println(raw[0:oldlength])
		// error handling is needed return values
	}
	datastring = raw[oldlength:(oldlength + length)]
	raw = raw[(oldlength + length):]
	return datastring, raw
}

func main() {
	fmt.Println("isoparser the begining")

	// initialing Data struct for 930000
	data := &Data{
		No:        iso8583.NewNumeric("930000"),
		Reserved1: iso8583.NewLllvar([]byte("201702120000000000000000099999999999999999                                10AB")),
	}

	// assigning default values
	data.GetDefaultParams()

	// assigning dynamic values
	data.GetDynamicValues()

	// assigning accountNumber
	data.GetAccountNumber("361103670064591")

	msg := iso8583.NewMessage("1200", data)
	msg.MtiEncode = iso8583.ASCII
	b, err := msg.Bytes()
	if err != nil {
		fmt.Println(err.Error())
	}
	// dial a connection
	ln, err := net.DialTimeout(CONN_TYPE, CONN_HOST+":"+CONN_PORT, time.Duration(10)*time.Second)
	if err != nil {
		fmt.Println("error while creating connection")
		os.Exit(1)
	}

	ln.Write(b)
	msgValue := make([]byte, 8192)

	// set readtimeout
	err = ln.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		fmt.Printf("/n cannot set time out/n")
	}
	_, err = ln.Read(msgValue)
	if err != nil {
		fmt.Printf("/n read time out/n")
	}
	fmt.Println(reflect.TypeOf(b))
	fmt.Println(reflect.TypeOf(msgValue))
	// Close the listener when the application closes.
	defer ln.Close()

	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

	// Assigning empty Data struct for parsing response
	data2 := NewData()

	msg2 := iso8583.NewMessage("1210", data2)
	msg2.MtiEncode = iso8583.ASCII
	_, fields, err := msg2.Convert(msgValue)

	if err != nil {
		fmt.Println("we have an error\n")
		fmt.Println(err)
	}

	si := make([]int, 0, len(fields))
	for i := range fields {
		si = append(si, i)
	}
	sort.Ints(si)

	dataString := string(msgValue[20:])
	i := 0
	for _, key := range si {

		var t interface{}
		t = reflect.TypeOf(fields[key].Field)

		val := reflect.Indirect(reflect.ValueOf(data2))
		val2 := val.Type().Field(i).Name

		switch t {
		default:
			fmt.Println(t)
			fmt.Printf("unexpected type ") // %T prints whatever type t has
		case reflect.TypeOf(data2.No):
			dataString = data2.UpdateNumericValues(val2, dataString, fields[key].Length)
		case reflect.TypeOf(data2.Ret2): // alphanumeric
			dataString = data2.UpdateAlphaValues(val2, dataString, fields[key].Length)
		case reflect.TypeOf(data2.AcqInsIdeCode): //aLlvar
			dataString = data2.UpdateLlValues(val2, dataString)
		case reflect.TypeOf(data2.AdditionalPrivateData): //Lllvar
			dataString = data2.UpdateLllValues(val2, dataString)
		}
		i = i + 1
	}

	// Display parsed reponse
	msg2.PrintValue(msgValue)
	// msg2.Transactions(msgValue)
	data2.Transactions()
	// fmt.Println("data2.N0", reflect.ValueOf(data2.No))
	// fmt.Println("data2.N02", reflect.ValueOf(data2.No2))
	// fmt.Println(data2.Ret2)
	// fmt.Println(data2.Rrn)
	// fmt.Println(data2.DateTime)
	// fmt.Println(data2.Capturedate)
	// fmt.Println(data2.ActionCode)
	// fmt.Println(data2.CaTerminalId)
	// fmt.Println(data2.CaId)
	// fmt.Println(data2.CurrencyCode)
	// fmt.Printf("%s\n", data2.AcqInsIdeCode)
	// fmt.Printf("%s\n", data2.AccountNumber)
	// fmt.Printf("%s\n", data2.AccountId)
	// fmt.Printf("%s\n", data2.Reserved1)
	a := reflect.ValueOf(data2.AdditionalPrivateData).Elem()
	fmt.Printf("%v", a)
}
