package printbox

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

// a terrible bare minimum object representation of the compose objet
// probably should have riped this from the actual project
type Composer struct {
	Version  string             `yaml:"version,omitempty"`
	Services map[string]Service `yaml:"services,omitempty"`
	Volumes  map[string]*Volume `yaml:"volumes,omitempty"`
	Networks map[string]Network `yaml:"networks,omitempty"`
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
	Networks    []string `yaml:"networks,omitempty"`
	Hostname    string   `yaml:"hostname,omitempty"`
}

type Network struct {
	External   bool   `yaml:"external,omitempty"`
	Name       string `yaml:"name,omitempty"`
	Driver     string `yaml:"driver,omitempty"`
	Attachable bool   `yaml:"attachable,omitempty"`
}

type Volume struct {
	External bool `json:"external,omitempty"`
}

func BuildComposeFile(ports map[string]string) ([]byte, error) {
	cs := &Composer{
		Version:  "3.9",
		Services: map[string]Service{},
		Volumes:  map[string]*Volume{},
		Networks: map[string]Network{},
	}

	// create the network
	cs.Networks["printbox"] = Network{}

	// config volume
	cs.Volumes["printbox"] = &Volume{
		External: true,
	}
	volS := fmt.Sprintf("printbox:%s", SharedPath)

	// create fluid
	cs.Services["fluidd"] = Service{
		Image:   "cadriel/fluidd",
		Ports:   []string{"80:80"},
		Restart: "unless-stopped",
		Volumes: []string{
			volS,
		},
		Networks: []string{"printbox"},
		Hostname: "fluidd",
	}

	// create the config-editor
	// cs.Services["config-editor"] = Service{
	// 	Image:   "linuxserver/code-server",
	// 	Ports:   []string{"8443:8443"},
	// 	Restart: "unless-stopped",
	// 	Environment: []string{
	// 		"PUID=0",
	// 		"GUID=0",
	// 		"TZ=America/Los_Angeles",
	// 	},
	// 	Volumes: []string{
	// 		volS,
	// 	},
	// 	Networks: []string{"printbox"},
	// }

	// create moonraker/klippers
	i := 0
	for alias, port := range ports {
		i++

		// creater the names
		portS := fmt.Sprintf("808%d:7125", i)
		nameS := fmt.Sprintf("printer_%s", alias)
		deviceS := fmt.Sprintf("%s:/dev/klipperserial", port)
		envDir := fmt.Sprintf("PRINTBOX_DIR=%s/%s", SharedPath, alias)
		envID := fmt.Sprintf("PRINTBOX_ID=%d", i)

		fmt.Printf("klipraker %s\n", nameS)
		fmt.Printf(" %-20s %s\n", "moonraker port", portS)
		fmt.Printf(" %-20s %s\n", "serial", port)
		fmt.Printf(" %-20s %s\n", "shared dir", envDir)
		fmt.Println()

		// create the klipraker service
		cs.Services[nameS] = Service{
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
			Networks: []string{"printbox"},
			Hostname: nameS,
		}
	}

	data, err := yaml.Marshal(cs)
	if err != nil {
		return []byte{}, err
	}
	return data, nil
}
