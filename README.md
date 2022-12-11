# MusGo_int
Provides serialization of the Golang's `int` type.

# How to use
Simply cast `int` to `musgo_int.Int`:
```go
	var n int = 5
	// Marshal
	buf := make([]byte, musgo_int.Int(n).SizeMUS())
	musgo_int.Int(n).MarshalMUS(buf)
	// Unmarshal
	_, err := (*musgo_int.Int)(&n).UnmarshalMUS(buf)
	if err != nil {
		panic(err)
	}
```

# More info
You can find at [github.com/ymz-ncnk/musgo](https://github.com/ymz-ncnk/musgo).

