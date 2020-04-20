package process

import (
	"bytes"
	"errors"
	"syscall"
	"unsafe"
)

const (
	MAX_MODULE_NAME32        = 255
	MAX_PATH                 = 260
	TH32CS_SNAPHEAPLIST      = 0x00000001
	TH32CS_SNAPPROCESS       = 0x00000002
	TH32CS_SNAPTHREAD        = 0x00000004
	TH32CS_SNAPMODULE        = 0x00000008
	TH32CS_SNAPMODULE32      = 0x00000010
	TH32CS_INHERIT           = 0x80000000
	TH32CS_SNAPALL           = TH32CS_SNAPHEAPLIST | TH32CS_SNAPMODULE | TH32CS_SNAPPROCESS | TH32CS_SNAPTHREAD
	STANDARD_RIGHTS_REQUIRED = 0x000F
	SYNCHRONIZE              = 0x00100000
	PROCESS_ALL_ACCESS       = STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0xffff
)

type (
// BOOL      int32
// DWORD     uint32
// ULONG_PTR uintptr
// HANDLE    uintptr
// LPVOID    unsafe.Pointer
// LPCVOID   unsafe.Pointer
// SIZE_T    uintptr
// HMODULE   uintptr
// BYTE      byte
)

var (
	kernel32                     = syscall.MustLoadDLL("kernel32.dll")
	procCloseHandle              = kernel32.MustFindProc("CloseHandle")
	procCreateToolhelp32Snapshot = kernel32.MustFindProc("CreateToolhelp32Snapshot")
	procGetLastError             = kernel32.MustFindProc("GetLastError")
	procGetModuleHandle          = kernel32.MustFindProc("GetModuleHandleW")
	procProcess32First           = kernel32.MustFindProc("Process32First")
	procProcess32Next            = kernel32.MustFindProc("Process32Next")
	procModule32First            = kernel32.MustFindProc("Module32First")
	procModule32Next             = kernel32.MustFindProc("Module32Next")
	procOpenProcess              = kernel32.MustFindProc("OpenProcess")
	procReadProcessMemory        = kernel32.MustFindProc("ReadProcessMemory")
	procWriteProcessMemory       = kernel32.MustFindProc("WriteProcessMemory")
	psapi                        = syscall.MustLoadDLL("psapi.dll") //kern32 didnt work
	procEnumProcessModules       = psapi.MustFindProc("EnumProcessModules")
)

// https://msdn.microsoft.com/9e2f7345-52bf-4bfc-9761-90b0b374c727
type ProcessEntry32 struct {
	DwSize              uint32
	CntUsage            uint32
	Th32ProcessID       uint32
	Th32DefaultHeapID   uintptr
	Th32ModuleID        uint32
	CntThreads          uint32
	Th32ParentProcessID uint32
	PcPriClassBase      uint32
	DwFlags             uint32
	SzExeFile           [MAX_PATH]uint8
}

// https://msdn.microsoft.com/305fab35-625c-42e3-a434-e2513e4c8870
type ModuleEntry32 struct {
	DwSize        uint32
	Th32ModuleID  uint32
	Th32ProcessID uint32
	GlblcntUsage  uint32
	ProccntUsage  uint32
	ModBaseAddr   *uintptr
	ModBaseSize   uint32
	HModule       uintptr
	SzModule      [MAX_MODULE_NAME32 + 1]uint8
	SzExePath     [MAX_PATH]uint8
}

// https://msdn.microsoft.com/8774e145-ee7f-44de-85db-0445b905f986
func ReadProcessMemory(hProcess uintptr, lpBaseAddress uintptr, lpBuffer *uintptr, nSize uintptr) (uintptr, error) {
	ret, _, err := procReadProcessMemory.Call(
		uintptr(hProcess),
		uintptr(lpBaseAddress),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(nSize),
		0,
	)

	if err.Error() != "The operation completed successfully." {
		return 0, err
	}

	return ret, nil
}

// https://msdn.microsoft.com/9cd91f1c-58ce-4adc-b027-45748543eb06
func WriteProcessMemory(hProcess uintptr, lpBaseAddress uintptr, lpBuffer *uintptr, nSize uintptr) (uintptr, error) {
	ret, _, err := procWriteProcessMemory.Call(
		uintptr(hProcess),
		uintptr(lpBaseAddress),
		uintptr(unsafe.Pointer(lpBuffer)),
		uintptr(nSize),
	)
	
	if err.Error() != "The operation completed successfully." {
		return 0, err
	}

	return ret, nil
}

// https://msdn.microsoft.com/8f695c38-19c4-49e4-97de-8b64ea536cb1
func OpenProcess(dwDesiredAccess uint32, bInheritHandle bool, dwProcessId uint32) (uintptr, error) {
	inHandle := 0
	if bInheritHandle {
		inHandle = 1
	}

	ret, _, _ := procOpenProcess.Call(
		uintptr(dwDesiredAccess),
		uintptr(inHandle),
		uintptr(dwProcessId),
	)

	if ret == 0 {
		return 0, errors.New("failed to open process")
	}

	return uintptr(ret), nil
}

func GetModule(module string, PID uint32) (uintptr, error) {
	var me32 ModuleEntry32
	var snap uintptr

	snap = createToolhelp32Snapshot(TH32CS_SNAPMODULE|TH32CS_SNAPMODULE32, PID)
	me32.DwSize = uint32(unsafe.Sizeof(me32))
	exit := module32First(snap, &me32)
	parsed := parseint8(me32.SzModule[:])

	if !exit {
		closeHandle(snap)

		return (uintptr)(unsafe.Pointer(me32.ModBaseAddr)), errors.New("unexpected exit")
	} else {
		for i := true; i; i = module32Next(snap, &me32) {
			parsed = parseint8(me32.SzModule[:])

			if parsed == module {
				return (uintptr)(unsafe.Pointer(me32.ModBaseAddr)), nil
			}
		}
	}
	return (uintptr)(unsafe.Pointer(me32.ModBaseAddr)), errors.New("not found")
}

func GetProcessID(process string) (uint32, error) {
	var handle uintptr
	var pe32 ProcessEntry32

	handle = createToolhelp32Snapshot(TH32CS_SNAPALL, 0)
	pe32.DwSize = uint32(unsafe.Sizeof(pe32))
	exit := process32First(handle, &pe32)
	parsed := parseint8(pe32.SzExeFile[:])

	if !exit {
		closeHandle(handle)

		return 0, errors.New("failed to get pid")
	} else {
		for i := true; i; i = process32Next(handle, &pe32) {
			parsed = parseint8(pe32.SzExeFile[:])

			if parsed == process {
				return pe32.Th32ProcessID, nil
			}
		}
	}

	return 0, errors.New("failed to get pid")
}

// https://msdn.microsoft.com/df643c25-7558-424c-b187-b3f86ba51358
func createToolhelp32Snapshot(dwFlags uintptr, th32ProcessID uint32) uintptr {
	ret, _, _ := procCreateToolhelp32Snapshot.Call(
		uintptr(dwFlags),
		uintptr(th32ProcessID),
	)

	return uintptr(ret)
}

// https://msdn.microsoft.com/097790e8-30c2-4b00-9256-fa26e2ceb893
func process32First(hSnapshot uintptr, pe *ProcessEntry32) bool {
	ret, _, _ := procProcess32First.Call(
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(pe)),
	)

	return ret != 0
}

// https://msdn.microsoft.com/843a95fd-27ae-4215-83d0-82fc402b82b6
func process32Next(hSnapshot uintptr, pe *ProcessEntry32) bool {
	ret, _, _ := procProcess32Next.Call(
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(pe)),
	)

	return ret != 0
}

// https://msdn.microsoft.com/bb41cab9-13a1-469d-bf76-68c172e982f6
func module32First(hSnapshot uintptr, me *ModuleEntry32) bool {
	ret, _, _ := procModule32First.Call(
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(me)),
	)

	return ret != 0
}

// https://msdn.microsoft.com/88ec1af4-bae7-4cd7-b830-97a98fb337f4
func module32Next(hSnapshot uintptr, me *ModuleEntry32) bool {
	ret, _, _ := procModule32Next.Call(
		uintptr(hSnapshot),
		uintptr(unsafe.Pointer(me)),
	)

	return ret != 0
}

// https://msdn.microsoft.com/9b84891d-62ca-4ddc-97b7-c4c79482abd9
func closeHandle(hObject uintptr) bool {
	ret, _, _ := procCloseHandle.Call(
		uintptr(hObject),
	)

	return ret != 0
}

func parseint8(arr []uint8) string {
	n := bytes.Index(arr, []uint8{0})

	return string(arr[:n])
}
