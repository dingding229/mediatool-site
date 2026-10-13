package mediatoolsite

// Refund markers charge normally until their tracker-side refund condition is
// satisfied. The host keeps future/unrecognized promotion markers unknown.
func promotionContract() Map {
	return Map{
		"normal_factor":   1,
		"marker_selector": "[class*='pro_'], [class*='torrent-pro-'], img[alt*='%'], img[alt*='Free'], img[alt*='free'], img[alt*='FREE'], .seed-box-policy-download",
		"known_markers": []any{
			"pro_free", "pro_free2up", "pro_50pctdown", "pro_50pctdown2up",
			"pro_30pctdown", "pro_2up", "pro_twoup",
			"torrent-pro-icon", "torrent-pro-free", "torrent-pro-twoupfree",
			"torrent-pro-halfdown", "torrent-pro-twouphalfdown", "torrent-pro-thirtypercent",
			"torrent-pro-refund-left", "torrent-pro-refund-right",
			"seed-box-policy-status", "seed-box-policy-download",
			"is-blocked", "is-limited", "is-unlimited",
		},
	}
}
func blockedSeedboxFactor() Map {
	return Map{
		"selector":  ".seed-box-policy-download.is-blocked",
		"attribute": "class",
		"filters": []any{
			Map{"name": "not_blank", "args": []any{1, ""}},
		},
	}
}
