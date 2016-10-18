// http://cryptopals.com/sets/1/challenges/1

package main

import (
  "encoding/base64"
  "encoding/hex"
)

func Hex2base64(hexStr string) (string, error) {
  bytes, err := hex.DecodeString(hexStr)
  if err != nil {
    return "", err
  }

  encoding := base64.NewEncoding("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/")
  return encoding.EncodeToString(bytes), nil
}
