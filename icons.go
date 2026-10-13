package mediatoolsite

import "embed"

// Identity icons are frozen public resources, never fetched with account credentials.
//
//go:embed assets/*.ico
var icons embed.FS

func Icon(id string) ([]byte, bool) {
	if _, ok := Get(id); !ok {
		return nil, false
	}
	raw, err := icons.ReadFile("assets/" + id + ".ico")
	return raw, err == nil
}
