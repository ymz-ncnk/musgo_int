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
	(*musgo_int.Int)(&n).UnmarshalMUS(buf)
```

# More info
You can find at [github.com/ymz-ncnk/musgo](https://github.com/ymz-ncnk/musgo).

