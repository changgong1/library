package util

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/url"
	"strconv"
	"strings"
)

type StringInt string

func (s *StringInt) ToInt() int {
	body := string(*s)
	if body == "" {
		return 0
	}
	d, err := strconv.Atoi(body)
	if err != nil {
		fmt.Println(body+": error", err.Error())
	}
	return d
}

func (s *StringInt) ToInt64() int64 {
	body := string(*s)
	d, err := strconv.ParseInt(body, 10, 64)
	if err != nil {
		fmt.Println(body+": error", err.Error())
	}
	return d
}

type StringF string

func (s *StringF) ToFloat() float64 {
	body := string(*s)
	if body == "" {
		return 0
	}
	d, err := strconv.ParseFloat(body, 64)
	if err != nil {
		fmt.Println(body+": error", err.Error())
	}
	return d
}

type StringBool string

func (s *StringBool) True() bool {
	body := string(*s)
	if body == "false" {
		return false
	} else {
		return true
	}
}

type StringZh string // 中文解码

func (s *StringZh) TryToZh() string {
	body := string(*s)
	if i := strings.Index(body, "\\u"); i > -1 {
		sUnicodev := strings.Split(body, "\\u")
		var context string
		for _, v := range sUnicodev {
			if len(v) < 1 {
				continue
			}
			temp, err := strconv.ParseInt(v, 16, 32)
			if err != nil {
				panic(err)
			}
			context += fmt.Sprintf("%c", temp)
		}
		return context
	} else if i := strings.Index(body, "%"); i > -1 {
		context, err := url.QueryUnescape(body)
		if err != nil {
			fmt.Println(body+": error", err.Error())
			return body
		}
		return context
	}
	return body
}

// GetMd5String 生成32位md5字串
func GetMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}

func FaceToBool(face interface{}) bool {
	if f, ok := face.(bool); ok {
		return f
	}
	if f, ok := face.(string); ok {
		if strings.ToLower(f) == "true" {
			return true
		}
	}
	return false
}

func AesEncryptByECB(plaintext, key string) (string, error) {
	if len(key) != 16 && len(key) != 24 && len(key) != 32 {
		return "", errors.New("key length must be 16, 24, or 32 bytes")
	}

	block, err := aes.NewCipher([]byte(key))
	if err != nil {
		return "", err
	}

	paddedText := pkcs7Pad([]byte(plaintext), aes.BlockSize)
	ciphertext := make([]byte, len(paddedText))

	// 初始化向量，固定值
	iv := []byte("1234500000054321")

	mode := cipher.NewCBCEncrypter(block, iv)
	mode.CryptBlocks(ciphertext, paddedText)

	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

// 补码
func pkcs7Pad(data []byte, blocksize int) []byte {
	padding := blocksize - len(data)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(data, padtext...)
}
