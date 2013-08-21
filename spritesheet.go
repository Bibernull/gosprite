package gosprite

import (
	"io/ioutil"
	"image"
	"image/png"
	"image/draw"
	"os"
	"fmt"
	"sort"
	"path"
	"strings"
    "strconv"
)

type ImageSlice []*Image

func (p ImageSlice) Less(i, j int) bool { return p[i].height > p[j].height }
func (p ImageSlice) Len() int { return len(p)}
func (p ImageSlice) Swap(i, j int) { p[i], p[j] = p[j], p[i] }

type Image struct{
	path string
	width int
	height int
	x int
	y int
}

type Spritesheet struct {
	selector string
    Source_dir string
    Sprite_path string
    rel_Sprite_path string
    css_path string
    sprite_name string
    image_paths []string
    sheet RectangleSheet
    images []*Image

}

func readImagesMetadata(files []string) (images []*Image) {

    images = make([]*Image, len(files))

    for key, v := range files {
    	file, _ := os.Open(v)
        img, err := png.Decode(file)
        file.Close()

        if err != nil{
        	fmt.Println(err)
        } else {
        	img_tmp := &Image{path: v, width: img.Bounds().Max.X, height: img.Bounds().Max.Y }
        	images[key] = img_tmp
        }

    }
    sort.Sort(ImageSlice(images))



    return
}

func Sprite_dir(source_dir string, sprite_path string) []*Image{

	var files []string;
   	file_list, _ := ioutil.ReadDir(source_dir)
	for _, f := range file_list {
        files = append(files, source_dir + "/" + f.Name())
	}

	images := readImagesMetadata(files)

	make_sprite(images, sprite_path)

	return images
}

func Sprite_images(image_paths []string, sprite_path string) []*Image{

	images := readImagesMetadata(image_paths)

	make_sprite(images, sprite_path)

	return images
}

func make_sprite(images []*Image, sprite_path string){

    sheet :=  RectangleSheet{}
    sheet.calculate(images)

    write_images(images, sheet.width, sheet.height, sprite_path)

}

func write_images(images []*Image, width int, height int, sprite_path string) {

	m := image.NewRGBA(image.Rect(0, 0, width, height))


	for _, v := range images {
		file, _ := os.Open(v.path)
        img, _ := png.Decode(file)
        file.Close()

        r := image.Rectangle{image.Point{v.x, v.y}, image.Point{v.x + v.width, v.y + v.height}}

	    draw.Draw(m, r, img, image.Point{0,0}, draw.Src)

	}

	toimg, _ := os.Create(sprite_path)
	//images = nil
	defer toimg.Close()

	png.Encode(toimg, m)
	fmt.Println("Saved " + sprite_path);

}

func Css(images []*Image, sprite_name, sprite_path string) string{

	css_str := ""
	//sprite_name := path.Base(sprite_path[0:strings.LastIndex(sprite_path, ".")])

	for _, v := range images{
		name := path.Base(v.path[: strings.LastIndex(v.path, ".")])

		css_str +=  "." + sprite_name + "." + name + " { \n";
		css_str += "\twidth: " + strconv.Itoa(v.width) + "px;\n";
		css_str += "\theight: " + strconv.Itoa(v.height) + "px;\n";
		css_str += "\tbackground-image: url(" + sprite_path + ");\n";
		css_str += "\tbackground-position: -" + strconv.Itoa(v.x) + "px -" + strconv.Itoa(v.y) + "px;\n";
		css_str += "}\n";
	}

	return css_str


}