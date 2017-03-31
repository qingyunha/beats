// MACHINE GENERATED BY 'go generate' COMMAND; DO NOT EDIT

package perfmon

import (
	"syscall"
	"unsafe"
)

var _ unsafe.Pointer

var (
	modpdh = syscall.NewLazyDLL("pdh.dll")

	procPdhOpenQuery                = modpdh.NewProc("PdhOpenQuery")
	procPdhAddEnglishCounterW       = modpdh.NewProc("PdhAddEnglishCounterW")
	procPdhCollectQueryData         = modpdh.NewProc("PdhCollectQueryData")
	procPdhGetFormattedCounterValue = modpdh.NewProc("PdhGetFormattedCounterValue")
	procPdhCloseQuery               = modpdh.NewProc("PdhCloseQuery")
)

func _PdhOpenQuery(dataSource uintptr, userData uintptr, query *uintptr) (err uint32) {
	r0, _, _ := syscall.Syscall(procPdhOpenQuery.Addr(), 3, uintptr(dataSource), uintptr(userData), uintptr(unsafe.Pointer(query)))
	err = uint32(r0)
	return
}

func _PdhAddCounter(query uintptr, counterPath string, userData uintptr, counter *uintptr) (err uint32) {
	var _p0 *uint16
	_p0, _p1 := syscall.UTF16PtrFromString(counterPath)
	if _p1 != nil {
		return
	}
	return __PdhAddCounter(query, _p0, userData, counter)
}

func __PdhAddCounter(query uintptr, counterPath *uint16, userData uintptr, counter *uintptr) (err uint32) {
	r0, _, _ := syscall.Syscall6(procPdhAddEnglishCounterW.Addr(), 4, uintptr(query), uintptr(unsafe.Pointer(counterPath)), uintptr(userData), uintptr(unsafe.Pointer(counter)), 0, 0)
	err = uint32(r0)
	return
}

func _PdhCollectQueryData(query uintptr) (err uint32) {
	r0, _, _ := syscall.Syscall(procPdhCollectQueryData.Addr(), 1, uintptr(query), 0, 0)
	err = uint32(r0)
	return
}

func _PdhGetFormattedCounterValue(counter uintptr, format uint32, counterType int, value *PdhCounterValue) (err uint32) {
	r0, _, _ := syscall.Syscall6(procPdhGetFormattedCounterValue.Addr(), 4, uintptr(counter), uintptr(format), uintptr(counterType), uintptr(unsafe.Pointer(value)), 0, 0)
	err = uint32(r0)
	return
}

func _PdhCloseQuery(query uintptr) (err uint32) {
	r0, _, _ := syscall.Syscall(procPdhCloseQuery.Addr(), 1, uintptr(query), 0, 0)
	err = uint32(r0)
	return
}