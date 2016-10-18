// http://cryptopals.com/sets/1/challenges/2

package main

import (
  "errors"
)

func FixedXor(buff1 []byte, buff2 []byte) ([]byte, error) {
  if len(buff1) != len(buff2) {
    return nil, errors.New("Length of buff1 not equal to length of buff2")
  }

  res := make([]byte, len(buff1))
  for i := 0; i < len(res); i++ {
    res[i] = buff1[i] ^ buff2[i]
  }

  return res, nil
}
