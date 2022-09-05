package fetch

type Progress struct {
	Reporter func(percent int64, current, total int64)
	Total    int64
	Current  int64
}

func (p *Progress) Write(b []byte) (n int, err error) {
	n = len(b)
	p.Current += int64(n)
	p.Reporter(p.Current/p.Total, p.Current, p.Total)
	return
}
