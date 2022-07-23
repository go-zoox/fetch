package fetch

import "strings"

type cookie struct {
	data []cookieItem
}

type cookieItem struct {
	Key   string
	Value string
}

func newCookie(origin ...string) *cookie {
	c := &cookie{}

	if len(origin) > 0 {
		if origin[0] != "" {
			c.Parse(origin[0])
		}
	}

	return c
}

func (c *cookie) Parse(str string) (err error) {
	str = strings.TrimSpace(str)
	str = strings.TrimSuffix(str, ";")
	str = strings.TrimSpace(str)

	items := strings.Split(str, ";")

	for _, item := range items {
		item = strings.TrimSpace(item)
		if item == "" {
			continue
		}

		kv := strings.Split(item, "=")
		if len(kv) != 2 {
			continue
		}

		if err := c.Add(kv[0], kv[1]); err != nil {
			return err
		}
	}

	return
}

func (c *cookie) Add(key, value string) (err error) {
	if key == "" {
		return ErrCookieEmptyKey
	}

	c.data = append(c.data, cookieItem{key, value})
	return nil
}

func (c *cookie) Get(key string) string {
	for _, item := range c.data {
		if item.Key == key {
			return item.Value
		}
	}
	return ""
}

func (c *cookie) Set(key, value string) (err error) {
	for i, item := range c.data {
		if item.Key == key {
			c.data[i].Value = value
			return nil
		}
	}

	return c.Add(key, value)
}

func (c *cookie) Remove(key string) (err error) {
	for i, item := range c.data {
		if item.Key == key {
			c.data = append(c.data[:i], c.data[i+1:]...)
			return
		}
	}

	return
}

func (c *cookie) Clear() error {
	c.data = nil
	return nil
}

func (c *cookie) Items() []cookieItem {
	return c.data
}

func (c *cookie) String() string {
	var ss []string
	for _, item := range c.data {
		ss = append(ss, item.Key+"="+item.Value)
	}

	return strings.Join(ss, "; ")
}
