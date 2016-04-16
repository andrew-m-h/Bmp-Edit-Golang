package bitmap

import (
	"errors"
)


func (bmp *Bitmap) Crop(up, down, left, right int) (int, int, error){
	const (
		bytesPerPixel = 3
	)
	i := 0
	w, h := int(bmp.InfoHead.Width) - left - right, int(bmp.InfoHead.Height) - up - down
	if w <= 0 || h <= 0 {
		return w, h, errors.New("Bounds error")
	}

	bNewWidth := int((float32(bmp.InfoHead.BitCount) * float32(int(bmp.InfoHead.Width) - left - right) + 31.0) / 32.0) * 4

	for y := down; y < int(bmp.InfoHead.Height) - up; y++{
		adjLeft := (y * int(bmp.InfoHead.Width) + left) * bytesPerPixel;
		copy(bmp.Data[i:i + bNewWidth], bmp.Data[adjLeft:adjLeft + bNewWidth])
		i += bNewWidth
	}

	bmp.Resize(w, h)
	return w, h, nil
}