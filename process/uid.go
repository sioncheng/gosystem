package main

import (
	"fmt"
	"os"
	"os/user"
	"strconv"
)

func main() {
	uid := os.Getuid()
	u, err := user.LookupId(strconv.Itoa(uid))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("User: %s (uid %s)\n", u.Username, u.Uid)

	gid := os.Getgid()
	g, err := user.LookupGroupId(strconv.Itoa(gid))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Group: %s (gid %s)\n", g.Name, g.Gid)
}
