package main

import "flag"

var flags struct {
	TerraformDir string
	Bootstrap    bool
}

func init() {
	flag.StringVar(&flags.TerraformDir, "terraform-dir", "terraform", "The directory where the terraform docker container context is")
	flag.BoolVar(&flags.TerraformDir, "terraform-dir", "terraform", "The directory where the terraform docker container context is")

	flag.Parse()
}
