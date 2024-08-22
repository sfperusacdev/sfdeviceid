package main

import (
	"fmt"
	"sfdeviceid"
)

func main() {
	fmt.Println(sfdeviceid.GenDeviceID())
	fmt.Println(sfdeviceid.DeviceID())
}
