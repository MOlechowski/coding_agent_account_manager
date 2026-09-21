//go:build !darwin

package authfile

const agyKeychainDescription = "unsupported keychain"

func agyKeychainSupported() bool { return false }

func agyKeychainToken() ([]byte, bool, error) { return nil, false, nil }

func setAgyKeychainToken([]byte) error { return nil }

func clearAgyKeychainToken() error { return nil }
