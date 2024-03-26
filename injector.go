package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"runtime"

	"github.com/google/uuid"
)

func main() {
	os := CheckOS()
	out, in := GetArgs()
	if os == "windows" {
		out += ".exe"
	}
	CheckArgs(out, in, os)
	bytes := ReadIn(in)
	Wrap(out)
	AddToWrap(out, bytes)
	log.Println("Done!")
}

func ReadIn(names []string) [][]byte {
	var fin [][]byte
	for _, name := range names {
		fin = append(fin, []byte("Ayy lmao"))
		data, err := ioutil.ReadFile(name)
		if err != nil {
			log.Fatal(errors.New("Failed to read file " + name + ":" + err.Error()))
		}
		fin = append(fin, data)
	}
	return fin
}

func OpenFile(name string) *os.File {
	file, err := os.OpenFile(name, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(errors.New("Failed to open file " + name + ":" + err.Error()))
	}
	return file
}

func WriteToFile(name string, file *os.File, data []byte) {
	_, err := file.Write(data)
	if err != nil {
		log.Fatal(errors.New("Failed to write to file " + name + ":" + err.Error()))
	}
}

func Wrap(name string) {
	wrap := uuid.NewString() + ".go"
	wrapF := OpenFile(wrap)
	wrapC := []byte(
		`package main

import (
	"os"
	"strings"
	"io/ioutil"
	"github.com/amenzhinsky/go-memexec"
	"fmt"
)

func main() {
	bytes, _ := ioutil.ReadFile(os.Args[0])
	xs := strings.Split(string(bytes), "Ayy lmao")
	for i := 2; i < len(xs); i++ {
		x, _ := memexec.New([]byte(xs[i]))
		defer x.Close()
		cmd := x.Command()
		r, _ := cmd.Output()
		fmt.Print(string(r))
	}
}`)
	WriteToFile(wrap, wrapF, wrapC)
	wrapF.Close()
	cmd := exec.Command("go", "build", "-o", "./"+name, wrap)
	err := cmd.Start()
	if err != nil {
		log.Fatal(errors.New("Failed to compile mother:" + err.Error()))
	}
	err = cmd.Wait()
	if err != nil {
		log.Fatal(errors.New("Failed to compile mother:" + err.Error()))
	}
	err = os.Remove(wrap)
	if err != nil {
		log.Fatal(errors.New("Failed to remove file " + name + ":" + err.Error()))
	}
}

func AddToWrap(name string, data [][]byte) {
	wrap := OpenFile(name)
	defer wrap.Close()
	for _, v := range data {
		WriteToFile(name, wrap, v)
	}
}

func CheckOS() string {
	switch runtime.GOOS {
	case "windows":
		return "windows"
	case "linux":
		return "linux"
	case "darwin":
		return "darwin"
	default:
		log.Fatal(errors.New("Unsupported OS: " + runtime.GOOS))
		return ""
	}
}

func GetArgs() (string, []string) {
	if len(os.Args) < 4 {
		log.Fatal(errors.New("Not enough arguments. Please provide a target name and at least one executable name."))
	}
	var args []string
	for i := 2; i < len(os.Args); i++ {
		args = append(args, os.Args[i])
	}
	return (os.Args[1]), args
}

func CheckArgs(out string, in []string, ops string) {
	for _, name := range in {
		if _, err := os.Stat(name); err != nil {
			log.Fatal(errors.New("File " + name + " does not exist."))
		}
	}
	if _, err := os.Stat(out); err == nil {
		log.Println("File " + out + " already exists. Overwrite? (y/n)")
		var yn string
		for yn != "y" && yn != "n" {
			fmt.Scanln(&yn)
		}
		if yn == "n" {
			log.Fatal(errors.New("Aborted."))
		}
	}
}
