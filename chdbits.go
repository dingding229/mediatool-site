package mediatoolsite

func chdbits() Template {
	t := basic("ptchdbits", "CHDBITS", "https://ptchdbits.co")
	t.Definition["schema"] = "NexusPHP"
	t.Definition["encoding"] = "UTF-8"
	t.Definition["additional_domain"] = []any{
		"https://bj.chddiy.xyz",
		"https://js.chddiy.xyz",
		"https://hn.chddiy.xyz",
	}
	t.Definition["requests"].(Map)["search"] = Map{
		"parser": "CssSelector",
		"method": "GET",
		"path":   "/torrents.php",
		"params": Map{
			"search": []any{
				"{keyword}",
			},
		},
		"list": Map{
			"selector":  "table.torrents tr:has(table.torrentname)",
			"container": "table.torrents",
			"next_page": Map{
				"selector":  "a[rel='next'][href], a[href]:contains('下一页'), a[href]:contains('下一頁'), a[href]:contains('Next'), a[href]:contains('next')",
				"attribute": "href",
			},
		},
		"fields": Map{
			"id": Map{
				"selector":  "a[href*='details.php?id=']",
				"attribute": "href",
				"filters": []any{
					Map{
						"name": "re_search",
						"args": []any{
							"\\d+",
							0,
						},
					},
				},
			},
			"title": Map{
				"selector": "table.torrentname tr > td.embedded > a > b",
			},
			"details": Map{
				"selector":  "a[href*='details.php?id=']",
				"attribute": "href",
			},
			"download": Map{
				"selector":  "a[href*='download.php?id=']",
				"attribute": "href",
			},
			"size": Map{
				"selector": "td:nth-child(5):not(.embedded)",
				"filters": []any{
					Map{
						"name": "byte_size",
					},
				},
			},
			"seeders": Map{
				"selector": "td:nth-child(6):not(.embedded)",
			},
			"leechers": Map{
				"selector": "td:nth-child(7):not(.embedded)",
			},
			"downloadvolumefactor": Map{
				"any": []any{
					Map{
						"selector":  "img.pro_free,img.pro_free2up",
						"attribute": "class",
						"filters": []any{
							Map{
								"name": "not_blank",
								"args": []any{
									0,
									"",
								},
							},
						},
					},
					Map{
						"selector":  "img.pro_50pctdown,img.pro_50pctdown2up",
						"attribute": "class",
						"filters": []any{
							Map{
								"name": "not_blank",
								"args": []any{
									0.5,
									"",
								},
							},
						},
					},
					Map{
						"selector":  "img.pro_30pctdown",
						"attribute": "class",
						"filters": []any{
							Map{
								"name": "not_blank",
								"args": []any{
									0.3,
									"",
								},
							},
						},
					},
				},
			},
			"description": Map{
				"selector": "font.subtitle",
				"remove":   "div",
				"filters": []any{
					Map{
						"name": "split",
						"args": []any{
							"\n",
							-1,
						},
					},
				},
			},
			"labels": Map{
				"selector": "font.subtitle > div[style] > div.tag",
				"array":    true,
			},
			"hr_days": Map{
				"selector": "div.circle > div.circle-text",
				"filters": []any{
					Map{
						"name": "re_search",
						"args": []any{
							"\\d+",
							0,
						},
					},
				},
			},
		},
	}
	t.Definition["requests"].(Map)["seeding_statistics"] = Map{
		"parser": "XPath",
		"method": "GET",
		"path":   "/getusertorrentlistajax.php?userid={userId}&type=seeding",
		"headers": Map{
			"Referer": "{domain}/userdetails.php?id={userId}",
		},
		"list": Map{
			"selector": "//tr[position()>1]",
			"next_page": Map{
				"selector": "//a[contains(.//text(), '下一页') or contains(.//text(), '下一頁')]/@href",
			},
		},
		"fields": Map{
			"size": Map{
				"selector": ".//td[3]",
				"filters": []any{
					Map{
						"name": "byte_size",
					},
				},
			},
		},
	}
	search := t.Definition["requests"].(Map)["search"].(Map)
	params := search["params"].(Map)
	for key, value := range map[string]string{"incldead": "0", "inclbookmarked": "0", "search_area": "0", "search_mode": "0", "spstate": "0"} {
		params[key] = []any{value}
	}

	return t
}
