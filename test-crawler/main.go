package main

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"

	c "testsuites/collector"
	ex "testsuites/extractor"
)

const OUTPUT_FOLDER = "outputs"

func main() {

	config := NewConfig()

	// for origin mode, pull or clone remote origin, for local just ignore and continue
	if config.Repo.Mode == "origin" {
		for i, repo := range config.Repo.Name {
			if _, err := os.Stat(fmt.Sprintf("%s/%s", config.Repo.Destination, repo)); os.IsNotExist(err) {
				_ = os.Mkdir(config.Repo.Destination, 7660)
				_, _, err := Shell(fmt.Sprintf("git clone %s %s", config.Repo.Origin[i], repo), config.Repo.Destination)
				if err != nil {
					fmt.Printf("error: %s", err.Error())
				}
			} else {
				Shell("git pull", config.Repo.Destination)
			}
		}
	}

	files, err := c.GetTestFiles(config.Repo.Destination, config.Language)
	if err != nil {
		fmt.Println(err)
		return
	}

	ctx := context.Background()

	for i, file := range files {
		scenarios, meta, err := ex.ExtractInfo(file, ctx, config.Language)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if scenarios == nil && meta == nil {
			continue
		}

		files[i].Package = meta.Package
		files[i].TestType = meta.TestType
		files[i].Ignore = meta.Ignore
		files[i].Scenarios = scenarios
	}

	Save(files, config.OutputMode)
}

func Save(files []c.TestFile, mode OutputMode) {

	content, err := json.Marshal(files)
	if err != nil {
		return
	}

	switch mode {
	case MODE_FILE:
		now := time.Now()
		timestamp := now.Unix()

		filename := fmt.Sprintf("%s/output_%d.json", OUTPUT_FOLDER, timestamp)

		file, err := os.Create(filename)
		if err != nil {
			return
		}
		defer file.Close()

		file.Write(content)
	case MODE_STDOUT:
		fmt.Print(string(content))
	default:
		fmt.Print(string(content))
	}

}

func Shell(command string, destination string) (string, string, error) {
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd := exec.Command("bash", "-c", command)
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	cmd.Dir = fmt.Sprintf("./%s", destination)
	err := cmd.Run()
	return stdout.String(), stderr.String(), err
}
