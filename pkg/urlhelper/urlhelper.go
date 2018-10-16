package urlhelper

import "net/url"

// URLHelper contains url helper methods.
type URLHelper struct {
	baseURL *url.URL
}

// New is a constructor for URLHelper struct.
func New() *URLHelper {
	return &URLHelper{}
}

// Get returns url string.
func (uh *URLHelper) Get() string {
	return uh.baseURL.String()
}

// Update changes baseURL value.
func (uh *URLHelper) Update(new string) error {
	u, err := url.Parse(new)
	if err != nil {
		return err
	}
	uh.baseURL = u
	return nil
}

// Map transforms string pair to map.
func (uh *URLHelper) Map(k, v string) map[string]string {
	return map[string]string{
		k: v,
	}
}

// Add creates new query param.
func (uh *URLHelper) Add(kv ...map[string]string) {
	q := uh.baseURL.Query()
	for _, mp := range kv {
		for k, v := range mp {
			q.Add(k, v)
		}
	}
	uh.baseURL.RawQuery = q.Encode()
}

// Set update query param value.
func (uh *URLHelper) Set(k, v string) {
	q := uh.baseURL.Query()
	q.Set(k, v)
	uh.baseURL.RawQuery = q.Encode()
}
