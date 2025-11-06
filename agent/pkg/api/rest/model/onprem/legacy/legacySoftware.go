package legacy

type LegacySoftware struct {
	Name            string   `json:"name"`
	Version         string   `json:"version"`
	NeededLibraries []string `json:"needed_libraries"`
	BinaryPath      string   `json:"binary_path"`
	//EnvFiles        []string `json:"env_files"`
	SystemdFiles    []string `json:"systemd_files"`
	CustomDataPaths []string `json:"custom_data_paths"`
	CustomConfigs   []string `json:"custom_configs"`
}
