// +build !windows

package fsutil

import (
	"fmt"
	"os/exec"
)

func CopyFile(dst, src string) error {
	b, err := exec.Command("/bin/cp", "-n", src, dst).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, string(b))
	}
	return nil
}
