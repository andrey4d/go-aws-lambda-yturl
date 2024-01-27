/*
 *   Copyright (c) 2024 Andrey andrey4d.dev@gmail.com
 *   All rights reserved.
 */
package ytclient

import (
	"github.com/kkdai/youtube/v2"
)

type Video struct {
	ID          string   `json:"id,omitempty"`
	Title       string   `json:"title,omitempty"`
	Author      string   `json:"author,omitempty"`
	Description string   `json:"description,omitempty"`
	Formats     []Format `json:"formats,omitempty"`
}

type Format struct {
	URL           string `json:"url,omitempty"`
	MimeType      string `json:"mime_type,omitempty"`
	ContentLength int64  `json:"content_length,string,omitempty"`
	QualityLabel  string `json:"quality_label,omitempty"`
	AudioChannels int    `json:"audio_channels,omitempty"`
	AudioQuality  string `json:"audio_quality,omitempty"`
}

type Client struct {
	ytClient youtube.Client
	Video    youtube.Video
}

func New() *Client {
	return &Client{
		ytClient: youtube.Client{},
		Video:    youtube.Video{},
	}
}

func (c *Client) SetVideoInfo(url string) error {
	yt_video, err := c.ytClient.GetVideo(url)
	if err != nil {
		return err
	}
	c.Video = *yt_video

	return nil
}
