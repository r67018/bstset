package bst

import (
	"bufio"
	"io/ioutil"
	"os"
	"reflect"
	"sort"
	"strings"
)

const (
	NAME = iota
	VALUE
)

const (
	DEFAULT_FILE_PERMISSION = 0664
	DEFAULT_CONFIG_PATH     = "C:\\ProgramData\\BlueStacks_nxt\\bluestacks.conf"
)

type Option map[string]string

type Config struct {
	m Option
}

func NewConfig() *Config {
	return &Config{
		m: Option{},
	}
}

func (c *Config) Load(path string) error {
	fp, err := os.Open(path)
	if err != nil {
		return err
	}
	defer fp.Close()

	scanner := bufio.NewScanner(fp)

	for scanner.Scan() {
		line := scanner.Text()
		token := strings.Split(line, "=")
		name := token[NAME]
		value := strings.ReplaceAll(token[VALUE], "\"", "")
		c.m[name] = value
	}
	return nil
}

func (c *Config) Apply(o Option) {
	for name, value := range o {
		c.m[name] = value
	}
}

func (c *Config) Write(path string) error {
	data := []byte{}
	for _, name := range sortedKeys(c.m) {
		value := c.m[name]
		line := connectToken(name, value)
		lineBytes := []byte(line)
		data = append(data, lineBytes...)
	}
	err := ioutil.WriteFile(path, data, DEFAULT_FILE_PERMISSION)
	if err != nil {
		return err
	}
	return nil
}

func sortedKeys(m any) []string {
	values := reflect.ValueOf(m).MapKeys()
	result := make([]string, len(values))
	for i, value := range values {
		result[i] = value.String()
	}
	sort.Strings(result)
	return result
}

func connectToken(name, value string) string {
	return name + "=\"" + value + "\"\n"
}
