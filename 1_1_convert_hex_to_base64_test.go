// http://cryptopals.com/sets/1/challenges/1

package main

import "testing"

func TestHex2base64(t *testing.T) {
  str, err := Hex2base64("49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d")

  if err != nil {
    t.Error(err)
  } else if str != "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t" {
    t.Error("failed http://cryptopals.com/sets/1/challenges/1")
  }
}
