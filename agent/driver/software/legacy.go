package software

import (
	"bufio"
	"bytes"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	model "github.com/cloud-barista/cm-honeybee/agent/pkg/api/rest/model/onprem/legacy"
)

// GetLegacySoftwareInfo collects information about running legacy software
// by scanning the /proc directory on the current VM.
func GetLegacySoftwareInfo() ([]model.LegacySoftware, error) {
	var result []model.LegacySoftware

	// 1. read proc
	procDirs, err := os.ReadDir("/proc")
	if err != nil {
		return nil, err
	}

	// 2. get proc
	for _, d := range procDirs {
		// 3. Skip if it's not a directory or not PID
		if !d.IsDir() || !isDigitDir(d.Name()) {
			continue
		}
		pid := d.Name()

		// 4. raad exec cmd
		cmdlineBytes, err := os.ReadFile(filepath.Join("/proc", pid, "cmdline"))
		if err != nil || len(cmdlineBytes) == 0 {
			continue
		}

		// 5. skip null, convert blank
		cmdline := strings.ReplaceAll(string(cmdlineBytes), "\x00", " ")

		if len(cmdline) == 0 {
			continue
		}

		// 6. read exec path
		exePath, err := os.Readlink(filepath.Join("/proc", pid, "exe"))
		if err != nil {
			continue
		}

		// 7. skip systemd processes
		if strings.Contains(exePath, "systemd") || strings.Contains(exePath, "kworker") {
			continue
		}

		// 8. Check dynamic libs using ldd
		cmd := exec.Command("ldd", exePath)
		var out bytes.Buffer
		cmd.Stdout = &out
		cmd.Stderr = &out
		_ = cmd.Run()

		var libs []string
		if !strings.Contains(out.String(), "not a dynamic executable") {
			scanner := bufio.NewScanner(&out)
			for scanner.Scan() {
				line := scanner.Text()
				if strings.Contains(line, "=>") {
					for _, p := range strings.Fields(line) {
						if strings.HasPrefix(p, "/") {
							libs = append(libs, p)
						}
					}
				}
			}
		}

		//envBytes, _ := os.ReadFile(filepath.Join("/proc", pid, "environ"))
		//envs := strings.Split(string(envBytes), "\x00")

		// 9. get service file path
		var sysdFiles []string
		for _, dir := range []string{"/etc/systemd/system", "/usr/lib/systemd/system", "/lib/systemd/system"} {
			matches, _ := filepath.Glob(filepath.Join(dir, "*"+filepath.Base(exePath)+"*.service"))
			sysdFiles = append(sysdFiles, matches...)
		}

		// 10. get version
		version := ""
		if _, err := os.Stat(exePath); err == nil {
			vout := new(bytes.Buffer)
			cmd := exec.Command(exePath, "-version")
			cmd.Stdout = vout
			cmd.Stderr = vout
			if err := cmd.Run(); err == nil {
				version = strings.Split(strings.TrimSpace(vout.String()), "\n")[0]
			}
		}

		info := model.LegacySoftware{
			Name:            filepath.Base(exePath),
			Version:         version,
			NeededLibraries: libs,
			BinaryPath:      exePath,
			//EnvFiles:        envs,
			SystemdFiles:    sysdFiles,
			CustomDataPaths: []string{},
			CustomConfigs:   []string{},
		}

		result = append(result, info)
	}

	return result, nil
}

// helper function to check if directory name is PID
func isDigitDir(name string) bool {
	for _, ch := range name {
		if ch < '0' || ch > '9' {
			return false
		}
	}
	return true
}
