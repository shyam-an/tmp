package main

import (
   "fmt"
   "io/ioutil"
   "os"
   "path"
   "strconv"
   "strings"
 )


func main() {
  base_path := "/sys/block/nbd"
  for i := 0; ; i++ {
    nbd_path := base_path + strconv.Itoa(i)
    fmt.Printf("checking %s\n", nbd_path)
    _, err := os.Lstat(nbd_path)
    if err != nil {
      break
    }
    pidBytes, err := ioutil.ReadFile(path.Join(nbd_path, "pid"))
    if err != nil {
      continue
    }

    cmdlineFile := path.Join("/proc", strings.TrimSpace(string(pidBytes)), "cmdline")
    fmt.Printf("reading %s\n", cmdlineFile)
    rawCmdline, err := ioutil.ReadFile(cmdlineFile)
    cmdLineArgs := strings.FieldsFunc(string(rawCmdline), func (r rune) bool {
      if (r == '\u0000') {
	return true
      }
      return false
    })
    if len(cmdLineArgs) < 3 {
      continue
    }

    if cmdLineArgs[0] != "rbd-nbd" || cmdLineArgs[1] != "map" {
    }
    image := cmdLineArgs[2]
    fmt.Println(image)

    // get the pool name out of the rest of the args
    pool := ""
    for n := 3; n < len(cmdLineArgs) - 1; n++ {
      if cmdLineArgs[n] == "--pool" {
	pool = cmdLineArgs[n + 1]
	break
      }
    }
    fmt.Println(pool)
  }
}
