//-----------------------------------------------------------------------------
/*

Test functions for libusb-1.0

*/
//-----------------------------------------------------------------------------

package libusb

import (
	"fmt"
	"testing"
)

//-----------------------------------------------------------------------------

func Test_Error_Name(t *testing.T) {
	err := Init()
	defer Exit()
	if err != nil {
		t.Error("FAIL")
	}
	fmt.Printf("%s\n", Error_Name(LIBUSB_ERROR_BUSY))
}

func Test_Version(t *testing.T) {
	err := Init()
	defer Exit()
	if err != nil {
		t.Error("FAIL")
	}
	fmt.Printf("%+v\n", Get_Version())
}

//-----------------------------------------------------------------------------
