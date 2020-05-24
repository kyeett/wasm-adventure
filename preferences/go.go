// +build !js

package preferences

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
)

const (
	fileName = "preferences"
)

func dirNotExists(path string) bool {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return true
	}
	return false
}

func (p *Preferences) init(namespace string) error {
	p.pref = map[string]interface{}{}

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	// Path to config file
	p.preferencesPath = path.Join(home, ".config", namespace)

	// Check if path exists
	if dirNotExists(p.preferencesPath) {
		if err := createDirectory(p.preferencesPath); err != nil {
			return err
		}
		return nil
	}

	return p.loadFromFile()
}

func (p *Preferences) loadFromFile() error {
	filePath := path.Join(p.preferencesPath, fileName)
	f, err := os.OpenFile(filePath, os.O_RDONLY|os.O_CREATE, 0755)
	if err != nil {
		return err
	}
	defer f.Close()

	var m map[string]interface{}
	if err := json.NewDecoder(f).Decode(&m); err != nil {
		fmt.Println(err)
		return nil
	}
	p.pref = m
	return nil
}

func createDirectory(path string) error {
	// Create directory
	if err := os.MkdirAll(path, 0755); err != nil {
		return err
	}

	return nil
}

func (p *Preferences) setItem(key string, value interface{}) error {
	p.pref[key] = value

	filePath := path.Join(p.preferencesPath, fileName)
	f, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0755)
	if err != nil {
		return err
	}

	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")

	if err := enc.Encode(p.pref); err != nil {
		return err
	}

	return nil
}

func (p *Preferences) getString(key string) (string, error) {
	val, found := p.pref[key]
	if !found {
		return "", fmt.Errorf("%q not found", key)
	}

	s, ok := val.(string)
	if !ok {
		return "", fmt.Errorf("%q is type %T, not string", key, val)
	}

	return s, nil
}

func (p *Preferences) getBool(key string) (bool, error) {
	val, found := p.pref[key]
	if !found {
		return false, fmt.Errorf("%q not found", key)
	}

	t, ok := val.(bool)
	if !ok {
		return false, fmt.Errorf("%q is type %T, not bool", key, val)
	}

	return t, nil
}

func (p *Preferences) getInt(key string) (int64, error) {
	val, found := p.pref[key]
	if !found {
		return 0, fmt.Errorf("%q not found", key)
	}

	switch v := val.(type) {
	case int64:
		return v, nil
	case int32:
		return int64(v), nil
	case int:
		return int64(v), nil
	case float64:
		return int64(v), nil
	case float32:
		return int64(v), nil
	default:
		return 0, fmt.Errorf("%q is type %T, not int64", key, val)
	}
}
