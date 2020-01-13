package tool

import (
	"DevTool/global"
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"runtime"
	"time"
)

// open the specified uri via browser
func OpenURI(uri string) error {
	run, ok := global.Commands[runtime.GOOS]["actions"]["openBrowser"]["cmd"]
	if !ok {
		return fmt.Errorf("Error opening url on %s platform", runtime.GOOS)
	}

	args := global.Commands[runtime.GOOS]["actions"]["openBrowser"]["args"]
	cmd := exec.Command(run, args+uri)
	return cmd.Start()
}

// judge the given path if exists
func Exists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}

// judge the give path if a directory
func IsDir(path string) bool {
	s, err := os.Stat(path)
	if err != nil {
		return false
	}
	return s.IsDir()
}

// judge the give path if a file
func IsFile(path string) bool {
	return !IsDir(path)
}

// generate random string for the given length
func GenRandomString(length int) string {

	bytes := []byte(global.Source)
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < length; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}
