package main

import (
	"fmt"
	"iso8583"
	"math/rand"
	"net"
	"os"
	"reflect"
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
	// dial a connection
	ln, err := net.DialTimeout(CONN_TYPE, CONN_HOST+":"+CONN_PORT, time.Duration(10)*time.Second)
	if err != nil {
		fmt.Println("error while creating connection")
		os.Exit(1)
	}

	fmt.Println("string is")
	ln.Write(b)
	parser := &iso8583.Parser{}
	parser.MtiEncode = iso8583.ASCII
	//msg2 := &Data{}
	err = parser.Register("1200", data)

	if err != nil {
		fmt.Printf("parser ssssssssssssssssssn\n")
	}
	m, err := parser.Parse(b)

	if err != nil {
		fmt.Printf("parser issue\n")
	}
	fmt.Printf("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\n")
	fmt.Println(reflect.ValueOf(m))
	fmt.Printf("\naaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa\n")
	msgValue := make([]byte, 8192)
	err = ln.SetReadDeadline(time.Now().Add(5 * time.Second))
	if err != nil {
		fmt.Printf("/n cannot set time out/n")
	}
	_, err = ln.Read(msgValue)
	if err != nil {
		fmt.Printf("/n read time out/n")
	}
	fmt.Printf(string(msgValue))
	// Close the listener when the application closes.
	defer ln.Close()

	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)

}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	buf := make([]byte, 8192)
	conn.Write([]byte("Message received."))
	// Read the incoming connection into the buffer.
	_, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err.Error())
	}
	fmt.Printf(string(buf))
	// Send a response back to person contacting us.

	// Close the connection when you're done with it.
	conn.Close()
}
