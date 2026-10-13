// Package mediatoolsite provides compiled Go PT-site templates. No remote code is executed at runtime.
package mediatoolsite

import (
	"encoding/json"
	"strings"
)

type Map = map[string]any
type Field struct {
	Key      string `json:"key"`
	Label    string `json:"label"`
	Secret   bool   `json:"secret"`
	Required bool   `json:"required"`
}
type RSS struct {
	Pages    []string
	Generate Map
}
type Template struct {
	ID         string  `json:"id"`
	Name       string  `json:"name"`
	Domain     string  `json:"domain"`
	Fields     []Field `json:"fields"`
	Definition Map     `json:"-"`
	RSS        RSS     `json:"-"`
}

func clone(m Map) Map { raw, _ := json.Marshal(m); var out Map; json.Unmarshal(raw, &out); return out }
func (t Template) Snapshot() ([]byte, error) {
	m := clone(t.Definition)
	m["templateVersion"] = 2
	m["id"] = t.ID
	m["name"] = t.Name
	m["domain"] = t.Domain
	return json.Marshal(m)
}
func Get(id string) (Template, bool) {
	for _, t := range All() {
		if t.ID == id {
			return t, true
		}
	}
	return Template{}, false
}
func All() []Template { return []Template{hdsky(), ourbits(), chdbits(), ssd(), ttg(), ubits()} }
func cookieFields() []Field {
	return []Field{{Key: "cookie", Label: "Cookie", Secret: true, Required: true}, {Key: "userAgent", Label: "User-Agent", Required: true}}
}
func basic(id, name, domain string) Template {
	return Template{ID: id, Name: name, Domain: domain, Fields: cookieFields(), Definition: Map{"schema": "NexusPHP", "encoding": "UTF-8", "requests": commonRequests()}, RSS: RSS{Pages: []string{"/getrss.php", "/userdetails.php?id={userId}"}, Generate: rssGenerator(1)}}
}
func rssGenerator(index int) Map {
	return Map{"parser": "XPath", "method": "POST", "path": "/getrss.php", "form_data": Map{"inclbookmarked": []any{"0"}, "paid": []any{"0"}, "itemcategory": []any{"1"}, "itemsmalldescr": []any{"1"}, "itemsize": []any{"1"}, "showrows": []any{"10"}, "search_mode": []any{"0"}}, "fields": Map{"content": Map{"selector": "(//a[@class='faqlink'])[" + strings.TrimSpace(string(rune('0'+index))) + "]/@href"}}}
}
