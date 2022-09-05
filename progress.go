package fetch

// inspired by:
//	https://github.com/schollz/progressbar/blob/master/progressbar.go
//  https://stackoverflow.com/questions/26050380/go-tracking-post-request-progress

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
