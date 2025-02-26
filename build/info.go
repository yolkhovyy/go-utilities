package info

import (
	"runtime/debug"
)

type Info struct {
	Commit  string
	Time    string
	Version string
}

const unknown = "unknown"

//nolint:gochecknoglobals
var tag = unknown

// ReadInfo extracts build metadata.
func ReadInfo() Info {
	info := Info{
		Commit:  unknown,
		Time:    unknown,
		Version: tag,
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
