package main

import (
  "fmt"
)

func main() {
  const rows uint = 4
  const cols uint = 6

  var colmask uint8
  colmask = 0xff>>(8-(cols%8))

  results := make(map[[rows+cols]uint8]int)
  var res [rows+cols]uint8
  var mat uint64;
  var i uint
  for mat=0; mat<1<<(rows*cols); mat++ {
    var r uint
    for i=0; i<rows+cols; i++ {
      res[i] = 0
    }
    for r=0; r<rows; r++ {
      var c uint
      for c=0; c<cols; c+=8 {
        var b uint8
        b = uint8((mat>>(r*cols+c)) & 0xff)
        if c+8>cols {
          b &= colmask
        }
        res[r] += popc[b]
        //TODO: make this a pocketed 64 bit addition
        for i=0; i<8 && c+i<cols; i++ {
          res[rows + c + i] += uint8(expanded[b]>>(8*i))
        }
      }
    }
    results[res] += 1
  }

  for res = range(results) {
    if results[res] == 29 {
      sum := 0
      for r := range res {
        sum += int(res[r])
      }
      fmt.Println(res, sum)
    }
  }
}
