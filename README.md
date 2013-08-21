# gosprite

This is a rewrite(hardly) of `spritesheetjs`.
I wanted to check out golang and aswell compare it's speed to nodejs in making spritesheets.
Did not notice any difference, and tested with 60 simple images so far, but i expected more.
`image/draw` package and it's `draw` func doesn't seem to be slower(nor faster) than the manual drawing to image from `spritesheetjs`.

# Example

```go
   package main

    import (
        "github.com/bibernull/gosprite"
        "os"
        "fmt"
    )

    func main() {

        images1 := gosprite.Sprite_dir("images_dir", "new.png")
        images2 := gosprite.Sprite_images([]string{"images_dir/image1.png", "images_dir/image2.png"}, "new2.png")

        css1 := gosprite.Css(images1, "sprite1", "new.png")
        css2 := gosprite.Css(images2, "sprite2", "new2.png")

        css_file, _ := os.Create("sprite1.css")
        css_file.WriteString(css1)
        css_file.Close()
        fmt.Println("Saved sprite1.css")


        css_file, _ = os.Create("sprite2.css")
        css_file.WriteString(css2)
        css_file.Close()
        fmt.Println("Saved sprite2.css")
    }
```
