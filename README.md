# Tinify API client for Golang

Golang client for the Tinify API, used for [TinyPNG](https://tinypng.com) and [TinyJPG](https://tinyjpg.com). Tinify compresses or resize your images intelligently. Read more at [http://tinify.com](http://tinify.com).

## Documentation

[Go to the documentation for the HTTP client](https://tinypng.com/developers/reference).

## Installation

Install the API client with `go get`.

```shell
go get -u github.com/gwpp/tinify-go
```

## Usage

- compress
    ```golang
    func TestCompressFromFile(t *testing.T) {
        Tinify.SetKey(Key)
        source, err := Tinify.FromFile("./test.jpg")
        if err != nil {
            t.Error(err)
            return
        }

        err = source.ToFile("./test_output/CompressFromFile.jpg")
        if err != nil {
            t.Error(err)
            return
        }
        t.Log("Compress successful")
    }
    ```

- resize
    ```golang
    func TestResizeFromBuffer(t *testing.T) {
        Tinify.SetKey(Key)

        buf, err := ioutil.ReadFile("./test.jpg")
        if err != nil {
            t.Error(err)
            return
        }
        source, err := Tinify.FromBuffer(buf)
        if err != nil {
            t.Error(err)
            return
        }

        err = source.Resize(&Tinify.ResizeOption{
            Method: Tinify.ResizeMethodScale,
            Width:  200,
        })
        if err != nil {
            t.Error(err)
            return
        }

        err = source.ToFile("./test_output/ResizesFromBuffer.jpg")
        if err != nil {
            t.Error(err)
            return
        }
        t.Log("Resize successful")
    }
    ```

- ***notice:***

    Tinify.ResizeMethod support `scale`, `fit` and `cover`. If used fit or cover, you must provide `both a width and a height`. But used scale, you must provide either a target width or a target height, `but not both`.


- More usage, please see [tinify_test.go](./tinify_test.go)

## Running tests

```shell
cd $GOPATH/src/github.com/gwpp/tinify-go
go test
```

## License

This software is licensed under the MIT License. [View the license](LICENSE).
