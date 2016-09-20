package sign

import (
	"crypto/des"
	"encoding/hex"
	"errors"

	"github.com/cespare/matasano/pkcs7"
)

func getDesKey() []byte {
	key := []byte(_secretKey)[:8] // only use the fist 8 characters of the secret key
	return key
}

func DesEncrypt(data string) (string, error) {
	src := []byte(data)
	key := getDesKey()
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = pkcs7.Pad(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

func DesDecrypt(data string) (string, error) {
	// convert string to []byte
	src, _ := hex.DecodeString(data)

	key := getDesKey()

	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out, _ = pkcs7.Unpad(out)
	return string(out), nil
}
