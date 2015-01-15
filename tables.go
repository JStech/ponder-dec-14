package main

var expanded = [...]uint64 {
  0x0000000000000000, 0x0000000000000001, 0x0000000000000100, 0x0000000000000101,
  0x0000000000010000, 0x0000000000010001, 0x0000000000010100, 0x0000000000010101,
  0x0000000001000000, 0x0000000001000001, 0x0000000001000100, 0x0000000001000101,
  0x0000000001010000, 0x0000000001010001, 0x0000000001010100, 0x0000000001010101,
  0x0000000100000000, 0x0000000100000001, 0x0000000100000100, 0x0000000100000101,
  0x0000000100010000, 0x0000000100010001, 0x0000000100010100, 0x0000000100010101,
  0x0000000101000000, 0x0000000101000001, 0x0000000101000100, 0x0000000101000101,
  0x0000000101010000, 0x0000000101010001, 0x0000000101010100, 0x0000000101010101,
  0x0000010000000000, 0x0000010000000001, 0x0000010000000100, 0x0000010000000101,
  0x0000010000010000, 0x0000010000010001, 0x0000010000010100, 0x0000010000010101,
  0x0000010001000000, 0x0000010001000001, 0x0000010001000100, 0x0000010001000101,
  0x0000010001010000, 0x0000010001010001, 0x0000010001010100, 0x0000010001010101,
  0x0000010100000000, 0x0000010100000001, 0x0000010100000100, 0x0000010100000101,
  0x0000010100010000, 0x0000010100010001, 0x0000010100010100, 0x0000010100010101,
  0x0000010101000000, 0x0000010101000001, 0x0000010101000100, 0x0000010101000101,
  0x0000010101010000, 0x0000010101010001, 0x0000010101010100, 0x0000010101010101,
  0x0001000000000000, 0x0001000000000001, 0x0001000000000100, 0x0001000000000101,
  0x0001000000010000, 0x0001000000010001, 0x0001000000010100, 0x0001000000010101,
  0x0001000001000000, 0x0001000001000001, 0x0001000001000100, 0x0001000001000101,
  0x0001000001010000, 0x0001000001010001, 0x0001000001010100, 0x0001000001010101,
  0x0001000100000000, 0x0001000100000001, 0x0001000100000100, 0x0001000100000101,
  0x0001000100010000, 0x0001000100010001, 0x0001000100010100, 0x0001000100010101,
  0x0001000101000000, 0x0001000101000001, 0x0001000101000100, 0x0001000101000101,
  0x0001000101010000, 0x0001000101010001, 0x0001000101010100, 0x0001000101010101,
  0x0001010000000000, 0x0001010000000001, 0x0001010000000100, 0x0001010000000101,
  0x0001010000010000, 0x0001010000010001, 0x0001010000010100, 0x0001010000010101,
  0x0001010001000000, 0x0001010001000001, 0x0001010001000100, 0x0001010001000101,
  0x0001010001010000, 0x0001010001010001, 0x0001010001010100, 0x0001010001010101,
  0x0001010100000000, 0x0001010100000001, 0x0001010100000100, 0x0001010100000101,
  0x0001010100010000, 0x0001010100010001, 0x0001010100010100, 0x0001010100010101,
  0x0001010101000000, 0x0001010101000001, 0x0001010101000100, 0x0001010101000101,
  0x0001010101010000, 0x0001010101010001, 0x0001010101010100, 0x0001010101010101,
  0x0100000000000000, 0x0100000000000001, 0x0100000000000100, 0x0100000000000101,
  0x0100000000010000, 0x0100000000010001, 0x0100000000010100, 0x0100000000010101,
  0x0100000001000000, 0x0100000001000001, 0x0100000001000100, 0x0100000001000101,
  0x0100000001010000, 0x0100000001010001, 0x0100000001010100, 0x0100000001010101,
  0x0100000100000000, 0x0100000100000001, 0x0100000100000100, 0x0100000100000101,
  0x0100000100010000, 0x0100000100010001, 0x0100000100010100, 0x0100000100010101,
  0x0100000101000000, 0x0100000101000001, 0x0100000101000100, 0x0100000101000101,
  0x0100000101010000, 0x0100000101010001, 0x0100000101010100, 0x0100000101010101,
  0x0100010000000000, 0x0100010000000001, 0x0100010000000100, 0x0100010000000101,
  0x0100010000010000, 0x0100010000010001, 0x0100010000010100, 0x0100010000010101,
  0x0100010001000000, 0x0100010001000001, 0x0100010001000100, 0x0100010001000101,
  0x0100010001010000, 0x0100010001010001, 0x0100010001010100, 0x0100010001010101,
  0x0100010100000000, 0x0100010100000001, 0x0100010100000100, 0x0100010100000101,
  0x0100010100010000, 0x0100010100010001, 0x0100010100010100, 0x0100010100010101,
  0x0100010101000000, 0x0100010101000001, 0x0100010101000100, 0x0100010101000101,
  0x0100010101010000, 0x0100010101010001, 0x0100010101010100, 0x0100010101010101,
  0x0101000000000000, 0x0101000000000001, 0x0101000000000100, 0x0101000000000101,
  0x0101000000010000, 0x0101000000010001, 0x0101000000010100, 0x0101000000010101,
  0x0101000001000000, 0x0101000001000001, 0x0101000001000100, 0x0101000001000101,
  0x0101000001010000, 0x0101000001010001, 0x0101000001010100, 0x0101000001010101,
  0x0101000100000000, 0x0101000100000001, 0x0101000100000100, 0x0101000100000101,
  0x0101000100010000, 0x0101000100010001, 0x0101000100010100, 0x0101000100010101,
  0x0101000101000000, 0x0101000101000001, 0x0101000101000100, 0x0101000101000101,
  0x0101000101010000, 0x0101000101010001, 0x0101000101010100, 0x0101000101010101,
  0x0101010000000000, 0x0101010000000001, 0x0101010000000100, 0x0101010000000101,
  0x0101010000010000, 0x0101010000010001, 0x0101010000010100, 0x0101010000010101,
  0x0101010001000000, 0x0101010001000001, 0x0101010001000100, 0x0101010001000101,
  0x0101010001010000, 0x0101010001010001, 0x0101010001010100, 0x0101010001010101,
  0x0101010100000000, 0x0101010100000001, 0x0101010100000100, 0x0101010100000101,
  0x0101010100010000, 0x0101010100010001, 0x0101010100010100, 0x0101010100010101,
  0x0101010101000000, 0x0101010101000001, 0x0101010101000100, 0x0101010101000101,
  0x0101010101010000, 0x0101010101010001, 0x0101010101010100, 0x0101010101010101, 
}

var popc = [...]uint8 {
  0, 1, 1, 2, 1, 2, 2, 3, 1, 2, 2, 3, 2, 3, 3, 4,
  1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
  1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
  2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
  1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
  2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
  2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
  3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
  1, 2, 2, 3, 2, 3, 3, 4, 2, 3, 3, 4, 3, 4, 4, 5,
  2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
  2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
  3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
  2, 3, 3, 4, 3, 4, 4, 5, 3, 4, 4, 5, 4, 5, 5, 6,
  3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
  3, 4, 4, 5, 4, 5, 5, 6, 4, 5, 5, 6, 5, 6, 6, 7,
  4, 5, 5, 6, 5, 6, 6, 7, 5, 6, 6, 7, 6, 7, 7, 8, 
}
