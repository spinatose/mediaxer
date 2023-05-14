package fileops

import (
	"path"
	"time"
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