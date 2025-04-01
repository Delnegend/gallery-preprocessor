export namespace backend {
	
	export enum TaskID {
	    Artefact = "Artefact",
	    ArtefactAvif = "ArtefactAvif",
	    CjxlLossless = "CjxlLossless",
	    AvifLossy = "AvifLossy",
	    Djxl = "Djxl",
	    Par2 = "Par2",
	}

}

export namespace main {
	
	export enum OtherEmitID {
	    Progress = "Progress",
	    Warning = "Warning",
	    CancelTask = "CancelTask",
	    TaskDone = "TaskDone",
	    TaskStart = "TaskStart",
	}

}

