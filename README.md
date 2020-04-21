
[![CircleCI](https://circleci.com/gh/go-ml-dev/dyl.svg?style=svg)](https://circleci.com/gh/go-ml-dev/dyl)
[![Maintainability](https://api.codeclimate.com/v1/badges/1e480a564c6ba1572581/maintainability)](https://codeclimate.com/github//go-ml-dev/dyl/maintainability)
[![Test Coverage](https://api.codeclimate.com/v1/badges/1e480a564c6ba1572581/test_coverage)](https://codeclimate.com/github/go-ml-dev/dyl/test_coverage)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-ml-dev/dyl)](https://goreportcard.com/report/github.com/go-ml-dev/dyl)
[![License](https://img.shields.io/badge/License-Apache%202.0-blue.svg)](https://opensource.org/licenses/Apache-2.0)

```golang
/*

int function(int);

#define DEFINE_JUMPER(x) \
        void *_dyl_##x = (void*)0; \
        __asm__(".global "#x"\n\t"#x":\n\tmovq _dyl_"#x"(%rip),%rax\n\tjmp *%rax\n")
  
DEFINE_JUMPER(function);

*/
import "C"

import (
	"go-ml.dev/dyl"
	"runtime"
	"unsafe"
)

func init() {
    urlbase := "https://github.com/go-ml-dev/nativelibs/releases/download/files/"
    if runtime.GOOS == "linux" && runtime.GOARCH == "amd64"{
        so := dyl.Load(
            dyl.Cache("go-ml/dyl/libfunction.so"),
            dyl.LzmaExternal(urlbase+"libfunction_lin64.xz"))
    } else if runtime.GOOS == "windows" && runtime.GOARCH == "amd64" {
        so := dyl.Load(
            dyl.Cache("go-ml/dyl/function.dll"),
            dyl.LzmaExternal(urlbase+"libfunction_win64.xz"))
    }
    so.Bind("function",unsafe.Pointer(&C._dyl_function))
}

func main() {
    C.function(0)
}
```
