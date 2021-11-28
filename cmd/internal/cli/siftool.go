// Copyright (c) 2021 Apptainer a Series of LF Projects LLC
//   For website terms of use, trademark policy, privacy policy and other
//   project policies see https://lfprojects.org/policies
// Copyright (c) 2019, Sylabs Inc. All rights reserved.
// This software is licensed under a 3-clause BSD license. Please consult the
// LICENSE.md file distributed with the sources of this project regarding your
// rights to use or distribute this software.

package cli

import (
	"github.com/hpcng/sif/pkg/siftool"
	"github.com/hpcng/singularity/pkg/cmdline"
)

// SiftoolCmd is easily set since the sif repo allows the cobra.Command struct to be
// easily accessed with Siftool(), we do not need to do anything but call that function.
var SiftoolCmd = siftool.Siftool()

func init() {
	addCmdInit(func(cmdManager *cmdline.CommandManager) {
		cmdManager.RegisterCmd(SiftoolCmd)
	})
}
