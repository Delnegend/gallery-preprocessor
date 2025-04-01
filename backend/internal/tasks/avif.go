package tasks

import (
	"context"
	"errors"
	"fmt"
	"gallery-preprocessor-go/backend/internal/utils"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
	"syscall"
)

func AvifLossy(
	ctx context.Context,
	files []string,
	poolSize int,
	updateProgressBase func(func() float64) func(),
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

	for _, inputFile := range pngFiles {
		outputFile := utils.ReplaceExt(inputFile, ".avif")
		if _, err := os.Stat(outputFile); err == nil {
			sendWarning(fmt.Errorf("possible output file '%s' already exists", outputFile))
			return
		}
	}

	processedFiles := 0
	var progressMutex sync.Mutex
	updateProgress := updateProgressBase(func() float64 {
		progressMutex.Lock()
		defer progressMutex.Unlock()
		processedFiles++
		return float64(processedFiles) / float64(len(pngFiles))
	})

	pool := utils.NewWorkerPool(ctx, poolSize)

	for _, inputFile := range pngFiles {
		pool.Run(func() {
			defer updateProgress()
			outputFile := utils.ReplaceExt(inputFile, ".avif")

			cmd := exec.CommandContext(ctx, "ffmpeg", "-i", inputFile, "-c:v", "libsvtav1", "-crf", "22", "-preset", "2", "-pix_fmt", "yuv420p10le", "-vf", "scale=ceil(iw/2)*2:ceil(ih/2)*2", "-svtav1-params", "avif=1", outputFile)
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			outputMsgBytes, err := cmd.CombinedOutput()
			outputMsgString := string(outputMsgBytes)
			switch {
			case err != nil && outputMsgString != "":
				sendWarning(fmt.Errorf("ffmpeg error: %s", outputMsgString))
				return
			case err != nil && outputMsgString == "":
				sendWarning(fmt.Errorf("ffmpeg error: %w", err))
				return
			}

			// check output file exists
			_, err = os.Stat(outputFile)
			if errors.Is(err, os.ErrNotExist) {
				sendWarning(fmt.Errorf("output file '%s' not created", outputFile))
			} else if err != nil {
				sendWarning(fmt.Errorf("can't check if output file exists: %w", err))
			}
		})
	}
}
