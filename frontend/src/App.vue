<script setup lang="tsx">
import { backend, main } from '../wailsjs/go/models'
import {
	EventsEmit,
	EventsOn,
	OnFileDrop,
	OnFileDropOff
} from '../wailsjs/runtime/runtime'
import HoverCard from './components/ui/hover-card/HoverCard.vue'
import HoverCardContent from './components/ui/hover-card/HoverCardContent.vue'
import HoverCardTrigger from './components/ui/hover-card/HoverCardTrigger.vue'
import { cn } from './lib/utils'
import { onMounted, onUnmounted, ref } from 'vue'

const progress = ref<number>(0.0)
const cmdOutputs = ref<string[]>([])
EventsOn(main.OtherEmitID.Progress, (data: number) => {
	progress.value = data
})
EventsOn(main.OtherEmitID.Warning, (data: string) =>
	cmdOutputs.value.push(data)
)

const tasks = (
	[
		[
			backend.TaskID.Artefact,
			'Artefact',
			'Remove JPEG artifacts and output PNG.<br /><br />Accepts: <code>.jpg</code>'
		],
		[
			backend.TaskID.ArtefactAvif,
			'Artefact + AVIF (Lossy)',
			'Remove JPEG artifacts and output PNG, then compress to AVIF (lossy).<br /><br />Accepts: <code>.jpg</code>'
		],
		[
			backend.TaskID.CjxlLossless,
			'CJXL (Lossless)',
			'Compress JPG/PNG to JXL (lossless).<br /><br />Accepts: <code>.jpg</code>, <code>.png</code>'
		],
		[
			backend.TaskID.AvifLossy,
			'AVIF (Lossy)',
			'Compress JPG/PNG to JXL (lossy).<br /><br />Accepts: <code>.jpg</code>, <code>.png</code>'
		],
		[
			backend.TaskID.Djxl,
			'DJXL',
			'Decompress JXL to JPG/PNG.<br /><br />Accepts: <code>.jxl</code>'
		],
		[
			backend.TaskID.Par2,
			'PAR2',
			'Create parity files for 7z.<br /><br />Accepts: <code>.7z</code>'
		],
		[
			backend.TaskID.DifferDiff,
			'Differ diff',
			'Generate diff images sequence.<br /><br />Accepts: <code>.png</code>'
		],
		[
			backend.TaskID.DifferJoin,
			'Differ join',
			'Reconstruct image from diff images sequence.<br /><br />Accepts: <code>.png</code>'
		]
	] satisfies Array<[backend.TaskID, string, string]>
).map(([ID, Label, Description]) => ({
	ID,
	Label,
	Description,
	Bounds: { X: 0, Y: 0, Width: 0, Height: 0 }
}))

const resizeObserver = new ResizeObserver((entries) => {
	for (const entry of entries) {
		const { target } = entry
		const task = tasks.find(
			(task) => task.ID === (target.id as backend.TaskID)
		)
		if (task) {
			const { width, height, x, y } = target.getBoundingClientRect()
			task.Bounds = { X: x, Y: y, Width: width, Height: height }
		}
	}
})

onMounted(() => {
	document.querySelector('html')?.classList.add('dark')

	for (const task of tasks
		.map((task) => document.getElementById(task.ID))
		.filter((task) => task !== null)) {
		resizeObserver.observe(task)
	}
})

// this exists just to takeover webview's drag and drop event
OnFileDrop(() => {
	/** */
}, false)
EventsOn('wails:file-drop', (x: number, y: number, paths: string[]) => {
	const droppedOn = tasks.find((task) => {
		const { X, Y, Width, Height } = task.Bounds
		return x >= X && x <= X + Width && y >= Y && y <= Y + Height
	})
	if (droppedOn) {
		cmdOutputs.value = []
		EventsEmit('process', droppedOn.ID, paths.sort())
	}
})

const runningTask = ref<backend.TaskID | null>(null)
EventsOn(main.OtherEmitID.TaskStart, (taskID: backend.TaskID) => {
	runningTask.value = taskID
})
EventsOn(main.OtherEmitID.TaskDone, () => {
	runningTask.value = null
	progress.value = 0.0
})

onUnmounted(() => {
	OnFileDropOff()
	resizeObserver.disconnect()
})
</script>

<template>
	<div class="grid h-dvh grid-rows-[auto_auto_1fr]">
		<button
			class="top-0 inline-flex h-10 h-12 w-full items-center justify-center gap-2 bg-destructive px-8 text-xl font-medium whitespace-nowrap text-destructive-foreground shadow-xs transition-colors hover:bg-destructive/90 focus-visible:ring-1 focus-visible:ring-ring focus-visible:outline-hidden disabled:pointer-events-none disabled:opacity-50 [&_svg]:pointer-events-none [&_svg]:size-4 [&_svg]:shrink-0"
			:onclick="
				() => EventsEmit(main.OtherEmitID.CancelTask, runningTask)
			"
			:disabled="runningTask === null"
		>
			Cancel Task
		</button>

		<div class="relative grid grid-cols-3 grid-rows-2">
			<div
				class="absolute top-0 left-0 -z-10 size-full bg-secondary transition-transform"
				:style="{
					transform: `translateX(-${(1 - progress) * 100}%)`
				}"
			/>
			<div
				v-for="task in tasks"
				:key="task.ID"
				:id="task.ID"
				:class="
					cn(
						'flex h-full items-center justify-center border border-secondary-foreground/30 px-3 py-2 text-center text-base backdrop-blur-xs',
						runningTask === task.ID && 'font-black text-blue-400',
						runningTask !== null &&
							runningTask !== task.ID &&
							'font-extralight text-secondary-foreground/50'
					)
				"
			>
				<HoverCard :openDelay="500" :closeDelay="100">
					<HoverCardTrigger class="underline-offset-4 select-none">
						{{ task.Label }}
					</HoverCardTrigger>
					<HoverCardContent
						class="px-3 py-2 text-sm text-balance text-primary/80"
					>
						<span v-html="task.Description" />
					</HoverCardContent>
				</HoverCard>
			</div>
		</div>

		<div class="flex flex-col gap-2 overflow-y-auto">
			<code v-for="[index, output] in cmdOutputs.entries()" :key="index">
				{{ output }}
			</code>
		</div>
	</div>
</template>
