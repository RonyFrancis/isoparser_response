package main

import (
	"fmt"
	"iso8583"
	"net"
	"os"
)

const (
	CONN_HOST = "localhost"
	CONN_PORT = "33221"
	CONN_TYPE = "tcp"
)

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

func main() {
	fmt.Println("isoparser the begining")
	// Listen for incoming connections.
	// l, err := net.Listen(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	// if err != nil {
	// 	fmt.Println("Error listening:", err.Error())
	// 	os.Exit(1)
	// }
	// // Close the listener when the application closes.
	// defer l.Close()
	// fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	// for {
	// 	// Listen for an incoming connection.
	// 	conn, err := l.Accept()
	// 	if err != nil {
	// 		fmt.Println("Error accepting: ", err.Error())
	// 		os.Exit(1)
	// 	}
	// 	// Handle connections in a new goroutine.
	// 	go handleRequest(conn)
	// }
	data := &Data{
		No:              iso8583.NewNumeric("930000"),
		No2:             iso8583.NewNumeric("0000000000000000"),
		Ret2:            iso8583.NewAlphanumeric("131968991659"),
		DateTime:        iso8583.NewNumeric("20170404144019"),
		Capturedate:     iso8583.NewNumeric("20170404"),
		Functioncode:    iso8583.NewNumeric("200"),
		AcqInsIdeCode:   iso8583.NewLlvar([]byte("012")),
		AccountNumber:   iso8583.NewLlvar([]byte("29040100007529")),
		Rrn:             iso8583.NewAlphanumeric("131968991659"),
		CaTerminalId:    iso8583.NewAlphanumeric("MPASSBOOK"),
		CaId:            iso8583.NewAlphanumeric("MPASSBOOK"),
		CurrencyCode:    iso8583.NewAlphanumeric("INR"),
		AccountId:       iso8583.NewLlvar([]byte("012        3611    361103670062573")),
		DeliveryChannel: iso8583.NewLllvar([]byte("MPB")),
		Reserved1:       iso8583.NewLllvar([]byte("201702120000000000000000099999999999999999                                10AB")),
	}
	msg := iso8583.NewMessage("1200", data)
	fmt.Println("msg values ===================>")
	fmt.Println(string(msg.Mti))
	fmt.Println(msg.SecondBitmap)
	msg.MtiEncode = iso8583.ASCII
	fmt.Printf("%s\n", msg)
	fmt.Println("msg values ===================>")
	b, err := msg.Bytes()
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("stringgggggggggggggggg  bbbbbbbbbbbbbbbbbbb\n")
	fmt.Printf(string(b))
	fmt.Printf("\n")
	// dst := new(bytes.Buffer)

	// src := []byte("30200040020c0001")

	// json.HTMLEscape(dst, src)

	// fmt.Println(dst)
	// fmt.Printf("heelllllllll")
	// //sample := "\xB00\x80\x01J\xC1\x80\x00\x00\x00\x00\x00\x00\x00\x00"
	// fmt.Printf("&&&&&&&&&&&&&&&&&&&&&\n\n")

	// this is working i think
	// sample := "B030810148C080000000000004000028"

	// bs, err := hex.DecodeString(sample)
	// if err != nil {
	// 	panic(err)
	// }

	//fmt.Printf("%x\n", bs)
	//fmt.Printf("%U\n", bs)
	//fmt.Printf("&&&&&&&&&&&&&&&&&&&&&\n\n")
	//fmt.Printf("\n%q\n", sample)
	// for i := 0; i < len(sample); i++ {
	// 	fmt.Printf("%c", sample[i])
	// 	//fmt.Printf("%U\n", sample[i])
	// }
	// fmt.Printf("%#U", sample)
	//fmt.Printf("\n%x\n", sample)
	// a := "b03080014ac1800000000000000000"
	// fmt.Printf("\n%q\n", a)
	// code := "1200"
	// //string := append(code, bs...) // + "93000000000000000000005352304303672017030418101820170304200030121429040100007529535230430367MPASSBOOK       MPASSBOOK      INR34012        3611    361103670062573003MPB078201702120000000000000000099999999999999999                                10AB"
	// string := "93000000000000000000001703172255912017040320543720170403200030121429040100007529170317225591MPASSBOOK       MPASSBOOK      INR34012        3611    361103670062573003MPB078201702120000000000000000099999999999999999                                10AB"
	// fmt.Printf("heelllllllll\n")
	// fmt.Printf("code is %s\n", code)
	// fmt.Printf("%s\n", bs)
	// appendString := append([]byte(code), bs...)
	// appendString = append(appendString, string...)
	// fmt.Printf("%s\n", string)
	// fmt.Printf("appended string: %s\n", appendString)
	// //string2 := append(string, dataparams...)
	// fmt.Printf("heelllllllll\n")
	//encodedStr := hex.EncodeToString(b)
	//fmt.Printf("%s\n", encodedStr)

	// dial a connection
	ln, err := net.Dial(CONN_TYPE, CONN_HOST+":"+CONN_PORT)
	if err != nil {
		fmt.Println("error while creating connection")
		os.Exit(1)
	}
	//for {

	fmt.Println("string is")
	//fmt.Println(string)
	//fmt.Println(len(appendString))
	ln.Write(b)
	msgValue := make([]byte, 8192)
	ln.Read(msgValue)
	fmt.Printf(string(msgValue))
	// Close the listener when the application closes.
	defer ln.Close()

	fmt.Println("Listening on " + CONN_HOST + ":" + CONN_PORT)
	// for {
	// 	// Listen for an incoming connection.
	// 	conn2, err := ln.Accept()
	// 	if err != nil {
	// 		fmt.Println("Error accepting: ", err.Error())
	// 		os.Exit(1)
	// 	}
	// 	// Handle connections in a new goroutine.
	// 	go handleRequest(conn2)
	// }
	//}

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
