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

type Device_Descriptor struct {
	bLength            uint8
	bDescriptorType    uint8
	bcdUSB             uint16
	bDeviceClass       uint8
	bDeviceSubClass    uint8
	bDeviceProtocol    uint8
	bMaxPacketSize0    uint8
	idVendor           uint16
	idProduct          uint16
	bcdDevice          uint16
	iManufacturer      uint8
	iProduct           uint8
	iSerialNumber      uint8
	bNumConfigurations uint8
}

/*

struct libusb_endpoint_descriptor {
	uint8_t  bLength;
	uint8_t  bDescriptorType;
	uint8_t  bEndpointAddress;
	uint8_t  bmAttributes;
	uint16_t wMaxPacketSize;
	uint8_t  bInterval;
	uint8_t  bRefresh;
	uint8_t  bSynchAddress;
	const unsigned char *extra;
	int extra_length;
};

struct libusb_interface_descriptor {
	uint8_t  bLength;
	uint8_t  bDescriptorType;
	uint8_t  bInterfaceNumber;
	uint8_t  bAlternateSetting;
	uint8_t  bNumEndpoints;
	uint8_t  bInterfaceClass;
	uint8_t  bInterfaceSubClass;
	uint8_t  bInterfaceProtocol;
	uint8_t  iInterface;
	const struct libusb_endpoint_descriptor *endpoint;
	const unsigned char *extra;
	int extra_length;
};

struct libusb_interface {
	const struct libusb_interface_descriptor *altsetting;
	int num_altsetting;
};

struct libusb_config_descriptor {
	uint8_t  bLength;
	uint8_t  bDescriptorType;
	uint16_t wTotalLength;
	uint8_t  bNumInterfaces;
	uint8_t  bConfigurationValue;
	uint8_t  iConfiguration;
	uint8_t  bmAttributes;
	uint8_t  MaxPower;
	const struct libusb_interface *interface;
	const unsigned char *extra;
	int extra_length;
};

struct libusb_ss_endpoint_companion_descriptor {
	uint8_t  bLength;
	uint8_t  bDescriptorType;
	uint8_t  bMaxBurst;
	uint8_t  bmAttributes;
	uint16_t wBytesPerInterval;
};

struct libusb_bos_dev_capability_descriptor {
	uint8_t bLength;
	uint8_t bDescriptorType;
	uint8_t bDevCapabilityType;
	uint8_t dev_capability_data
};

struct libusb_bos_descriptor {
	uint8_t  bLength;
	uint8_t  bDescriptorType;
	uint16_t wTotalLength;
	uint8_t  bNumDeviceCaps;
	struct libusb_bos_dev_capability_descriptor *dev_capability
};

struct libusb_usb_2_0_extension_descriptor {
	uint8_t  bLength;
	uint8_t  bDescriptorType;
	uint8_t  bDevCapabilityType;
	uint32_t  bmAttributes;
};

struct libusb_ss_usb_device_capability_descriptor {
	uint8_t  bLength;
	uint8_t  bDescriptorType;
	uint8_t  bDevCapabilityType;
	uint8_t  bmAttributes;
	uint16_t wSpeedSupported;
	uint8_t  bFunctionalitySupport;
	uint8_t  bU1DevExitLat;
	uint16_t bU2DevExitLat;
};

struct libusb_container_id_descriptor {
	uint8_t  bLength;
	uint8_t  bDescriptorType;
	uint8_t  bDevCapabilityType;
	uint8_t bReserved;
	uint8_t  ContainerID[16];
};

struct libusb_control_setup {
	uint8_t  bmRequestType;
	uint8_t  bRequest;
	uint16_t wValue;
	uint16_t wIndex;
	uint16_t wLength;
};

*/

type Version struct {
	major    uint16
	minor    uint16
	micro    uint16
	nano     uint16
	rc       string
	describe string
}

//-----------------------------------------------------------------------------
// errors

type libusb_error_t struct {
	name string
	code int
}

func (e *libusb_error_t) Error() string {
	return fmt.Sprintf("libusb_error: %s returned %d(%s)", e.name, e.code, Error_Name(e.code))
}

func libusb_error(name string, code int) error {
	return &libusb_error_t{
		name: name,
		code: code,
	}
}

//-----------------------------------------------------------------------------

type Context *C.struct_libusb_context

func Init(ctx *Context) error {
	rc := int(C.libusb_init((**C.struct_libusb_context)(ctx)))
	if rc != LIBUSB_SUCCESS {
		return libusb_error("libusb_init", rc)
	}
	return nil
}

func Exit(ctx Context) {
	C.libusb_exit(ctx)
}

func Set_Debug(ctx Context, level int) {
	C.libusb_set_debug(ctx, C.int(level))
}

func Get_Version() *Version {
	ver := (*C.struct_libusb_version)(unsafe.Pointer(C.libusb_get_version()))
	return &Version{
		major:    uint16(ver.major),
		minor:    uint16(ver.minor),
		micro:    uint16(ver.micro),
		nano:     uint16(ver.nano),
		rc:       C.GoString(ver.rc),
		describe: C.GoString(ver.describe),
	}
}

func Error_Name(code int) string {
	return C.GoString(C.libusb_error_name(C.int(code)))
}

//func Get_Device_List() ([]Device, error) {
//ssize_t LIBUSB_CALL libusb_get_device_list(libusb_context *ctx, libusb_device ***list);
//}

//func Free_Device_List() ([]Device, error) {
//void libusb_free_device_list	(	libusb_device ** 	list, int 	unref_devices )
//}

//-----------------------------------------------------------------------------
