package metadata

import (
	"errors"
	"io"
	"regexp"
	"strconv"
	"time"

	"github.com/PuerkitoBio/goquery"
)

// CurseforgeReader reader struct for addon page
type CurseforgeReader struct {
	doc *goquery.Document
}

// NewCurseforgeReader create a reader for curseforge addon page
func NewCurseforgeReader(r io.Reader) (*CurseforgeReader, error) {
	doc, err := goquery.NewDocumentFromReader(r)
	if err != nil {
		return nil, err
	}

	return &CurseforgeReader{
		doc: doc,
	}, nil
}

// LastMod of the addon
func (r *CurseforgeReader) LastMod() (time.Time, error) {
	stdDate := r.doc.Find("div.infobox__content abbr.standard-date").First()
	epochStr, exists := stdDate.Attr("data-epoch")
	if !exists {
		return time.Time{}, errors.New("lastmod attribute not found")
	}
	epoch, err := strconv.ParseInt(epochStr, 10, 64)
	if err != nil {
		return time.Time{}, errors.New("lastmod could not convert epoch to int")
	}
	t := time.Unix(epoch, 0)
	return t, nil
}

// Description of he addon
func (r *CurseforgeReader) Description() (string, error) {
	descNode := r.doc.Find(`meta[property="og:description"]`).First()
	descContent, exists := descNode.Attr("content")
	if !exists {
		return "", errors.New("description attribute not found")
	}
	return descContent, nil
}

// Name of the addon
func (r *CurseforgeReader) Name() (string, error) {
	titleNode := r.doc.Find(`meta[property="og:title"]`).First()
	titleContent, exists := titleNode.Attr("content")
	if !exists {
		return "", errors.New("name attribute not found")
	}
	return titleContent, nil
}

// Upstream return the url of the project
func (r *CurseforgeReader) Upstream() (string, error) {
	upstreamNode := r.doc.Find("p.infobox__cta a").First()
	upstreamHref, exists := upstreamNode.Attr("href")
	if !exists {
		return "", errors.New("upstream attribute not found")
	}
	return upstreamHref, nil
}

// GameVersion Return the addon game version
func (r *CurseforgeReader) GameVersion() (string, error) {
	re := regexp.MustCompile(`\d{1,2}\.\d{1,2}.\d{1,2}`)
	spanText := r.doc.Find("span.stats--game-version").Text()
	return re.FindString(spanText), nil
}
