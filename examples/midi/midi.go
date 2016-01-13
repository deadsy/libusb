package main

import (
	"fmt"
	"github.com/deadsy/libusb"
	"os"
)

func display_String_Descriptor(hdl libusb.Device_Handle, name string, idx uint8) {
	if idx != 0 {
		str := make([]byte, 128)
		str, err := libusb.Get_String_Descriptor_ASCII(hdl, idx, str)
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("%s(%d) %s\n", name, idx, str)
	}
}

func midi_device(ctx libusb.Context, vid uint16, pid uint16) {

	fmt.Printf("Opening device %04X:%04X ", vid, pid)
	hdl := libusb.Open_Device_With_VID_PID(ctx, vid, pid)
	if hdl == nil {
		fmt.Printf("failed (do you have permission?)\n")
		return
	}
	fmt.Printf("ok\n")
	defer libusb.Close(hdl)

	dev := libusb.Get_Device(hdl)

	fmt.Printf("Device Descriptor:\n")
	dd, err := libusb.Get_Device_Descriptor(dev)
	if err != nil {
		fmt.Printf("%s\n", err)
		return
	}

	fmt.Printf("%s\n", libusb.Device_Descriptor_str(dd))

	fmt.Printf("String Descriptors:\n")
	display_String_Descriptor(hdl, "  iManufacturer", dd.IManufacturer)
	display_String_Descriptor(hdl, "  iProduct", dd.IProduct)
	display_String_Descriptor(hdl, "  iSerialNumber", dd.ISerialNumber)

	for i := 0; i < int(dd.BNumConfigurations); i++ {
		fmt.Printf("Configuration Descriptor %d:\n", i)
		cd, err := libusb.Get_Config_Descriptor(dev, uint8(i))
		if err != nil {
			fmt.Printf("%s\n", err)
			return
		}
		fmt.Printf("%s\n", libusb.Config_Descriptor_str(cd))
		libusb.Free_Config_Descriptor(cd)
	}
}

func main() {

	var ctx libusb.Context

	err := libusb.Init(&ctx)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(-1)
	}
	defer libusb.Exit(ctx)

	midi_device(ctx, 0x944, 0x115)

	os.Exit(0)
}
