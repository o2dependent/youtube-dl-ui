package main

import (
	"context"
	"fmt"
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
	fmt.Println(video.Author)

	if err != nil {
		panic(err)
	}

	videoFileName := dir + "/" + video.Title + "-tmp-video." + fileExt
	audioFileName := dir + "/" + video.Title + "-tmp-audio." + fileExt
	outputFileName := dir + "/" + video.Title + "." + fileExt

	// Video File
	fmt.Println("---- VIDEO ONLY ----")
	formats := video.Formats
	formats = formats.Select(func(f youtube.Format) bool {
		fmt.Println(f.AudioQuality)
		fmt.Println(f.AudioSampleRate)
		fmt.Println(f.Quality)
		fmt.Println(f.MimeType)

		return f.Quality == quality && strings.Contains(f.MimeType, "video/"+fileExt)
	})
	fmt.Println(formats)

	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	file, err := os.Create(videoFileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = io.Copy(file, stream)
	if err != nil {
		panic(err)
	}

	// Audio File
	fmt.Println("---- AUDIO ONLY ----")
	formats = video.Formats.WithAudioChannels() // only get videos with audio
	formats = formats.Select(func(f youtube.Format) bool {
		fmt.Println(f.AudioQuality)
		fmt.Println(f.AudioSampleRate)
		fmt.Println(f.Quality)
		fmt.Println(f.MimeType)

		return strings.Contains(f.MimeType, "audio/"+fileExt)
	})

	audioStream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer audioStream.Close()

	audioFile, err := os.Create(audioFileName)
	if err != nil {
		panic(err)
	}
	defer audioFile.Close()

	_, err = io.Copy(audioFile, audioStream)
	if err != nil {
		panic(err)
	}

	// Concat audio and video using ffmpeg
	cmd := exec.Command("ffmpeg", "-i", videoFileName, "-i", audioFileName, "-c:v", "copy", "-c:a", "aac", outputFileName)
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	// Delete video and audio file
	os.Remove(videoFileName)
	os.Remove(audioFileName)

	return true
}
