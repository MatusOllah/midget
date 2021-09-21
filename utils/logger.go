package utils

import (
	"fmt"
	"os"

	"github.com/fatih/color"
)

func CheckError(a interface{}) {
	if a != nil {
		LogError(a)
	}
}

func LogError(a interface{}) {
	r := color.New(color.FgRed, color.Bold).SprintFunc()

	fmt.Printf("\a%s: %v\n", r("Error"), a)
	os.Exit(1)
}

func LogWarning(a interface{}) {
	y := color.New(color.FgYellow, color.Bold).SprintFunc()

	fmt.Printf("\a%s: %v\n", y("Warning"), a)
}

func PrintTriggers(trigs []string) {
	y := color.New(color.FgYellow, color.Bold).SprintFunc()

	if trigs != nil {
		fmt.Printf("\a%s: \n", y("Trigger warning"))
		for _, trig := range trigs {
			fmt.Printf("\t%s\n", trig)
		}
	}
}
