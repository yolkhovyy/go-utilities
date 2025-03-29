package buildinfo

import (
	"runtime/debug"
)

type Data struct {
	Commit  string
	Time    string
	Version string
}

const unknown = "unknown"

//nolint:gochecknoglobals
var Version string

// ReadData extracts build metadata.
func ReadData() Data {
	info := Data{
		Commit:  unknown,
		Time:    unknown,
		Version: Version,
	}

	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range buildInfo.Settings {
			switch setting.Key {
			case vcsKeyRevision:
				info.Commit = setting.Value
			case vcsKeyTime:
				info.Time = setting.Value
			case vcsKeyModified:
				if setting.Value == "true" {
					info.Commit += " (modified)"
				}
			}
		}
	}

	return info
}

type vcsKey = string

const (
	vcsKeyRevision vcsKey = "vcs.revision"
	vcsKeyTime     vcsKey = "vcs.time"
	vcsKeyModified vcsKey = "vcs.modified"
)
