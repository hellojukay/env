package main

import (
	"flag"
	"os"
	"strings"
)

var (
	clean  bool
	unsets arrayFlags
)

func init() {
	flag.BoolVar(&clean, "i", false, "clear the environment then run command")
	flag.Var(&unsets, "u", "remove variable from the environment and run a program:")
	flag.Parse()
}

type Env map[string]string

func GetEnv() *Env {
	var e Env = make(map[string]string)
	envs := os.Environ()
	for _, item := range envs {
		s := strings.Split(item, "=")
		key := s[0]
		value := strings.Join(s[1:], "=")
		e[key] = value
	}
	return &e
}
func (e *Env) String() string {
	var s string
	for key, value := range *e {
		s = s + key + "=" + value + "\n"
	}
	return strings.TrimRight(s, "\n")
}

func (e *Env) Remove(key string) {
	delete(*e, key)
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
	Exec(*e, cmd)
}
