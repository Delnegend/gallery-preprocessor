package main

import (
	"context"
	"embed"
	"fmt"
	"gallery-preprocessor-go/backend"
	"sync"

	"github.com/gen2brain/beeep"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

type App struct {
	ctx context.Context
}

type OtherEmitID string

const (
	ProgressEmitID   OtherEmitID = "Progress"
	WarningEmitID    OtherEmitID = "Warning"
	CancelTaskEmitID OtherEmitID = "CancelTask"
	TaskDoneEmitID   OtherEmitID = "TaskDone"
	TaskStartEmitID  OtherEmitID = "TaskStart"
)

func main() {
	app := App{}

	// Create application with options
	err := wails.Run(&options.App{
		Title: "gallery-preprocessor-go",

		Width:         320,
		Height:        500,
		DisableResize: false,

		AlwaysOnTop:      true,
		AssetServer:      &assetserver.Options{Assets: assets},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.ctx = ctx
		},
		Bind: []any{&app},
		EnumBind: []any{
			backend.AllTasks,
			[]struct {
				Value  OtherEmitID
				TSName string
			}{
				{ProgressEmitID, string(ProgressEmitID)},
				{WarningEmitID, string(WarningEmitID)},
				{CancelTaskEmitID, string(CancelTaskEmitID)},
				{TaskDoneEmitID, string(TaskDoneEmitID)},
				{TaskStartEmitID, string(TaskStartEmitID)},
			},
		},
		DragAndDrop: &options.DragAndDrop{EnableFileDrop: true},
		OnDomReady: func(ctx context.Context) {
			progressChan := make(chan float64)
			warnChan := make(chan error)

			go func() {
				for progress := range progressChan {
					runtime.EventsEmit(ctx, string(ProgressEmitID), progress)
				}
			}()
			go func() {
				for warn := range warnChan {
					runtime.EventsEmit(ctx, string(WarningEmitID), warn.Error())
				}
			}()

			var taskCancel context.CancelFunc

			isProcessing := false
			var isProcessingMu sync.Mutex
			runtime.EventsOn(ctx, "process", func(data ...any) {
				if isProcessing {
					warnChan <- fmt.Errorf("cannot process while another task is running")
					return
				}

				isProcessingMu.Lock()
				defer isProcessingMu.Unlock()
				isProcessing = true
				defer func() { isProcessing = false }()

				var taskId backend.TaskID
				if parsed, ok := data[0].(string); ok {
					tasks := backend.AllTasks
					for _, task := range tasks {
						if string(task.Value) == parsed {
							taskId = task.Value
							break
						}
					}
					if taskId == "" {
						warnChan <- fmt.Errorf("unknown task %s", parsed)
						return
					}
				} else {
					warnChan <- fmt.Errorf("expect first argument of \"process\" event to be string, got %v", data[0])
					return
				}

				inputsAny, ok := data[1].([]any)
				if !ok {
					warnChan <- fmt.Errorf("expect second argument of \"process\" event to be []string, got %v", data[1])
					return
				}
				inputStrings := make([]string, len(inputsAny))
				for i, inputAny := range inputsAny {
					inputString, ok := inputAny.(string)
					if !ok {
						warnChan <- fmt.Errorf("expect string from frontend, got %v", inputAny)
						return
					}
					inputStrings[i] = inputString
				}

				runtime.EventsEmit(ctx, string(TaskStartEmitID), taskId)
				backend.PerformTask(
					ctx,
					backend.TaskInput{Inputs: inputStrings, TaskID: taskId}, progressChan,
					warnChan,
				)
				runtime.EventsEmit(ctx, string(TaskDoneEmitID), taskId)
				err := beeep.Notify("Gallery Preprocessor", fmt.Sprintf("Task %s finished", taskId), "")
				if err != nil {
					panic(err)
				}
			})

			runtime.EventsOn(ctx, string(CancelTaskEmitID), func(data ...any) {
				if taskCancel != nil {
					taskCancel()
				}
				runtime.EventsEmit(ctx, string(TaskDoneEmitID))
			})
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
