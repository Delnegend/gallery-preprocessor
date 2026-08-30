export namespace backend {
	export enum TaskID {
		Artefact = 'Artefact',
		ArtefactAvif = 'ArtefactAvif',
		AvifLossy = 'AvifLossy',
		CjxlLossless = 'CjxlLossless',
		DifferDiff = 'DifferDiff',
		DifferJoin = 'DifferJoin',
		Djxl = 'Djxl',
		Par2 = 'Par2'
	}
}

export namespace main {
	export enum OtherEmitID {
		CancelTask = 'CancelTask',
		Progress = 'Progress',
		TaskDone = 'TaskDone',
		TaskStart = 'TaskStart',
		Warning = 'Warning'
	}
}
