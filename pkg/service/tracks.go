package service

import (
	"github.com/wellWINeo/MusicPlayerBackend"
	"github.com/wellWINeo/MusicPlayerBackend/pkg/repository"
	"os"
	"path"
	"strconv"

	ffmpeg "github.com/u2takey/ffmpeg-go"
)

type TracksService struct {
	repo repository.Tracks
}

func NewTracksService(repo repository.Tracks) *TracksService {
	return &TracksService{repo: repo}
}

func (t *TracksService) CreateTrack(userId int, track MusicPlayerBackend.Track) (int, error) {
	return t.repo.CreateTrack(userId, track)
}

func (t *TracksService) GetTrack(trackId int) (MusicPlayerBackend.Track, error) {
	return t.repo.GetTrack(trackId)
}

func (t *TracksService) UpdateTrack(trackId int, track MusicPlayerBackend.Track) error {
	return t.repo.UpdateTrack(trackId, track)
}

func (t *TracksService) DeleteTrack(trackId int) error {
	return t.repo.DeleteTrack(trackId)
}

func (t *TracksService) GetAllTracks(userId int) ([]MusicPlayerBackend.Track, error) {
	return t.repo.GetAllTracks(userId)
}

func (t *TracksService) GetAllTracksId(userId int) ([]int, error) {
	return t.repo.GetAllTracksId(userId)
}

func (t *TracksService) SetLike(trackId int) error {
	return t.repo.SetLike(trackId)
}

func (t *TracksService) GetAllLikes(userId int) ([]int, error) {
	return t.repo.GetAllLikes(userId)
}

func (t *TracksService) DownloadTrack(trackId int) ([]byte, error) {
	return t.repo.DownloadTrack(trackId)
}

func (t *TracksService) UploadTrack(trackId int, fileName, dataPath string) error {
	// create directory with name as trackId
	savePath := path.Join(dataPath, strconv.Itoa(trackId))
	_, err := os.Stat(savePath)

	// directory already exists
	if err == nil {
		// delete, to remove all contains file
		err = os.RemoveAll(savePath)
		if err != nil {
			return err
		}
	}

	// create new directory
	err = os.Mkdir(savePath, os.ModePerm)
	if err != nil {
		return err
	}

	// convert file
	// TODO move hardcoded settings
	err = ffmpeg.Input(fileName, nil).Output(path.Join(savePath, "%03d"),
		ffmpeg.KwArgs{
			"c:a":            "libmp3lame",
			"b:a":            "128k",
			"map":            "0:0",
			"f":              "segment",
			"segment_time":   10,
			"segment_list":   path.Join(savePath, "index.m3u8"),
			"segment_format": "mpegts",
		}).Run()
	if err != nil {
		return err
	}

	// remove temp file
	err = os.Remove(fileName)
	if err != nil {
		return err
	}

	// rename index file
	err = os.Rename(path.Join(savePath, "index.m3u8"),
		path.Join(savePath, "index"))
	if err != nil {
		return err
	}

	return nil
}
