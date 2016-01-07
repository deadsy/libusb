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
	"reflect"
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

const (
	LIBUSB_ENDPOINT_IN  = C.LIBUSB_ENDPOINT_IN  // In: device-to-host.
	LIBUSB_ENDPOINT_OUT = C.LIBUSB_ENDPOINT_OUT // Out: host-to-device.
)

const LIBUSB_API_VERSION = C.LIBUSB_API_VERSION

//-----------------------------------------------------------------------------
// structures

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

type Device_Descriptor C.struct_libusb_device_descriptor

type Context *C.struct_libusb_context
type Device *C.struct_libusb_device
type Device_Handle *C.struct_libusb_device_handle
type Hotplug_Callback *C.struct_libusb_hotplug_callback

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

func Error_Name(code int) string {
	return C.GoString(C.libusb_error_name(C.int(code)))
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

//-----------------------------------------------------------------------------
// Library initialization/deinitialization

func Set_Debug(ctx Context, level int) {
	C.libusb_set_debug(ctx, C.int(level))
}

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

//-----------------------------------------------------------------------------
// Device handling and enumeration

func Get_Device_List(ctx Context) ([]Device, error) {
	var hdl **C.struct_libusb_device
	rc := int(C.libusb_get_device_list(ctx, (***C.struct_libusb_device)(&hdl)))
	if rc < 0 {
		return nil, libusb_error("libusb_get_device_list", rc)
	}
	// turn the c array into a slice of device pointers
	var list []Device
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&list))
	hdr.Cap = rc
	hdr.Len = rc
	hdr.Data = uintptr(unsafe.Pointer(hdl))
	return list, nil
}

func Free_Device_List(list []Device, unref_devices int) {
	if list == nil {
		return
	}
	C.libusb_free_device_list((**C.struct_libusb_device)(&list[0]), C.int(unref_devices))
}

func Get_Bus_Number(dev Device) uint8 {
	return uint8(C.libusb_get_bus_number(dev))
}

func Get_Port_Number(dev Device) uint8 {
	return uint8(C.libusb_get_port_number(dev))
}

func Get_Port_Numbers(dev Device) ([]uint8, error) {
	ports := make([]uint8, 16)
	rc := int(C.libusb_get_port_numbers(dev, (*C.uint8_t)(&ports[0]), (C.int)(len(ports))))
	if rc < 0 {
		return nil, libusb_error("libusb_get_port_numbers", rc)
	}
	return ports[:rc], nil
}

func Get_Parent(dev Device) Device {
	return C.libusb_get_parent(dev)
}

func Get_Device_Address(dev Device) uint8 {
	return uint8(C.libusb_get_device_address(dev))
}

func Get_Device_Speed(dev Device) int {
	return int(C.libusb_get_device_speed(dev))
}

func Get_Max_Packet_Size(dev Device, endpoint uint8) int {
	return int(C.libusb_get_max_packet_size(dev, (C.uchar)(endpoint)))
}

func Get_Max_ISO_Packet_Size(dev Device, endpoint uint8) int {
	return int(C.libusb_get_max_iso_packet_size(dev, (C.uchar)(endpoint)))
}

func Ref_Device(dev Device) Device {
	return C.libusb_ref_device(dev)
}

func Unref_Device(dev Device) {
	C.libusb_unref_device(dev)
}

func Open(dev Device) (Device_Handle, error) {
	var hdl Device_Handle
	rc := int(C.libusb_open(dev, (**C.struct_libusb_device_handle)(&hdl)))
	if rc < 0 {
		return nil, libusb_error("libusb_open", rc)
	}
	return hdl, nil
}

func Open_Device_With_VID_PID(ctx Context, vendor_id uint16, product_id uint16) Device_Handle {
	return C.libusb_open_device_with_vid_pid(ctx, (C.uint16_t)(vendor_id), (C.uint16_t)(product_id))
}

func Close(hdl Device_Handle) {
	C.libusb_close(hdl)
}

func Get_Device(hdl Device_Handle) Device {
	return C.libusb_get_device(hdl)
}

//int 	libusb_get_configuration (libusb_device_handle *dev, int *config)
//int 	libusb_set_configuration (libusb_device_handle *dev, int configuration)
//int 	libusb_claim_interface (libusb_device_handle *dev, int interface_number)
//int 	libusb_release_interface (libusb_device_handle *dev, int interface_number)
//int 	libusb_set_interface_alt_setting (libusb_device_handle *dev, int interface_number, int alternate_setting)
//int 	libusb_clear_halt (libusb_device_handle *dev, unsigned char endpoint)
//int 	libusb_reset_device (libusb_device_handle *dev)
//int 	libusb_kernel_driver_active (libusb_device_handle *dev, int interface_number)
//int 	libusb_detach_kernel_driver (libusb_device_handle *dev, int interface_number)
//int 	libusb_attach_kernel_driver (libusb_device_handle *dev, int interface_number)
//int 	libusb_set_auto_detach_kernel_driver (libusb_device_handle *dev, int enable)

//-----------------------------------------------------------------------------
// USB descriptors

func Get_Device_Descriptor(dev Device) (*Device_Descriptor, error) {
	var dd Device_Descriptor
	rc := int(C.libusb_get_device_descriptor(dev, (*C.struct_libusb_device_descriptor)(&dd)))
	if rc != LIBUSB_SUCCESS {
		return nil, libusb_error("libusb_get_device_descriptor", rc)
	}
	return &dd, nil
}

//-----------------------------------------------------------------------------
// Synchronous device I/O

func Control_Transfer(hdl Device_Handle, bmRequestType uint8, bRequest uint8, wValue uint16, wIndex uint16, data []byte, timeout uint) ([]byte, error) {
	//int 	libusb_control_transfer (libusb_device_handle *dev_handle, uint8_t bmRequestType, uint8_t bRequest, uint16_t wValue, uint16_t wIndex, unsigned char *data, uint16_t wLength, unsigned int timeout)
	return nil, nil
}

func Bulk_Transfer(hdl Device_Handle, endpoint uint8, data []byte, timeout uint) ([]byte, error) {
	var length int
	var transferred C.int
	if endpoint&LIBUSB_ENDPOINT_IN != 0 {
		// read device
		length = cap(data)
	} else {
		// write device
		length = len(data)
	}
	rc := int(C.libusb_bulk_transfer(hdl, (C.uchar)(endpoint), (*C.uchar)(&data[0]), (C.int)(length), &transferred, (C.uint)(timeout)))
	if rc != LIBUSB_SUCCESS {
		return nil, libusb_error("libusb_bulk_transfer", rc)
	}
	return data[:int(transferred)], nil
}

func Interrupt_Transfer(hdl Device_Handle, endpoint uint8, data []byte, timeout uint) ([]byte, error) {
	var length int
	var transferred C.int
	if endpoint&LIBUSB_ENDPOINT_IN != 0 {
		// read device
		length = cap(data)
	} else {
		// write device
		length = len(data)
	}
	rc := int(C.libusb_interrupt_transfer(hdl, (C.uchar)(endpoint), (*C.uchar)(&data[0]), (C.int)(length), &transferred, (C.uint)(timeout)))
	if rc != LIBUSB_SUCCESS {
		return nil, libusb_error("libusb_interrupt_transfer", rc)
	}
	return data[:int(transferred)], nil
}

//-----------------------------------------------------------------------------
