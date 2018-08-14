package mycnf

import (
	"github.com/go-ini/ini"
	"github.com/mitchellh/go-homedir"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

// ReadDefault reads configuration from default location (~/.my.cnf)
func ReadDefault() ([]Config, error) {
	path, err := homedir.Expand("~/.my.cnf")
	if err != nil {
		return nil, err
	}
	return Read(path)
}

// Read reads .my.cnf file
func Read(filename string) ([]Config, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return nil, nil
}

func parse(r io.Reader) ([]Config, error) {
	f, err := ini.Load(ioutil.NopCloser(r))
	if err != nil {
		return nil, err
	}

	var cc []Config
	for _, section := range f.Sections() {
		name := strings.ToLower(section.Name())
		if !strings.HasPrefix(name, "client") {
			continue
		}

		c := Config{ConnectionName: name[6:]}
		c.User = orElse(section, "user", "")
		c.Passwd = orElse(section, "password", "")
		c.DBName = orElse(section, "database", "")
		host := orElse(section, "host", "localhost")
		port := orElse(section, "port", "3306")
		c.Addr = host + ":" + port

		cc = append(cc, c)
	}

	return cc, nil
}

func orElse(s *ini.Section, key, els string) string {
	v, err := s.GetKey(key)
	if err != nil || v == nil || len(v.Value()) == 0 {
		return els
	}

	return v.Value()
}
