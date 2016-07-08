// +build !windows

package fsutil

import (
	"fmt"
	"os/exec"
)

func MoveFile(dst, src string) error {
	b, err := exec.Command("/bin/mv", "-n", src, dst).CombinedOutput()
	if err != nil {
		return fmt.Errorf("%s: %s", err, string(b))
	}
	return nil
}
