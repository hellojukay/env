package main

import (
	"os"
	"sort"
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
	var keys []string
	for key := range *e {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var s string
	for _, key := range keys {
		s = s + key + "=" + (*e)[key] + "\n"
	}

	return strings.TrimRight(s, "\n")
}

func (e *Env) Set(key string, value string) {
	(*e)[key] = value
}

func (e *Env) Remove(key string) {
	delete(*e, key)
}
