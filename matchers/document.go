package matchers

var (
	TypeXls = newType("xls", "application/vnd.ms-excel")
	TypePpt = newType("xls", "application/vnd.ms-powerpoint")
	TypeFdf = newType("pdf", "application/vnd.fdf")
	TypeAi  = newType("pdf", "application/postscript")
	TypeTsa = newType("tsa", "text/tab-separated-values")
	TypeWim = newType("wim", "application/octet-stream")
)

var Document = Map{
	TypeXls: Xls,
	TypePpt: Ppt,
	TypeFdf: Fdf,
	TypeAi:  Ai,
	TypeTsa: Tsa,
	TypeWim: Wim,
}

func Xls(buf []byte) bool {
	return len(buf) > 7 &&
		buf[0] == 0xD0 && buf[1] == 0xCF &&
		buf[2] == 0x11 && buf[3] == 0xE0 &&
		buf[4] == 0xA1 && buf[5] == 0xB1 &&
		buf[6] == 0x1A && buf[7] == 0xE1
}

func Ppt(buf []byte) bool {
	return Xls(buf)
}

func Fdf(buf []byte) bool {
	return Pdf(buf)
}

func Ai(buf []byte) bool {
	return Pdf(buf)
}

func Tsa(buf []byte) bool {
	return len(buf) > 0 &&
		buf[0] == 47
}

func Wim(buf []byte) bool {
	return len(buf) > 4 &&
		buf[0] == 0x4D && buf[1] == 0x53 &&
		buf[2] == 0x57 && buf[3] == 0x49 &&
		buf[4] == 0x4D
}
