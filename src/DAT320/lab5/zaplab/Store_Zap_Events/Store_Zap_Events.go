// Store_Zap_Events

package Store_Zap_Events

import (
	"fmt"
	"net"
	"strings"
	"time"
)

const timeLayout = "2006/01/02, 15:04:05"

type StoreZapEvents struct {
	Date   time.Time //time.Time can take both date and time
	IP     net.IP
	FromCh string
	ToCh   string
}

//Constructor for struct function
func NewStoreZapEvents(received_data string) *StoreZapEvents {
	received_data_slice := strings.Split(received_data, ", ")
	date_and_time := strings.Join(received_data_slice[0:2], ", ")
	parsed_date_and_time, _ := time.Parse(timeLayout, date_and_time)
	ip := net.ParseIP(received_data_slice[2])
	fromchan := received_data_slice[3]
	tochan := received_data_slice[4]
	//fmt.Printf("INGENTING\n")
	//fmt.Printf("%s %s %s %s\n", parsed_date_and_time, ip, fromchan, tochan)
	return &StoreZapEvents{parsed_date_and_time, ip, fromchan, tochan}
}

//This method allow Go's fmt package print methods will use it to print StoreZapEvents values as stirng
func (SZE *StoreZapEvents) String() string {
	Date := SZE.Date.Format(timeLayout)
	IP := SZE.IP.String()
	FromCh := SZE.FromCh
	ToCh := SZE.ToCh
	return fmt.Sprintf("%s %s %s %s", Date, IP, FromCh, ToCh)
}

func (SZE StoreZapEvents) Duration(provided StoreZapEvents) time.Duration {
	return SZE.Date.Sub(provided.Date)
}
