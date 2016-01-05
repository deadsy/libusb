//-----------------------------------------------------------------------------
/*

Golang wrapper for libusb-1.0

*/
//-----------------------------------------------------------------------------

package libusb

/*
#cgo LDFLAGS: -lusb-1.0
#include <libusb-1.0/libusb.h>
*/
import "C"

import (
	"fmt"
	"path"
	"runtime"
	"unsafe"
)

//-----------------------------------------------------------------------------

const (
	LIBUSB_LOG_LEVEL_NONE    = C.LIBUSB_LOG_LEVEL_NONE
	LIBUSB_LOG_LEVEL_ERROR   = C.LIBUSB_LOG_LEVEL_ERROR
	LIBUSB_LOG_LEVEL_WARNING = C.LIBUSB_LOG_LEVEL_WARNING
	LIBUSB_LOG_LEVEL_INFO    = C.LIBUSB_LOG_LEVEL_INFO
	LIBUSB_LOG_LEVEL_DEBUG   = C.LIBUSB_LOG_LEVEL_DEBUG
)

const (
	LIBUSB_SUCCESS             = C.LIBUSB_SUCCESS
	LIBUSB_ERROR_IO            = C.LIBUSB_ERROR_IO
	LIBUSB_ERROR_INVALID_PARAM = C.LIBUSB_ERROR_INVALID_PARAM
	LIBUSB_ERROR_ACCESS        = C.LIBUSB_ERROR_ACCESS
	LIBUSB_ERROR_NO_DEVICE     = C.LIBUSB_ERROR_NO_DEVICE
	LIBUSB_ERROR_NOT_FOUND     = C.LIBUSB_ERROR_NOT_FOUND
	LIBUSB_ERROR_BUSY          = C.LIBUSB_ERROR_BUSY
	LIBUSB_ERROR_TIMEOUT       = C.LIBUSB_ERROR_TIMEOUT
	LIBUSB_ERROR_OVERFLOW      = C.LIBUSB_ERROR_OVERFLOW
	LIBUSB_ERROR_PIPE          = C.LIBUSB_ERROR_PIPE
	LIBUSB_ERROR_INTERRUPTED   = C.LIBUSB_ERROR_INTERRUPTED
	LIBUSB_ERROR_NO_MEM        = C.LIBUSB_ERROR_NO_MEM
	LIBUSB_ERROR_NOT_SUPPORTED = C.LIBUSB_ERROR_NOT_SUPPORTED
	LIBUSB_ERROR_OTHER         = C.LIBUSB_ERROR_OTHER
)

const LIBUSB_API_VERSION = C.LIBUSB_API_VERSION

//-----------------------------------------------------------------------------
// structures

type Version struct {
	major    int
	minor    int
	micro    int
	nano     int
	rc       string
	describe string
}

//-----------------------------------------------------------------------------
// errors

type libusb_error_t struct {
	code int
	file string
	line int
}

func (e *libusb_error_t) Error() string {
	return fmt.Sprintf("libusb_error(%s, line %d): %s(%d)", path.Base(e.file), e.line,
		Error_Name(e.code), e.code)
}

func libusb_error(code int) error {
	var e libusb_error_t
	_, file, line, ok := runtime.Caller(1)
	if ok {
		e.file = file
		e.line = line
	}
	e.code = code
	return &e
}

//-----------------------------------------------------------------------------

func Init() error {
	rc := int(C.libusb_init(nil))
	if rc != LIBUSB_SUCCESS {
		return libusb_error(rc)
	}
	return nil
}

func Exit() {
	C.libusb_exit(nil)
}

func Set_Debug(level int) {
	C.libusb_set_debug(nil, C.int(level))
}

func Get_Version() *Version {
	ver := (*C.struct_libusb_version)(unsafe.Pointer(C.libusb_get_version()))
	return &Version{
		major:    int(ver.major),
		minor:    int(ver.minor),
		micro:    int(ver.micro),
		nano:     int(ver.nano),
		rc:       C.GoString(ver.rc),
		describe: C.GoString(ver.describe),
	}
}

func Error_Name(code int) string {
	return C.GoString(C.libusb_error_name(C.int(code)))
}

//-----------------------------------------------------------------------------
