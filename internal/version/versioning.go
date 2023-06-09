package version

import "fmt"

type Versioning struct {
	Major, Minor, Patch uint
}

func (v Versioning) String() string {
	return fmt.Sprintf("%d.%d.%d", v.Major, v.Minor, v.Patch)
}
