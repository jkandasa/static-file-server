package version

import (
	"fmt"
	"runtime"
)

var (
	gitCommit string
	version   string
	buildDate string
)

// Version holds version data
type Version struct {
	Version   string `json:"version"`
	GitCommit string `json:"gitCommit"`
	BuildDate string `json:"buildDate"`
	GoLang    string `json:"goLang"`
	Platform  string `json:"platform"`
	Arch      string `json:"arch"`
}

func (v Version) ToString() string {
	return fmt.Sprintf("version:%s, goLang:%s, platform:%s, arch:%s, gitCommit:%s, buildDate:%s",
		v.Version, v.GoLang, v.Platform, v.Arch, v.GitCommit, v.BuildDate)
}

// Get returns the Version object
func Get() Version {
	return Version{
		GitCommit: gitCommit,
		Version:   version,
		BuildDate: buildDate,
		GoLang:    runtime.Version(),
		Platform:  runtime.GOOS,
		Arch:      runtime.GOARCH,
	}
}
