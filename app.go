package main

import (
	"context"
	_ "embed"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	goRuntime "runtime"
	"strings"

	"github.com/kkdai/youtube/v2"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed embeds/ffmpeg
var ffmpegBinary []byte

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

func (a *App) GetDirectory() string {
	_dir, err := runtime.OpenDirectoryDialog(a.ctx,
		runtime.OpenDialogOptions{DefaultDirectory: dir,
			Title:                "Select Directory to download",
			ShowHiddenFiles:      true,
			CanCreateDirectories: true,
		})

	if err != nil || _dir == "" {
		return dir
	}

	dir = _dir

	return dir
}

func getFFmpegPath() (string, error) {
	// Determine OS-specific filename
	filename := "ffmpeg"
	if goRuntime.GOOS == "windows" {
		filename += ".exe"
	}

	// Create a temporary directory
	tempDir := os.TempDir()
	ffmpegPath := filepath.Join(tempDir, filename)

	// Write embedded binary to temp file
	if err := os.WriteFile(ffmpegPath, ffmpegBinary, 0755); err != nil {
		return "", fmt.Errorf("failed to write FFmpeg: %v", err)
	}

	// Set executable permissions (non-Windows)
	if goRuntime.GOOS != "windows" {
		if err := os.Chmod(ffmpegPath, 0755); err != nil {
			return "", fmt.Errorf("failed to set permissions: %v", err)
		}
	}

	return ffmpegPath, nil
}

func downloadAudioOnly(client youtube.Client, video *youtube.Video, audioQuality string, fileExt string, filePath string) bool {
	// Audio File
	formats := video.Formats.WithAudioChannels() // only get videos with audio
	formats = formats.Select(func(f youtube.Format) bool {
		return f.AudioQuality == audioQuality && strings.Contains(f.MimeType, "audio/"+fileExt)
	})

	audioStream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer audioStream.Close()

	audioFile, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer audioFile.Close()

	_, err = io.Copy(audioFile, audioStream)
	if err != nil {
		panic(err)
	}

	return true
}

func downloadVideoOnly(client youtube.Client, video *youtube.Video, quality string, fileExt string, filePath string) bool {

	formats := video.Formats
	formats = formats.Select(func(f youtube.Format) bool {
		return f.Quality == quality && strings.Contains(f.MimeType, "video/"+fileExt)
	})

	stream, _, err := client.GetStream(video, &formats[0])
	if err != nil {
		panic(err)
	}
	defer stream.Close()

	videoFile, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}
	defer videoFile.Close()

	_, err = io.Copy(videoFile, stream)
	if err != nil {
		panic(err)
	}

	return true
}

func downloadAudioVideo(client youtube.Client, video *youtube.Video, quality string, audioQuality string, fileExt string, filePath string) bool {

	videoFileName := "tmp-video"
	audioFileName := "tmp-audio"

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
		return f.AudioQuality == audioQuality && strings.Contains(f.MimeType, "audio/"+fileExt)
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
	ffmpegPath, err := getFFmpegPath()
	if err != nil {
		panic(err)
	}
	cmd := exec.Command(ffmpegPath, "-i", videoFile.Name(), "-i", audioFile.Name(), "-c:v", "copy", "-c:a", "aac", filePath)
	if err := cmd.Run(); err != nil {
		panic(err)
	}

	return true
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

	fmt.Println(videoID)
	video, err := client.GetVideo(videoID)
	fmt.Println(video)
	if err != nil {
		panic(err)
	}

	filePath, err := getUniqueFileName(dir, video.Title+"."+fileExt)
	if err != nil {
		panic(err)
	}

	if quality != "" && audioQuality == "" { // VIDEO ONLY
		return downloadVideoOnly(client, video, quality, fileExt, filePath)
	} else if quality == "" && audioQuality != "" { // AUDIO ONLY
		return downloadAudioOnly(client, video, audioQuality, fileExt, filePath)
	} else if quality != "" && audioQuality != "" { // AUDIO & VIDEO
		return downloadAudioVideo(client, video, quality, audioQuality, fileExt, filePath)
	} else {
		return false
	}
}

/* ---- INSTALL FFMPEG ---- */

func (a *App) CheckFFMPEG() bool {
	ffmpegPath, err := getFFmpegPath()
	if err != nil {
		panic(err)
	}
	if err := exec.Command(ffmpegPath, "-version").Run(); err != nil {
		return false
	}

	return true
}
