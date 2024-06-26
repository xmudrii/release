/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by counterfeiter. DO NOT EDIT.
package specsfakes

import (
	"io/fs"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"sync"

	semver "github.com/blang/semver/v4"
	"k8s.io/release/pkg/obs/metadata"
	"k8s.io/release/pkg/release"
)

type FakeImpl struct {
	CompressStub        func(string, string, ...*regexp.Regexp) error
	compressMutex       sync.RWMutex
	compressArgsForCall []struct {
		arg1 string
		arg2 string
		arg3 []*regexp.Regexp
	}
	compressReturns struct {
		result1 error
	}
	compressReturnsOnCall map[int]struct {
		result1 error
	}
	CreateFileStub        func(string) (*os.File, error)
	createFileMutex       sync.RWMutex
	createFileArgsForCall []struct {
		arg1 string
	}
	createFileReturns struct {
		result1 *os.File
		result2 error
	}
	createFileReturnsOnCall map[int]struct {
		result1 *os.File
		result2 error
	}
	ExtractStub        func(string, string) error
	extractMutex       sync.RWMutex
	extractArgsForCall []struct {
		arg1 string
		arg2 string
	}
	extractReturns struct {
		result1 error
	}
	extractReturnsOnCall map[int]struct {
		result1 error
	}
	GCSCopyToLocalStub        func(string, string) error
	gCSCopyToLocalMutex       sync.RWMutex
	gCSCopyToLocalArgsForCall []struct {
		arg1 string
		arg2 string
	}
	gCSCopyToLocalReturns struct {
		result1 error
	}
	gCSCopyToLocalReturnsOnCall map[int]struct {
		result1 error
	}
	GetKubeVersionStub        func(release.VersionType) (string, error)
	getKubeVersionMutex       sync.RWMutex
	getKubeVersionArgsForCall []struct {
		arg1 release.VersionType
	}
	getKubeVersionReturns struct {
		result1 string
		result2 error
	}
	getKubeVersionReturnsOnCall map[int]struct {
		result1 string
		result2 error
	}
	GetRequestStub        func(string) (*http.Response, error)
	getRequestMutex       sync.RWMutex
	getRequestArgsForCall []struct {
		arg1 string
	}
	getRequestReturns struct {
		result1 *http.Response
		result2 error
	}
	getRequestReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	HeadRequestStub        func(string) (*http.Response, error)
	headRequestMutex       sync.RWMutex
	headRequestArgsForCall []struct {
		arg1 string
	}
	headRequestReturns struct {
		result1 *http.Response
		result2 error
	}
	headRequestReturnsOnCall map[int]struct {
		result1 *http.Response
		result2 error
	}
	IsExistStub        func(error) bool
	isExistMutex       sync.RWMutex
	isExistArgsForCall []struct {
		arg1 error
	}
	isExistReturns struct {
		result1 bool
	}
	isExistReturnsOnCall map[int]struct {
		result1 bool
	}
	LoadPackageMetadataStub        func(string) (metadata.PackageMetadataList, error)
	loadPackageMetadataMutex       sync.RWMutex
	loadPackageMetadataArgsForCall []struct {
		arg1 string
	}
	loadPackageMetadataReturns struct {
		result1 metadata.PackageMetadataList
		result2 error
	}
	loadPackageMetadataReturnsOnCall map[int]struct {
		result1 metadata.PackageMetadataList
		result2 error
	}
	MkdirStub        func(string, fs.FileMode) error
	mkdirMutex       sync.RWMutex
	mkdirArgsForCall []struct {
		arg1 string
		arg2 fs.FileMode
	}
	mkdirReturns struct {
		result1 error
	}
	mkdirReturnsOnCall map[int]struct {
		result1 error
	}
	MkdirAllStub        func(string, fs.FileMode) error
	mkdirAllMutex       sync.RWMutex
	mkdirAllArgsForCall []struct {
		arg1 string
		arg2 fs.FileMode
	}
	mkdirAllReturns struct {
		result1 error
	}
	mkdirAllReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveAllStub        func(string) error
	removeAllMutex       sync.RWMutex
	removeAllArgsForCall []struct {
		arg1 string
	}
	removeAllReturns struct {
		result1 error
	}
	removeAllReturnsOnCall map[int]struct {
		result1 error
	}
	RemoveFileStub        func(string) error
	removeFileMutex       sync.RWMutex
	removeFileArgsForCall []struct {
		arg1 string
	}
	removeFileReturns struct {
		result1 error
	}
	removeFileReturnsOnCall map[int]struct {
		result1 error
	}
	StatStub        func(string) (fs.FileInfo, error)
	statMutex       sync.RWMutex
	statArgsForCall []struct {
		arg1 string
	}
	statReturns struct {
		result1 fs.FileInfo
		result2 error
	}
	statReturnsOnCall map[int]struct {
		result1 fs.FileInfo
		result2 error
	}
	TagStringToSemverStub        func(string) (semver.Version, error)
	tagStringToSemverMutex       sync.RWMutex
	tagStringToSemverArgsForCall []struct {
		arg1 string
	}
	tagStringToSemverReturns struct {
		result1 semver.Version
		result2 error
	}
	tagStringToSemverReturnsOnCall map[int]struct {
		result1 semver.Version
		result2 error
	}
	TrimTagPrefixStub        func(string) string
	trimTagPrefixMutex       sync.RWMutex
	trimTagPrefixArgsForCall []struct {
		arg1 string
	}
	trimTagPrefixReturns struct {
		result1 string
	}
	trimTagPrefixReturnsOnCall map[int]struct {
		result1 string
	}
	WalkStub        func(string, filepath.WalkFunc) error
	walkMutex       sync.RWMutex
	walkArgsForCall []struct {
		arg1 string
		arg2 filepath.WalkFunc
	}
	walkReturns struct {
		result1 error
	}
	walkReturnsOnCall map[int]struct {
		result1 error
	}
	WriteFileStub        func(string, []byte, fs.FileMode) error
	writeFileMutex       sync.RWMutex
	writeFileArgsForCall []struct {
		arg1 string
		arg2 []byte
		arg3 fs.FileMode
	}
	writeFileReturns struct {
		result1 error
	}
	writeFileReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *FakeImpl) Compress(arg1 string, arg2 string, arg3 ...*regexp.Regexp) error {
	fake.compressMutex.Lock()
	ret, specificReturn := fake.compressReturnsOnCall[len(fake.compressArgsForCall)]
	fake.compressArgsForCall = append(fake.compressArgsForCall, struct {
		arg1 string
		arg2 string
		arg3 []*regexp.Regexp
	}{arg1, arg2, arg3})
	stub := fake.CompressStub
	fakeReturns := fake.compressReturns
	fake.recordInvocation("Compress", []interface{}{arg1, arg2, arg3})
	fake.compressMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3...)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) CompressCallCount() int {
	fake.compressMutex.RLock()
	defer fake.compressMutex.RUnlock()
	return len(fake.compressArgsForCall)
}

func (fake *FakeImpl) CompressCalls(stub func(string, string, ...*regexp.Regexp) error) {
	fake.compressMutex.Lock()
	defer fake.compressMutex.Unlock()
	fake.CompressStub = stub
}

func (fake *FakeImpl) CompressArgsForCall(i int) (string, string, []*regexp.Regexp) {
	fake.compressMutex.RLock()
	defer fake.compressMutex.RUnlock()
	argsForCall := fake.compressArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImpl) CompressReturns(result1 error) {
	fake.compressMutex.Lock()
	defer fake.compressMutex.Unlock()
	fake.CompressStub = nil
	fake.compressReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) CompressReturnsOnCall(i int, result1 error) {
	fake.compressMutex.Lock()
	defer fake.compressMutex.Unlock()
	fake.CompressStub = nil
	if fake.compressReturnsOnCall == nil {
		fake.compressReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.compressReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) CreateFile(arg1 string) (*os.File, error) {
	fake.createFileMutex.Lock()
	ret, specificReturn := fake.createFileReturnsOnCall[len(fake.createFileArgsForCall)]
	fake.createFileArgsForCall = append(fake.createFileArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.CreateFileStub
	fakeReturns := fake.createFileReturns
	fake.recordInvocation("CreateFile", []interface{}{arg1})
	fake.createFileMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) CreateFileCallCount() int {
	fake.createFileMutex.RLock()
	defer fake.createFileMutex.RUnlock()
	return len(fake.createFileArgsForCall)
}

func (fake *FakeImpl) CreateFileCalls(stub func(string) (*os.File, error)) {
	fake.createFileMutex.Lock()
	defer fake.createFileMutex.Unlock()
	fake.CreateFileStub = stub
}

func (fake *FakeImpl) CreateFileArgsForCall(i int) string {
	fake.createFileMutex.RLock()
	defer fake.createFileMutex.RUnlock()
	argsForCall := fake.createFileArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) CreateFileReturns(result1 *os.File, result2 error) {
	fake.createFileMutex.Lock()
	defer fake.createFileMutex.Unlock()
	fake.CreateFileStub = nil
	fake.createFileReturns = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) CreateFileReturnsOnCall(i int, result1 *os.File, result2 error) {
	fake.createFileMutex.Lock()
	defer fake.createFileMutex.Unlock()
	fake.CreateFileStub = nil
	if fake.createFileReturnsOnCall == nil {
		fake.createFileReturnsOnCall = make(map[int]struct {
			result1 *os.File
			result2 error
		})
	}
	fake.createFileReturnsOnCall[i] = struct {
		result1 *os.File
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) Extract(arg1 string, arg2 string) error {
	fake.extractMutex.Lock()
	ret, specificReturn := fake.extractReturnsOnCall[len(fake.extractArgsForCall)]
	fake.extractArgsForCall = append(fake.extractArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.ExtractStub
	fakeReturns := fake.extractReturns
	fake.recordInvocation("Extract", []interface{}{arg1, arg2})
	fake.extractMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) ExtractCallCount() int {
	fake.extractMutex.RLock()
	defer fake.extractMutex.RUnlock()
	return len(fake.extractArgsForCall)
}

func (fake *FakeImpl) ExtractCalls(stub func(string, string) error) {
	fake.extractMutex.Lock()
	defer fake.extractMutex.Unlock()
	fake.ExtractStub = stub
}

func (fake *FakeImpl) ExtractArgsForCall(i int) (string, string) {
	fake.extractMutex.RLock()
	defer fake.extractMutex.RUnlock()
	argsForCall := fake.extractArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) ExtractReturns(result1 error) {
	fake.extractMutex.Lock()
	defer fake.extractMutex.Unlock()
	fake.ExtractStub = nil
	fake.extractReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) ExtractReturnsOnCall(i int, result1 error) {
	fake.extractMutex.Lock()
	defer fake.extractMutex.Unlock()
	fake.ExtractStub = nil
	if fake.extractReturnsOnCall == nil {
		fake.extractReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.extractReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) GCSCopyToLocal(arg1 string, arg2 string) error {
	fake.gCSCopyToLocalMutex.Lock()
	ret, specificReturn := fake.gCSCopyToLocalReturnsOnCall[len(fake.gCSCopyToLocalArgsForCall)]
	fake.gCSCopyToLocalArgsForCall = append(fake.gCSCopyToLocalArgsForCall, struct {
		arg1 string
		arg2 string
	}{arg1, arg2})
	stub := fake.GCSCopyToLocalStub
	fakeReturns := fake.gCSCopyToLocalReturns
	fake.recordInvocation("GCSCopyToLocal", []interface{}{arg1, arg2})
	fake.gCSCopyToLocalMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) GCSCopyToLocalCallCount() int {
	fake.gCSCopyToLocalMutex.RLock()
	defer fake.gCSCopyToLocalMutex.RUnlock()
	return len(fake.gCSCopyToLocalArgsForCall)
}

func (fake *FakeImpl) GCSCopyToLocalCalls(stub func(string, string) error) {
	fake.gCSCopyToLocalMutex.Lock()
	defer fake.gCSCopyToLocalMutex.Unlock()
	fake.GCSCopyToLocalStub = stub
}

func (fake *FakeImpl) GCSCopyToLocalArgsForCall(i int) (string, string) {
	fake.gCSCopyToLocalMutex.RLock()
	defer fake.gCSCopyToLocalMutex.RUnlock()
	argsForCall := fake.gCSCopyToLocalArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) GCSCopyToLocalReturns(result1 error) {
	fake.gCSCopyToLocalMutex.Lock()
	defer fake.gCSCopyToLocalMutex.Unlock()
	fake.GCSCopyToLocalStub = nil
	fake.gCSCopyToLocalReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) GCSCopyToLocalReturnsOnCall(i int, result1 error) {
	fake.gCSCopyToLocalMutex.Lock()
	defer fake.gCSCopyToLocalMutex.Unlock()
	fake.GCSCopyToLocalStub = nil
	if fake.gCSCopyToLocalReturnsOnCall == nil {
		fake.gCSCopyToLocalReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.gCSCopyToLocalReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) GetKubeVersion(arg1 release.VersionType) (string, error) {
	fake.getKubeVersionMutex.Lock()
	ret, specificReturn := fake.getKubeVersionReturnsOnCall[len(fake.getKubeVersionArgsForCall)]
	fake.getKubeVersionArgsForCall = append(fake.getKubeVersionArgsForCall, struct {
		arg1 release.VersionType
	}{arg1})
	stub := fake.GetKubeVersionStub
	fakeReturns := fake.getKubeVersionReturns
	fake.recordInvocation("GetKubeVersion", []interface{}{arg1})
	fake.getKubeVersionMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) GetKubeVersionCallCount() int {
	fake.getKubeVersionMutex.RLock()
	defer fake.getKubeVersionMutex.RUnlock()
	return len(fake.getKubeVersionArgsForCall)
}

func (fake *FakeImpl) GetKubeVersionCalls(stub func(release.VersionType) (string, error)) {
	fake.getKubeVersionMutex.Lock()
	defer fake.getKubeVersionMutex.Unlock()
	fake.GetKubeVersionStub = stub
}

func (fake *FakeImpl) GetKubeVersionArgsForCall(i int) release.VersionType {
	fake.getKubeVersionMutex.RLock()
	defer fake.getKubeVersionMutex.RUnlock()
	argsForCall := fake.getKubeVersionArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) GetKubeVersionReturns(result1 string, result2 error) {
	fake.getKubeVersionMutex.Lock()
	defer fake.getKubeVersionMutex.Unlock()
	fake.GetKubeVersionStub = nil
	fake.getKubeVersionReturns = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) GetKubeVersionReturnsOnCall(i int, result1 string, result2 error) {
	fake.getKubeVersionMutex.Lock()
	defer fake.getKubeVersionMutex.Unlock()
	fake.GetKubeVersionStub = nil
	if fake.getKubeVersionReturnsOnCall == nil {
		fake.getKubeVersionReturnsOnCall = make(map[int]struct {
			result1 string
			result2 error
		})
	}
	fake.getKubeVersionReturnsOnCall[i] = struct {
		result1 string
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) GetRequest(arg1 string) (*http.Response, error) {
	fake.getRequestMutex.Lock()
	ret, specificReturn := fake.getRequestReturnsOnCall[len(fake.getRequestArgsForCall)]
	fake.getRequestArgsForCall = append(fake.getRequestArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.GetRequestStub
	fakeReturns := fake.getRequestReturns
	fake.recordInvocation("GetRequest", []interface{}{arg1})
	fake.getRequestMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) GetRequestCallCount() int {
	fake.getRequestMutex.RLock()
	defer fake.getRequestMutex.RUnlock()
	return len(fake.getRequestArgsForCall)
}

func (fake *FakeImpl) GetRequestCalls(stub func(string) (*http.Response, error)) {
	fake.getRequestMutex.Lock()
	defer fake.getRequestMutex.Unlock()
	fake.GetRequestStub = stub
}

func (fake *FakeImpl) GetRequestArgsForCall(i int) string {
	fake.getRequestMutex.RLock()
	defer fake.getRequestMutex.RUnlock()
	argsForCall := fake.getRequestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) GetRequestReturns(result1 *http.Response, result2 error) {
	fake.getRequestMutex.Lock()
	defer fake.getRequestMutex.Unlock()
	fake.GetRequestStub = nil
	fake.getRequestReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) GetRequestReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.getRequestMutex.Lock()
	defer fake.getRequestMutex.Unlock()
	fake.GetRequestStub = nil
	if fake.getRequestReturnsOnCall == nil {
		fake.getRequestReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.getRequestReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) HeadRequest(arg1 string) (*http.Response, error) {
	fake.headRequestMutex.Lock()
	ret, specificReturn := fake.headRequestReturnsOnCall[len(fake.headRequestArgsForCall)]
	fake.headRequestArgsForCall = append(fake.headRequestArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.HeadRequestStub
	fakeReturns := fake.headRequestReturns
	fake.recordInvocation("HeadRequest", []interface{}{arg1})
	fake.headRequestMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) HeadRequestCallCount() int {
	fake.headRequestMutex.RLock()
	defer fake.headRequestMutex.RUnlock()
	return len(fake.headRequestArgsForCall)
}

func (fake *FakeImpl) HeadRequestCalls(stub func(string) (*http.Response, error)) {
	fake.headRequestMutex.Lock()
	defer fake.headRequestMutex.Unlock()
	fake.HeadRequestStub = stub
}

func (fake *FakeImpl) HeadRequestArgsForCall(i int) string {
	fake.headRequestMutex.RLock()
	defer fake.headRequestMutex.RUnlock()
	argsForCall := fake.headRequestArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) HeadRequestReturns(result1 *http.Response, result2 error) {
	fake.headRequestMutex.Lock()
	defer fake.headRequestMutex.Unlock()
	fake.HeadRequestStub = nil
	fake.headRequestReturns = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) HeadRequestReturnsOnCall(i int, result1 *http.Response, result2 error) {
	fake.headRequestMutex.Lock()
	defer fake.headRequestMutex.Unlock()
	fake.HeadRequestStub = nil
	if fake.headRequestReturnsOnCall == nil {
		fake.headRequestReturnsOnCall = make(map[int]struct {
			result1 *http.Response
			result2 error
		})
	}
	fake.headRequestReturnsOnCall[i] = struct {
		result1 *http.Response
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) IsExist(arg1 error) bool {
	fake.isExistMutex.Lock()
	ret, specificReturn := fake.isExistReturnsOnCall[len(fake.isExistArgsForCall)]
	fake.isExistArgsForCall = append(fake.isExistArgsForCall, struct {
		arg1 error
	}{arg1})
	stub := fake.IsExistStub
	fakeReturns := fake.isExistReturns
	fake.recordInvocation("IsExist", []interface{}{arg1})
	fake.isExistMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) IsExistCallCount() int {
	fake.isExistMutex.RLock()
	defer fake.isExistMutex.RUnlock()
	return len(fake.isExistArgsForCall)
}

func (fake *FakeImpl) IsExistCalls(stub func(error) bool) {
	fake.isExistMutex.Lock()
	defer fake.isExistMutex.Unlock()
	fake.IsExistStub = stub
}

func (fake *FakeImpl) IsExistArgsForCall(i int) error {
	fake.isExistMutex.RLock()
	defer fake.isExistMutex.RUnlock()
	argsForCall := fake.isExistArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) IsExistReturns(result1 bool) {
	fake.isExistMutex.Lock()
	defer fake.isExistMutex.Unlock()
	fake.IsExistStub = nil
	fake.isExistReturns = struct {
		result1 bool
	}{result1}
}

func (fake *FakeImpl) IsExistReturnsOnCall(i int, result1 bool) {
	fake.isExistMutex.Lock()
	defer fake.isExistMutex.Unlock()
	fake.IsExistStub = nil
	if fake.isExistReturnsOnCall == nil {
		fake.isExistReturnsOnCall = make(map[int]struct {
			result1 bool
		})
	}
	fake.isExistReturnsOnCall[i] = struct {
		result1 bool
	}{result1}
}

func (fake *FakeImpl) LoadPackageMetadata(arg1 string) (metadata.PackageMetadataList, error) {
	fake.loadPackageMetadataMutex.Lock()
	ret, specificReturn := fake.loadPackageMetadataReturnsOnCall[len(fake.loadPackageMetadataArgsForCall)]
	fake.loadPackageMetadataArgsForCall = append(fake.loadPackageMetadataArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.LoadPackageMetadataStub
	fakeReturns := fake.loadPackageMetadataReturns
	fake.recordInvocation("LoadPackageMetadata", []interface{}{arg1})
	fake.loadPackageMetadataMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) LoadPackageMetadataCallCount() int {
	fake.loadPackageMetadataMutex.RLock()
	defer fake.loadPackageMetadataMutex.RUnlock()
	return len(fake.loadPackageMetadataArgsForCall)
}

func (fake *FakeImpl) LoadPackageMetadataCalls(stub func(string) (metadata.PackageMetadataList, error)) {
	fake.loadPackageMetadataMutex.Lock()
	defer fake.loadPackageMetadataMutex.Unlock()
	fake.LoadPackageMetadataStub = stub
}

func (fake *FakeImpl) LoadPackageMetadataArgsForCall(i int) string {
	fake.loadPackageMetadataMutex.RLock()
	defer fake.loadPackageMetadataMutex.RUnlock()
	argsForCall := fake.loadPackageMetadataArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) LoadPackageMetadataReturns(result1 metadata.PackageMetadataList, result2 error) {
	fake.loadPackageMetadataMutex.Lock()
	defer fake.loadPackageMetadataMutex.Unlock()
	fake.LoadPackageMetadataStub = nil
	fake.loadPackageMetadataReturns = struct {
		result1 metadata.PackageMetadataList
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) LoadPackageMetadataReturnsOnCall(i int, result1 metadata.PackageMetadataList, result2 error) {
	fake.loadPackageMetadataMutex.Lock()
	defer fake.loadPackageMetadataMutex.Unlock()
	fake.LoadPackageMetadataStub = nil
	if fake.loadPackageMetadataReturnsOnCall == nil {
		fake.loadPackageMetadataReturnsOnCall = make(map[int]struct {
			result1 metadata.PackageMetadataList
			result2 error
		})
	}
	fake.loadPackageMetadataReturnsOnCall[i] = struct {
		result1 metadata.PackageMetadataList
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) Mkdir(arg1 string, arg2 fs.FileMode) error {
	fake.mkdirMutex.Lock()
	ret, specificReturn := fake.mkdirReturnsOnCall[len(fake.mkdirArgsForCall)]
	fake.mkdirArgsForCall = append(fake.mkdirArgsForCall, struct {
		arg1 string
		arg2 fs.FileMode
	}{arg1, arg2})
	stub := fake.MkdirStub
	fakeReturns := fake.mkdirReturns
	fake.recordInvocation("Mkdir", []interface{}{arg1, arg2})
	fake.mkdirMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) MkdirCallCount() int {
	fake.mkdirMutex.RLock()
	defer fake.mkdirMutex.RUnlock()
	return len(fake.mkdirArgsForCall)
}

func (fake *FakeImpl) MkdirCalls(stub func(string, fs.FileMode) error) {
	fake.mkdirMutex.Lock()
	defer fake.mkdirMutex.Unlock()
	fake.MkdirStub = stub
}

func (fake *FakeImpl) MkdirArgsForCall(i int) (string, fs.FileMode) {
	fake.mkdirMutex.RLock()
	defer fake.mkdirMutex.RUnlock()
	argsForCall := fake.mkdirArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) MkdirReturns(result1 error) {
	fake.mkdirMutex.Lock()
	defer fake.mkdirMutex.Unlock()
	fake.MkdirStub = nil
	fake.mkdirReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) MkdirReturnsOnCall(i int, result1 error) {
	fake.mkdirMutex.Lock()
	defer fake.mkdirMutex.Unlock()
	fake.MkdirStub = nil
	if fake.mkdirReturnsOnCall == nil {
		fake.mkdirReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mkdirReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) MkdirAll(arg1 string, arg2 fs.FileMode) error {
	fake.mkdirAllMutex.Lock()
	ret, specificReturn := fake.mkdirAllReturnsOnCall[len(fake.mkdirAllArgsForCall)]
	fake.mkdirAllArgsForCall = append(fake.mkdirAllArgsForCall, struct {
		arg1 string
		arg2 fs.FileMode
	}{arg1, arg2})
	stub := fake.MkdirAllStub
	fakeReturns := fake.mkdirAllReturns
	fake.recordInvocation("MkdirAll", []interface{}{arg1, arg2})
	fake.mkdirAllMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) MkdirAllCallCount() int {
	fake.mkdirAllMutex.RLock()
	defer fake.mkdirAllMutex.RUnlock()
	return len(fake.mkdirAllArgsForCall)
}

func (fake *FakeImpl) MkdirAllCalls(stub func(string, fs.FileMode) error) {
	fake.mkdirAllMutex.Lock()
	defer fake.mkdirAllMutex.Unlock()
	fake.MkdirAllStub = stub
}

func (fake *FakeImpl) MkdirAllArgsForCall(i int) (string, fs.FileMode) {
	fake.mkdirAllMutex.RLock()
	defer fake.mkdirAllMutex.RUnlock()
	argsForCall := fake.mkdirAllArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) MkdirAllReturns(result1 error) {
	fake.mkdirAllMutex.Lock()
	defer fake.mkdirAllMutex.Unlock()
	fake.MkdirAllStub = nil
	fake.mkdirAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) MkdirAllReturnsOnCall(i int, result1 error) {
	fake.mkdirAllMutex.Lock()
	defer fake.mkdirAllMutex.Unlock()
	fake.MkdirAllStub = nil
	if fake.mkdirAllReturnsOnCall == nil {
		fake.mkdirAllReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.mkdirAllReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) RemoveAll(arg1 string) error {
	fake.removeAllMutex.Lock()
	ret, specificReturn := fake.removeAllReturnsOnCall[len(fake.removeAllArgsForCall)]
	fake.removeAllArgsForCall = append(fake.removeAllArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.RemoveAllStub
	fakeReturns := fake.removeAllReturns
	fake.recordInvocation("RemoveAll", []interface{}{arg1})
	fake.removeAllMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) RemoveAllCallCount() int {
	fake.removeAllMutex.RLock()
	defer fake.removeAllMutex.RUnlock()
	return len(fake.removeAllArgsForCall)
}

func (fake *FakeImpl) RemoveAllCalls(stub func(string) error) {
	fake.removeAllMutex.Lock()
	defer fake.removeAllMutex.Unlock()
	fake.RemoveAllStub = stub
}

func (fake *FakeImpl) RemoveAllArgsForCall(i int) string {
	fake.removeAllMutex.RLock()
	defer fake.removeAllMutex.RUnlock()
	argsForCall := fake.removeAllArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) RemoveAllReturns(result1 error) {
	fake.removeAllMutex.Lock()
	defer fake.removeAllMutex.Unlock()
	fake.RemoveAllStub = nil
	fake.removeAllReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) RemoveAllReturnsOnCall(i int, result1 error) {
	fake.removeAllMutex.Lock()
	defer fake.removeAllMutex.Unlock()
	fake.RemoveAllStub = nil
	if fake.removeAllReturnsOnCall == nil {
		fake.removeAllReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeAllReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) RemoveFile(arg1 string) error {
	fake.removeFileMutex.Lock()
	ret, specificReturn := fake.removeFileReturnsOnCall[len(fake.removeFileArgsForCall)]
	fake.removeFileArgsForCall = append(fake.removeFileArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.RemoveFileStub
	fakeReturns := fake.removeFileReturns
	fake.recordInvocation("RemoveFile", []interface{}{arg1})
	fake.removeFileMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) RemoveFileCallCount() int {
	fake.removeFileMutex.RLock()
	defer fake.removeFileMutex.RUnlock()
	return len(fake.removeFileArgsForCall)
}

func (fake *FakeImpl) RemoveFileCalls(stub func(string) error) {
	fake.removeFileMutex.Lock()
	defer fake.removeFileMutex.Unlock()
	fake.RemoveFileStub = stub
}

func (fake *FakeImpl) RemoveFileArgsForCall(i int) string {
	fake.removeFileMutex.RLock()
	defer fake.removeFileMutex.RUnlock()
	argsForCall := fake.removeFileArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) RemoveFileReturns(result1 error) {
	fake.removeFileMutex.Lock()
	defer fake.removeFileMutex.Unlock()
	fake.RemoveFileStub = nil
	fake.removeFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) RemoveFileReturnsOnCall(i int, result1 error) {
	fake.removeFileMutex.Lock()
	defer fake.removeFileMutex.Unlock()
	fake.RemoveFileStub = nil
	if fake.removeFileReturnsOnCall == nil {
		fake.removeFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.removeFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) Stat(arg1 string) (fs.FileInfo, error) {
	fake.statMutex.Lock()
	ret, specificReturn := fake.statReturnsOnCall[len(fake.statArgsForCall)]
	fake.statArgsForCall = append(fake.statArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.StatStub
	fakeReturns := fake.statReturns
	fake.recordInvocation("Stat", []interface{}{arg1})
	fake.statMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) StatCallCount() int {
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	return len(fake.statArgsForCall)
}

func (fake *FakeImpl) StatCalls(stub func(string) (fs.FileInfo, error)) {
	fake.statMutex.Lock()
	defer fake.statMutex.Unlock()
	fake.StatStub = stub
}

func (fake *FakeImpl) StatArgsForCall(i int) string {
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	argsForCall := fake.statArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) StatReturns(result1 fs.FileInfo, result2 error) {
	fake.statMutex.Lock()
	defer fake.statMutex.Unlock()
	fake.StatStub = nil
	fake.statReturns = struct {
		result1 fs.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) StatReturnsOnCall(i int, result1 fs.FileInfo, result2 error) {
	fake.statMutex.Lock()
	defer fake.statMutex.Unlock()
	fake.StatStub = nil
	if fake.statReturnsOnCall == nil {
		fake.statReturnsOnCall = make(map[int]struct {
			result1 fs.FileInfo
			result2 error
		})
	}
	fake.statReturnsOnCall[i] = struct {
		result1 fs.FileInfo
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) TagStringToSemver(arg1 string) (semver.Version, error) {
	fake.tagStringToSemverMutex.Lock()
	ret, specificReturn := fake.tagStringToSemverReturnsOnCall[len(fake.tagStringToSemverArgsForCall)]
	fake.tagStringToSemverArgsForCall = append(fake.tagStringToSemverArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.TagStringToSemverStub
	fakeReturns := fake.tagStringToSemverReturns
	fake.recordInvocation("TagStringToSemver", []interface{}{arg1})
	fake.tagStringToSemverMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *FakeImpl) TagStringToSemverCallCount() int {
	fake.tagStringToSemverMutex.RLock()
	defer fake.tagStringToSemverMutex.RUnlock()
	return len(fake.tagStringToSemverArgsForCall)
}

func (fake *FakeImpl) TagStringToSemverCalls(stub func(string) (semver.Version, error)) {
	fake.tagStringToSemverMutex.Lock()
	defer fake.tagStringToSemverMutex.Unlock()
	fake.TagStringToSemverStub = stub
}

func (fake *FakeImpl) TagStringToSemverArgsForCall(i int) string {
	fake.tagStringToSemverMutex.RLock()
	defer fake.tagStringToSemverMutex.RUnlock()
	argsForCall := fake.tagStringToSemverArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) TagStringToSemverReturns(result1 semver.Version, result2 error) {
	fake.tagStringToSemverMutex.Lock()
	defer fake.tagStringToSemverMutex.Unlock()
	fake.TagStringToSemverStub = nil
	fake.tagStringToSemverReturns = struct {
		result1 semver.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) TagStringToSemverReturnsOnCall(i int, result1 semver.Version, result2 error) {
	fake.tagStringToSemverMutex.Lock()
	defer fake.tagStringToSemverMutex.Unlock()
	fake.TagStringToSemverStub = nil
	if fake.tagStringToSemverReturnsOnCall == nil {
		fake.tagStringToSemverReturnsOnCall = make(map[int]struct {
			result1 semver.Version
			result2 error
		})
	}
	fake.tagStringToSemverReturnsOnCall[i] = struct {
		result1 semver.Version
		result2 error
	}{result1, result2}
}

func (fake *FakeImpl) TrimTagPrefix(arg1 string) string {
	fake.trimTagPrefixMutex.Lock()
	ret, specificReturn := fake.trimTagPrefixReturnsOnCall[len(fake.trimTagPrefixArgsForCall)]
	fake.trimTagPrefixArgsForCall = append(fake.trimTagPrefixArgsForCall, struct {
		arg1 string
	}{arg1})
	stub := fake.TrimTagPrefixStub
	fakeReturns := fake.trimTagPrefixReturns
	fake.recordInvocation("TrimTagPrefix", []interface{}{arg1})
	fake.trimTagPrefixMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) TrimTagPrefixCallCount() int {
	fake.trimTagPrefixMutex.RLock()
	defer fake.trimTagPrefixMutex.RUnlock()
	return len(fake.trimTagPrefixArgsForCall)
}

func (fake *FakeImpl) TrimTagPrefixCalls(stub func(string) string) {
	fake.trimTagPrefixMutex.Lock()
	defer fake.trimTagPrefixMutex.Unlock()
	fake.TrimTagPrefixStub = stub
}

func (fake *FakeImpl) TrimTagPrefixArgsForCall(i int) string {
	fake.trimTagPrefixMutex.RLock()
	defer fake.trimTagPrefixMutex.RUnlock()
	argsForCall := fake.trimTagPrefixArgsForCall[i]
	return argsForCall.arg1
}

func (fake *FakeImpl) TrimTagPrefixReturns(result1 string) {
	fake.trimTagPrefixMutex.Lock()
	defer fake.trimTagPrefixMutex.Unlock()
	fake.TrimTagPrefixStub = nil
	fake.trimTagPrefixReturns = struct {
		result1 string
	}{result1}
}

func (fake *FakeImpl) TrimTagPrefixReturnsOnCall(i int, result1 string) {
	fake.trimTagPrefixMutex.Lock()
	defer fake.trimTagPrefixMutex.Unlock()
	fake.TrimTagPrefixStub = nil
	if fake.trimTagPrefixReturnsOnCall == nil {
		fake.trimTagPrefixReturnsOnCall = make(map[int]struct {
			result1 string
		})
	}
	fake.trimTagPrefixReturnsOnCall[i] = struct {
		result1 string
	}{result1}
}

func (fake *FakeImpl) Walk(arg1 string, arg2 filepath.WalkFunc) error {
	fake.walkMutex.Lock()
	ret, specificReturn := fake.walkReturnsOnCall[len(fake.walkArgsForCall)]
	fake.walkArgsForCall = append(fake.walkArgsForCall, struct {
		arg1 string
		arg2 filepath.WalkFunc
	}{arg1, arg2})
	stub := fake.WalkStub
	fakeReturns := fake.walkReturns
	fake.recordInvocation("Walk", []interface{}{arg1, arg2})
	fake.walkMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) WalkCallCount() int {
	fake.walkMutex.RLock()
	defer fake.walkMutex.RUnlock()
	return len(fake.walkArgsForCall)
}

func (fake *FakeImpl) WalkCalls(stub func(string, filepath.WalkFunc) error) {
	fake.walkMutex.Lock()
	defer fake.walkMutex.Unlock()
	fake.WalkStub = stub
}

func (fake *FakeImpl) WalkArgsForCall(i int) (string, filepath.WalkFunc) {
	fake.walkMutex.RLock()
	defer fake.walkMutex.RUnlock()
	argsForCall := fake.walkArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *FakeImpl) WalkReturns(result1 error) {
	fake.walkMutex.Lock()
	defer fake.walkMutex.Unlock()
	fake.WalkStub = nil
	fake.walkReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) WalkReturnsOnCall(i int, result1 error) {
	fake.walkMutex.Lock()
	defer fake.walkMutex.Unlock()
	fake.WalkStub = nil
	if fake.walkReturnsOnCall == nil {
		fake.walkReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.walkReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) WriteFile(arg1 string, arg2 []byte, arg3 fs.FileMode) error {
	var arg2Copy []byte
	if arg2 != nil {
		arg2Copy = make([]byte, len(arg2))
		copy(arg2Copy, arg2)
	}
	fake.writeFileMutex.Lock()
	ret, specificReturn := fake.writeFileReturnsOnCall[len(fake.writeFileArgsForCall)]
	fake.writeFileArgsForCall = append(fake.writeFileArgsForCall, struct {
		arg1 string
		arg2 []byte
		arg3 fs.FileMode
	}{arg1, arg2Copy, arg3})
	stub := fake.WriteFileStub
	fakeReturns := fake.writeFileReturns
	fake.recordInvocation("WriteFile", []interface{}{arg1, arg2Copy, arg3})
	fake.writeFileMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *FakeImpl) WriteFileCallCount() int {
	fake.writeFileMutex.RLock()
	defer fake.writeFileMutex.RUnlock()
	return len(fake.writeFileArgsForCall)
}

func (fake *FakeImpl) WriteFileCalls(stub func(string, []byte, fs.FileMode) error) {
	fake.writeFileMutex.Lock()
	defer fake.writeFileMutex.Unlock()
	fake.WriteFileStub = stub
}

func (fake *FakeImpl) WriteFileArgsForCall(i int) (string, []byte, fs.FileMode) {
	fake.writeFileMutex.RLock()
	defer fake.writeFileMutex.RUnlock()
	argsForCall := fake.writeFileArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *FakeImpl) WriteFileReturns(result1 error) {
	fake.writeFileMutex.Lock()
	defer fake.writeFileMutex.Unlock()
	fake.WriteFileStub = nil
	fake.writeFileReturns = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) WriteFileReturnsOnCall(i int, result1 error) {
	fake.writeFileMutex.Lock()
	defer fake.writeFileMutex.Unlock()
	fake.WriteFileStub = nil
	if fake.writeFileReturnsOnCall == nil {
		fake.writeFileReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.writeFileReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *FakeImpl) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.compressMutex.RLock()
	defer fake.compressMutex.RUnlock()
	fake.createFileMutex.RLock()
	defer fake.createFileMutex.RUnlock()
	fake.extractMutex.RLock()
	defer fake.extractMutex.RUnlock()
	fake.gCSCopyToLocalMutex.RLock()
	defer fake.gCSCopyToLocalMutex.RUnlock()
	fake.getKubeVersionMutex.RLock()
	defer fake.getKubeVersionMutex.RUnlock()
	fake.getRequestMutex.RLock()
	defer fake.getRequestMutex.RUnlock()
	fake.headRequestMutex.RLock()
	defer fake.headRequestMutex.RUnlock()
	fake.isExistMutex.RLock()
	defer fake.isExistMutex.RUnlock()
	fake.loadPackageMetadataMutex.RLock()
	defer fake.loadPackageMetadataMutex.RUnlock()
	fake.mkdirMutex.RLock()
	defer fake.mkdirMutex.RUnlock()
	fake.mkdirAllMutex.RLock()
	defer fake.mkdirAllMutex.RUnlock()
	fake.removeAllMutex.RLock()
	defer fake.removeAllMutex.RUnlock()
	fake.removeFileMutex.RLock()
	defer fake.removeFileMutex.RUnlock()
	fake.statMutex.RLock()
	defer fake.statMutex.RUnlock()
	fake.tagStringToSemverMutex.RLock()
	defer fake.tagStringToSemverMutex.RUnlock()
	fake.trimTagPrefixMutex.RLock()
	defer fake.trimTagPrefixMutex.RUnlock()
	fake.walkMutex.RLock()
	defer fake.walkMutex.RUnlock()
	fake.writeFileMutex.RLock()
	defer fake.writeFileMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *FakeImpl) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}
