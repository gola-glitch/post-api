package utils

import (
	"context"
	"github.com/gola-glitch/gola-utils/logging"
	"post-api/story/models"
	"strings"
)

func CountImageReadTime(imageCount int, readTime *int) {
	imageReadStartsAt := 12
	if imageCount > 10 {
		for i := imageCount; i > 0; i-- {
			if i < 10 {
				*readTime = *readTime + 3
				continue
			}
			if imageReadStartsAt < 3 {
				*readTime = *readTime + 3
				continue
			}
			*readTime = *readTime + imageReadStartsAt
			imageReadStartsAt--
		}
		return
	}
	for i := imageCount; i > 0; i-- {
		*readTime = *readTime + imageReadStartsAt
		imageReadStartsAt--
	}
}

func CountContentReadTime(contentWordsCount int) int {
	return int((0.0036 * float64(contentWordsCount)) * 60)
}

func GetNumberOfWords(content models.JSONString, wordsCount *int, ctx context.Context, imageCount *int, extractedTagline, titleString, previewImage *string) error {
	logger := logging.GetLogger(ctx).WithField("class", "StoryUtils").WithField("method", "GetNumberOfWords")
	var postData []interface{}
	err := content.Unmarshal(&postData)

	if err != nil {
		logger.Errorf("something went wrong %v", err)
		return err
	}

	for _, data := range postData {
		singleData := data.(map[string]interface{})
		value := singleData["type"]
		if value != nil {
			if value == "image" {
				if singleData["url"] != "" && *previewImage == "" {
					*previewImage = singleData["url"].(string)
				}
				*imageCount = *imageCount + 1
			}
		}
		singleChildren := singleData["children"].([]interface{})
		for _, childrenData := range singleChildren {
			data := childrenData.(map[string]interface{})
			textString := data["text"].(string)
			if textString != "" {
				individual := strings.Split(textString, " ")
				*wordsCount = len(individual) + *wordsCount
				if *titleString == "" {
					if len([]rune(textString)) > 100 {
						*titleString = string([]rune(textString)[:100])
						continue
					}
					*titleString = textString
				} else if *extractedTagline == "" {
					if len(textString) > 100 {
						tamilRunes := string([]rune(textString))
						*extractedTagline = tamilRunes[:100]
						continue
					}
					*extractedTagline = textString
				}
			}
		}
	}
	return nil
}

func GetTitleFromSlateJson(ctx context.Context, titleJson models.JSONString) (string, error) {
	logger := logging.GetLogger(ctx).WithField("class", "StoryUtils").WithField("method", "GetNumberOfWords")
	var postData []interface{}
	err := titleJson.Unmarshal(&postData)

	if err != nil {
		logger.Errorf("Error occurred while unmarshalling title text from slate json %v", err)
		return "", err
	}

	var titleString string
	for _, data := range postData {
		singleData := data.(map[string]interface{})
		singleChildren := singleData["children"].([]interface{})
		for _, childrenData := range singleChildren {
			data := childrenData.(map[string]interface{})
			textString := data["text"].(string)
			if textString != "" {
				if len(textString) > 100 {
					titleString = string([]rune(textString)[:100])
					break
				}
				titleString = textString
				break
			}
		}
	}

	return titleString, nil
}
