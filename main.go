package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"runtime/debug"
	"strings"
)

var (
	clean       bool
	unsets      arrayFlags
	showVersion bool
)

func init() {
	flag.BoolVar(&showVersion, "v", false, "show program version than exit")
	flag.BoolVar(&clean, "i", false, "clear the environment then run command")
	flag.Var(&unsets, "u", "remove variable from the environment and run a program:")
	flag.Parse()

	if showVersion {
		printVersion()
	}
}

func main() {
	e := GetEnv()
	if len(os.Args) == 1 {
		println(e.String())
		os.Exit(0)
	}
	cmd := os.Args[len(os.Args)-1]
	if clean {
		os.Clearenv()
		e = GetEnv()
	}
	for _, env := range unsets {
		e.Remove(env)
	}
	r := regexp.MustCompile(`.+=.+`)
	for _, arg := range os.Args {
		if !r.Match([]byte(arg)) {
			continue
		}
		arr := strings.Split(arg, "=")
		e.Set(arr[0], arr[1])
	}
	Exec(*e, cmd)
}

func printVersion() {
	info, _ := debug.ReadBuildInfo()
	fmt.Printf("%s build with %s\n", info.Main.Version, info.GoVersion)
	os.Exit(0)
}
