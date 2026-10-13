package mediatoolsite

func hdsky() Template {
	t := basic("hdsky", "HDSKY", "https://hdsky.me")
	t.Definition["schema"] = "NexusPHP"
	t.Definition["encoding"] = "UTF-8"
	search := t.Definition["requests"].(Map)["search"].(Map)
	search["list"].(Map)["empty_evidence"] = Map{
		"selector":      `td.text:contains("没有种子。请用准确的关键字重试。")`,
		"authenticated": `a[href*="logout"], a[href*="usercp"]`,
	}
	return t
}
