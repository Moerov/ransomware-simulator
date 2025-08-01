package elastic

import (
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func Simulate() error {
	log.Println("Searching for Elastic canary directory in C:\\")

	root := `C:\`
	var canaryDir string

	// Step 1: Find a folder that matches the pattern
	entries, err := os.ReadDir(root)
	if err != nil {
		return fmt.Errorf("error reading root directory: %w", err)
	}

	for _, entry := range entries {
	if entry.IsDir() && strings.HasPrefix(entry.Name(), "aaAntiRansomElastic-DO-NOT-TOUCH-") {
		canaryDir = filepath.Join(root, entry.Name())
		break
		}
	}

	if canaryDir == "" {
		return fmt.Errorf("no Elastic canary folder found in %s", root)
	}

	log.Printf("Found canary directory: %s", canaryDir)

	// Step 2: Find a matching file in that folder
	var originalFile string
	entries, err = os.ReadDir(canaryDir)
	if err != nil {
		return fmt.Errorf("error reading canary directory: %w", err)
	}
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasPrefix(entry.Name(), "AntiRansomElastic-DO-NOT-TOUCH-") && strings.HasSuffix(entry.Name(), ".txt") {
			originalFile = filepath.Join(canaryDir, entry.Name())
			break
		}
	}
	if originalFile == "" {
		return fmt.Errorf("no Elastic canary file found in folder")
	}
	log.Printf("Found canary file: %s", originalFile)

	// Step 3: Backup the file
	backupFile := originalFile + `.bcp`
	if err := copyFile(originalFile, backupFile); err != nil {
		return fmt.Errorf("failed to back up file: %w", err)
	}
	log.Printf("Created backup: %s", backupFile)

	// Step 4: Overwrite the original file
	if err := os.WriteFile(originalFile, []byte("Possehl ransomware simulation"), 0644); err != nil {
		return fmt.Errorf("failed to overwrite canary file: %w", err)
	}
	log.Println("Overwrote canary file with simulation string")

	// Step 5: Wait 1 minute
	log.Println("Waiting 1 minute before restoring...")
	time.Sleep(1 * time.Minute)

	// Step 6: Restore from backup
	if err := copyFile(backupFile, originalFile); err != nil {
		return fmt.Errorf("failed to restore original file: %w", err)
	}
	log.Println("Successfully restored original canary file")

	return nil
}

// copyFile copies src to dst (overwrites if exists)
func copyFile(src, dst string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()

	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer func() {
		_ = out.Close()
	}()

	if _, err := io.Copy(out, in); err != nil {
		return err
	}
	return out.Sync()
}
