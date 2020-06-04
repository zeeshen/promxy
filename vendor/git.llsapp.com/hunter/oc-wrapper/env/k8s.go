package env

import (
	"github.com/golang/glog"
	"io/ioutil"
	"os"
	"strings"
)

const DefaultConfigPath = "/etc/podinfo/labels"

var DefaultExportEnvKeyMap = map[string]string{"cluster": SERVICE_NAME}

func init() {
	// if we were in k8s, export pods info to ENV
	if _, err := os.Stat(DefaultConfigPath); !os.IsNotExist(err) {
		exportEnvsFromFile(DefaultConfigPath, DefaultExportEnvKeyMap)
	}
}

func exportEnvsFromFile(filePath string, keyMap map[string]string) {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		glog.Errorf("read podinfo labels file(path:%s) failed, err: %s", filePath, err)
		return
	}

	exportEnvsFromKVStrs(string(data), keyMap)
}

func exportEnvsFromKVStrs(s string, keyMap map[string]string) {
	lines := strings.Split(s, "\n")

	for _, line := range lines {
		kv := strings.Split(line, "=")
		if len(kv) == 2 {
			k := strings.TrimSpace(kv[0])
			v := strings.Trim(strings.TrimSpace(kv[1]), "\"")
			if envName := keyMap[k]; envName != "" {
				os.Setenv(envName, v)
				glog.Infof("export env for key[%s] to %s=%s\n", k, envName, v)
			}
		}
	}
}
