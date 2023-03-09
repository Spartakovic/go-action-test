package main

import (
	"context"
	"fmt"
	"github.com/hashicorp/terraform-exec/tfexec"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
)

// write code to print something to the console
func main() {
	argsWithoutProg := os.Args[1:]
	workingDir := argsWithoutProg[0]
	tfCommand := argsWithoutProg[1]

	tfPath, err := exec.LookPath("terraform")

	if err != nil {
		log.Fatalf("error running LookPath: %s", err)
	}

	tf, err := tfexec.NewTerraform(workingDir, tfPath)
	if err != nil {
		log.Fatalf("error running NewTerraform: %s", err)
	}

	err = tf.Init(context.Background(), tfexec.Upgrade(false))
	if err != nil {
		log.Fatalf("error running Init: %s", err)
	}

	if tfCommand == "plan" {
		_, err := tf.Plan(context.Background(), tfexec.Out("plan.tfplan"))
		if err != nil {
			log.Fatalf("error running Plan: %s", err)
		}
		data, err := ioutil.ReadFile("plan.tfplan")
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(string(data))

	} else if tfCommand == "apply" {
		err := tf.Apply(context.Background())
		if err != nil {
			log.Fatalf("error running Apply: %s", err)
		}
	}
}
