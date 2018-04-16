package matchers

var (
	TypeMidi = newType("mid", "audio/midi")
	TypeMp3  = newType("mp3", "audio/mpeg")
	TypeM4a  = newType("m4a", "audio/m4a")
	TypeOgg  = newType("ogg", "audio/ogg")
	TypeFlac = newType("flac", "audio/x-flac")
	TypeWav  = newType("wav", "audio/x-wav")
	TypeAmr  = newType("amr", "audio/amr")
	TypeAac  = newType("aac", "audio/aac")
	TypeWma  = newType("wmv", "audio/x-ms-wma")
	TypeCaf  = newType("caf", "audio/x-caf")
	TypeCda  = newType("cda", "application/x-cdf")
	TypeQcp  = newType("qcp", "audio/vnd.qcelp")
	TypeRa   = newType("ra", "audio/x-realaudio")
	TypeRm   = newType("rm", "audio/x-pn-realaudio")
	TypeRmi  = newType("rmi", "audio/mid")
	TypeVoc  = newType("voc", "audio/voc")
)

var Audio = Map{
	TypeMidi: Midi,
	TypeMp3:  Mp3,
	TypeM4a:  M4a,
	TypeOgg:  Ogg,
	TypeFlac: Flac,
	TypeWav:  Wav,
	TypeAmr:  Amr,

	TypeAac: Aac,
	TypeWma: Wma,
	TypeCaf: Caf,
	TypeCda: Cda,
	TypeQcp: Qcp,
	TypeRa:  Ra,
	TypeRm:  Rm,
	TypeRmi: Rmi,
	TypeVoc: Voc,
}

func Midi(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x4D && buf[1] == 0x54 &&
		buf[2] == 0x68 && buf[3] == 0x64
}

func Mp3(buf []byte) bool {
	return len(buf) > 2 &&
		((buf[0] == 0x49 && buf[1] == 0x44 && buf[2] == 0x33) ||
			(buf[0] == 0xFF && buf[1] == 0xfb))
}

func M4a(buf []byte) bool {
	return len(buf) > 10 &&
		((buf[4] == 0x66 && buf[5] == 0x74 && buf[6] == 0x79 &&
			buf[7] == 0x70 && buf[8] == 0x4D && buf[9] == 0x34 && buf[10] == 0x41) ||
			(buf[0] == 0x4D && buf[1] == 0x34 && buf[2] == 0x41 && buf[3] == 0x20))
}

func Ogg(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x4F && buf[1] == 0x67 &&
		buf[2] == 0x67 && buf[3] == 0x53
}

func Flac(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x66 && buf[1] == 0x4C &&
		buf[2] == 0x61 && buf[3] == 0x43
}

func Wav(buf []byte) bool {
	return len(buf) > 11 &&
		buf[0] == 0x52 && buf[1] == 0x49 &&
		buf[2] == 0x46 && buf[3] == 0x46 &&
		buf[8] == 0x57 && buf[9] == 0x41 &&
		buf[10] == 0x56 && buf[11] == 0x45
}

func Amr(buf []byte) bool {
	return len(buf) > 11 &&
		buf[0] == 0x23 && buf[1] == 0x21 &&
		buf[2] == 0x41 && buf[3] == 0x4D &&
		buf[4] == 0x52 && buf[5] == 0x0A
}

func Aac(buf []byte) bool {
	return len(buf) > 1 &&
		((buf[0] == 0xFF && buf[1] == 0xF1) ||
			(buf[0] == 0xFF && buf[1] == 0xF9))
}

func Wma(buf []byte) bool {
	return len(buf) > 15 &&
		buf[0] == 0x30 && buf[1] == 0x26 &&
		buf[2] == 0xB2 && buf[3] == 0x75 &&
		buf[4] == 0x8E && buf[5] == 0x66 &&
		buf[6] == 0xCF && buf[7] == 0x11 &&
		buf[8] == 0xA6 && buf[9] == 0xD9 &&
		buf[10] == 0x00 && buf[11] == 0xAA &&
		buf[12] == 0x00 && buf[13] == 0x62 &&
		buf[14] == 0xCE && buf[15] == 0x6C
}

func Caf(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x63 && buf[1] == 0x61 &&
		buf[2] == 0x66 && buf[3] == 0x66
}

func Cda(buf []byte) bool {
	return len(buf) > 15 &&
		buf[0] == 0x52 && buf[1] == 0x49 &&
		buf[2] == 0x46 && buf[3] == 0x46 &&
		buf[8] == 0x43 && buf[9] == 0x44 &&
		buf[10] == 0x44 && buf[11] == 0x41 &&
		buf[12] == 0x66 && buf[13] == 0x6D &&
		buf[14] == 0x74 && buf[15] == 0x20
}

func Qcp(buf []byte) bool {
	return len(buf) > 15 &&
		buf[0] == 0x52 && buf[1] == 0x49 &&
		buf[2] == 0x46 && buf[3] == 0x46 &&
		buf[8] == 0x51 && buf[9] == 0x4C &&
		buf[10] == 0x43 && buf[11] == 0x4D &&
		buf[12] == 0x66 && buf[13] == 0x6D &&
		buf[14] == 0x74 && buf[15] == 0x20
}

func Ra(buf []byte) bool {
	return len(buf) > 4 &&
		((buf[0] == 0x2E && buf[1] == 0x52 &&
			buf[2] == 0x4D && buf[3] == 0x46 &&
			buf[4] == 0x00 && buf[5] == 0x00 &&
			buf[6] == 0x00 && buf[7] == 0x12 &&
			buf[8] == 0x00) ||
			(buf[0] == 0x2E && buf[1] == 0x72 &&
				buf[2] == 0x61 && buf[3] == 0xFD &&
				buf[4] == 0x00))
}

func Rm(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x2E && buf[1] == 0x52 &&
		buf[2] == 0x4D && buf[3] == 0x46
}

func Rmi(buf []byte) bool {
	return len(buf) > 15 &&
		buf[0] == 0x52 && buf[1] == 0x49 &&
		buf[2] == 0x46 && buf[3] == 0x46 &&
		buf[8] == 0x52 && buf[9] == 0x4D &&
		buf[10] == 0x49 && buf[11] == 0x44 &&
		buf[12] == 0x64 && buf[13] == 0x61 &&
		buf[14] == 0x74 && buf[15] == 0x61
}
func Voc(buf []byte) bool {
	return len(buf) > 15 &&
		buf[0] == 0x43 && buf[1] == 0x72 &&
		buf[2] == 0x65 && buf[3] == 0x61 &&
		buf[4] == 0x74 && buf[5] == 0x69 &&
		buf[6] == 0x76 && buf[7] == 0x65 &&
		buf[8] == 0x20 && buf[9] == 0x56 &&
		buf[10] == 0x6F && buf[11] == 0x69 &&
		buf[12] == 0x63 && buf[13] == 0x65 &&
		buf[14] == 0x20 && buf[15] == 0x46
}
