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
  
// TODO get image from command line
  reader, err := os.Open("/Users/majagunna/Desktop/a.jpg")
  
  if err != nil {
      log.Fatal(err)
  }
  defer reader.Close()

  m, _, err := image.Decode(reader)
  if err != nil {
    log.Fatal(err)
  }

//   TODO resize image 

  bounds := m.Bounds()

  // asciiLength := 70
//   TODO refactor 
  for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
    for x := bounds.Min.X; x < bounds.Max.X; x++ {
      r, g, b, _ := m.At(x, y).RGBA()
      total := uint32(r) + uint32(g) + uint32(b)
      asciiIndex  := uint32(float32(total) / float32(199999) * float32(55))
      
      fmt.Print(string("`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"[asciiIndex ]))
      fmt.Print(string("`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"[asciiIndex ]))
      fmt.Print(string("`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"[asciiIndex ]))
    }
    fmt.Println("\n")
  }


}