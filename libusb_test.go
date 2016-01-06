//-----------------------------------------------------------------------------
/*

Test functions for libusb-1.0

*/
//-----------------------------------------------------------------------------

package libusb

import (
	"log"
	"os"
	"testing"
)

//-----------------------------------------------------------------------------

var logger = log.New(os.Stdout, "", log.Lshortfile)

//-----------------------------------------------------------------------------

func Test_Error_Name(t *testing.T) {
	if Error_Name(LIBUSB_ERROR_BUSY) != "LIBUSB_ERROR_BUSY" {
		t.Error("FAIL")
	}
}

func Test_Device_List(t *testing.T) {
	var ctx Context
	err := Init(&ctx)
	defer Exit(ctx)
	if err != nil {
		t.Error("FAIL")
	}
	list, err := Get_Device_List(ctx)
	if err != nil {
		t.Error("FAIL")
	}
	logger.Printf("%d devices %v\n", len(list), list)
	Free_Device_List(list, 1)
}

func Test_Version(t *testing.T) {
	logger.Printf("%+v\n", Get_Version())
}

func Test_Init_Exit(t *testing.T) {
	var ctx Context
	err := Init(&ctx)
	defer Exit(ctx)
	if err != nil {
		t.Error("FAIL")
	}
}

//-----------------------------------------------------------------------------
