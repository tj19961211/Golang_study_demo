package struct_func_and_interface


func (r *Rectangle) Perimeter() float64 {
	return 2*(r.Width + r.Height)
}

func Area(width float64, height float64) float64 {
	return width * height
}

type Rectangle struct {
	Width float64
	Height float64
}

