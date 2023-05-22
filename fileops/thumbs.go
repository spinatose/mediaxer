package fileops

import (
	"path"
	"time"

	"golang.org/x/exp/slices"
)

type Thumbnail struct {
	CreatedDate time.Time
	Name        string
	OriginPath   string 
	Size        int64
}

type Thumbnails struct {
	Items []Thumbnail
}

func NewThumbnails() *Thumbnails {
	return &Thumbnails{
		Items: []Thumbnail{},
	}
}

func (thumbs *Thumbnails) AddItem(thumb Thumbnail) {
	thumbs.Items = append(thumbs.Items, thumb)
}

func (thumbs *Thumbnails) ToString() string {
	returnString := ""

	for i, thumb := range thumbs.Items {
		if (i == 0) {
			returnString = path.Join(thumb.OriginPath, thumb.Name)
		} else {
			returnString += "\n" + path.Join(thumb.OriginPath, thumb.Name)
		}
	}

	return returnString
}

func (thumbs *Thumbnails) GetCreatedDays() []string {
	days := []string {}

	for _, thumb := range thumbs.Items {
		created := thumb.CreatedDate.Format(time.DateOnly)
		idx := slices.IndexFunc(days, func(d string) bool { return d == created })
		if (idx == -1) {
			days = append(days, created)
		}
	}

	return days 
}