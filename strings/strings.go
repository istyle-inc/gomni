package strings

import (
	"bytes"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"golang.org/x/crypto/sha3"
)

// NVL return sub when s is nil
// Params s *string
// Params sub string
// Return string
func NVL(s *string, sub string) string {
	if s == nil || len(*s) == 0 {
		return sub
	}
	return *s
}

// BuildSHA256String 与えられた複数の文字列を結合し、SHA256ハッシュを取得
// BuildSHA256String combine passed parameters
func BuildSHA256String(src ...interface{}) (string, error) {
	ts, err := Combine("::", src...)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha3.Sum256(bytes.NewBufferString(ts).Bytes())), nil
}

// BuildSHA512String 与えられた複数の文字列を結合し、SHA512ハッシュを取得
// BuildSHA512String combine passed parameters
func BuildSHA512String(src ...interface{}) (string, error) {
	ts, err := Combine("::", src...)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", sha3.Sum512(bytes.NewBufferString(ts).Bytes())), nil
}

// Combine combine passed parameters
// sep separator charactor
// src ...interface{} parameters for combining
func Combine(sep string, src ...interface{}) (str string, err error) {
	arr := make([]string, len(src))
	for i := range src {
		arr[i], err = ToString(src[i])
		if err != nil {
			return "", err
		}
	}
	str = strings.Join(arr, sep)
	return
}

// ErrUnConvertableInterface Error object uncoverted
var ErrUnConvertableInterface = errors.New("This is not unconvertable interface{}")

// ToString will try to convert interface{} to string
// Warning: rune will be judged int32
func ToString(src interface{}) (string, error) {
	switch st := src.(type) {
	case string:
		return src.(string), nil
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
		return fmt.Sprintf("%d", src), nil
	case float32, float64:
		return fmt.Sprintf("%f", src), nil
	case bool:
		return strconv.FormatBool(src.(bool)), nil
	default:
		str, ok := src.(fmt.Stringer)
		if ok {
			return str.String(), nil
		}
		gostr, ok := src.(fmt.GoStringer)
		if ok {
			return gostr.GoString(), nil
		}
		log.Println("Couldn't find any types to be convertable to string : ", st)
	}
	return "", ErrUnConvertableInterface
}

// HasEmpty 引数として与えられたstringが全て空でないかチェックする
func HasEmpty(target ...string) bool {
	for _, s := range target {
		if len(s) <= 0 {
			return true
		}
	}
	return false
}
