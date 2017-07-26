package strings

import "testing"

// Test for NVL
func TestNVL(t *testing.T) {
	// Nil渡した場合にsubを返却すること
	sub := "sub-value"

	if testee := NVL(nil, sub); testee != sub {
		t.Errorf("NVL returned %s not %s", testee, sub)
	}

	// 空文字渡した場合にsubを返却すること
	val := ""

	if testee := NVL(&val, sub); testee != sub {
		t.Errorf("NVL returned %s not %s", testee, sub)
	}

	// Nil以外が渡った場合はその値が返却されること
	val = "valuable-one"
	if testee := NVL(&val, sub); testee == sub {
		t.Errorf("NVL returned %s not %s", sub, val)
	}

}

// BuildSHA256Stringテスト
func TestBuildSHA256String(t *testing.T) {
	// 有効な文字列以外が渡った場合ハッシュ化された文字列が返却されること
	val := "test"
	sub := "36f028580bb02cc8272a9a020f4200e346e276ae664e45ee80745574e2f5ab80"
	if testee, _ := BuildSHA256String(val); testee != sub {
		t.Errorf("BuildSHA256String returned %s not %s", testee, sub)
	}

	// Error
	if testee, err := BuildSHA256String(errorStruct{}); testee != "" || err == nil {
		t.Error("Expected return error but not returned", testee, err)
	}
}

// BuildSHA512Stringテスト
func TestBuildSHA512String(t *testing.T) {
	// 有効な文字列以外が渡った場合ハッシュ化された文字列が返却されること
	val := "test"
	sub := "9ece086e9bac491fac5c1d1046ca11d737b92a2b2ebd93f005d7b710110c0a678288166e7fbe796883a4f2e9b3ca9f484f521d0ce464345cc1aec96779149c14"
	if testee, _ := BuildSHA512String(val); testee != sub {
		t.Errorf("BuildSHA512String returned %s not %s", testee, sub)
	}

	// Error
	if testee, err := BuildSHA512String(errorStruct{}); testee != "" || err == nil {
		t.Error("Expected return error but not returned", testee, err)
	}
}

// HasEmpty テスト
func TestHasEmpty(t *testing.T) {
	// すべて空文字でない場合にfalseが返ること
	if HasEmpty("a", "b", "c") {
		t.Error("HasEmpty should return false when all arguments is not empty")
	}

	// 空文字が含まれる場合にtrueが返ること
	if !HasEmpty("", "b", "c") {
		t.Error("HasEmpty should return true when arguments contains empty")
	}

	// すべて空文字の場合にtrueが返ること
	if !HasEmpty("", "", "") {
		t.Error("HasEmpty should return true when arguments are all empty")
	}

}

func TestToString(t *testing.T) {
	var s string
	var err error
	// string
	s, err = ToString("string")
	if err != nil || s != "string" {
		t.Error("type conversion occured error ", s, err)
	}
	// int
	s, err = ToString(9999)
	if err != nil || s != "9999" {
		t.Error("type conversion occured error ", s, err)
	}
	// int8
	s, err = ToString(int8(16))
	if err != nil || s != "16" {
		t.Error("type conversion occured error ", s, err)
	}
	// int16
	s, err = ToString(int16(128))
	if err != nil || s != "128" {
		t.Error("type conversion occured error ", s, err)
	}
	// int32
	s, err = ToString(int32(65535))
	if err != nil || s != "65535" {
		t.Error("type conversion occured error ", s, err)
	}
	// int64
	s, err = ToString(int64(255486129307))
	if err != nil || s != "255486129307" {
		t.Error("type conversion occured error ", s, err)
	}
	// uint
	s, err = ToString(9999)
	if err != nil || s != "9999" {
		t.Error("type conversion occured error ", s, err)
	}
	// uint8
	s, err = ToString(uint8(16))
	if err != nil || s != "16" {
		t.Error("type conversion occured error ", s, err)
	}
	// uint16
	s, err = ToString(uint16(128))
	if err != nil || s != "128" {
		t.Error("type conversion occured error ", s, err)
	}
	// uint32
	s, err = ToString(uint32(65535))
	if err != nil || s != "65535" {
		t.Error("type conversion occured error ", s, err)
	}
	// uint64
	s, err = ToString(uint64(255486129307))
	if err != nil || s != "255486129307" {
		t.Error("type conversion occured error ", s, err)
	}
	// float32
	s, err = ToString(float32(1.010101))
	if err != nil || s != "1.010101" {
		t.Error("type conversion occured error ", s, err)
	}
	// float64
	s, err = ToString(float64(1.000001))
	if err != nil || s != "1.000001" {
		t.Error("type conversion occured error ", s, err)
	}
	// bool
	s, err = ToString(true)
	if err != nil || s != "true" {
		t.Error("type conversion occured error ", s, err)
	}
	// Strniger
	ts := &testStringer{}
	s, err = ToString(ts)
	if err != nil || s != "test-stringer" {
		t.Error("type conversion occured error ", s, err)
	}
	// GoStrniger
	tgs := &testGoStringer{}
	s, err = ToString(tgs)
	if err != nil || s != "test-gostringer" {
		t.Error("type conversion occured error ", s, err)
	}
	// Error
	s, err = ToString(errorStruct{})
	if err != ErrUnConvertableInterface {
		t.Error("Convert must fail and return expected error", s, err)
	}
}

// testStringer Stringer implemented
type testStringer struct{}

// String Stringer.String()
func (t *testStringer) String() string {
	return "test-stringer"
}

// testStringer Stringer implemented
type testGoStringer struct{}

// String Stringer.String()
func (t *testGoStringer) GoString() string {
	return "test-gostringer"
}

// errorStruct will occur Error in ToString
type errorStruct struct{}
