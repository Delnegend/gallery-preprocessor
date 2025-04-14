package backend

import (
	"context"
	"fmt"
	"gallery-preprocessor-go/backend/internal/tasks"
	"os"
	"path/filepath"
	"sync"
)

type TaskID string

var (
	TaskArtefact     TaskID = "Artefact"
	TaskArtefactAvif TaskID = "ArtefactAvif"
	TaskCjxlLossLess TaskID = "CjxlLossless"
	TaskAvifLossy    TaskID = "AvifLossy"
	TaskDjxl         TaskID = "Djxl"
	TaskPar2         TaskID = "Par2"
	TaskDifferDiff   TaskID = "DifferDiff"
	TaskDifferJoin   TaskID = "DifferJoin"

	AllTasks = []struct {
		Value  TaskID
		TSName string
	}{
		{TaskArtefact, string(TaskArtefact)},
		{TaskArtefactAvif, string(TaskArtefactAvif)},
		{TaskCjxlLossLess, string(TaskCjxlLossLess)},
		{TaskAvifLossy, string(TaskAvifLossy)},
		{TaskDjxl, string(TaskDjxl)},
		{TaskPar2, string(TaskPar2)},
		{TaskDifferDiff, string(TaskDifferDiff)},
		{TaskDifferJoin, string(TaskDifferJoin)},
	}
)

type TaskInput struct {
	TaskID TaskID
	Inputs []string
}

func PerformTask(taskCtx context.Context, taskInput TaskInput, progressChan chan<- float64, warnChan chan<- error) {
	var taskMutex sync.Mutex
	taskMutex.Lock()

	updateProgressBase := func(f func() float64) func() {
		return func() { go func() { progressChan <- f() }() }
	}
	sendWarning := func(err error) { go func() { warnChan <- err }() }

	files := []string{}
	for _, input := range taskInput.Inputs {
		info, err := os.Stat(input)
		if err != nil {
			sendWarning(fmt.Errorf("can't read input file info: %w", err))
			continue
		}
		if info.IsDir() {
			entries, err := os.ReadDir(input)
			if err != nil {
				sendWarning(fmt.Errorf("can't read input directory: %w", err))
				continue
			}
			for _, entry2 := range entries {
				if entry2.IsDir() {
					continue
				}
				files = append(files, filepath.Join(input, entry2.Name()))
			}
			continue
		}
		files = append(files, input)
	}

	switch taskInput.TaskID {
	case TaskArtefact:
		tasks.Artefact(taskCtx, files, 2, updateProgressBase, sendWarning)
	case TaskArtefactAvif:
		tasks.ArtefactAvif(taskCtx, files, 2, updateProgressBase, sendWarning)
	case TaskCjxlLossLess:
		tasks.Cjxl(taskCtx, files, 2, false, updateProgressBase, sendWarning)
	case TaskAvifLossy:
		tasks.AvifLossy(taskCtx, files, 2, updateProgressBase, sendWarning)
	case TaskDjxl:
		tasks.Djxl(taskCtx, files, 2, updateProgressBase, sendWarning)
	case TaskPar2:
		tasks.Par2(taskCtx, files, 2, updateProgressBase, sendWarning)
	case TaskDifferDiff:
		tasks.Differ(taskCtx, files, false, sendWarning)
	case TaskDifferJoin:
		tasks.Differ(taskCtx, files, true, sendWarning)
	default:
		sendWarning(fmt.Errorf("internal error: unknown task %s", taskInput.TaskID))
	}
}
