# Go Epub Tools
Tools for simple manipulations of epub files. Right now it's only renamer, so that epub filename has the pattern "Title - Author.epub", with both title and author read from epub's metadata. This is for more convenient importing into Kindle library, as Kindle tools do not fetch those fields properly from metadata and suggest using filename as the book title.

## Rationale
More on the background in my blogpost about a Ruby script that is the direct prototype of this: [Quick Hack: ePub renamer](https://tomash.wrug.eu/blog/2024/08/13/epub-renamer/).

It's also a way for me to learn basic Golang.

## Usage

```
  go build -o bin/renamer
  ./bin/renamer /path/to/file.epub
```

## License

MIT