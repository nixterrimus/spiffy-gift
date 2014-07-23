package main

import(
  "image"
  "image/jpeg"
  "fmt"
  "os"

  "github.com/disintegration/gift"
)

func main(){
  if (len(os.Args) != 2){
    fmt.Println("Usage:\tspiffy <file>")
    os.Exit(1)
  }

  srcFileName := os.Args[1]
  srcFile, _ := os.Open(srcFileName)
  src, _, _ := image.Decode(srcFile)


  // 1. Create a new GIFT and add some filters:
  g := gift.New(
    gift.Grayscale(),
    gift.UnsharpMask(1.0, 1.0, 0.0),
  )

  // 2. Create a new image of the corresponding size.
  // dst is a new target image, src is the original image
  dst := image.NewRGBA(g.Bounds(src.Bounds()))

  // 3. Use Draw func to apply the filters to src and store the result in dst:
  g.Draw(dst, src)

  outFileName := srcFileName + ".spiffy.jpg"
  toimg, _ := os.Create(outFileName)
  defer toimg.Close()

  jpeg.Encode(toimg, dst, &jpeg.Options{jpeg.DefaultQuality})
}
