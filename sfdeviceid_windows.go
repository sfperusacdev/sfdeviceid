//go:build windows
// +build windows

package sfdeviceid

import (
	"log/slog"

	"golang.org/x/sys/windows/registry"

	"github.com/google/uuid"
)

const regPath = `SOFTWARE\xhrydhslwlls`
const regName = "dhid"

func getRegistry() (*registry.Key, error) {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, regPath, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		if err == registry.ErrNotExist {
			key, _, err = registry.CreateKey(registry.LOCAL_MACHINE, regPath, registry.SET_VALUE)
			if err != nil {
				slog.Error("CreateKey registry", "error", err)
				return nil, err
			}
		} else {
			slog.Error("opening registry", "error", err)
			return nil, err
		}
	}
	return &key, nil
}
func GenDeviceID() (string, error) {
	key, err := getRegistry()
	if err != nil {
		slog.Error("LOCAL_MACHINE registry fail", "error", err)
		return "", err
	}
	defer key.Close()

	// Intenta leer el valor DeviceID
	deviceID, _, err := key.GetStringValue(regName)
	if err == registry.ErrNotExist {
		deviceID = uuid.New().String()
		err = key.SetStringValue(regName, deviceID)
		if err != nil {
			return "", err
		}
		slog.Info("device id was generated", "ID", deviceID)
	} else if err != nil {
		slog.Error("device id set registry value", "error", err)
		return "", err
	}
	return deviceID, nil
}

func DeviceID() string {
	key, err := getRegistry()
	if err != nil {
		slog.Error("LOCAL_MACHINE registry fail", "error", err)
		return ""
	}
	defer key.Close()
	deviceID, _, err := key.GetStringValue(regName)
	if err != registry.ErrNotExist {
		slog.Error("reading value from registry device_id", "error", err)
		return ""
	}
	return deviceID
}
