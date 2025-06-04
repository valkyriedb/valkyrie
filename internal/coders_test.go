package internal

import (
	"bytes"
	"math"
	"testing"

	decode "github.com/valkyriedb/valkyrie/internal/decoder"
	encode "github.com/valkyriedb/valkyrie/internal/encoder"
)

func TestBool(t *testing.T) {
	bools := []bool{true, false, true, false, true}

	t.Run("multiple bools", func(t *testing.T) {
		var data []byte
		for _, input := range bools {
			data = encode.AppendAny(data, input)
		}
		buf := bytes.NewBuffer(data)
		for _, want := range bools {
			result, err := decode.Bool(buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result != want {
				t.Errorf("got: %t, want: %t", result, want)
			}
		}
	})

	t.Run("single bool", func(t *testing.T) {
		for _, v := range bools {
			data := encode.AppendAny(nil, v)
			buf := bytes.NewBuffer(data)
			result, err := decode.Bool(buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result != v {
				t.Errorf("got: %t, want: %t", result, v)
			}
		}
	})
}

func TestInt64(t *testing.T) {
	int64s := []int64{0, 1, -1, 42, -42, 1000, -1000, math.MaxInt64, math.MinInt64}

	t.Run("multiple int64s", func(t *testing.T) {
		var data []byte
		for _, input := range int64s {
			data = encode.AppendAny(data, input)
		}
		buf := bytes.NewBuffer(data)
		for _, want := range int64s {
			result, err := decode.Int64(buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result != want {
				t.Errorf("got: %d, want: %d", result, want)
			}
		}
	})

	t.Run("single int64", func(t *testing.T) {
		for _, v := range int64s {
			data := encode.AppendAny(nil, v)
			buf := bytes.NewBuffer(data)
			result, err := decode.Int64(buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if result != v {
				t.Errorf("got: %d, want: %d", result, v)
			}
		}
	})
}

func TestFloat64(t *testing.T) {
	float64s := []float64{0.0, 1.0, -1.0, 3.14159, -2.71828, 1e10, -1e10, math.MaxFloat64, math.SmallestNonzeroFloat64, math.Inf(1), math.Inf(-1), math.NaN()}

	t.Run("multiple float64s", func(t *testing.T) {
		var data []byte
		for _, input := range float64s {
			data = encode.AppendAny(data, input)
		}
		buf := bytes.NewBuffer(data)
		for _, want := range float64s {
			result, err := decode.Float64(buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if math.IsNaN(want) {
				if !math.IsNaN(result) {
					t.Errorf("got: %f, want: NaN", result)
				}
			} else if result != want {
				t.Errorf("got: %f, want: %f", result, want)
			}
		}
	})

	t.Run("single float64", func(t *testing.T) {
		for _, v := range float64s {
			data := encode.AppendAny(nil, v)
			buf := bytes.NewBuffer(data)
			result, err := decode.Float64(buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if math.IsNaN(v) {
				if !math.IsNaN(result) {
					t.Errorf("got: %f, want: NaN", result)
				}
			} else if result != v {
				t.Errorf("got: %f, want: %f", result, v)
			}
		}
	})
}

func TestString(t *testing.T) {
	strings := []string{"", "test", "AaAaAaA", "Тест", "Привіт", "こんにちは", "\"\n\t\r"}

	t.Run("multiple strings", func(t *testing.T) {
		var data []byte
		for _, input := range strings {
			data = encode.AppendAny(data, input)
		}
		buf := bytes.NewBuffer(data)
		for _, want := range strings {
			s, err := decode.String(buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if s != want {
				t.Errorf("got: %s, want: %s", s, want)
			}
		}
	})

	t.Run("single string", func(t *testing.T) {
		for _, v := range strings {
			data := encode.AppendAny(nil, v)
			buf := bytes.NewBuffer(data)
			s, err := decode.String(buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if s != v {
				t.Errorf("got: %s, want: %s", s, v)
			}
		}
	})
}

func TestBlob(t *testing.T) {
	blobs := [][]byte{
		{},
		{0},
		{1, 2, 3},
		{255, 0, 128, 64},
		make([]byte, 1000), 
		{0x00, 0xFF, 0x00, 0xFF, 0xAA, 0x55},
	}

	for i := range blobs[4] {
		blobs[4][i] = byte(i % 256)
	}

	t.Run("multiple blobs", func(t *testing.T) {
		var data []byte
		for _, input := range blobs {
			data = encode.AppendAny(data, input)
		}
		buf := bytes.NewBuffer(data)
		for _, want := range blobs {
			result, err := decode.Blob(buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !bytes.Equal(result, want) {
				t.Errorf("got: %v, want: %v", result, want)
			}
		}
	})

	t.Run("single blob", func(t *testing.T) {
		for _, v := range blobs {
			data := encode.AppendAny(nil, v)
			buf := bytes.NewBuffer(data)
			result, err := decode.Blob(buf)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
			}
			if !bytes.Equal(result, v) {
				t.Errorf("got: %v, want: %v", result, v)
			}
		}
	})
}

func TestBoolSlice(t *testing.T) {
	boolSlices := [][]bool{
		{},
		{true},
		{false},
		{true, false},
		{false, true, false, true},
		{true, true, true, false, false, false},
	}

	t.Run("multiple bool slices", func(t *testing.T) {
		var data []byte
		for _, input := range boolSlices {
			data = encode.AppendAny(data, input)
		}
		buf := bytes.NewBuffer(data)
		for _, want := range boolSlices {
			for _, expectedBool := range want {
				result, err := decode.Bool(buf)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != expectedBool {
					t.Errorf("got: %t, want: %t", result, expectedBool)
				}
			}
		}
	})

	t.Run("single bool slice", func(t *testing.T) {
		for _, v := range boolSlices {
			data := encode.AppendAny(nil, v)
			buf := bytes.NewBuffer(data)
			for _, expectedBool := range v {
				result, err := decode.Bool(buf)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != expectedBool {
					t.Errorf("got: %t, want: %t", result, expectedBool)
				}
			}
		}
	})
}

func TestInt64Slice(t *testing.T) {
	int64Slices := [][]int64{
		{},
		{0},
		{1, 2, 3},
		{-1, -2, -3},
		{math.MaxInt64, math.MinInt64},
		{42, -42, 0, 1000, -1000},
	}

	t.Run("multiple int64 slices", func(t *testing.T) {
		var data []byte
		for _, input := range int64Slices {
			data = encode.AppendAny(data, input)
		}
		buf := bytes.NewBuffer(data)
		for _, want := range int64Slices {
			for _, expectedInt := range want {
				result, err := decode.Int64(buf)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != expectedInt {
					t.Errorf("got: %d, want: %d", result, expectedInt)
				}
			}
		}
	})

	t.Run("single int64 slice", func(t *testing.T) {
		for _, v := range int64Slices {
			data := encode.AppendAny(nil, v)
			buf := bytes.NewBuffer(data)
			for _, expectedInt := range v {
				result, err := decode.Int64(buf)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != expectedInt {
					t.Errorf("got: %d, want: %d", result, expectedInt)
				}
			}
		}
	})
}

func TestFloat64Slice(t *testing.T) {
	float64Slices := [][]float64{
		{},
		{0.0},
		{1.0, 2.0, 3.14159},
		{-1.0, -2.71828, -3.0},
		{math.Inf(1), math.Inf(-1), math.NaN()},
		{math.MaxFloat64, math.SmallestNonzeroFloat64},
	}

	t.Run("multiple float64 slices", func(t *testing.T) {
		var data []byte
		for _, input := range float64Slices {
			data = encode.AppendAny(data, input)
		}
		buf := bytes.NewBuffer(data)
		for _, want := range float64Slices {
			for _, expectedFloat := range want {
				result, err := decode.Float64(buf)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if math.IsNaN(expectedFloat) {
					if !math.IsNaN(result) {
						t.Errorf("got: %f, want: NaN", result)
					}
				} else if result != expectedFloat {
					t.Errorf("got: %f, want: %f", result, expectedFloat)
				}
			}
		}
	})

	t.Run("single float64 slice", func(t *testing.T) {
		for _, v := range float64Slices {
			data := encode.AppendAny(nil, v)
			buf := bytes.NewBuffer(data)
			for _, expectedFloat := range v {
				result, err := decode.Float64(buf)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if math.IsNaN(expectedFloat) {
					if !math.IsNaN(result) {
						t.Errorf("got: %f, want: NaN", result)
					}
				} else if result != expectedFloat {
					t.Errorf("got: %f, want: %f", result, expectedFloat)
				}
			}
		}
	})
}

func TestStringSlice(t *testing.T) {
	stringSlices := [][]string{
		{},
		{""},
		{"hello"},
		{"hello", "world"},
		{"", "test", "", "another"},
		{"Тест", "Привіт", "こんにちは"},
		{"\"\n\t\r", "special\x00chars"},
	}

	t.Run("multiple string slices", func(t *testing.T) {
		var data []byte
		for _, input := range stringSlices {
			data = encode.AppendAny(data, input)
		}
		buf := bytes.NewBuffer(data)
		for _, want := range stringSlices {
			for _, expectedString := range want {
				result, err := decode.String(buf)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != expectedString {
					t.Errorf("got: %s, want: %s", result, expectedString)
				}
			}
		}
	})

	t.Run("single string slice", func(t *testing.T) {
		for _, v := range stringSlices {
			data := encode.AppendAny(nil, v)
			buf := bytes.NewBuffer(data)
			for _, expectedString := range v {
				result, err := decode.String(buf)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if result != expectedString {
					t.Errorf("got: %s, want: %s", result, expectedString)
				}
			}
		}
	})
}

func TestBlobSlice(t *testing.T) {
	blobSlices := [][][]byte{
		{},
		{{}},
		{{1, 2, 3}},
		{{}, {1}, {2, 3}},
		{{255, 0}, {128, 64, 32}, {}},
		{{0x00, 0xFF, 0xAA, 0x55}},
	}

	t.Run("multiple blob slices", func(t *testing.T) {
		var data []byte
		for _, input := range blobSlices {
			data = encode.AppendAny(data, input)
		}
		buf := bytes.NewBuffer(data)
		for _, want := range blobSlices {
			for _, expectedBlob := range want {
				result, err := decode.Blob(buf)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if !bytes.Equal(result, expectedBlob) {
					t.Errorf("got: %v, want: %v", result, expectedBlob)
				}
			}
		}
	})

	t.Run("single blob slice", func(t *testing.T) {
		for _, v := range blobSlices {
			data := encode.AppendAny(nil, v)
			buf := bytes.NewBuffer(data)
			for _, expectedBlob := range v {
				result, err := decode.Blob(buf)
				if err != nil {
					t.Errorf("unexpected error: %v", err)
				}
				if !bytes.Equal(result, expectedBlob) {
					t.Errorf("got: %v, want: %v", result, expectedBlob)
				}
			}
		}
	})
}
