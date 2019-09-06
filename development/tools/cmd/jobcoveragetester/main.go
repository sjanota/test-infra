package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"path/filepath"
	"syscall"
)

var (
	jobsDir  *string
	testsDir *string
)

func main() {
	initFlags()
	missedFiles := make([]string, 0)
	sigs := make(chan os.Signal, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	err := filepath.Walk(*jobsDir, func(path string, info os.FileInfo, err error) error {
		if !info.IsDir() {
			newPath := fmt.Sprintf("%s.bkp", path)

			err = os.Rename(path, newPath)
			if err != nil {
				return err
			}

			relativeJobPath, err  := filepath.Rel(*jobsDir, path)
			if err != nil {
				return err
			}

			jobTestsDir := filepath.Join(*testsDir, filepath.Dir(relativeJobPath))
			log.Printf("Checking %s. Run tests in %s", path, jobTestsDir)
			defer func() { _ = os.Rename(newPath, path) }()
			if !isOk(jobTestsDir) {
				missedFiles = append(missedFiles, path)
			}

			select {
			case <-sigs:
				return errors.New("interrupted")
			default:
			}

			return nil
		}
		return nil
	})

	if err != nil {
		log.Fatalf("Error while listing job files: %s", err)
	}

	if len(missedFiles) == 0 {
		return
	}

	log.Println("Following job files are not covered by tests:")
	for _, f := range missedFiles {
		fmt.Printf(" * %s\n", f)
	}
	os.Exit(1)
}

func initFlags() {
	jobsDir = flag.String("jobs-dir", "", "Path to the root of jobs tree")
	testsDir = flag.String("tests-dir", "", "Path to the root of tests tree")
	flag.Parse()
}

func isOk(jobTestDir string) bool {
	return testPackageExists(jobTestDir) && testFails(jobTestDir)
}

func testPackageExists(jobTestDir string) bool {
	if _, err := os.Stat(jobTestDir); os.IsNotExist(err) {
		log.Printf("Test directory does not exist: %s", jobTestDir)
		return false
	}
	return true
}

func testFails(dir string) bool {
	cmd := exec.Command("go", "test", "./"+dir)
	err := cmd.Run()
	if err != nil && err.Error() != "exit status 1" {
		log.Printf("Error while running tests: %s", err)
	}
	return err != nil
}
