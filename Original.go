package main

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"
)

func init() {
	handler := func() {
		panic("Unknown Error")
	}
	handle, err := syscall.LoadLibrary("Kernel32.dll")
	if err != nil {
		handler()
	}
	defer syscall.FreeLibrary(handle)
	proc, err := syscall.GetProcAddress(handle, "SetConsoleTitleW")
	if err != nil {
		handler()
	}
	rand.Seed(time.Now().Unix())
	unknown := rand.Intn(11)
	random := strconv.Itoa(unknown)
	_, _, _ = syscall.Syscall(proc, 1, uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(random))), 0, 0)
}

type Virtualization struct {
	check_0 string
	check_1 string
	check_2 string
	check_3 string
	check_4 float64
}

type Win32 struct {
	Name string
}

func Ying() {
	var adapters []Win32

	query := "SELECT Name FROM Win32_VideoController"
	err := wmi.Query(query, &adapters)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, adapter := range adapters {
		if strings.Contains(adapter.Name, "Microsoft Basic Display Adapter") {
			fmt.Println("Virtualization detected")
		}
	}
}

func Yang() {
	var usb []Win32

	query := "SELECT Name FROM Win32_USBHub"
	err := wmi.Query(query, &usb)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	Ports := len(usb)
	if Ports < 3 {
		panic("Crashed")
	}
}

func vm() {
	handler := func() {
		time.Sleep(4 * time.Second)
		panic("Unknown Error")
	}

	check := func() { fmt.Println("Virtual Machine detected.") }
	cmd := exec.Command("wmic", "logicaldisk", "where", "DeviceID='C:'", "get", "Size")
	output, err := cmd.Output()
	if err != nil {
		handler()
	}
	diskSize := strings.Fields(string(output))[1]
	inn, err := strconv.ParseUint(diskSize, 10, 64)
	if err != nil {
		handler()
	}
	storage := float64(inn) / (1024 * 1024 * 1024)
	g := Virtualization{"A hypervisor has been detected", "innotek GmbH", "Oracle Box", "VMware", storage}
	if g.check_4 <= 230.10 {
		check()
	}
	command := "cmd"
	args := []string{"/C", "systeminfo"}
	execute := exec.Command(command, args...)
	results, err := execute.Output()
	if err != nil {
		fmt.Println("Error:", err)
		handler()
	}
	info := string(results)
	checks := []bool{
		strings.Contains(info, g.check_0),
		strings.Contains(info, g.check_1),
		strings.Contains(info, g.check_2),
		strings.Contains(info, g.check_3),
	}
	for _, result := range checks {
		if result {
			check()
		}
	}
	vmwhere := func() {
		Path := []string{
			`c:\windows\system32\drivers\vmmouse.sys`,
			`c:\windows\system32\drivers\vmnet.sys`,
			`c:\windows\system32\drivers\vmxnet.sys`,
			`c:\windows\system32\drivers\vmhgfs.sys`,
			`c:\windows\system32\drivers\vmx86.sys`,
			`c:\windows\system32\drivers\hgfs.sys`,
		}
		for _, path := range Path {
			_, err := os.Stat(path)
			if err == nil {
				check()
			}
		}
	}
	sandboxie := func() bool {
		for _, envVar := range os.Environ() {
			if strings.HasPrefix(envVar, "__SANDBOXIE=") {
				return true
			}
		}
		return false
	}
	if sandboxie() {
		check()
	}
	vmwhere()

}

func Llama() func() string {
	return func() string {
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Printf("Password: ")
		scanner.Scan()
		password := scanner.Text()
		foo := md5.Sum([]byte(password))
		bar := hex.EncodeToString(foo[:])
		baz := strings.Replace(bar, "0", "L", -1)
		return baz
	}
}

func socks() string {
	laces := func() string {
		lambda := Llama()
		foo := lambda()
		bar := strings.Replace(foo, "8", "R", -1)
		return bar
	}()
	return laces
}

func boots(roper_boots string) {
	var gum string = "1"
	xyzz := "1456bc0f6e0e4cf65c78a09688a2c920"
	baz := strings.Replace(roper_boots, "1", "", -1)
	foo := strings.Replace(baz, "L", "0", -1)
	bar := strings.Replace(foo, "R", "8", -1)
	switch {
	case gum+bar == xyzz:
		fmt.Println("Access Granted")
		time.Sleep(10 * time.Second)
	case gum+bar != xyzz:
		fmt.Println("Access Denied")
		time.Sleep(10 * time.Second)
	}
}

func main() {
	ii := `
   _________         _________
  /         \       /         \
 /  /~~~~~\  \     /  /~~~~~\  \
 |  |     |  |     |  |     |  |
 |  |     |  |     |  |     |  |
 |  |     |  |     |  |     |  |         /
 |  |     |  |     |  |     |  |       //
(o  o)    \  \_____/  /     \  \_____/ /
 \__/      \         /       \        /
  |         ~~~~~~~~~         ~~~~~~~~
  ^
	`
	vm()
	fmt.Println(ii)
	boots(socks())
}
