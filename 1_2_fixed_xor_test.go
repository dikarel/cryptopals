// http://cryptopals.com/sets/1/challenges/2

package main

import (
  "encoding/hex"
  "testing"
)

func TestFixedXor(t *testing.T) {
  buff1, err1 := hex.DecodeString("1c0111001f010100061a024b53535009181c")
  if err1 != nil {
    t.Error(err1)
  }

  buff2, err2 := hex.DecodeString("686974207468652062756c6c277320657965")
  if err2 != nil {
    t.Error(err2)
  }

  res, err3 := FixedXor(buff1, buff2)
  if err3 != nil {
    t.Error(err3)
  } else if hex.EncodeToString(res) != "746865206b696420646f6e277420706c6179" {
    t.Error("failed http://cryptopals.com/sets/1/challenges/2")
  }
}
