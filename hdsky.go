package mediatoolsite

func hdsky() Template {
	t := basic("hdsky", "HDSKY", "https://hdsky.me")
	t.Definition["schema"] = "NexusPHP"
	t.Definition["encoding"] = "UTF-8"
	return t
}
