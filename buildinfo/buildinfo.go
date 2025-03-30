package buildinfo

import (
	"runtime/debug"
)

type Data struct {
	Revision string
	Time     string
}

const unknown = "unknown"

// ReadData extracts build metadata.
func ReadData() Data {
	info := Data{
		Revision: unknown,
		Time:     unknown,
	}

	if buildInfo, ok := debug.ReadBuildInfo(); ok {
		for _, setting := range buildInfo.Settings {
			switch setting.Key {
			case vcsKeyRevision:
				info.Revision = setting.Value
			case vcsKeyTime:
				info.Time = setting.Value
			case vcsKeyModified:
				if setting.Value == "true" {
					info.Revision += " (modified)"
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
