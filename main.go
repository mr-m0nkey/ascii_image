package main


import (
  "fmt"
  "image"
  "os"
  "log"
  _ "image/gif"
  _ "image/png"
  _ "image/jpeg"
)


func main() {
  
  reader, err := os.Open("/Users/majagunna/Desktop/index.png")
  
  if err != nil {
      log.Fatal(err)
  }
  defer reader.Close()

  m, _, err := image.Decode(reader)
  if err != nil {
    log.Fatal(err)
  }

  bounds := m.Bounds()

  // ascii_length := 70
  for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
    for x := bounds.Min.X; x < bounds.Max.X; x++ {
      r, g, b, _ := m.At(x, y).RGBA()
      total := uint32(r) + uint32(g) + uint32(b)
      ascii_index  := uint32(float32(total) / float32(199999) * float32(55))
      
      fmt.Print(string("`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"[ascii_index ]))
      fmt.Print(string("`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"[ascii_index ]))
      fmt.Print(string("`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"[ascii_index ]))
    }
    fmt.Println("\n")
  }


}