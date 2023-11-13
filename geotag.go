package main

import (
	"fmt"
	"os/exec"
)

// file := "m3.jpg"
//
//	latitude := 21.015852769331687
//	longitude := 105.85046224237817
//	SetGeotag(file, latitude, longitude)
func SetGeotag(file string, latitude float64, longitude float64) {
	cmd := exec.Command("exiftool",
		"-GPSLatitude="+fmt.Sprintf("%f", latitude),
		"-GPSLongitude="+fmt.Sprintf("%f", longitude),
		file)
	output, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error setting GPS coordinates: %v\n", err)
		return
	}
	fmt.Printf("exiftool output: %s\n", output)
}
