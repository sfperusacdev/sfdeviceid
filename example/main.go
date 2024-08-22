package main

import (
	"fmt"

	"github.com/sfperusacdev/sfdeviceid"
)

func main() {
	fmt.Println(sfdeviceid.GenDeviceID())
	fmt.Println(sfdeviceid.DeviceID())
}
