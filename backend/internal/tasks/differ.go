package tasks

import (
	"context"
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"
)

func Differ(
	ctx context.Context,
	files []string,
	isJoin bool,
	sendWarning func(error),
) {
	pngFiles := []string{}
	for _, file := range files {
		fileExt := strings.ToLower(filepath.Ext(file))
		if fileExt == ".png" {
			pngFiles = append(pngFiles, file)
		}
	}

	if len(pngFiles) == 0 {
		sendWarning(fmt.Errorf("no png files found"))
		return
	}

	kind := "-diff"
	if isJoin {
		kind = "-join"
	}

	args := []string{kind}
	args = append(args, pngFiles...)
	cmd := exec.CommandContext(ctx, "differ", args...)
	cmd.SysProcAttr = getSysProcAttr()
	outputMsgBytes, err := cmd.CombinedOutput()
	outputMsgString := string(outputMsgBytes)
	switch {
	case err != nil && outputMsgString != "":
		sendWarning(fmt.Errorf("differ error: %s", outputMsgString))
		return
	case err != nil && outputMsgString == "":
		sendWarning(fmt.Errorf("differ error: %w", err))
		return
	}
}
