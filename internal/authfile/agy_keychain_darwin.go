//go:build darwin

package authfile

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

const (
	agyKeychainService     = "gemini"
	agyKeychainAccount     = "antigravity"
	agyKeychainDescription = "macOS Keychain: service=gemini account=antigravity"
)

func agyKeychainSupported() bool { return os.Getenv("GEMINI_HOME") == "" }

func agyKeychainToken() ([]byte, bool, error) {
	if !agyKeychainSupported() {
		return nil, false, nil
	}
	output, err := exec.Command("/usr/bin/security", "find-generic-password", "-s", agyKeychainService, "-a", agyKeychainAccount, "-w").Output()
	if err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 44 {
			return nil, false, nil
		}
		return nil, false, fmt.Errorf("find-generic-password: %w", err)
	}
	return bytes.TrimSpace(output), true, nil
}

func setAgyKeychainToken(token []byte) error {
	if !agyKeychainSupported() {
		return nil
	}
	cmd := exec.Command("/usr/bin/security", "add-generic-password", "-U", "-s", agyKeychainService, "-a", agyKeychainAccount, "-w", string(token))
	if output, err := cmd.CombinedOutput(); err != nil {
		return fmt.Errorf("add-generic-password: %w: %s", err, bytes.TrimSpace(output))
	}
	return nil
}

func clearAgyKeychainToken() error {
	if !agyKeychainSupported() {
		return nil
	}
	cmd := exec.Command("/usr/bin/security", "delete-generic-password", "-s", agyKeychainService, "-a", agyKeychainAccount)
	if output, err := cmd.CombinedOutput(); err != nil {
		if exitErr, ok := err.(*exec.ExitError); ok && exitErr.ExitCode() == 44 {
			return nil
		}
		return fmt.Errorf("delete-generic-password: %w: %s", err, bytes.TrimSpace(output))
	}
	return nil
}
