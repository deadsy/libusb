//-----------------------------------------------------------------------------
/*

Golang wrapper for libusb-1.0

*/
//-----------------------------------------------------------------------------

package libusb

/*
#cgo LDFLAGS: -lusb-1.0
#include <libusb-1.0/libusb.h>

// cgo has trouble seeing x.dev_capability_data because it is a trailing []/[0] field
static uint8_t *dev_capability_data_ptr(struct libusb_bos_dev_capability_descriptor *x) {
  return &x->dev_capability_data[0];
}

// cgo has trouble seeing x.dev_capability because it is a trailing []/[0] field
static struct libusb_bos_dev_capability_descriptor **dev_capability_ptr(struct libusb_bos_descriptor *x) {
  return &x->dev_capability[0];
}

*/
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"
)

//-----------------------------------------------------------------------------

const (
	LIBUSB_CLASS_PER_INTERFACE       = C.LIBUSB_CLASS_PER_INTERFACE
	LIBUSB_CLASS_AUDIO               = C.LIBUSB_CLASS_AUDIO
	LIBUSB_CLASS_COMM                = C.LIBUSB_CLASS_COMM
	LIBUSB_CLASS_HID                 = C.LIBUSB_CLASS_HID
	LIBUSB_CLASS_PHYSICAL            = C.LIBUSB_CLASS_PHYSICAL
	LIBUSB_CLASS_PRINTER             = C.LIBUSB_CLASS_PRINTER
	LIBUSB_CLASS_PTP                 = C.LIBUSB_CLASS_PTP
	LIBUSB_CLASS_IMAGE               = C.LIBUSB_CLASS_IMAGE
	LIBUSB_CLASS_MASS_STORAGE        = C.LIBUSB_CLASS_MASS_STORAGE
	LIBUSB_CLASS_HUB                 = C.LIBUSB_CLASS_HUB
	LIBUSB_CLASS_DATA                = C.LIBUSB_CLASS_DATA
	LIBUSB_CLASS_SMART_CARD          = C.LIBUSB_CLASS_SMART_CARD
	LIBUSB_CLASS_CONTENT_SECURITY    = C.LIBUSB_CLASS_CONTENT_SECURITY
	LIBUSB_CLASS_VIDEO               = C.LIBUSB_CLASS_VIDEO
	LIBUSB_CLASS_PERSONAL_HEALTHCARE = C.LIBUSB_CLASS_PERSONAL_HEALTHCARE
	LIBUSB_CLASS_DIAGNOSTIC_DEVICE   = C.LIBUSB_CLASS_DIAGNOSTIC_DEVICE
	LIBUSB_CLASS_WIRELESS            = C.LIBUSB_CLASS_WIRELESS
	LIBUSB_CLASS_APPLICATION         = C.LIBUSB_CLASS_APPLICATION
	LIBUSB_CLASS_VENDOR_SPEC         = C.LIBUSB_CLASS_VENDOR_SPEC
)

const (
	LIBUSB_DT_DEVICE                = C.LIBUSB_DT_DEVICE
	LIBUSB_DT_CONFIG                = C.LIBUSB_DT_CONFIG
	LIBUSB_DT_STRING                = C.LIBUSB_DT_STRING
	LIBUSB_DT_INTERFACE             = C.LIBUSB_DT_INTERFACE
	LIBUSB_DT_ENDPOINT              = C.LIBUSB_DT_ENDPOINT
	LIBUSB_DT_BOS                   = C.LIBUSB_DT_BOS
	LIBUSB_DT_DEVICE_CAPABILITY     = C.LIBUSB_DT_DEVICE_CAPABILITY
	LIBUSB_DT_HID                   = C.LIBUSB_DT_HID
	LIBUSB_DT_REPORT                = C.LIBUSB_DT_REPORT
	LIBUSB_DT_PHYSICAL              = C.LIBUSB_DT_PHYSICAL
	LIBUSB_DT_HUB                   = C.LIBUSB_DT_HUB
	LIBUSB_DT_SUPERSPEED_HUB        = C.LIBUSB_DT_SUPERSPEED_HUB
	LIBUSB_DT_SS_ENDPOINT_COMPANION = C.LIBUSB_DT_SS_ENDPOINT_COMPANION
)

const (
	LIBUSB_REQUEST_TYPE_STANDARD = C.LIBUSB_REQUEST_TYPE_STANDARD
	LIBUSB_REQUEST_TYPE_CLASS    = C.LIBUSB_REQUEST_TYPE_CLASS
	LIBUSB_REQUEST_TYPE_VENDOR   = C.LIBUSB_REQUEST_TYPE_VENDOR
	LIBUSB_REQUEST_TYPE_RESERVED = C.LIBUSB_REQUEST_TYPE_RESERVED
)

const (
	LIBUSB_REQUEST_GET_STATUS        = C.LIBUSB_REQUEST_GET_STATUS
	LIBUSB_REQUEST_CLEAR_FEATURE     = C.LIBUSB_REQUEST_CLEAR_FEATURE
	LIBUSB_REQUEST_SET_FEATURE       = C.LIBUSB_REQUEST_SET_FEATURE
	LIBUSB_REQUEST_SET_ADDRESS       = C.LIBUSB_REQUEST_SET_ADDRESS
	LIBUSB_REQUEST_GET_DESCRIPTOR    = C.LIBUSB_REQUEST_GET_DESCRIPTOR
	LIBUSB_REQUEST_SET_DESCRIPTOR    = C.LIBUSB_REQUEST_SET_DESCRIPTOR
	LIBUSB_REQUEST_GET_CONFIGURATION = C.LIBUSB_REQUEST_GET_CONFIGURATION
	LIBUSB_REQUEST_SET_CONFIGURATION = C.LIBUSB_REQUEST_SET_CONFIGURATION
	LIBUSB_REQUEST_GET_INTERFACE     = C.LIBUSB_REQUEST_GET_INTERFACE
	LIBUSB_REQUEST_SET_INTERFACE     = C.LIBUSB_REQUEST_SET_INTERFACE
	LIBUSB_REQUEST_SYNCH_FRAME       = C.LIBUSB_REQUEST_SYNCH_FRAME
	LIBUSB_REQUEST_SET_SEL           = C.LIBUSB_REQUEST_SET_SEL
	LIBUSB_SET_ISOCH_DELAY           = C.LIBUSB_SET_ISOCH_DELAY
)

const (
	LIBUSB_RECIPIENT_DEVICE    = C.LIBUSB_RECIPIENT_DEVICE
	LIBUSB_RECIPIENT_INTERFACE = C.LIBUSB_RECIPIENT_INTERFACE
	LIBUSB_RECIPIENT_ENDPOINT  = C.LIBUSB_RECIPIENT_ENDPOINT
	LIBUSB_RECIPIENT_OTHER     = C.LIBUSB_RECIPIENT_OTHER
)

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

const LIBUSB_TRANSFER_TYPE_MASK = C.LIBUSB_TRANSFER_TYPE_MASK

const (
	LIBUSB_TRANSFER_TYPE_CONTROL     = C.LIBUSB_TRANSFER_TYPE_CONTROL
	LIBUSB_TRANSFER_TYPE_ISOCHRONOUS = C.LIBUSB_TRANSFER_TYPE_ISOCHRONOUS
	LIBUSB_TRANSFER_TYPE_BULK        = C.LIBUSB_TRANSFER_TYPE_BULK
	LIBUSB_TRANSFER_TYPE_INTERRUPT   = C.LIBUSB_TRANSFER_TYPE_INTERRUPT
	LIBUSB_TRANSFER_TYPE_BULK_STREAM = C.LIBUSB_TRANSFER_TYPE_BULK_STREAM
)

const LIBUSB_API_VERSION = C.LIBUSB_API_VERSION

//-----------------------------------------------------------------------------
// structures

type Endpoint_Descriptor struct {
	ptr              *C.struct_libusb_endpoint_descriptor
	BLength          uint8
	BDescriptorType  uint8
	BEndpointAddress uint8
	BmAttributes     uint8
	WMaxPacketSize   uint16
	BInterval        uint8
	BRefresh         uint8
	BSynchAddress    uint8
	Extra            []byte
}

func c2go_Endpoint_Descriptor(x *C.struct_libusb_endpoint_descriptor) *Endpoint_Descriptor {
	return &Endpoint_Descriptor{
		ptr:              x,
		BLength:          uint8(x.bLength),
		BDescriptorType:  uint8(x.bDescriptorType),
		BEndpointAddress: uint8(x.bEndpointAddress),
		BmAttributes:     uint8(x.bmAttributes),
		WMaxPacketSize:   uint16(x.wMaxPacketSize),
		BInterval:        uint8(x.bInterval),
		BRefresh:         uint8(x.bRefresh),
		BSynchAddress:    uint8(x.bSynchAddress),
		Extra:            C.GoBytes(unsafe.Pointer(x.extra), x.extra_length),
	}
}

type Interface_Descriptor struct {
	ptr                *C.struct_libusb_interface_descriptor
	BLength            uint8
	BDescriptorType    uint8
	BInterfaceNumber   uint8
	BAlternateSetting  uint8
	BNumEndpoints      uint8
	BInterfaceClass    uint8
	BInterfaceSubClass uint8
	BInterfaceProtocol uint8
	IInterface         uint8
	Endpoint           []*Endpoint_Descriptor
	Extra              []byte
}

func c2go_Interface_Descriptor(x *C.struct_libusb_interface_descriptor) *Interface_Descriptor {
	var list []C.struct_libusb_endpoint_descriptor
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&list))
	hdr.Cap = int(x.bNumEndpoints)
	hdr.Len = int(x.bNumEndpoints)
	hdr.Data = uintptr(unsafe.Pointer(x.endpoint))
	endpoints := make([]*Endpoint_Descriptor, x.bNumEndpoints)
	for i, _ := range endpoints {
		endpoints[i] = c2go_Endpoint_Descriptor(&list[i])
	}
	return &Interface_Descriptor{
		ptr:                x,
		BLength:            uint8(x.bLength),
		BDescriptorType:    uint8(x.bDescriptorType),
		BInterfaceNumber:   uint8(x.bInterfaceNumber),
		BAlternateSetting:  uint8(x.bAlternateSetting),
		BNumEndpoints:      uint8(x.bNumEndpoints),
		BInterfaceClass:    uint8(x.bInterfaceClass),
		BInterfaceSubClass: uint8(x.bInterfaceSubClass),
		BInterfaceProtocol: uint8(x.bInterfaceProtocol),
		IInterface:         uint8(x.iInterface),
		Endpoint:           endpoints,
		Extra:              C.GoBytes(unsafe.Pointer(x.extra), x.extra_length),
	}
}

type Interface struct {
	ptr            *C.struct_libusb_interface
	Num_altsetting int
	Altsetting     []*Interface_Descriptor
}

func c2go_Interface(x *C.struct_libusb_interface) *Interface {
	var list []C.struct_libusb_interface_descriptor
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&list))
	hdr.Cap = int(x.num_altsetting)
	hdr.Len = int(x.num_altsetting)
	hdr.Data = uintptr(unsafe.Pointer(x.altsetting))
	altsetting := make([]*Interface_Descriptor, x.num_altsetting)
	for i, _ := range altsetting {
		altsetting[i] = c2go_Interface_Descriptor(&list[i])
	}
	return &Interface{
		ptr:            x,
		Num_altsetting: int(x.num_altsetting),
		Altsetting:     altsetting,
	}
}

type Config_Descriptor struct {
	ptr                 *C.struct_libusb_config_descriptor
	BLength             uint8
	BDescriptorType     uint8
	WTotalLength        uint16
	BNumInterfaces      uint8
	BConfigurationValue uint8
	IConfiguration      uint8
	BmAttributes        uint8
	MaxPower            uint8
	Interface           []*Interface
	Extra               []byte
}

func c2go_Config_Descriptor(x *C.struct_libusb_config_descriptor) *Config_Descriptor {
	var list []C.struct_libusb_interface
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&list))
	hdr.Cap = int(x.bNumInterfaces)
	hdr.Len = int(x.bNumInterfaces)
	hdr.Data = uintptr(unsafe.Pointer(x._interface))
	interfaces := make([]*Interface, x.bNumInterfaces)
	for i, _ := range interfaces {
		interfaces[i] = c2go_Interface(&list[i])
	}
	return &Config_Descriptor{
		ptr:                 x,
		BLength:             uint8(x.bLength),
		BDescriptorType:     uint8(x.bDescriptorType),
		WTotalLength:        uint16(x.wTotalLength),
		BNumInterfaces:      uint8(x.bNumInterfaces),
		BConfigurationValue: uint8(x.bConfigurationValue),
		IConfiguration:      uint8(x.iConfiguration),
		BmAttributes:        uint8(x.bmAttributes),
		MaxPower:            uint8(x.MaxPower),
		Interface:           interfaces,
		Extra:               C.GoBytes(unsafe.Pointer(x.extra), x.extra_length),
	}
}

type SS_Endpoint_Companion_Descriptor struct {
	ptr               *C.struct_libusb_ss_endpoint_companion_descriptor
	BLength           uint8
	BDescriptorType   uint8
	BMaxBurst         uint8
	BmAttributes      uint8
	WBytesPerInterval uint16
}

func c2go_SS_Endpoint_Companion_Descriptor(x *C.struct_libusb_ss_endpoint_companion_descriptor) *SS_Endpoint_Companion_Descriptor {
	return &SS_Endpoint_Companion_Descriptor{
		ptr:               x,
		BLength:           uint8(x.bLength),
		BDescriptorType:   uint8(x.bDescriptorType),
		BMaxBurst:         uint8(x.bMaxBurst),
		BmAttributes:      uint8(x.bmAttributes),
		WBytesPerInterval: uint16(x.wBytesPerInterval),
	}
}

type BOS_Dev_Capability_Descriptor struct {
	ptr                 *C.struct_libusb_bos_dev_capability_descriptor
	BLength             uint8
	BDescriptorType     uint8
	BDevCapabilityType  uint8
	Dev_capability_data []byte
}

func c2go_BOS_Dev_Capability_Descriptor(x *C.struct_libusb_bos_dev_capability_descriptor) *BOS_Dev_Capability_Descriptor {
	return &BOS_Dev_Capability_Descriptor{
		ptr:                 x,
		BLength:             uint8(x.bLength),
		BDescriptorType:     uint8(x.bDescriptorType),
		BDevCapabilityType:  uint8(x.bDevCapabilityType),
		Dev_capability_data: C.GoBytes(unsafe.Pointer(C.dev_capability_data_ptr(x)), C.int(x.bLength-3)),
	}
}

type BOS_Descriptor struct {
	ptr             *C.struct_libusb_bos_descriptor
	BLength         uint8
	BDescriptorType uint8
	WTotalLength    uint16
	Dev_capability  []*BOS_Dev_Capability_Descriptor
}

func c2go_BOS_Descriptor(x *C.struct_libusb_bos_descriptor) *BOS_Descriptor {
	var list []*C.struct_libusb_bos_dev_capability_descriptor
	hdr := (*reflect.SliceHeader)(unsafe.Pointer(&list))
	hdr.Cap = int(x.bNumDeviceCaps)
	hdr.Len = int(x.bNumDeviceCaps)
	hdr.Data = uintptr(unsafe.Pointer(C.dev_capability_ptr(x)))
	dev_capability := make([]*BOS_Dev_Capability_Descriptor, x.bNumDeviceCaps)
	for i, _ := range dev_capability {
		dev_capability[i] = c2go_BOS_Dev_Capability_Descriptor(list[i])
	}
	return &BOS_Descriptor{
		ptr:             x,
		BLength:         uint8(x.bLength),
		BDescriptorType: uint8(x.bDescriptorType),
		WTotalLength:    uint16(x.wTotalLength),
		Dev_capability:  dev_capability,
	}
}

type USB_2_0_Extension_Descriptor struct {
	ptr                *C.struct_libusb_usb_2_0_extension_descriptor
	BLength            uint8
	BDescriptorType    uint8
	BDevCapabilityType uint8
	BmAttributes       uint32
}

func c2go_USB_2_0_Extension_Descriptor(x *C.struct_libusb_usb_2_0_extension_descriptor) *USB_2_0_Extension_Descriptor {
	return &USB_2_0_Extension_Descriptor{
		ptr:                x,
		BLength:            uint8(x.bLength),
		BDescriptorType:    uint8(x.bDescriptorType),
		BDevCapabilityType: uint8(x.bDevCapabilityType),
		BmAttributes:       uint32(x.bmAttributes),
	}
}

type SS_USB_Device_Capability_Descriptor struct {
	ptr                   *C.struct_libusb_ss_usb_device_capability_descriptor
	BLength               uint8
	BDescriptorType       uint8
	BDevCapabilityType    uint8
	BmAttributes          uint8
	WSpeedSupported       uint16
	BFunctionalitySupport uint8
	BU1DevExitLat         uint8
	BU2DevExitLat         uint16
}

func c2go_SS_USB_Device_Capability_Descriptor(x *C.struct_libusb_ss_usb_device_capability_descriptor) *SS_USB_Device_Capability_Descriptor {
	return &SS_USB_Device_Capability_Descriptor{
		ptr:                   x,
		BLength:               uint8(x.bLength),
		BDescriptorType:       uint8(x.bDescriptorType),
		BDevCapabilityType:    uint8(x.bDevCapabilityType),
		BmAttributes:          uint8(x.bmAttributes),
		WSpeedSupported:       uint16(x.wSpeedSupported),
		BFunctionalitySupport: uint8(x.bFunctionalitySupport),
		BU1DevExitLat:         uint8(x.bU1DevExitLat),
		BU2DevExitLat:         uint16(x.bU2DevExitLat),
	}
}

type Container_ID_Descriptor struct {
	ptr                *C.struct_libusb_container_id_descriptor
	BLength            uint8
	BDescriptorType    uint8
	BDevCapabilityType uint8
	BReserved          uint8
	ContainerID        []byte
}

func c2go_Container_ID_Descriptor(x *C.struct_libusb_container_id_descriptor) *Container_ID_Descriptor {
	return &Container_ID_Descriptor{
		ptr:                x,
		BLength:            uint8(x.bLength),
		BDescriptorType:    uint8(x.bDescriptorType),
		BDevCapabilityType: uint8(x.bDevCapabilityType),
		BReserved:          uint8(x.bReserved),
		ContainerID:        C.GoBytes(unsafe.Pointer(&x.ContainerID[0]), 16),
	}
}

/*
struct libusb_control_setup {
	uint8_t  bmRequestType;
	uint8_t  bRequest;
	uint16_t wValue;
	uint16_t wIndex;
	uint16_t wLength;
};
*/

type Device_Descriptor struct {
	ptr                *C.struct_libusb_device_descriptor
	BLength            uint8
	BDescriptorType    uint8
	BcdUSB             uint16
	BDeviceClass       uint8
	BDeviceSubClass    uint8
	BDeviceProtocol    uint8
	BMaxPacketSize0    uint8
	IdVendor           uint16
	IdProduct          uint16
	BcdDevice          uint16
	IManufacturer      uint8
	IProduct           uint8
	ISerialNumber      uint8
	BNumConfigurations uint8
}

func c2go_Device_Descriptor(x *C.struct_libusb_device_descriptor) *Device_Descriptor {
	return &Device_Descriptor{
		ptr:                x,
		BLength:            uint8(x.bLength),
		BDescriptorType:    uint8(x.bDescriptorType),
		BcdUSB:             uint16(x.bcdUSB),
		BDeviceClass:       uint8(x.bDeviceClass),
		BDeviceSubClass:    uint8(x.bDeviceSubClass),
		BDeviceProtocol:    uint8(x.bDeviceProtocol),
		BMaxPacketSize0:    uint8(x.bMaxPacketSize0),
		IdVendor:           uint16(x.idVendor),
		IdProduct:          uint16(x.idProduct),
		BcdDevice:          uint16(x.bcdDevice),
		IManufacturer:      uint8(x.iManufacturer),
		IProduct:           uint8(x.iProduct),
		ISerialNumber:      uint8(x.iSerialNumber),
		BNumConfigurations: uint8(x.bNumConfigurations),
	}
}

type Version struct {
	ptr      *C.struct_libusb_version
	Major    uint16
	Minor    uint16
	Micro    uint16
	Nano     uint16
	Rc       string
	Describe string
}

func c2go_Version(x *C.struct_libusb_version) *Version {
	return &Version{
		ptr:      x,
		Major:    uint16(x.major),
		Minor:    uint16(x.minor),
		Micro:    uint16(x.micro),
		Nano:     uint16(x.nano),
		Rc:       C.GoString(x.rc),
		Describe: C.GoString(x.describe),
	}
}

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
	return fmt.Sprintf("libusb_error: %s() returned %s(%d)", e.name, Error_Name(e.code), e.code)
}

func libusb_error(name string, code int) error {
	return &libusb_error_t{
		name: name,
		code: code,
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

func Get_Port_Numbers(dev Device, ports []byte) ([]byte, error) {
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

func Get_Configuration(hdl Device_Handle) (int, error) {
	var config C.int
	rc := int(C.libusb_get_configuration(hdl, &config))
	if rc < 0 {
		return 0, libusb_error("libusb_get_configuration", rc)
	}
	return int(config), nil
}

func Set_Configuration(hdl Device_Handle, configuration int) error {
	rc := int(C.libusb_set_configuration(hdl, (C.int)(configuration)))
	if rc < 0 {
		return libusb_error("libusb_set_configuration", rc)
	}
	return nil
}

func Claim_Interface(hdl Device_Handle, interface_number int) error {
	rc := int(C.libusb_claim_interface(hdl, (C.int)(interface_number)))
	if rc < 0 {
		return libusb_error("libusb_claim_interface", rc)
	}
	return nil
}

func Release_Interface(hdl Device_Handle, interface_number int) error {
	rc := int(C.libusb_release_interface(hdl, (C.int)(interface_number)))
	if rc < 0 {
		return libusb_error("libusb_release_interface", rc)
	}
	return nil
}

func Set_Interface_Alt_Setting(hdl Device_Handle, interface_number int, alternate_setting int) error {
	rc := int(C.libusb_set_interface_alt_setting(hdl, (C.int)(interface_number), (C.int)(alternate_setting)))
	if rc < 0 {
		return libusb_error("libusb_set_interface_alt_setting", rc)
	}
	return nil
}

func Clear_Halt(hdl Device_Handle, endpoint uint8) error {
	rc := int(C.libusb_clear_halt(hdl, (C.uchar)(endpoint)))
	if rc < 0 {
		return libusb_error("libusb_clear_halt", rc)
	}
	return nil
}

func Reset_Device(hdl Device_Handle) error {
	rc := int(C.libusb_reset_device(hdl))
	if rc < 0 {
		return libusb_error("libusb_reset_device", rc)
	}
	return nil
}

func Kernel_Driver_Active(hdl Device_Handle, interface_number int) (bool, error) {
	rc := int(C.libusb_kernel_driver_active(hdl, (C.int)(interface_number)))
	if rc < 0 {
		return false, libusb_error("libusb_kernel_driver_active", rc)
	}
	return rc != 0, nil
}

func Detach_Kernel_Driver(hdl Device_Handle, interface_number int) error {
	rc := int(C.libusb_detach_kernel_driver(hdl, (C.int)(interface_number)))
	if rc < 0 {
		return libusb_error("libusb_detach_kernel_driver", rc)
	}
	return nil
}

func Attach_Kernel_Driver(hdl Device_Handle, interface_number int) error {
	rc := int(C.libusb_attach_kernel_driver(hdl, (C.int)(interface_number)))
	if rc < 0 {
		return libusb_error("libusb_attach_kernel_driver", rc)
	}
	return nil
}

func Set_Auto_Detach_Kernel_Driver(hdl Device_Handle, enable bool) error {
	enable_int := 0
	if enable {
		enable_int = 1
	}
	rc := int(C.libusb_set_auto_detach_kernel_driver(hdl, (C.int)(enable_int)))
	if rc < 0 {
		return libusb_error("libusb_set_auto_detach_kernel_driver", rc)
	}
	return nil
}

//-----------------------------------------------------------------------------
// Miscellaneous

func Has_Capability(capability uint32) bool {
	rc := int(C.libusb_has_capability((C.uint32_t)(capability)))
	return rc != 0
}

func Error_Name(code int) string {
	return C.GoString(C.libusb_error_name(C.int(code)))
}

func Get_Version() *Version {
	ver := (*C.struct_libusb_version)(unsafe.Pointer(C.libusb_get_version()))
	return c2go_Version(ver)
}

func CPU_To_LE16(x uint16) uint16 {
	return uint16(C.libusb_cpu_to_le16((C.uint16_t)(x)))
}

func Setlocale(locale string) error {
	cstr := C.CString(locale)
	rc := int(C.libusb_setlocale(cstr))
	if rc < 0 {
		return libusb_error("libusb_setlocale", rc)
	}
	return nil
}

func Strerror(errcode int) string {
	return C.GoString(C.libusb_strerror(int32(errcode)))
}

//-----------------------------------------------------------------------------
// USB descriptors

func Get_Device_Descriptor(dev Device) (*Device_Descriptor, error) {
	var desc C.struct_libusb_device_descriptor
	rc := int(C.libusb_get_device_descriptor(dev, &desc))
	if rc != 0 {
		return nil, libusb_error("libusb_get_device_descriptor", rc)
	}
	return c2go_Device_Descriptor(&desc), nil
}

func Get_Active_Config_Descriptor(dev Device) (*Config_Descriptor, error) {
	var desc *C.struct_libusb_config_descriptor
	rc := int(C.libusb_get_active_config_descriptor(dev, &desc))
	if rc != 0 {
		return nil, libusb_error("libusb_get_active_config_descriptor", rc)
	}
	return c2go_Config_Descriptor(desc), nil
}

func Get_Config_Descriptor(dev Device, config_index uint8) (*Config_Descriptor, error) {
	var desc *C.struct_libusb_config_descriptor
	rc := int(C.libusb_get_config_descriptor(dev, (C.uint8_t)(config_index), &desc))
	if rc != 0 {
		return nil, libusb_error("libusb_get_config_descriptor", rc)
	}
	return c2go_Config_Descriptor(desc), nil
}

func Get_Config_Descriptor_By_Value(dev Device, bConfigurationValue uint8) (*Config_Descriptor, error) {
	var desc *C.struct_libusb_config_descriptor
	rc := int(C.libusb_get_config_descriptor_by_value(dev, (C.uint8_t)(bConfigurationValue), &desc))
	if rc != 0 {
		return nil, libusb_error("libusb_get_config_descriptor_by_value", rc)
	}
	return c2go_Config_Descriptor(desc), nil
}

func Free_Config_Descriptor(config *Config_Descriptor) {
	C.libusb_free_config_descriptor(config.ptr)
}

func Get_SS_Endpoint_Companion_Descriptor(ctx Context, endpoint *Endpoint_Descriptor) (*SS_Endpoint_Companion_Descriptor, error) {
	var desc *C.struct_libusb_ss_endpoint_companion_descriptor
	rc := int(C.libusb_get_ss_endpoint_companion_descriptor(ctx, endpoint.ptr, &desc))
	if rc != 0 {
		return nil, libusb_error("libusb_get_ss_endpoint_companion_descriptor", rc)
	}
	return c2go_SS_Endpoint_Companion_Descriptor(desc), nil
}

func Free_SS_Endpoint_Companion_Descriptor(ep_comp *SS_Endpoint_Companion_Descriptor) {
	C.libusb_free_ss_endpoint_companion_descriptor(ep_comp.ptr)
}

func Get_BOS_Descriptor(hdl Device_Handle) (*BOS_Descriptor, error) {
	var desc *C.struct_libusb_bos_descriptor
	rc := int(C.libusb_get_bos_descriptor(hdl, &desc))
	if rc != 0 {
		return nil, libusb_error("libusb_get_bos_descriptor", rc)
	}
	return c2go_BOS_Descriptor(desc), nil
}

func Free_BOS_Descriptor(bos *BOS_Descriptor) {
	C.libusb_free_bos_descriptor(bos.ptr)
}

func Get_USB_2_0_Extension_Descriptor(ctx Context, dev_cap *BOS_Dev_Capability_Descriptor) (*USB_2_0_Extension_Descriptor, error) {
	var desc *C.struct_libusb_usb_2_0_extension_descriptor
	rc := int(C.libusb_get_usb_2_0_extension_descriptor(ctx, dev_cap.ptr, &desc))
	if rc != 0 {
		return nil, libusb_error("libusb_get_usb_2_0_extension_descriptor", rc)
	}
	return c2go_USB_2_0_Extension_Descriptor(desc), nil
}

func Free_USB_2_0_Extension_Descriptor(usb_2_0_extension *USB_2_0_Extension_Descriptor) {
	C.libusb_free_usb_2_0_extension_descriptor(usb_2_0_extension.ptr)
}

func Get_SS_USB_Device_Capability_Descriptor(ctx Context, dev_cap *BOS_Dev_Capability_Descriptor) (*SS_USB_Device_Capability_Descriptor, error) {
	var desc *C.struct_libusb_ss_usb_device_capability_descriptor
	rc := int(C.libusb_get_ss_usb_device_capability_descriptor(ctx, dev_cap.ptr, &desc))
	if rc != 0 {
		return nil, libusb_error("libusb_get_ss_usb_device_capability_descriptor", rc)
	}
	return c2go_SS_USB_Device_Capability_Descriptor(desc), nil
}

func Free_SS_USB_Device_Capability_Descriptor(ss_usb_device_cap *SS_USB_Device_Capability_Descriptor) {
	C.libusb_free_ss_usb_device_capability_descriptor(ss_usb_device_cap.ptr)
}

func Get_Container_ID_Descriptor(ctx Context, dev_cap *BOS_Dev_Capability_Descriptor) (*Container_ID_Descriptor, error) {
	var desc *C.struct_libusb_container_id_descriptor
	rc := int(C.libusb_get_container_id_descriptor(ctx, dev_cap.ptr, &desc))
	if rc != 0 {
		return nil, libusb_error("libusb_get_container_id_descriptor", rc)
	}
	return c2go_Container_ID_Descriptor(desc), nil
}

func Free_Container_ID_Descriptor(container_id *Container_ID_Descriptor) {
	C.libusb_free_container_id_descriptor(container_id.ptr)
}

func Get_String_Descriptor_ASCII(dev Device_Handle, desc_index uint8, data []byte) ([]byte, error) {
	rc := int(C.libusb_get_string_descriptor_ascii(dev, (C.uint8_t)(desc_index), (*C.uchar)(&data[0]), (C.int)(len(data))))
	if rc < 0 {
		return nil, libusb_error("libusb_get_string_descriptor_ascii", rc)
	}
	return data[:rc], nil
}

func Get_Descriptor(dev Device_Handle, desc_type uint8, desc_index uint8, data []byte) ([]byte, error) {
	rc := int(C.libusb_get_descriptor(dev, (C.uint8_t)(desc_type), (C.uint8_t)(desc_index), (*C.uchar)(&data[0]), (C.int)(len(data))))
	if rc < 0 {
		return nil, libusb_error("libusb_get_descriptor", rc)
	}
	return data[:rc], nil
}

func Get_String_Descriptor(dev Device_Handle, desc_index uint8, langid uint16, data []byte) ([]byte, error) {
	rc := int(C.libusb_get_string_descriptor(dev, (C.uint8_t)(desc_index), (C.uint16_t)(langid), (*C.uchar)(&data[0]), (C.int)(len(data))))
	if rc < 0 {
		return nil, libusb_error("libusb_get_string_descriptor", rc)
	}
	return data[:rc], nil
}

//-----------------------------------------------------------------------------
// Device hotplug event notification

//int 	libusb_hotplug_register_callback (libusb_context *ctx, libusb_hotplug_event events, libusb_hotplug_flag flags, int vendor_id, int product_id, int dev_class, libusb_hotplug_callback_fn cb_fn, void *user_data, libusb_hotplug_callback_handle *handle)
//void 	libusb_hotplug_deregister_callback (libusb_context *ctx, libusb_hotplug_callback_handle handle)

//-----------------------------------------------------------------------------
//Asynchronous device I/O

// int 	libusb_alloc_streams (libusb_device_handle *dev, uint32_t num_streams, unsigned char *endpoints, int num_endpoints)
// int 	libusb_free_streams (libusb_device_handle *dev, unsigned char *endpoints, int num_endpoints)
// struct libusb_transfer * 	libusb_alloc_transfer (int iso_packets)
// void 	libusb_free_transfer (struct libusb_transfer *transfer)
// int 	libusb_submit_transfer (struct libusb_transfer *transfer)
// int 	libusb_cancel_transfer (struct libusb_transfer *transfer)
// void 	libusb_transfer_set_stream_id (struct libusb_transfer *transfer, uint32_t stream_id)
// uint32_t 	libusb_transfer_get_stream_id (struct libusb_transfer *transfer)
// static unsigned char * 	libusb_control_transfer_get_data (struct libusb_transfer *transfer)
// static struct libusb_control_setup * 	libusb_control_transfer_get_setup (struct libusb_transfer *transfer)
// static void 	libusb_fill_control_setup (unsigned char *buffer, uint8_t bmRequestType, uint8_t bRequest, uint16_t wValue, uint16_t wIndex, uint16_t wLength)
// static void 	libusb_fill_control_transfer (struct libusb_transfer *transfer, libusb_device_handle *dev_handle, unsigned char *buffer, libusb_transfer_cb_fn callback, void *user_data, unsigned int timeout)
// static void 	libusb_fill_bulk_transfer (struct libusb_transfer *transfer, libusb_device_handle *dev_handle, unsigned char endpoint, unsigned char *buffer, int length, libusb_transfer_cb_fn callback, void *user_data, unsigned int timeout)
// static void 	libusb_fill_bulk_stream_transfer (struct libusb_transfer *transfer, libusb_device_handle *dev_handle, unsigned char endpoint, uint32_t stream_id, unsigned char *buffer, int length, libusb_transfer_cb_fn callback, void *user_data, unsigned int timeout)
// static void 	libusb_fill_interrupt_transfer (struct libusb_transfer *transfer, libusb_device_handle *dev_handle, unsigned char endpoint, unsigned char *buffer, int length, libusb_transfer_cb_fn callback, void *user_data, unsigned int timeout)
// static void 	libusb_fill_iso_transfer (struct libusb_transfer *transfer, libusb_device_handle *dev_handle, unsigned char endpoint, unsigned char *buffer, int length, int num_iso_packets, libusb_transfer_cb_fn callback, void *user_data, unsigned int timeout)
// static void 	libusb_set_iso_packet_lengths (struct libusb_transfer *transfer, unsigned int length)
// static unsigned char * 	libusb_get_iso_packet_buffer (struct libusb_transfer *transfer, unsigned int packet)
// static unsigned char * 	libusb_get_iso_packet_buffer_simple (struct libusb_transfer *transfer, unsigned int packet)

//-----------------------------------------------------------------------------
// Polling and timing

// int 	libusb_try_lock_events (libusb_context *ctx)
// void 	libusb_lock_events (libusb_context *ctx)
// void 	libusb_unlock_events (libusb_context *ctx)
// int 	libusb_event_handling_ok (libusb_context *ctx)
// int 	libusb_event_handler_active (libusb_context *ctx)
// void 	libusb_lock_event_waiters (libusb_context *ctx)
// void 	libusb_unlock_event_waiters (libusb_context *ctx)
// int 	libusb_wait_for_event (libusb_context *ctx, struct timeval *tv)
// int 	libusb_handle_events_timeout_completed (libusb_context *ctx, struct timeval *tv, int *completed)
// int 	libusb_handle_events_timeout (libusb_context *ctx, struct timeval *tv)
// int 	libusb_handle_events (libusb_context *ctx)
// int 	libusb_handle_events_completed (libusb_context *ctx, int *completed)
// int 	libusb_handle_events_locked (libusb_context *ctx, struct timeval *tv)
// int 	libusb_pollfds_handle_timeouts (libusb_context *ctx)
// int 	libusb_get_next_timeout (libusb_context *ctx, struct timeval *tv)
// void 	libusb_set_pollfd_notifiers (libusb_context *ctx, libusb_pollfd_added_cb added_cb, libusb_pollfd_removed_cb removed_cb, void *user_data)
// const struct libusb_pollfd ** 	libusb_get_pollfds (libusb_context *ctx)
// void 	libusb_free_pollfds (const struct libusb_pollfd **pollfds)

//-----------------------------------------------------------------------------
// Synchronous device I/O

func Control_Transfer(hdl Device_Handle, bmRequestType uint8, bRequest uint8, wValue uint16, wIndex uint16, data []byte, timeout uint) ([]byte, error) {
	rc := int(C.libusb_control_transfer(hdl, (C.uint8_t)(bmRequestType), (C.uint8_t)(bRequest), (C.uint16_t)(wValue), (C.uint16_t)(wIndex),
		(*C.uchar)(&data[0]), (C.uint16_t)(len(data)), (C.uint)(timeout)))
	if rc < 0 {
		return nil, libusb_error("libusb_control_transfer", rc)
	}
	return data[:rc], nil
}

func Bulk_Transfer(hdl Device_Handle, endpoint uint8, data []byte, timeout uint) ([]byte, error) {
	var transferred C.int
	rc := int(C.libusb_bulk_transfer(hdl, (C.uchar)(endpoint), (*C.uchar)(&data[0]), (C.int)(len(data)), &transferred, (C.uint)(timeout)))
	if rc != LIBUSB_SUCCESS {
		return nil, libusb_error("libusb_bulk_transfer", rc)
	}
	return data[:int(transferred)], nil
}

func Interrupt_Transfer(hdl Device_Handle, endpoint uint8, data []byte, timeout uint) ([]byte, error) {
	var transferred C.int
	rc := int(C.libusb_interrupt_transfer(hdl, (C.uchar)(endpoint), (*C.uchar)(&data[0]), (C.int)(len(data)), &transferred, (C.uint)(timeout)))
	if rc != LIBUSB_SUCCESS {
		return nil, libusb_error("libusb_interrupt_transfer", rc)
	}
	return data[:int(transferred)], nil
}

//-----------------------------------------------------------------------------
