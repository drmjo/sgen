package main

import (
	"log"
)

const DEPLOYER_IMAGE_NAME = "sgen/deployer:tf"

func main() {

	dockerContext, err := buildTar(flags.TerraformDir)
	if err != nil {
		log.Printf("ERROR: %v", err.Error())
	}

	buildImage(dockerContext, "Dockerfile", []string{DEPLOYER_IMAGE_NAME})

	getImages()
}
