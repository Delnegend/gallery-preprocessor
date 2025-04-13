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

func ArtefactAvif(
	ctx context.Context,
	files []string,
	poolSize int,
	updateProgressBase func(func() float64) func(),
	sendWarning func(error),
) {
	jpgFiles := []string{}
	for _, entry := range files {
		if strings.ToLower(filepath.Ext(entry)) == ".jpg" {
			jpgFiles = append(jpgFiles, entry)
		}
	}

	for _, inputJpgFile := range jpgFiles {
		fmt.Println("processing", inputJpgFile)
	}

	if len(jpgFiles) == 0 {
		sendWarning(fmt.Errorf("no jpg files found"))
		return
	}

	// output file already exists
	for _, inputJpgFile := range jpgFiles {
		outputAvifFile := utils.ReplaceExt(inputJpgFile, ".avif")
		if _, err := os.Stat(outputAvifFile); err == nil {
			sendWarning(fmt.Errorf("possible output file '%s' already exists", outputAvifFile))
			return
		}
	}

	processedFiles := 0
	var progressMutex sync.Mutex
	updateProgress := updateProgressBase(func() float64 {
		progressMutex.Lock()
		defer progressMutex.Unlock()
		processedFiles++
		return float64(processedFiles) / float64(len(jpgFiles))
	})

	tmpDir, err := os.MkdirTemp("", "artefact-jxl-tmp")
	if err != nil {
		sendWarning(fmt.Errorf("can't create temp dir: %w", err))
		return
	}

	pool := utils.NewWorkerPool(ctx, poolSize)
	defer pool.WaitAndClose()

	for i, inputJpgFile := range jpgFiles {
		pool.Run(func() {
			defer updateProgress()

			outputPngFile := filepath.Join(tmpDir, fmt.Sprintf("%d.png", i))
			outputAvifFile := utils.ReplaceExt(inputJpgFile, ".avif")

			// jpg --artefact--> png
			cmd := exec.CommandContext(ctx, "artefact-cli", inputJpgFile, "-o", outputPngFile, "-i", "50")
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			outputMsgBytes, err := cmd.CombinedOutput()
			outputMsgString := string(outputMsgBytes)
			switch {
			case err != nil && outputMsgString != "":
				sendWarning(fmt.Errorf("artefact error: %s", outputMsgString))
				return
			case err != nil && outputMsgString == "":
				sendWarning(fmt.Errorf("artefact error: %w", err))
				return
			}

			// png -> avif
			cmd = exec.CommandContext(ctx, "ffmpeg", "-i", outputPngFile, "-c:v", "libsvtav1", "-crf", "22", "-preset", "2", "-pix_fmt", "yuv420p10le", "-vf", "scale=ceil(iw/2)*2:ceil(ih/2)*2", "-svtav1-params", "avif=1", outputAvifFile)
			cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
			outputMsgBytes, err = cmd.CombinedOutput()
			outputMsgString = string(outputMsgBytes)
			switch {
			case err != nil && outputMsgString != "":
				sendWarning(fmt.Errorf("ffmpeg error: %s", outputMsgString))
				return
			case err != nil && outputMsgString == "":
				sendWarning(fmt.Errorf("ffmpeg error: %w", err))
				return
			}

			// check if output exists
			if _, err := os.Stat(outputAvifFile); err != nil {
				if errors.Is(err, os.ErrNotExist) {
					sendWarning(fmt.Errorf("output file '%s' does not exist", outputAvifFile))
					return
				}
				sendWarning(fmt.Errorf("can't check output file '%s': %w", outputAvifFile, err))
			}
		})
	}

	if err := os.RemoveAll(tmpDir); err != nil {
		sendWarning(fmt.Errorf("can't remove temp dir: %w", err))
	}
}
