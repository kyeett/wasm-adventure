//+build js,wasm

package preferences

import (
	"fmt"
	"strconv"
	"syscall/js"
)

const (
	setFunc = "setPreferences"
	getFunc = "getPreferences"
)

func (p *Preferences) init(_ string) error {
	setAvailable := js.Global().Get(setFunc).Type() == js.TypeFunction
	if !setAvailable {
		return fmt.Errorf("javascript function %q missing", setFunc)
	}

	getAvailable := js.Global().Get(getFunc).Type() == js.TypeFunction
	if !getAvailable {
		return fmt.Errorf("javascript function %q missing", getFunc)
	}

	return nil
}

func (p *Preferences) setItem(key string, value interface{}) error{
	js.Global().Call(setFunc, key, value)
	return nil
}

func (p *Preferences) getString(key string) (string, error) {
	v := js.Global().Call(getFunc, key)
	if v.IsNull() {
		return "", fmt.Errorf("%q not found", key)
	}

	if v.Type() != js.TypeString {
		return "", fmt.Errorf("%q is type %q, not string", key, v.Type())
	}

	return v.String(), nil
}

func (p *Preferences) getBool(key string) (bool, error) {
	s, err := p.getString(key)
	if err != nil {
		return false, err
	}

	t, err := strconv.ParseBool(s)
	if err != nil {
		return false, err
	}

	return t, nil
}

func (p *Preferences) getInt(key string) (int64, error) {
	s, err := p.getString(key)
	if err != nil {
		return 0, err
	}

	i, err := strconv.ParseInt(s, 0, 64)
	if err != nil {
		return 0, err
	}

	return i, nil
}