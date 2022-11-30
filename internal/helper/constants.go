package helper

import (
	"path/filepath"
	"runtime"
)

const FAILEDGETDATA = "Failed to GET data"
const SUCCEEDGETDATA = "Succeed to GET data"
const FAILEDPOSTDATA = "Failed to POST data"
const SUCCEEDPOSTDATA = "Succeed to POST data"
const FAILEDUPDATEDATA = "Failed to UPDATE data"
const SUCCEEDUPDATEDATA = "Succeed to UPDATE data"

var (
	// Get current file full path from runtime
	_, b, _, _ = runtime.Caller(0)

	// Root folder of this project
	ProjectRootPath = filepath.Join(filepath.Dir(b), "../../")
)
