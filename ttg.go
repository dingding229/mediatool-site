package mediatoolsite

func ttg() Template {
	t := basic("ttg", "TTG", "https://totheglory.im")
	t.Definition["schema"] = "ttg"
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
				"selector": "//font[contains(text(), '上传量')]/following-sibling::font[1]/a/text()",
				"filters": []any{
					Map{
						"name": "byte_size",
					},
				},
			},
			"downloaded": Map{
				"selector": "//font[contains(text(), '下载量')]/following-sibling::font[1]/a/text()",
				"filters": []any{
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
				"any": []any{
					Map{
						"selector": "//a[contains(@href,'mybonus') or contains(@href,'bonus.php')]/text()",
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
					Map{
						"selector":  "*",
						"selection": "html",
						"filters": []any{
							Map{
								"name": "re_search",
								"args": []any{
									"mybonus.[\\[\\]:：<>/a-zA-Z_\\-=''\\s#;.(使用魔力值豆]+\\s*([\\d,.]+)[<()&\\s\\[]",
									1,
								},
							},
						},
					},
					Map{
						"selector":  "*",
						"selection": "html",
						"filters": []any{
							Map{
								"name": "re_search",
								"args": []any{
									"[魔力值|\\]][\\[\\]:：<>/a-zA-Z_\\-=''\\s#;]+\\s*([\\d,.]+|\\'[\\d,.]+')[<>()&\\s\\[]",
									1,
								},
							},
						},
					},
				},
			},
		},
	}
	t.Definition["requests"].(Map)["search"] = Map{
		"parser": "CssSelector",
		"method": "GET",
		"path":   "/browse.php",
		"params": Map{
			"c": []any{
				"M",
			},
			"search_field": []any{
				"{keyword}",
			},
		},
		"list": Map{
			"selector":  "table#torrent_table tr[id]",
			"container": "table#torrent_table",
			"next_page": Map{
				"selector":  "a[rel='next'][href], a[href]:contains('下一页'), a[href]:contains('下一頁'), a[href]:contains('Next'), a[href]:contains('next')",
				"attribute": "href",
			},
		},
		"fields": Map{
			"id": Map{
				"selector":  "div.name_left > a",
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
				"selector": "div.name_left > a > b",
				"remove":   "span",
				"filters": []any{
					Map{
						"name": "split",
						"args": []any{
							"\n",
							0,
						},
					},
				},
			},
			"details": Map{
				"selector":  "div.name_left > a",
				"attribute": "href",
			},
			"download": Map{
				"selector":  "a.dl_a",
				"attribute": "href",
			},
			"size": Map{
				"selector": "td:nth-child(7)",
				"filters": []any{
					Map{
						"name": "byte_size",
					},
				},
			},
			"seeders": Map{
				"selector": "td:nth-child(9) > b > a >font",
			},
			"leechers": Map{
				"selector": "td:nth-child(9) > b:nth-child(2) > a",
			},
			"downloadvolumefactor": Map{
				"any": []any{
					Map{
						"selector":  "img[alt='free']",
						"attribute": "alt",
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
						"selector":  "img[alt='50%']",
						"attribute": "alt",
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
						"selector":  "img[alt='30%']",
						"attribute": "alt",
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
				"remove": "span,a,img,font,b",
				"filters": []any{
					Map{
						"name": "split",
						"args": []any{
							"\n",
							-1,
						},
					},
				},
				"any": []any{
					Map{
						"selector": "div.name_left > a > b > font > span",
					},
					Map{
						"selector": "div.name_left > a > b > span",
					},
				},
			},
			"labels": Map{
				"selector": "div.name_left > span > b",
				"array":    true,
			},
			"hr_days": Map{
				"selector":  "td:nth-child(2) > div.name_left > img:nth-child(6)",
				"attribute": "title",
				"filters": []any{
					Map{
						"name": "re_search",
						"args": []any{
							"(?i)(?:\\bH\\s*&?\\s*R\\b|\\bhit[ -]*and[ -]*run\\b|保种|做种|考核)[^\\d]{0,80}(\\d+)",
							1,
						},
					},
				},
			},
		},
	}
	t.Definition["requests"].(Map)["seeding_statistics"] = Map{
		"parser": "XPath",
		"method": "GET",
		"path":   "/userdetails.php?id={userId}",
		"headers": Map{
			"Referer": "{domain}/userdetails.php?id={userId}",
		},
		"list": Map{
			"selector": "//td[contains(text(), '当前上传')]/following-sibling::td//table//tr[position()>1]",
			"next_page": Map{
				"selector": "//a[contains(text(), '下一页') or contains(text(), '下一頁')]/@href",
			},
		},
		"fields": Map{
			"size": Map{
				"selector": "/td[4]",
				"filters": []any{
					Map{
						"name": "byte_size",
					},
				},
			},
		},
	}
	t.RSS = RSS{
		Pages: []string{
			"/rss.php",
			"/userdetails.php?id={userId}",
		},
	}
	basic := t.Definition["requests"].(Map)["user_basic_info"].(Map)
	basic["fields"].(Map)["seeding"] = Map{"selector": "//img[@alt='做种中']/following-sibling::font[1]"}
	basic["fields"].(Map)["ratio"] = Map{"selector": "//font[contains(text(),'分享率')]/following-sibling::font[1]"}
	basic["seeding_zero_is_complete"] = true

	t.RSS.Pages = []string{"/rsstools.php"}
	t.RSS.Generate = Map{"path": "/rsstools.php", "method": "GET", "parser": "XPath", "check_all_prefix": "c"}
	return t
}
