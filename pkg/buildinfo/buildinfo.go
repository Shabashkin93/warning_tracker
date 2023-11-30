package buildinfo

import (
	"fmt"
	"os"
)

var (
	version    string
	buildTime  string
	commitHash string
)

type Buildinfo struct {
	Version    string
	BuildTime  string
	CommitHash string
}

func PrintBuildInfo(args []string) {
	if len(args) == 2 && args[1] == "version" {
		fmt.Println("Version:    ", version)
		fmt.Printf("Build Time: %s\n", buildTime)
		fmt.Printf("Commit Hash: %s\n\n", commitHash)
		os.Exit(0)
	}
}

func GetBuildInfo() (buildinfo *Buildinfo) {
	buildinfo = &Buildinfo{Version: version, BuildTime: buildTime, CommitHash: commitHash}
	return buildinfo
}
