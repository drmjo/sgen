package main

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/client"
)

func getClient() *client.Client {
	cli, err := client.NewEnvClient()
	if err != nil {
		log.Fatal(err, ":unable create client")
	}
	return cli
}

func buildImage(dockerContext *bytes.Buffer, dockerfile string, tags []string) {
	ctx := context.Background()
	cli := getClient()
	options := types.ImageBuildOptions{
		Tags:       tags,
		Dockerfile: dockerfile,
		Remove:     true}

	res, err := cli.ImageBuild(ctx, dockerContext, options)
	if err != nil {
		log.Fatal(err, " :unable to build docker image")
	}
	defer res.Body.Close()

	if _, err := io.Copy(os.Stdout, res.Body); err != nil {
		log.Fatal(err, " :unable to read image build response")
	}
}

func getImages() {
	cli, err := client.NewEnvClient()
	if err != nil {
		panic(err)
	}

	images, err := cli.ImageList(context.Background(), types.ImageListOptions{})
	if err != nil {
		panic(err)
	}

	for _, image := range images {
		fmt.Println(image.ID)
	}
}
