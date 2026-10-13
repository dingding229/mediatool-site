package mediatoolsite

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestCompiledSiteDefinitions(t *testing.T) {
	seen := map[string]bool{}
	for _, s := range All() {
		if seen[s.ID] || s.Name != strings.ToUpper(s.Name) || s.Domain == "" {
			t.Fatal(s.ID)
		}
		seen[s.ID] = true
		raw, e := s.Snapshot()
		if e != nil {
			t.Fatal(e)
		}
		var v Map
		json.Unmarshal(raw, &v)
		if v["templateVersion"] != float64(2) || len(v["requests"].(map[string]any)) != 4 {
			t.Fatal(string(raw))
		}
		if len(s.Fields) != 2 || !s.Fields[0].Required {
			t.Fatal(s.ID)
		}
	}
	if len(seen) != 6 {
		t.Fatal(seen)
	}
}
