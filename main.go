package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-exec/tfexec"
	"log"
	"os/exec"
)

// write code to print something to the console
func main() {
	workingDir := "."

	tfPath, err := exec.LookPath("terraform")

	if err != nil {
		log.Fatalf("error running LookPath: %s", err)
	}

	tf, err := tfexec.NewTerraform(workingDir, tfPath)
	if err != nil {
		log.Fatalf("error running NewTerraform: %s", err)
	}

	err = tf.Init(context.Background(), tfexec.Upgrade(true))
	if err != nil {
		log.Fatalf("error running Init: %s", err)
	}

	state, err := tf.Show(context.Background())
	if err != nil {
		log.Fatalf("error running Show: %s", err)
	}

	fmt.Println(state.FormatVersion)
}
