package main

import (
	"fmt"
	"os"
	"os/exec"
)

func Exec(e Env, cmd string) error {
	c := exec.Command(cmd)

	c.Env = nil
	var envs []string
	for key, value := range e {
		env := fmt.Sprintf("%s=%s", key, value)
		envs = append(envs, env)
	}
	c.Env = envs

	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	return c.Run()
}
