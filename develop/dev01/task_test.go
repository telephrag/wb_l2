package develop

import (
	"fmt"
	"testing"
	"time"
)

func TestGetPreciseTime(t *testing.T) {
	if t, err := GetPreciseTime(); err == nil {
		ntp := t.UnixMicro()
		os := time.Now().UnixMicro()
		fmt.Printf("NTP  : %18d\n", ntp)
		fmt.Printf("OS   : %18d\n", os)
		fmt.Printf("DIFF : %18d\n", ntp-os)
	} else {
		fmt.Println(err)
	}
}
