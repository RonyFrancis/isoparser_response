package iso8583

import (
	"fmt"
	"reflect"
	"sort"
)

// PrintValue ... print the values of the message
func (m *Message) PrintValue(raw []byte) {
	fields := parseFields(m.Data)
	fmt.Println(reflect.ValueOf(fields))
	si := make([]int, 0, len(fields))
	for i := range fields {
		si = append(si, i)
	}
	sort.Ints(si)
	fmt.Println(reflect.ValueOf(m.Data))
	for _, key := range si {
		fmt.Println("Field:", reflect.ValueOf(fields[key].Index),
			reflect.ValueOf(fields[key].LenEncode), "Value:",
			reflect.ValueOf(fields[key].Field))
	}
}
