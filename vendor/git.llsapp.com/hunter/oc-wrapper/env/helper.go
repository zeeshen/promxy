package env

import (
	"os"
)

func GetHostIp() string {
	// the vm instance ip
	return os.Getenv(HOST_IP)
}

func GetServiceName() string {
	return os.Getenv(SERVICE_NAME)
}

func GetHostName() string {
	return os.Getenv(HOST_NAME)
}
