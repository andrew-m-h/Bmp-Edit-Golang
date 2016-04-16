package bitmap

import "sync"

func memset(arr []byte, start, count int, value byte){
	for i := 0; i < count; i++ {
		arr[start+i] = value
	}
}

func (bmp *Bitmap) Threshold(threshold float32){
	const (
		bytePerPixel int = 3
		black byte = 255
		white byte = 0
	)
	var wg sync.WaitGroup
	wg.Add(int(bmp.InfoHead.Height))
	for y := 0; y < int(bmp.InfoHead.Height); y++ {
		go func(curY int) {
			defer wg.Done()
			for x := 0; x < int(bmp.InfoHead.Width) * bytePerPixel; x += 3 {

				var avg float32 = 0;
				for j := 0; j < bytePerPixel; j++ {
					avg += float32(bmp.Data[x + bytePerPixel * int(bmp.InfoHead.Width) * curY + j])
				}
				avg /= float32(bytePerPixel * int(black))

				if avg > threshold {
					memset(bmp.Data, x + bytePerPixel * int(bmp.InfoHead.Width) * curY, bytePerPixel, black)
				} else {
					memset(bmp.Data, x + bytePerPixel * int(bmp.InfoHead.Width) * curY, bytePerPixel, white)
				}
			}
		}(y)
	}
	wg.Wait()
}

