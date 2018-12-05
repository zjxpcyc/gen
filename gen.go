package gen

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
)

// MD5 returns md5 checksum of data
func MD5(data string, salt ...string) string {
	dt := []byte(data)
	for _, v := range salt {
		dt = append(dt, []byte(v)...)
	}

	return fmt.Sprintf("%x", md5.Sum(dt))
}

// SHA1 returns SHA1 checksum of data
func SHA1(data string) string {
	return fmt.Sprintf("%x", sha1.Sum([]byte(data)))
}

// SHA256 returns SHA256 checksum of data
func SHA256(data string) string {
	return fmt.Sprintf("%x", sha256.Sum256([]byte(data)))
}

// HmacSHA256 returns HMAC with sha256 hash of data
func HmacSHA256(data, key string) string {
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(data))
	return fmt.Sprintf("%x", h.Sum(nil))
}

// Base64 return base64 encoding of data
func Base64(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

// Base64Decode returns the bytes represented by the base64 string data
func Base64Decode(data string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(data)
}

// Map2XML can transform map[string]string to xml string
func Map2XML(data map[string]string) ([]byte, error) {
	return xml.Marshal(XMLMap(data))
}

// XML2Map can transform xml string to map[string]string
func XML2Map(data []byte) (map[string]string, error) {
	var m XMLMap
	err := xml.Unmarshal(data, &m)
	return map[string]string(m), err
}

// CBC7Decrypt AES-128-CBC PKCS#7
// 参考: https://github.com/medivhzhan/weapp/blob/master/util/crypto.go
// https://golang.google.cn/pkg/crypto/cipher/#example_NewCBCDecrypter
func CBC7Decrypt(data, key, iv []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	dtLen := len(data)
	if dtLen < aes.BlockSize {
		return nil, errors.New("data is too short")
	}

	mode := cipher.NewCBCDecrypter(block, iv[:aes.BlockSize])

	dist := make([]byte, dtLen)
	mode.CryptBlocks(dist, data)

	return pkcs7UnPadding(dist), nil
}

func pkcs7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func pkcs7UnPadding(dt []byte) []byte {
	length := len(dt)
	unpadding := int(dt[length-1])
	return dt[:(length - unpadding)]
}
