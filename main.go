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
  
  image_name := os.Args[1]

  // TODO throw error if argument is missing
  
  reader, err := os.Open(image_name)
  
  if err != nil {
      log.Fatal(err)
  }
  defer reader.Close()

  m, _, err := image.Decode(reader)
  if err != nil {
    log.Fatal(err)
  }

//   TODO resize image so it fits into the terminal

  bounds := m.Bounds()

  asciiCharacters := "`^\",:;Il!i~+_-?][}{1)(|\\/tfjrxnuvczXYUJCLQ0OZmwqpdbkhao*#MW&8%B@$"
  asciiLength := len(asciiCharacters)

  for y := bounds.Min.Y; y < bounds.Max.Y; y++ {
    for x := bounds.Min.X; x < bounds.Max.X; x++ {
      r, g, b, _ := m.At(x, y).RGBA()
      total := uint32(r) + uint32(g) + uint32(b)
      asciiIndex  := uint32(float32(total) / float32(199999) * float32(asciiLength))
      
      for thickness := 0; thickness < 3; thickness++ {
        fmt.Print(string(asciiCharacters[asciiIndex ]))
      }
    }
    fmt.Println("\n")
  }


}