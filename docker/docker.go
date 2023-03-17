package docker

import "os"

const DockerEnvFile = "/.dockerenv"

func IsRunInDocker() bool {
	_, err := os.Stat(DockerEnvFile)
	return err == nil
}
