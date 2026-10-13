package mediatoolsite

func ubits() Template {
	t := basic("ubits", "UBITS", "https://ubits.club")
	t.Definition["schema"] = "NexusPHP"
	t.Definition["encoding"] = "UTF-8"
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
				"any": []any{
					Map{
						"selector":  "a[title][href*='details.php?id=']",
						"attribute": "title",
					},
					Map{
						"selector": "a[href*='details.php?id=']",
					},
				},
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
				"selector": "td:nth-child(2) table.torrentname tr > td:nth-child(1)",
				"remove":   "span,a,img,font,b",
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
				"selector": "table.torrentname tr > td:nth-child(1) > span[title='']",
				"array":    true,
			},
			"hr": Map{
				"selector":  "img.hitandrun",
				"attribute": "class",
				"filters": []any{
					Map{
						"name": "not_blank",
						"args": []any{
							true,
							"",
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
		"fields": Map{
			"count": Map{
				"selector": ".",
				"filters": []any{
					Map{
						"name": "re_search",
						"args": []any{
							"^\\s*(?:共\\s*)?(\\d+)\\s*条记录[\\s,，;；]*总大小[:：]\\s*((?i:[0-9][0-9,.]*\\s*[KMGTPE]?i?B))\\s*$",
							1,
						},
					},
				},
			},
			"size": Map{
				"selector": ".",
				"filters": []any{
					Map{
						"name": "re_search",
						"args": []any{
							"^\\s*(?:共\\s*)?(\\d+)\\s*条记录[\\s,，;；]*总大小[:：]\\s*((?i:[0-9][0-9,.]*\\s*[KMGTPE]?i?B))\\s*$",
							2,
						},
					},
					Map{
						"name": "byte_size",
					},
				},
			},
		},
		"headers": Map{
			"Referer": "{domain}/userdetails.php?id={userId}",
		},
		"summary": Map{
			"selector": "//div[contains(., '条记录') and contains(., '总大小') and not(.//div[contains(., '条记录') and contains(., '总大小')])]",
		},
	}
	t.RSS.Generate = rssGenerator(3)
	return t
}
