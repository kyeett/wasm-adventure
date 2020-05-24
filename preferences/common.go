package preferences

import (
	"sync"
)

type Preferences struct {
	pref            map[string]interface{}
	mutex           *sync.RWMutex
	preferencesPath string
}

func New(namespace string) (*Preferences,error) {
	p := &Preferences{
		mutex: &sync.RWMutex{},
	}

	if err := p.init(namespace); err != nil {
		return nil, err
	}

	return p, nil
}

func (p *Preferences) SetItem(key string, value interface{}) error {
	p.mutex.Lock()
	defer p.mutex.Unlock()

	return p.setItem(key, value)
}

func (p *Preferences) GetString(key string) (string, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return p.getString(key)
}

func (p *Preferences) GetInt(key string) (int64, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return p.getInt(key)
}

func (p *Preferences) GetBool(key string) (bool, error) {
	p.mutex.RLock()
	defer p.mutex.RUnlock()

	return p.getBool(key)
}