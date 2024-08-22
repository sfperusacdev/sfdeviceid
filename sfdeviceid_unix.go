//go:build linux || darwin
// +build linux darwin

package sfdeviceid

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const (
	deviceIDFileName = ".device_id"
)

func GetDeviceID() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	deviceIDPath := filepath.Join(homeDir, deviceIDFileName)
	deviceID, err := os.ReadFile(deviceIDPath)
	if err != nil {
		if os.IsNotExist(err) {
			newDeviceID := uuid.New().String()
			err = os.WriteFile(deviceIDPath, []byte(newDeviceID), 0600)
			if err != nil {
				return "", err
			}
			fmt.Println("Device ID was generated:", newDeviceID)
			return newDeviceID, nil
		}
		return "", err
	}
	return string(deviceID), nil
}
