package fetch

// Session is remembered between requests
func Session() *Fetch {
	f := New()

	f.config.IsSession = true

	return f
}
