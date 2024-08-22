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

func GetDeviceID() (string, error) {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, regPath, registry.QUERY_VALUE|registry.SET_VALUE)
	if err != nil {
		if err == registry.ErrNotExist {
			key, _, err = registry.CreateKey(registry.LOCAL_MACHINE, regPath, registry.SET_VALUE)
			if err != nil {
				return "", err
			}
		} else {
			return "", err
		}
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
		return "", err
	}
	return deviceID, nil
}
