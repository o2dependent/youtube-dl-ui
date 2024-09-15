package main

import (
	"context"
	"io"
	"os"
	"os/exec"
	"strings"

	"github.com/kkdai/youtube/v2"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

type QualityInfo struct {
	Quality      string `json:"quality" ts_type:"string"`
	AudioQuality string `json:"audioQuality" ts_type:"string"`
	MimeType     string `json:"mimeType" ts_type:"string"`
}

type Info struct {
	Author      string             `json:"author" ts_type:"string"`
	Title       string             `json:"title" ts_type:"string"`
	Duration    string             `json:"duration" ts_type:"string"`
	QualityInfo []QualityInfo      `json:"qualityInfo" ts_type:"QualityInfo[]"`
	Thumbnails  youtube.Thumbnails `json:"thumbnails" ts_type:"{URL: string,Width: number,Height: number}[]"`
}

var dir string = "/"

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func (a *App) GetImportantInfo(videoUrl string) (Info, error) {
	videoID, err := youtube.ExtractVideoID(videoUrl)
	if err != nil {
		return Info{}, err
	}

	client := youtube.Client{}

	video, err := client.GetVideo(videoID)
	if err != nil {
		return Info{}, err
	}
	var qualityInfo []QualityInfo

	formats := video.Formats
	formats.Sort()

	for i := 0; i < len(video.Formats); i++ {
		f := formats[i]

		qualityInfo = append(qualityInfo, QualityInfo{
			Quality:      f.Quality,
			AudioQuality: f.AudioQuality,
			MimeType:     f.MimeType,
		})
	}

	info := Info{
		Author:      video.Author,
		Title:       video.Title,
		Duration:    video.Duration.String(),
		QualityInfo: qualityInfo,
		Thumbnails:  video.Thumbnails,
	}

	return info, err
}

func (a *App) GetDirectory() (string, bool) {
	_dir, err := runtime.OpenDirectoryDialog(a.ctx,
		runtime.OpenDialogOptions{DefaultDirectory: dir,
			Title:                "Select Directory to download",
			ShowHiddenFiles:      true,
			CanCreateDirectories: true,
		})

	dir = _dir

	if err != nil {
		return "", true
	}

	return dir, false
}

func (a *App) Download(_dir string, videoUrl string, quality string, audioQuality string, fileExt string) bool {
	if _dir != dir {
		return false
	}
	videoID, err := youtube.ExtractVideoID(videoUrl)
	if err != nil {
		panic(err)
	}

	client := youtube.Client{}

	video, err := client.GetVideo(videoID)

	if err != nil {
		panic(err)
	}

	videoFileName := "tmp-video"
	audioFileName := "tmp-audio"
	outputFileName := dir + "/" + video.Title + "." + fileExt

	// Video File
	formats := video.Formats
	formats = formats.Select(func(f youtube.Format) bool {
		return f.Quality == quality && strings.Contains(f.MimeType, "video/"+fileExt)
	})

	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	videoFile, err := os.CreateTemp("", videoFileName)
	if err != nil {
		panic(err)
	}
	defer os.Remove(videoFile.Name())

	_, err = io.Copy(videoFile, stream)
	if err != nil {
		panic(err)
	}

	// Audio File
	formats = video.Formats.WithAudioChannels() // only get videos with audio
	formats = formats.Select(func(f youtube.Format) bool {
		return strings.Contains(f.MimeType, "audio/"+fileExt)
	})

	audioStream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer audioStream.Close()

	audioFile, err := os.CreateTemp("", audioFileName)
	if err != nil {
		panic(err)
	}
	defer os.Remove(audioFile.Name())

	_, err = io.Copy(audioFile, audioStream)
	if err != nil {
		panic(err)
	}

	// Concat audio and video using ffmpeg
	cmd := exec.Command("ffmpeg", "-i", videoFile.Name(), "-i", audioFile.Name(), "-c:v", "copy", "-c:a", "aac", outputFileName)
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	return true
}
