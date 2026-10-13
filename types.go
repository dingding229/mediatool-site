// Package mediatoolsite provides compiled Go PT-site templates. No remote code is executed at runtime.
package mediatoolsite

import (
	"encoding/json"
	"strconv"
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
	Version    string  `json:"version"`
	IconURL    string  `json:"iconUrl"`
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
func All() []Template {
	templates := []Template{hdsky(), ourbits(), chdbits(), ssd(), ttg(), ubits()}
	for i := range templates {
		rule := templates[i].Definition["requests"].(Map)["search"].(Map)
		rule["promotion"] = Map{"normal_factor": 1, "marker_selector": "[class*='pro_'], [class*='torrent-pro-'], img[alt*='%'], img[alt*='Free'], img[alt*='free'], img[alt*='FREE'], .seed-box-policy-download", "known_markers": []any{"pro_free", "pro_free2up", "pro_50pctdown", "pro_50pctdown2up", "pro_30pctdown", "pro_2up", "pro_twoup", "torrent-pro-icon", "torrent-pro-free", "torrent-pro-twoupfree", "torrent-pro-halfdown", "torrent-pro-twouphalfdown", "torrent-pro-thirtypercent", "torrent-pro-refund-left", "torrent-pro-refund-right", "seed-box-policy-status", "seed-box-policy-download", "is-blocked", "is-limited", "is-unlimited"}}
		if templates[i].ID == "ubits" {
			factor := rule["fields"].(Map)["downloadvolumefactor"].(Map)
			factor["any"] = append([]any{Map{"selector": ".seed-box-policy-download.is-blocked", "filters": []any{Map{"name": "not_blank", "args": []any{1, ""}}}}}, factor["any"].([]any)...)
		}
	}
	return templates
}
func cookieFields() []Field {
	return []Field{
		{
			Key:      "cookie",
			Label:    "Cookie",
			Secret:   true,
			Required: true,
		},
		{
			Key:      "userAgent",
			Label:    "User-Agent",
			Required: true,
		},
	}
}
func basic(id, name, domain string) Template {
	return Template{
		ID:      id,
		Name:    name,
		Domain:  domain,
		Version: "1.2.0",
		IconURL: domain + "/favicon.ico",
		Fields:  cookieFields(),
		Definition: Map{
			"schema":   "NexusPHP",
			"encoding": "UTF-8",
			"requests": commonRequests(),
		},
		RSS: RSS{
			Pages: []string{
				"/getrss.php",
				"/userdetails.php?id={userId}",
			},
			Generate: rssGenerator(1),
		},
	}
}
func rssGenerator(index int) Map {
	return Map{
		"parser": "XPath",
		"method": "POST",
		"path":   "/getrss.php",
		"form_data": Map{
			"inclbookmarked": []any{
				"0",
			},
			"paid": []any{
				"0",
			},
			"itemcategory": []any{
				"1",
			},
			"itemsmalldescr": []any{
				"1",
			},
			"itemsize": []any{
				"1",
			},
			"showrows": []any{
				"10",
			},
			"search_mode": []any{
				"0",
			},
		},
		"fields": Map{
			"content": Map{
				"selector": "(//a[@class='faqlink'])[" + strconv.Itoa(index) + "]/@href",
			},
		},
	}
}
