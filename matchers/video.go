package matchers

var (
	TypeMp4  = newType("mp4", "video/mp4")
	TypeM4v  = newType("m4v", "video/x-m4v")
	TypeMkv  = newType("mkv", "video/x-matroska")
	TypeWebm = newType("webm", "video/webm")
	TypeMov  = newType("mov", "video/quicktime")
	TypeAvi  = newType("avi", "video/x-msvideo")
	TypeWmv  = newType("wmv", "video/x-ms-wmv")
	TypeMpeg = newType("mpeg", "video/mpeg")
	TypeVob  = newType("mpg", "video/x-ms-vob")
	TypeFlv  = newType("flv", "video/x-flv")
	TypeRmvb = newType("rmvb", "application/vnd.rn-realmedia")
	TypeVcd  = newType("vcd", "application/x-cdlink")
	TypeVmd  = newType("vmd", "chemical/x-vmd")
	TypeWpl  = newType("wpl", "application/vnd.ms-wpl")
	TypeDvr  = newType("dvr", "video/x-ms-dvr")
	TypeIvr  = newType("ivr", "video/vnd.rn-realvideo")
	Type3gg  = newType("3gg", "video/3gpp")
	Type3gp  = newType("3gp", "video/3gpp")
	Type3g2  = newType("3g2", "video/3gpp")
	TypeAsf  = newType("asf", "video/x-ms-asf")
)

var Video = Map{
	TypeMp4:  Mp4,
	TypeM4v:  M4v,
	TypeMkv:  Mkv,
	TypeWebm: Webm,
	TypeMov:  Mov,
	TypeAvi:  Avi,
	TypeWmv:  Wmv,
	TypeMpeg: Mpeg,
	TypeFlv:  Flv,

	TypeVob:  Vob,
	TypeRmvb: Rmvb,
	TypeVcd:  Vcd,
	TypeVmd:  Vmd,
	TypeWpl:  Wpl,
	TypeDvr:  Dvr,
	TypeIvr:  Ivr,
	Type3gg:  V3gg,
	Type3g2:  V3g2,
	Type3gp:  V3gp,
	TypeAsf:  Asf,
}

//M4v ...
func M4v(buf []byte) bool {
	return len(buf) > 11 &&
		((buf[0] == 0x0 && buf[1] == 0x0 &&
			buf[2] == 0x0 && buf[3] == 0x20 &&
			buf[4] == 0x66 && buf[5] == 0x74 &&
			buf[6] == 0x79 && buf[7] == 0x70 &&
			buf[8] == 0x4D && buf[9] == 0x34 &&
			buf[10] == 0x56 && buf[11] == 0x20) ||
			(buf[4] == 0x66 && buf[5] == 0x74 &&
				buf[6] == 0x79 && buf[7] == 0x70 &&
				buf[8] == 0x6D && buf[9] == 0x70 &&
				buf[10] == 0x34 && buf[11] == 0x32) ||
			(buf[4] == 0x66 && buf[5] == 0x74 &&
				buf[6] == 0x79 && buf[7] == 0x70 &&
				buf[8] == 0x4d && buf[9] == 0x34 &&
				buf[10] == 0x56 && buf[11] == 0x20))
}

//Mkv ...
func Mkv(buf []byte) bool {
	return (len(buf) > 15 &&
		buf[0] == 0x1A && buf[1] == 0x45 &&
		buf[2] == 0xDF && buf[3] == 0xA3 &&
		buf[4] == 0x93 && buf[5] == 0x42 &&
		buf[6] == 0x82 && buf[7] == 0x88 &&
		buf[8] == 0x6D && buf[9] == 0x61 &&
		buf[10] == 0x74 && buf[11] == 0x72 &&
		buf[12] == 0x6F && buf[13] == 0x73 &&
		buf[14] == 0x6B && buf[15] == 0x61) ||
		(len(buf) > 38 &&
			buf[31] == 0x6D && buf[32] == 0x61 &&
			buf[33] == 0x74 && buf[34] == 0x72 &&
			buf[35] == 0x6f && buf[36] == 0x73 &&
			buf[37] == 0x6B && buf[38] == 0x61)
}

//Mov ...
func Mov(buf []byte) bool {
	return len(buf) > 7 &&
		buf[0] == 0x0 && buf[1] == 0x0 &&
		buf[2] == 0x0 && buf[3] == 0x14 &&
		buf[4] == 0x66 && buf[5] == 0x74 &&
		buf[6] == 0x79 && buf[7] == 0x70
}

//Wmv ...
func Wmv(buf []byte) bool {
	return len(buf) > 9 &&
		buf[0] == 0x30 && buf[1] == 0x26 &&
		buf[2] == 0xB2 && buf[3] == 0x75 &&
		buf[4] == 0x8E && buf[5] == 0x66 &&
		buf[6] == 0xCF && buf[7] == 0x11 &&
		buf[8] == 0xA6 && buf[9] == 0xD9
}

//Mpeg ...
func Mpeg(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x0 && buf[1] == 0x0 &&
		buf[2] == 0x1 && buf[3] >= 0xb0 &&
		buf[3] <= 0xbf
}

//Flv ...
func Flv(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x46 && buf[1] == 0x4C &&
		buf[2] == 0x56 && buf[3] == 0x01
}

//Mp4 ...
func Mp4(buf []byte) bool {
	return len(buf) > 24 &&
		((buf[0] == 0x00 && buf[1] == 0x00 && buf[2] == 0x00 && buf[3] == 0x14 &&
			buf[4] == 0x66 && buf[5] == 0x74 && buf[6] == 0x79 && buf[7] == 0x70 &&
			buf[8] == 0x69 && buf[9] == 0x73 && buf[10] == 0x6F && buf[11] == 0x6D) ||
			(buf[0] == 0x00 && buf[1] == 0x00 && buf[2] == 0x00 && buf[3] == 0x18 &&
				buf[4] == 0x66 && buf[5] == 0x74 && buf[6] == 0x79 && buf[7] == 0x70 &&
				buf[8] == 0x33 && buf[9] == 0x67 && buf[10] == 0x70 && buf[11] == 0x35) ||
			(buf[0] == 0x00 && buf[1] == 0x00 && buf[2] == 0x00 && buf[3] == 0x1C &&
				buf[4] == 0x66 && buf[5] == 0x74 && buf[6] == 0x79 && buf[7] == 0x70 &&
				buf[8] == 0x4D && buf[9] == 0x53 && buf[10] == 0x4E && buf[11] == 0x56 &&
				buf[12] == 0x01 && buf[13] == 0x29 && buf[14] == 0x00 && buf[15] == 0x46 &&
				buf[16] == 0x4D && buf[17] == 0x53 && buf[18] == 0x4E && buf[19] == 0x56 &&
				buf[20] == 0x6D && buf[21] == 0x70 && buf[22] == 0x34 && buf[23] == 0x32) ||
			(buf[4] == 0x66 && buf[5] == 0x74 && buf[6] == 0x79 && buf[7] == 0x70 &&
				buf[8] == 0x33 && buf[9] == 0x67 && buf[10] == 0x70 && buf[11] == 0x35) ||
			(buf[4] == 0x66 && buf[5] == 0x74 && buf[6] == 0x79 && buf[7] == 0x70 &&
				buf[8] == 0x4D && buf[9] == 0x53 && buf[10] == 0x4E && buf[11] == 0x56) ||
			(buf[4] == 0x66 && buf[5] == 0x74 && buf[6] == 0x79 && buf[7] == 0x70 &&
				buf[8] == 0x69 && buf[9] == 0x73 && buf[10] == 0x6F && buf[11] == 0x6D))
}

func Webm(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x1A && buf[1] == 0x45 &&
		buf[2] == 0xDF && buf[3] == 0xA3
}

func Avi(buf []byte) bool {
	return len(buf) > 10 &&
		buf[0] == 0x52 && buf[1] == 0x49 &&
		buf[2] == 0x46 && buf[3] == 0x46 &&
		buf[8] == 0x41 && buf[9] == 0x56 &&
		buf[10] == 0x49
}

func Vob(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x00 && buf[1] == 0x00 &&
		buf[2] == 0x01 && buf[3] == 0xBA
}

func Rmvb(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x2E && buf[1] == 0x52 &&
		buf[2] == 0x4D && buf[3] == 0x46
}

func Vcd(buf []byte) bool {
	return len(buf) > 15 &&
		buf[0] == 0x45 && buf[1] == 0x4E &&
		buf[2] == 0x54 && buf[3] == 0x52 &&
		buf[4] == 0x59 && buf[5] == 0x56 &&
		buf[6] == 0x43 && buf[7] == 0x44 &&
		buf[8] == 0x02 && buf[9] == 0x00 &&
		buf[10] == 0x00 && buf[11] == 0x01 &&
		buf[12] == 0x02 && buf[13] == 0x00 &&
		buf[14] == 0x18 && buf[15] == 0x58
}

func Vmd(buf []byte) bool {
	return len(buf) > 4 &&
		buf[0] == 0x5B && buf[1] == 0x56 &&
		buf[2] == 0x4D && buf[3] == 0x44 &&
		buf[4] == 0x5D
}
func Wpl(buf []byte) bool {
	return len(buf) > 117 &&
		buf[84] == 0x4D && buf[85] == 0x69 &&
		buf[86] == 0x63 && buf[87] == 0x72 &&
		buf[88] == 0x6F && buf[89] == 0x73 &&
		buf[90] == 0x6F && buf[91] == 0x66 &&
		buf[92] == 0x74 && buf[93] == 0x20 &&
		buf[94] == 0x57 && buf[95] == 0x69 &&
		buf[96] == 0x6E && buf[97] == 0x64 &&
		buf[98] == 0x6F && buf[99] == 0x77 &&
		buf[100] == 0x73 && buf[101] == 0x20 &&
		buf[102] == 0x4D && buf[103] == 0x65 &&
		buf[104] == 0x64 && buf[105] == 0x69 &&
		buf[106] == 0x61 && buf[107] == 0x20 &&
		buf[108] == 0x50 && buf[109] == 0x6C &&
		buf[110] == 0x61 && buf[111] == 0x79 &&
		buf[112] == 0x65 && buf[113] == 0x72 &&
		buf[114] == 0x20 && buf[115] == 0x2D &&
		buf[116] == 0x2D && buf[117] == 0x20
}
func Dvr(buf []byte) bool {
	return len(buf) > 2 &&
		buf[0] == 0x44 && buf[1] == 0x56 &&
		buf[2] == 0x44
}
func Ivr(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x2E && buf[1] == 0x52 &&
		buf[2] == 0x45 && buf[3] == 0x43
}

func V3gg(buf []byte) bool {
	return len(buf) > 10 &&
		buf[4] == 0x66 && buf[5] == 0x74 &&
		buf[6] == 0x79 && buf[7] == 0x70 &&
		buf[8] == 0x33 && buf[9] == 0x67 &&
		buf[10] == 0x70
}
func V3gp(buf []byte) bool {
	return V3gg(buf)
}

func V3g2(buf []byte) bool {
	return V3gg(buf)
}

func Asf(buf []byte) bool {
	return Wmv(buf)
}
