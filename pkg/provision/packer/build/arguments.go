package build

import (
	"fmt"
	"github.com/funtimecoding/soil/pkg/system/constant"
)

func (b *Build) arguments(headless bool) [][]string {
	var processor string

	switch b.architecture {
	case constant.AMD64:
		processor = "host"
	case constant.ARM64:
		processor = "cortex-a72"
	}

	result := [][]string{
		{"-cpu", processor},
		{"-device", "qemu-xhci"},
		{"-device", "usb-kbd"},
		{"-device", "virtio-rng-pci"},
	}

	switch b.architecture {
	case constant.ARM64:
		result = append(result, []string{"-machine", "virt"})
		result = append(result, []string{"-boot", "strict=off"})
		result = append(result, []string{"-cpu", "cortex-a72"})
		result = append(result, []string{"-bios", "edk2-aarch64-code.fd"})
	}

	if headless {
		result = append(
			result,
			[]string{"-monitor", "telnet::6666,server=on,wait=off"},
		)
		result = append(
			result,
			[]string{
				"-chardev",
				fmt.Sprintf(
					"socket,id=char0,host=localhost,port=%d,server=on,wait=off",
					b.CharacterDevicePort,
				),
			},
		)
		result = append(result, []string{"-serial", "chardev:char0"})
		result = append(result, []string{"-nographic"})
	} else {
		result = append(
			result,
			[]string{"-device", "virtio-gpu-pci"},
		)
	}

	return result
}
