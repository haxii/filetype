package matchers

var (
	TypeZip    = newType("zip", "application/zip")
	TypeTar    = newType("tar", "application/x-tar")
	TypeRar    = newType("rar", "application/x-rar-compressed")
	TypeGz     = newType("gz", "application/gzip")
	TypeBz2    = newType("bz2", "application/x-bzip2")
	Type7z     = newType("7z", "application/x-7z-compressed")
	TypeXz     = newType("xz", "application/x-xz")
	TypePdf    = newType("pdf", "application/pdf")
	TypeExe    = newType("exe", "application/x-msdownload")
	TypeSwf    = newType("swf", "application/x-shockwave-flash")
	TypeRtf    = newType("rtf", "application/rtf")
	TypePs     = newType("ps", "application/postscript")
	TypeSqlite = newType("sqlite", "application/x-sqlite3")
	TypeNes    = newType("nes", "application/x-nintendo-nes-rom")
	TypeCrx    = newType("crx", "application/x-google-chrome-extension")
	TypeCab    = newType("cab", "application/vnd.ms-cab-compressed")
	TypeAr     = newType("ar", "application/x-unix-archive")
	TypeZ      = newType("Z", "application/x-compress")
	TypeLz     = newType("lz", "application/x-lzip")
	TypeRpm    = newType("rpm", "application/x-rpm")
)

var Archive = Map{
	TypeZip:    Zip,
	TypeTar:    Tar,
	TypeRar:    Rar,
	TypeGz:     Gz,
	TypeBz2:    Bz2,
	Type7z:     SevenZ,
	TypeXz:     Xz,
	TypePdf:    Pdf,
	TypeExe:    Exe,
	TypeSwf:    Swf,
	TypeRtf:    Rtf,
	TypePs:     Ps,
	TypeSqlite: Sqlite,
	TypeNes:    Nes,
	TypeCrx:    Crx,
	TypeCab:    Cab,
	TypeAr:     Ar,
	TypeZ:      Z,
	TypeLz:     Lz,
	TypeRpm:    Rpm,
}

func Zip(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x50 && buf[1] == 0x4B &&
		(buf[2] == 0x3 || buf[2] == 0x5 || buf[2] == 0x7) &&
		(buf[3] == 0x4 || buf[3] == 0x6 || buf[3] == 0x8)
}

func Tar(buf []byte) bool {
	return len(buf) > 261 &&
		buf[257] == 0x75 && buf[258] == 0x73 &&
		buf[259] == 0x74 && buf[260] == 0x61 &&
		buf[261] == 0x72
}

func Rar(buf []byte) bool {
	return len(buf) > 6 &&
		buf[0] == 0x52 && buf[1] == 0x61 && buf[2] == 0x72 &&
		buf[3] == 0x21 && buf[4] == 0x1A && buf[5] == 0x7 &&
		(buf[6] == 0x0 || buf[6] == 0x1)
}

func Gz(buf []byte) bool {
	return len(buf) > 2 &&
		buf[0] == 0x1F && buf[1] == 0x8B && buf[2] == 0x8
}

func Bz2(buf []byte) bool {
	return len(buf) > 2 &&
		buf[0] == 0x42 && buf[1] == 0x5A && buf[2] == 0x68
}

func SevenZ(buf []byte) bool {
	return len(buf) > 5 &&
		buf[0] == 0x37 && buf[1] == 0x7A && buf[2] == 0xBC &&
		buf[3] == 0xAF && buf[4] == 0x27 && buf[5] == 0x1C
}

func Pdf(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x25 && buf[1] == 0x50 &&
		buf[2] == 0x44 && buf[3] == 0x46
}

func Exe(buf []byte) bool {
	return len(buf) > 1 &&
		buf[0] == 0x4D && buf[1] == 0x5A
}

func Swf(buf []byte) bool {
	return len(buf) > 2 &&
		(buf[0] == 0x43 || buf[0] == 0x46) &&
		buf[1] == 0x57 && buf[2] == 0x53
}

func Rtf(buf []byte) bool {
	return len(buf) > 4 &&
		buf[0] == 0x7B && buf[1] == 0x5C &&
		buf[2] == 0x72 && buf[3] == 0x74 &&
		buf[4] == 0x66
}

func Nes(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x4E && buf[1] == 0x45 &&
		buf[2] == 0x53 && buf[3] == 0x1A
}

func Crx(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x43 && buf[1] == 0x72 &&
		buf[2] == 0x32 && buf[3] == 0x34
}

func Cab(buf []byte) bool {
	return len(buf) > 3 &&
		((buf[0] == 0x4D && buf[1] == 0x53 && buf[2] == 0x43 && buf[3] == 0x46) ||
			(buf[0] == 0x49 && buf[1] == 0x53 && buf[2] == 0x63 && buf[3] == 0x28))
}

func Ps(buf []byte) bool {
	return len(buf) > 1 &&
		buf[0] == 0x25 && buf[1] == 0x21
}

func Xz(buf []byte) bool {
	return len(buf) > 5 &&
		buf[0] == 0xFD && buf[1] == 0x37 &&
		buf[2] == 0x7A && buf[3] == 0x58 &&
		buf[4] == 0x5A && buf[5] == 0x00
}

func Sqlite(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x53 && buf[1] == 0x51 &&
		buf[2] == 0x4C && buf[3] == 0x69
}

func Ar(buf []byte) bool {
	return len(buf) > 6 &&
		buf[0] == 0x21 && buf[1] == 0x3C &&
		buf[2] == 0x61 && buf[3] == 0x72 &&
		buf[4] == 0x63 && buf[5] == 0x68 &&
		buf[6] == 0x3E
}

func Z(buf []byte) bool {
	return len(buf) > 1 &&
		((buf[0] == 0x1F && buf[1] == 0xA0) ||
			(buf[0] == 0x1F && buf[1] == 0x9D))
}

func Lz(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x4C && buf[1] == 0x5A &&
		buf[2] == 0x49 && buf[3] == 0x50
}

func Rpm(buf []byte) bool {
	return len(buf) > 96 &&
		buf[0] == 0xED && buf[1] == 0xAB &&
		buf[2] == 0xEE && buf[3] == 0xDB
}
