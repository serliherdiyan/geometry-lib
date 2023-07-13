package shape

type Retangle struct {
	Width  float32
	Length float32
}

func (retangle *Retangle) Area() float32 {
	return retangle.Width * retangle.Length
}

func (retangle *Retangle) Perimeter() float32 {
	return 2 * (retangle.Width + retangle.Length)
}

// Push projeck ke github

// ubah file.mod kode module https://github.com/serliherdiyan/geometry-lib
// git init
// git add .
// git commit -m 'message'
// git remote add origin https://github.com/serliherdiyan/geometri-lib.git
// git push origin master

// Membuat versi
// git tag v1.0.0
// git push origin v1.0.0
