package utils

import (
	"fmt"
	"time"

	"github.com/cavaliercoder/grab"
)

func DownloadFile(url string, path string) error {
	// create client
	client := grab.NewClient()
	req, _ := grab.NewRequest(path, url)

	// start download
	fmt.Printf("Downloading file from %v...\n", req.URL())
	resp := client.Do(req)
	fmt.Printf("Sent request, response: %v\n", resp.HTTPResponse.Status)

	// start UI loop
	t := time.NewTicker(100 * time.Millisecond)
	defer t.Stop()

Loop:
	for {
		select {
		case <-t.C:
			updateBarDL("Downloading file from "+req.URL().String()+"...", int(100*resp.Progress()), resp.BytesComplete(), resp.Size)

		case <-resp.Done:
			fmt.Println()
			// if the download is complete then break out of this loop
			break Loop
		}
	}

	// check for errors
	if err := resp.Err(); err != nil {
		return err
	}

	return nil
}
