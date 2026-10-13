package mediatoolsite

func ssd() Template {
	t := basic("ssd", "SSD", "https://springsunday.net")
	t.Definition["schema"] = "NexusPHP"
	t.Definition["encoding"] = "UTF-8"
	t.Definition["meta"] = Map{
		"has_hr": true,
	}
	t.Definition["requests"].(Map)["user_basic_info"] = Map{
		"parser": "XPath",
		"method": "GET",
		"path":   "",
		"fields": Map{
			"is_login": Map{
				"any": []any{
					Map{
						"selector": "//a[contains(@href, 'logout')]/text()",
					},
					Map{
						"selector": "//a[contains(@data-url, 'logout')]/text()",
					},
					Map{
						"selector": "//a[contains(@href, 'mybonus')]/text()",
					},
					Map{
						"selector": "//a[contains(@onclick, 'logout')]/text()",
					},
					Map{
						"selector": "//a[contains(@href, 'usercp')]/text()",
					},
					Map{
						"selector": "//form[contains(@action, 'logout')]/text()",
					},
					Map{
						"selector": "//div[@class='user-info-side']/text()",
					},
				},
				"filters": []any{
					Map{
						"name": "not_blank",
					},
				},
			},
			"id": Map{
				"selector":  "*",
				"selection": "html",
				"filters": []any{
					Map{
						"name": "re_search",
						"args": []any{
							"userdetails.php\\?id=(\\d+)",
							1,
						},
					},
				},
			},
			"name": Map{
				"any": []any{
					Map{
						"selector": "//a[contains(@href, 'userdetails')]//b//text()",
					},
					Map{
						"selector": "//a[contains(@href, 'userdetails')]//text()",
					},
					Map{
						"selector": "//a[contains(@href, 'userdetails')]//strong//text()",
					},
				},
			},
			"uploaded": Map{
				"selector":  "*",
				"selection": "html",
				"filters": []any{
					Map{
						"name": "re_search",
						"args": []any{
							"[^总]上[传傳]量?[:：_<>/a-zA-Z-=''\\s#;]+([\\d,.\\s]+(?i)[KMGTPI]*B)",
							1,
						},
					},
					Map{
						"name": "byte_size",
					},
				},
			},
			"downloaded": Map{
				"selector":  "*",
				"selection": "html",
				"filters": []any{
					Map{
						"name": "re_search",
						"args": []any{
							"[^总子影力]下[载載]量?[:：_<>/a-zA-Z-=''\\s#;]+([\\d,.\\s]+(?i)[KMGTPI]*B)",
							1,
						},
					},
					Map{
						"name": "byte_size",
					},
				},
			},
			"ratio": Map{
				"selector":  "*",
				"selection": "html",
				"filters": []any{
					Map{
						"name": "replace",
						"args": []any{
							"无限",
							"",
						},
					},
					Map{
						"name": "replace",
						"args": []any{
							"无穷",
							"",
						},
					},
					Map{
						"name": "replace",
						"args": []any{
							"∞",
							"",
						},
					},
					Map{
						"name": "replace",
						"args": []any{
							"Infinity",
							"",
						},
					},
					Map{
						"name": "replace",
						"args": []any{
							"infinity",
							"",
						},
					},
					Map{
						"name": "replace",
						"args": []any{
							"Inf",
							"",
						},
					},
					Map{
						"name": "replace",
						"args": []any{
							"inf",
							"",
						},
					},
					Map{
						"name": "re_search",
						"args": []any{
							"分享率[:：_<>/a-zA-Z-=\\\"'\\s#;]+([\\d,.\\s]+)",
							1,
						},
					},
				},
			},
			"bonus": Map{
				"selector": "//a[contains(@href,'mybonus')]/@title",
				"filters": []any{
					Map{
						"name": "re_search",
						"args": []any{
							"([\\d,.]+)",
							1,
						},
					},
				},
			},
		},
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
						"selector":  "span.torrent-pro-icon.torrent-pro-free,span.torrent-pro-icon.torrent-pro-twoupfree",
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
						"selector":  "span.torrent-pro-icon.torrent-pro-halfdown,span.torrent-pro-icon.torrent-pro-twouphalfdown",
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
						"selector":  "span.torrent-pro-icon.torrent-pro-thirtypercent",
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
				"selector": "table.torrentname td.embedded > div.torrent-smalldescr",
				"remove":   ".torrent-tag,a,img,font,b",
			},
			"labels": Map{
				"selector": "td.embedded > div.torrent-smalldescr > a > span",
				"array":    true,
			},
		},
	}
	t.Definition["requests"].(Map)["seeding_statistics"] = Map{
		"parser": "XPath",
		"method": "GET",
		"path":   "/getusertorrentlistajax.php?userid={userId}&type=seeding",
		"fields": Map{
			"count": Map{
				"selector": "(.//b)[1]",
				"filters": []any{
					Map{
						"name": "strip",
					},
					Map{
						"name": "re_search",
						"args": []any{
							"^\\d+$",
							0,
						},
					},
				},
			},
			"size": Map{
				"selector": "(.//b)[2]",
				"filters": []any{
					Map{
						"name": "strip",
					},
					Map{
						"name": "re_search",
						"args": []any{
							"(?i)^[0-9][0-9,.]*\\s*[KMGTPE]?i?B$",
							0,
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
			"selector": "//body[count(.//b)=2]",
		},
	}
	t.RSS.Generate = rssGenerator(2)
	return t
}
