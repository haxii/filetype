package matchers

var (
	TypeJpeg = newType("jpg", "image/jpeg")
	TypePng  = newType("png", "image/png")
	TypeGif  = newType("gif", "image/gif")
	TypeWebp = newType("webp", "image/webp")
	TypeCR2  = newType("cr2", "image/x-canon-cr2")
	TypeTiff = newType("tif", "image/tiff")
	TypeBmp  = newType("bmp", "image/bmp")
	TypeJxr  = newType("jxr", "image/vnd.ms-photo")
	TypePsd  = newType("psd", "image/vnd.adobe.photoshop")
	TypeIco  = newType("ico", "image/x-icon")
	TypeBgp  = newType("bgp", "")
	TypeImg  = newType("img", "")
	TypeJb2  = newType("jb2", "image/x-jb2")
	TypeJfif = newType("jpg", "image/jpeg")
	TypeJpe  = newType("jpg", "image/jpeg")
	TypeJpg  = newType("jpg", "image/jpeg")
	TypeJp2  = newType("jp2", "image/jp2")
	TypeTif  = newType("tif", "image/tiff")
	TypeVhd  = newType("vhd", "application/octet-stream")
	TypeWmf  = newType("wmf", "windows/metafile")
)

var Image = Map{
	TypeJpeg: Jpeg,
	TypePng:  Png,
	TypeGif:  Gif,
	TypeWebp: Webp,
	TypeCR2:  CR2,
	TypeTiff: Tiff,
	TypeBmp:  Bmp,
	TypeJxr:  Jxr,
	TypePsd:  Psd,
	TypeIco:  Ico,

	TypeBgp:  Bgp,
	TypeImg:  Img,
	TypeJb2:  Jb2,
	TypeJfif: Jfif,
	TypeJpe:  Jpe,
	TypeJpg:  Jpg,
	TypeJp2:  Jp2,
	TypeTif:  Tif,
	TypeVhd:  Vhd,
	TypeWmf:  Wmf,
}

func Webp(buf []byte) bool {
	return len(buf) > 11 &&
		buf[8] == 0x57 && buf[9] == 0x45 &&
		buf[10] == 0x42 && buf[11] == 0x50
}

func CR2(buf []byte) bool {
	return len(buf) > 9 &&
		((buf[0] == 0x49 && buf[1] == 0x49 && buf[2] == 0x2A && buf[3] == 0x0) ||
			(buf[0] == 0x4D && buf[1] == 0x4D && buf[2] == 0x0 && buf[3] == 0x2A)) &&
		buf[8] == 0x43 && buf[9] == 0x52
}

func Tiff(buf []byte) bool {
	return len(buf) > 3 &&
		((buf[0] == 0x49 && buf[1] == 0x49 && buf[2] == 0x2A && buf[3] == 0x0) ||
			(buf[0] == 0x4D && buf[1] == 0x4D && buf[2] == 0x0 && buf[3] == 0x2A) ||
			(buf[0] == 0x49 && buf[1] == 0x20 && buf[2] == 0x49))
}

func Jxr(buf []byte) bool {
	return len(buf) > 2 &&
		buf[0] == 0x49 &&
		buf[1] == 0x49 &&
		buf[2] == 0xBC
}

func Psd(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x38 && buf[1] == 0x42 &&
		buf[2] == 0x50 && buf[3] == 0x53
}

func Ico(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x00 && buf[1] == 0x00 &&
		buf[2] == 0x01 && buf[3] == 0x00
}

func Bgp(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x42 && buf[1] == 0x50 &&
		buf[2] == 0x47 && buf[3] == 0xFB
}
func Img(buf []byte) bool {
	return len(buf) > 8 &&
		buf[0] == 0x00 && buf[1] == 0x01 &&
		buf[2] == 0x00 && buf[3] == 0x08 &&
		buf[4] == 0x00 && buf[5] == 0x01 &&
		buf[6] == 0x00 && buf[7] == 0x01 &&
		buf[8] == 0x01
}
func Jb2(buf []byte) bool {
	return len(buf) > 7 &&
		buf[0] == 0x97 && buf[1] == 0x4A &&
		buf[2] == 0x42 && buf[3] == 0x32 &&
		buf[4] == 0x0D && buf[5] == 0x0A &&
		buf[6] == 0x1A && buf[7] == 0x0A
}
func Jfif(buf []byte) bool {
	return Jpeg(buf)
}

func Jpe(buf []byte) bool {
	return Jpeg(buf)
}

func Jp2(buf []byte) bool {
	return len(buf) > 9 &&
		buf[0] == 0x00 && buf[1] == 0x00 &&
		buf[2] == 0x00 && buf[3] == 0x0C &&
		buf[4] == 0x6A && buf[5] == 0x50 &&
		buf[6] == 0x20 && buf[7] == 0x20 &&
		buf[8] == 0x0D && buf[9] == 0x0A
}

func Tif(buf []byte) bool {
	return len(buf) > 2 &&
		((buf[0] == 0x43 && buf[1] == 0x57 &&
			buf[2] == 0x53) ||
			(buf[0] == 0x49 && buf[1] == 0x20 &&
				buf[2] == 0x49) ||
			(buf[0] == 0x49 && buf[1] == 0x49 &&
				buf[2] == 0x2A && buf[3] == 0x00) ||
			(buf[0] == 0x4D && buf[1] == 0x4D &&
				buf[2] == 0x00 && buf[3] == 0x2A))
}

func Vhd(buf []byte) bool {
	return len(buf) > 7 &&
		buf[0] == 0x63 && buf[1] == 0x6F &&
		buf[2] == 0x6E && buf[3] == 0x65 &&
		buf[4] == 0x63 && buf[5] == 0x74 &&
		buf[6] == 0x69 && buf[7] == 0x78
}
func Wmf(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0xD7 && buf[1] == 0xCD &&
		buf[2] == 0xC6 && buf[3] == 0x9A
}

func Jpeg(buf []byte) bool {
	return len(buf) > 2 &&
		buf[0] == 0xFF &&
		buf[1] == 0xD8 &&
		buf[2] == 0xFF
}

//Jpg ...
func Jpg(buf []byte) bool {
	return Jpeg(buf)
}

//Png ...
func Png(buf []byte) bool {
	return len(buf) > 3 &&
		buf[0] == 0x89 && buf[1] == 0x50 &&
		buf[2] == 0x4E && buf[3] == 0x47
}

//Gif ...
func Gif(buf []byte) bool {
	return len(buf) > 2 &&
		buf[0] == 0x47 && buf[1] == 0x49 && buf[2] == 0x46
}

//Bmp ...
func Bmp(buf []byte) bool {
	return len(buf) > 1 &&
		buf[0] == 0x42 &&
		buf[1] == 0x4D
}
