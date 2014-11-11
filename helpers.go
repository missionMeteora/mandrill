package mandrill

import (
	"encoding/json"
)

func getMessageStruct(html, subject, fromEmail, fromName, toEmail, toName, subAccount string,
	tags []string, attachments []*MessageAttachment) *Message {
	return &Message{
		Html:               html,
		Subject:            subject,
		FromEmail:          fromEmail,
		FromName:           fromName,
		To:                 getMessageTo(toEmail, toName, "to"),
		Important:          false,
		TrackOpens:         false,
		TrackClicks:        false,
		AutoText:           true,
		AutoHtml:           true,
		InlineCss:          true,
		UrlStripQs:         false,
		PreserveRecipients: false,
		ViewContentLink:    false,
		Merge:              true,
		Tags:               tags,
		SubAccount:         "Meteora Dashboard",
		Attachments:        attachments,
	}
}

func getMessageTo(email, name, messageType string) []*MessageTo {
	var returnMap []*MessageTo
	messageTo := &MessageTo{
		Email: email,
		Name:  name,
		Type:  messageType,
	}
	returnMap = append(returnMap, messageTo)
	return returnMap
}

func getMessageGlobalMergeVars(name, content string) []*MessageMergeItem {
	var returnMap []*MessageMergeItem
	mergeItem := &MessageMergeItem{
		Name:    name,
		Content: content,
	}
	returnMap = append(returnMap, mergeItem)
	return returnMap
}

func getSendRequestData(apiKey, html, subject, fromEmail, fromName, toEmail, toName, subAccount string,
	tags []string, attachments []*MessageAttachment) (string, error) {
	request := &SendRequest{
		Key:     apiKey,
		Message: getMessageStruct(html, subject, fromEmail, fromName, toEmail, toName, subAccount, tags, attachments),
		Async:   false,
	}

	requestData, err := json.Marshal(request)
	if err != nil {
		return ``, err
	}

	return string(requestData), nil
}
