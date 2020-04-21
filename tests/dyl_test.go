package tests

import (
	"go-ml.dev/pkg/dyl"
	"gotest.tools/assert"
	"testing"
)

const baseurl = "https://github.com/go-ml-dev/nativelibs/releases/download/files/"
const localLibSoName = "/tmp/go-ml-dyl-test/" + libSoName
const externalLibSoLzma = baseurl + libSoLzma
const externalLibSoGzip = baseurl + libSoGzip
const externalLibSo = baseurl + libSoName

func init() {
	dyl.Custom(localLibSoName).Preload(
		dyl.LzmaExternal(externalLibSoLzma))
}

func Test_LoadCustom(t *testing.T) {
	so := dyl.Load(
		dyl.OnError(func(err error) {
			assert.NilError(t, err)
		}),
		dyl.Custom(localLibSoName))
	assert.Assert(t, so.Ok())
	so.Bind("function", functionPtr())
	assert.Assert(t, function(1) == 2)
}

func Test_LoadLzmaExternal(t *testing.T) {
	err := dyl.Cached("go-ml/dyl/.dyl-test/loadlzmaexternal.so").Remove()
	assert.NilError(t, err)
	so := dyl.Load(
		dyl.OnError(func(err error) {
			assert.NilError(t, err)
		}),
		dyl.Cached("go-ml/dyl/.dyl-test/loadlzmaexternal.so"),
		dyl.LzmaExternal(externalLibSoLzma))
	assert.Assert(t, so.Ok())
	*(*uintptr)(functionPtr()) = 0
	so.Bind("function", functionPtr())
	assert.Assert(t, function(1) == 2)
}

func Test_LoadGzipExternal(t *testing.T) {
	err := dyl.Cached("go-ml/dyl/.dyl-test/loadgzipexternal" + SoExt).Remove()
	assert.NilError(t, err)
	so := dyl.Load(
		dyl.OnError(func(err error) {
			assert.NilError(t, err)
		}),
		dyl.Cached("go-ml/dyl/.dyl-test/loadgzipexternal"+SoExt),
		dyl.GzipExternal(externalLibSoGzip))
	assert.Assert(t, so.Ok())
	*(*uintptr)(functionPtr()) = 0
	so.Bind("function", functionPtr())
	assert.Assert(t, function(1) == 2)
}

func Test_LoadUncompressedExternal(t *testing.T) {
	err := dyl.Cached("go-ml/dyl/.dyl-test/loadexternal" + SoExt).Remove()
	assert.NilError(t, err)
	so := dyl.Load(
		dyl.OnError(func(err error) {
			assert.NilError(t, err)
		}),
		dyl.Cached("go-ml/dyl/.dyl-test/loadexternal"+SoExt),
		dyl.External(externalLibSo))
	assert.Assert(t, so.Ok())
	*(*uintptr)(functionPtr()) = 0
	so.Bind("function", functionPtr())
	assert.Assert(t, function(1) == 2)
}

func Test_LoadCached(t *testing.T) {
	err := dyl.Cached("go-ml/dyl/" + libSoName).Remove()
	assert.NilError(t, err)
	dyl.Cached("go-ml/dyl/"+libSoName).Preload(
		dyl.LzmaExternal(externalLibSoLzma),
		dyl.OnError(func(err error) {
			assert.NilError(t, err)
		}))
	so := dyl.Load(
		dyl.OnError(func(err error) {
			assert.NilError(t, err)
		}),
		dyl.Cached("go-ml/dyl/"+libSoName))
	assert.Assert(t, so.Ok())
	*(*uintptr)(functionPtr()) = 0
	so.Bind("function", functionPtr())
	assert.Assert(t, function(1) == 2)
}
