// Copyright 2020 Axel Etcheverry. All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package version

import (
	"fmt"
	"runtime"
	"strings"

	"github.com/coreos/go-semver/semver"
	"github.com/pkg/errors"
)

var version string
var gitCommit string
var gitTreeState string
var buildDate string
var platform = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)

// Info version
type Info struct {
	Version      string
	GitVersion   string
	GitCommit    string
	GitTreeState string
	BuildDate    string
	GoVersion    string
	Compiler     string
	Platform     string
}

var instance *Info

// Get returns the version and buildtime information about the binary
func Get() *Info {
	if instance == nil {
		// These variables typically come from -ldflags settings to `go build`
		instance = &Info{
			Version:      version,
			GitCommit:    gitCommit,
			GitTreeState: gitTreeState,
			BuildDate:    buildDate,
			GoVersion:    runtime.Version(),
			Compiler:     runtime.Compiler,
			Platform:     platform,
		}
	}

	return instance
}

// Parse version string to semver.Version
func Parse(version string) (*semver.Version, error) {
	// Strip the leading 'v' in our version strings
	version = strings.TrimSpace(version)
	v, err := semver.NewVersion(strings.TrimLeft(version, "v"))
	if err != nil {
		return nil, errors.Wrap(err, "parsing semver")
	}

	return v, nil
}
