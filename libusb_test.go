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

func Test_Error_Name(t *testing.T) {
	if Error_Name(LIBUSB_ERROR_BUSY) != "LIBUSB_ERROR_BUSY" {
		t.Error("FAIL")
	}
}

func Test_Version(t *testing.T) {
	var ctx Context
	err := Init(&ctx)
	defer Exit(ctx)
	if err != nil {
		t.Error("FAIL")
	}
	logger.Printf("%+v\n", Get_Version())
}

//-----------------------------------------------------------------------------
