# file

Package putridparrot/file implements some file based utility methods/functionality. 

---

* [Install](#install)
* [Examples](#examples)

---

## Install

```sh
go get -u github.com/putridparrot/file
```

## Examples

```go
import fc "github.com/putridparrot/file"

func main() {
  f := fc.NewFileCopy()
  f.Copy("c:\\sourceFileOrFolder", "c:\\destinationFolder")
}
```
