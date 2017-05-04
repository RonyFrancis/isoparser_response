package iso8583

import (
	"fmt"
	//"reflect"
)

// // Transactions ... extracting transactions
// func (m *Message) Transactions(raw []byte) {
// 	fields := parseFields(m.Data)
// 	fieldValue := reflect.ValueOf(fields[48].Field)
// 	field48 := reflect.Indirect(fieldValue).FieldByName("Value")
// 	fmt.Println(field48)
// }

// transaction ... New function for Data
func (data *main.Data) Transactions() {
	fmt.Println("aaaaaaaaaaa")
	fmt.Println(data.AdditionalPrivateData)
}
