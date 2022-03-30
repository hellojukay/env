package main

import (
	"os"
	"strings"
)

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

func (e *Env) Set(key string, value string) {
	(*e)[key] = value
}

func (e *Env) Remove(key string) {
	delete(*e, key)
}
