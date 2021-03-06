package oz

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	ProfileDir      string   `json:"profile_dir" desc:"Directory containing the sandbox profiles"`
	ShellPath       string   `json:"shell_path" desc:"Path of the shell used when entering a sandbox"`
	PrefixPath      string   `json:"prefix_path" desc:"Prefix path containing the oz executables"`
	EtcPrefix       string   `json:"etc_prefix" desc:"Prefix for configuration files"`
	SandboxPath     string   `json:"sandbox_path" desc:"Path of the sandboxes base"`
	BridgeMACAddr   string   `json:"bridge_mac" desc:"MAC Address of the bridge interface"`
	DivertSuffix    string   `json:"divert_suffix" desc:"Suffix using for dpkg-divert of application executables"`
	NMIgnoreFile    string   `json:"nm_ignore_file" desc:"Path to the NetworkManager ignore config file, disables the warning if empty"`
	UseFullDev      bool     `json:"use_full_dev" desc:"Give sandboxes full access to devices instead of a restricted set"`
	AllowRootShell  bool     `json:"allow_root_shell" desc:"Allow entering a sandbox shell as root"`
	LogXpra         bool     `json:"log_xpra" desc:"Log output of Xpra"`
	EnvironmentVars []string `json:"environment_vars" desc:"Default environment variables passed to sandboxes"`
	DefaultGroups  []string  `json:"default_groups" desc:"List of default group names that can be used inside the sandbox"`
}

const OzVersion = "0.0.1"
const DefaultConfigPath = "/etc/oz/oz.conf"

func NewDefaultConfig() *Config {
	return &Config{
		ProfileDir:      "/var/lib/oz/cells.d",
		ShellPath:       "/bin/bash",
		PrefixPath:      "/usr/local",
		EtcPrefix:       "/etc/oz",
		SandboxPath:     "/srv/oz",
		NMIgnoreFile:    "/etc/NetworkManager/conf.d/oz.conf",
		BridgeMACAddr:   "6A:A8:2E:56:E8:9C",
		DivertSuffix:    "unsafe",
		UseFullDev:      false,
		AllowRootShell:  false,
		LogXpra:         false,
		EnvironmentVars: []string{
			"USER", "USERNAME", "LOGNAME",
			"LANG", "LANGUAGE", "_", "TZ=UTC",
		},
		DefaultGroups:   []string{
			"audio", "video",
		},
	}
}

func LoadConfig(cpath string) (*Config, error) {
	if _, err := os.Stat(cpath); os.IsNotExist(err) {
		return nil, err
	}
	if err := checkConfigPermissions(cpath); err != nil {
		return nil, err
	}

	bs, err := ioutil.ReadFile(cpath)
	if err != nil {
		return nil, err
	}
	c := NewDefaultConfig()
	if err := json.Unmarshal(bs, c); err != nil {
		return nil, err
	}
	return c, nil
}
