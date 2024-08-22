//go:build linux || darwin
// +build linux darwin

package sfdeviceid

import (
	"fmt"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

const (
	deviceIDFileName = ".device_id"
)

func filePath() (string, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	deviceIDPath := filepath.Join(homeDir, deviceIDFileName)
	return deviceIDPath, err
}

func GenDeviceID() (string, error) {
	deviceIDPath, err := filePath()
	if err != nil {
		return "", err
	}
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

func DeviceID() string {
	deviceIDPath, err := filePath()
	if err != nil {
		slog.Error("file path .device_id", "error", err)
		return ""
	}
	deviceID, err := os.ReadFile(deviceIDPath)
	if err != nil {
		if !os.IsNotExist(err) {
			slog.Error("reading file .device_id", "error", err)
		}
	}
	return string(deviceID)
}
