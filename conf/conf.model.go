package conf

import (
	"strconv"
)

type Conf map[string]string

func (c *Conf) Int(name string) (int, error) {
	return strconv.Atoi((*c)[name])
}

func (c *Conf) Str(name string) string {
	return (*c)[name]
}
