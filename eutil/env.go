package eutil

import (
	"fmt"
	"os"
	"strings"
)

func GetEnvs() {
	envs := os.Environ()
	for _, env := range envs {
		cache := strings.Split(env, "=")
		fmt.Printf("%v = %v\n", cache[0], cache[1])
	}
}
