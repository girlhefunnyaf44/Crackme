package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"strings"
	"syscall"
	"time"
	"unsafe"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/yusufpapurcu/wmi"
)

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
		log.Println("Error:", err)
		return
	}

	for _, adapter := range adapters {
		if strings.Contains(adapter.Name, "Microsoft Basic Display Adapter") {
			log.Println("Virtualization detected")
		}
	}
}

func Yang() {
	var usb []Win32

	query := "SELECT Name FROM Win32_USBHub"
	err := wmi.Query(query, &usb)
	if err != nil {
		log.Println("Error:", err)
		return
	}

	Ports := len(usb)
	if Ports < 3 {
		log.Println("Virtual Machine Detected: Type YG")
	}
}

func vom() {
	arch := runtime.GOARCH
	log.Println("System Architecture:", arch)
	handler := func() {
		time.Sleep(4 * time.Second)
		panic("Unknown Error")
	}

	check := func() { log.Println("Virtual Machine detected.") }
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
		log.Println("Error:", err)
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

var (
	user32      = syscall.NewLazyDLL("user32.dll")
	messageBoxW = user32.NewProc("MessageBoxW")
)

type Dps struct {
	NULL        uintptr
	MB_OKCANCEL uintptr
	MB_ICONHAND uintptr
	title       string
	message     string
}

func incoherence() {
	var Msg = Dps{0, 0x1, 0x10, "Wrong Password", "Skill Issue"}
	lpText := syscall.StringToUTF16Ptr(Msg.message)
	lpCaption := syscall.StringToUTF16Ptr(Msg.title)
	uType := Msg.MB_OKCANCEL | Msg.MB_ICONHAND

	result, _, _ := messageBoxW.Call(Msg.NULL, uintptr(unsafe.Pointer(lpText)), uintptr(unsafe.Pointer(lpCaption)), uintptr(uType))
	log.Printf("User clicked: %d\n", result)
}

func gui() {
	a := app.New()
	r, _ := fyne.LoadResourceFromPath("lock.png")
	a.SetIcon(r)
	w2 := a.NewWindow("687")
	w2.SetContent(widget.NewLabel("Find the password."))
	w2.Resize(fyne.NewSize(370, 600))
	w2.SetFixedSize(true)
	img := canvas.NewImageFromFile("444.png")
	input := widget.NewEntry()
	img.FillMode = canvas.ImageFillOriginal
	input.SetPlaceHolder("Enter Password")

	content := container.NewVBox(img, input, widget.NewButton("Save", func() {
		pw := input.Text
		foo := md5.Sum([]byte(pw))
		bar := hex.EncodeToString(foo[:])
		baz := strings.Replace(bar, "0", "L", -1)
		fooo := baz
		barr := strings.Replace(fooo, "8", "R", -1)
		var gum string = "1"
		xyzz := "1456bc0f6e0e4cf65c78a09688a2c920"
		ropers := strings.Replace(barr, "1", "", -1)
		buckaroo := strings.Replace(ropers, "L", "0", -1)
		riding := strings.Replace(buckaroo, "R", "8", -1)
		switch {
		case gum+riding == xyzz:
			w := a.NewWindow("CONGRATS!")
			imgg := canvas.NewImageFromFile("61.jpg")
			w.SetContent(imgg)
			w.Show()
			w.Resize(fyne.NewSize(500, 200))
		case gum+riding != xyzz:
			incoherence()
			time.Sleep(2 * time.Second)
		}
	}))

	w2.SetContent(content)
	w2.Show()
	a.Run()
}

func main() {
	vom()
	Ying()
	Yang()
	gui()
}
