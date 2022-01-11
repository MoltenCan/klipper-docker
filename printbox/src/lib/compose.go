package lib

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v3"
	"path/filepath"
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
		Version:  "2.4",
		Services: map[string]Service{},
		Volumes:  map[string]*Volume{},
	}

	// config volume
	cs.Volumes["shared"] = nil

	// create fluid
	svc := Service{
		Image:   "cadriel/fluidd",
		Ports:   []string{"80:80"},
		Restart: "unless-stopped",
		Volumes: []string{
			"shared:/shared",
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
		envS := fmt.Sprintf("KLIPPER_ID=%d", li)

		// create the service
		svc := Service{
			Build:   ".",
			Ports:   []string{portS},
			Restart: "unless-stopped",
			Volumes: []string{
				"shared:/shared",
			},
			Devices:     []string{deviceS},
			Environment: []string{envS},
		}
		cs.Services[nameS] = svc
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
			"shared:/shared",
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
