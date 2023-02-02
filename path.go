package main

const ReplayRoute = "/replay"

// for upload
const (
	BeginReplay   = "/{session_name}"
	UploadHeader  = "/{session_name}/file/{file_name}"
	AddEvent      = "/{session_name}/event"
	UpdateEvent   = "/{session_name}/event/{event_full_name}"
	StopStreaming = "/{session_name}/stopUploading"
	PostUser      = "/{session_name}/users"
)

// for download
const (
	RequestStartDownload  = "/{session_name}/startDownloading"
	RequestDownload       = "/{session_name}/file/{file_name}"
	RequestCheckpointData = "/{session_name}/event"
	RefreshViewer         = "/{session_name}/viewer/{viewer_name}"
)

// for search
const (
	SearchEvent        = "/event"
	SearchEventByNames = "/event/{event_name}"
)
