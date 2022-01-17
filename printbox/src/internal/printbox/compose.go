package printbox

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// a terrible bare minimum object representation of the compose objet
// probably should have riped this from the actual project
type Composer struct {
	Version  string             `yaml:"version,omitempty"`
	Services map[string]Service `yaml:"services,omitempty"`
	Volumes  map[string]*Volume `yaml:"volumes,omitempty"`
}

type Service struct {
	Image       string   `yaml:"image,omitempty"`
	Ports       []string `yaml:"ports,omitempty"`
	DependsOn   []string `yaml:"depends_on,omitempty"`
	Restart     string   `yaml:"restart,omitempty"`
	Environment []string `yaml:"environment,omitempty"`
	Volumes     []string `yaml:"volumes,omitempty"`
	Build       string   `yaml:"build,omitempty"`
	Devices     []string `yaml:"devices,omitempty"`
}

type Volume struct {
	External bool `json:"external,omitempty"`
}

func BuildComposeFile(bi *BoardInfo) error {
	cs := &Composer{
		Version:  "3.9",
		Services: map[string]Service{},
		Volumes:  map[string]*Volume{},
	}

	// config volume
	cs.Volumes["printbox"] = &Volume{
		External: true,
	}
	volS := fmt.Sprintf("printbox:%s", SharedPath)

	// create fluid
	svc := Service{
		Image:   "cadriel/fluidd",
		Ports:   []string{"80:80"},
		Restart: "unless-stopped",
		Volumes: []string{
			volS,
		},
	}
	cs.Services["fluidd"] = svc

	// create moonraker/klippers
	for i, port := range bi.USB {
		li := i + 1
		if !port.Connected {
			continue
		}
		// creater the names
		portS := fmt.Sprintf("808%d:7125", li)
		nameS := fmt.Sprintf("printer_%d", li)
		deviceS := fmt.Sprintf("%s:/dev/klipperserial", port.Device)
		envDir := fmt.Sprintf("PRINTBOX_DIR=%s/%d", SharedPath, li)
		envID := fmt.Sprintf("PRINTBOX_ID=%d", li)

		// create the klipper service
		svc := Service{
			Image:   "moltencan/klipraker",
			Restart: "unless-stopped",
			Ports:   []string{portS},
			Volumes: []string{
				volS,
			},
			Devices: []string{deviceS},
			Environment: []string{
				envDir,
				envID,
			},
		}
		cs.Services[nameS+"klipraker"] = svc

		// // create the moonraker service
		// svc = Service{
		// 	Image:   "moltencan/moonraker",
		// 	Ports:   []string{portS},
		// 	Restart: "unless-stopped",
		// 	Volumes: []string{
		// 		volS,
		// 	},
		// 	Environment: []string{
		// 		envDir,
		// 		envID,
		// 	},
		// }
		// cs.Services[nameS+"_moonraker"] = svc
	}

	// create the config-editor
	svc = Service{
		Image:   "linuxserver/code-server",
		Ports:   []string{"8443:8443"},
		Restart: "unless-stopped",
		Environment: []string{
			"PUID=0",
			"GUID=0",
			"TZ=America/Los_Angeles",
		},
		Volumes: []string{
			volS,
		},
	}
	cs.Services["config-editor"] = svc

	data, err := yaml.Marshal(cs)
	if err != nil {
		return err
	}
	dcFile := filepath.Join(SharedPath, "docker-compose.yml")
	Logf("writing compose to %s", dcFile)
	return ioutil.WriteFile(dcFile, data, 0644)
}
