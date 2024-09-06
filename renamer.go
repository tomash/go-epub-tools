package main

import (
	"fmt"

	//"io"
	"os"
	"regexp"

	//"strings"
	"path/filepath"

	"github.com/taylorskalyo/goreader/epub"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func to_ascii_brutal(s string) string {
	// it's stupid but it works

	s = regexp.MustCompile("[ĄÀ�?ÂÃ]").ReplaceAllString(s, "A")
	s = regexp.MustCompile("[âäàãáäå?ăąǎǟǡǻ�?ȃȧẵặ]").ReplaceAllString(s, "a")
	s = regexp.MustCompile("[Ę]").ReplaceAllString(s, "E")
	s = regexp.MustCompile("[ëêéèẽēĕẻȅȇẹȩęḙḛ�?ếễểḕḗệ�?]").ReplaceAllString(s, "e")
	s = regexp.MustCompile("[Ì�?ÎĨ]").ReplaceAllString(s, "I")
	s = regexp.MustCompile("[�?iìíîĩĭïỉ�?ịįȉȋḭɨḯ]").ReplaceAllString(s, "i")
	s = regexp.MustCompile("[ÒÓÔÕÖ]").ReplaceAllString(s, "O")
	s = regexp.MustCompile("[òóôõ�?�ȯö�?őǒ�?�?ơǫ�?ɵøồốỗổȱȫȭ�?�?ṑṓ�?ớỡởợǭộǿ]").ReplaceAllString(s, "o")
	s = regexp.MustCompile("[ÙÚÛŨÜ]").ReplaceAllString(s, "U")
	s = regexp.MustCompile("[ùúûũūŭüůűǔȕȗưụṳųṷṵṹṻǖǜǘǖǚừứữửự]").ReplaceAllString(s, "u")
	s = regexp.MustCompile("[ỳýŷỹȳ�?ỷẙƴỵ]").ReplaceAllString(s, "y")
	s = regexp.MustCompile("[Ń]").ReplaceAllString(s, "N")
	s = regexp.MustCompile("[ñǹń]").ReplaceAllString(s, "n")
	s = regexp.MustCompile("[ÇĆČ]").ReplaceAllString(s, "C")
	s = regexp.MustCompile("[çćč]").ReplaceAllString(s, "c")
	s = regexp.MustCompile("[ß]").ReplaceAllString(s, "ss")
	s = regexp.MustCompile("[œ]").ReplaceAllString(s, "oe")
	s = regexp.MustCompile("[ÆǼǢæ]").ReplaceAllString(s, "ae")
	s = regexp.MustCompile("[ĳ]").ReplaceAllString(s, "ij")
	s = regexp.MustCompile("[Ł]").ReplaceAllString(s, "L")
	s = regexp.MustCompile("[�?ł]").ReplaceAllString(s, "l")
	s = regexp.MustCompile("[ŚŠ]").ReplaceAllString(s, "S")
	s = regexp.MustCompile("[śš]").ReplaceAllString(s, "s")
	s = regexp.MustCompile("[Ț]").ReplaceAllString(s, "T")
	s = regexp.MustCompile("[ț]").ReplaceAllString(s, "t")
	s = regexp.MustCompile("[ŹŻ]").ReplaceAllString(s, "Z")
	s = regexp.MustCompile("[źż]").ReplaceAllString(s, "z")

	return s
}

func main() {
	pathname := os.Args[1]
	// fmt.Printf("pathname: %v\n", pathname)

	rc, err := epub.OpenReader(pathname)
	check(err)
	// defer rc.Close()

	book := rc.Rootfiles[0]

	fmt.Println(("Book title:"))
	fmt.Println(book.Title)
	// fmt.Println(("Language:"))
	// fmt.Println(book.Language)
	// fmt.Println(("Identifier	:"))
	// fmt.Println(book.Identifier)
	fmt.Println(("Book Author (Creator):"))
	fmt.Println(book.Creator)
	// fmt.Println(("Contributor:"))
	// fmt.Println(book.Contributor)
	// fmt.Println(("Publisher:"))
	// fmt.Println(book.Publisher)
	// fmt.Println(("Description:"))
	// fmt.Println(book.Description)

	//fmt.Printf("to_ascii_brutal(zażółć gęślą jaźń): %v\n", to_ascii_brutal("zażółć gęślą jaźń"))

	rc.Close()

	title_latin1 := to_ascii_brutal(book.Title)
	author_latin1 := to_ascii_brutal(book.Creator)
	new_title := fmt.Sprintf("%v - %v", title_latin1, author_latin1)
	new_filename := fmt.Sprintf("%v.epub", new_title)

	if filepath.Base(pathname) != new_filename {
		// renaming time!
		fmt.Printf("renaming to: %v\n", new_filename)
		new_path := filepath.Join(filepath.Dir(pathname), new_filename)
		fmt.Printf("new path: %v\n", new_path)
		err := os.Rename(pathname, new_path)
		check(err)
	} else {
		fmt.Println("Target file already has Title - Author.epub name, skippping")
	}
}
