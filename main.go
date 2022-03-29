package main

import (
	"flag"
	"os"
	"strings"
)

var (
	clean bool
)

func init() {
	flag.BoolVar(&clean, "i", false, "clear the environment then run command")
	flag.Parse()
}

type Env map[string]string

func GetEnv() Env {
	var e Env = make(map[string]string)
	envs := os.Environ()
	for _, item := range envs {
		s := strings.Split(item, "=")
		key := s[0]
		value := strings.Join(s[1:], "=")
		e[key] = value
	}
	return e
}
func (e Env) String() string {
	var s string
	for key, value := range e {
		s = s + key + "=" + value + "\n"
	}
	return strings.TrimRight(s, "\n")
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
	Exec(e, cmd)
}
