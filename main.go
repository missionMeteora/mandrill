package mandrill

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"mime"
	"net/http"
	"path/filepath"
	"strings"
)

const (
	MANDRILL_LOCATION string = "https://mandrillapp.com/api/1.0/"
	SEND_LOCATION     string = "messages/send.json"
)

func New(apiKey, subAccount, fromEmail, fromName string) *Client {
	return &Client{
		apiKey:     apiKey,
		subAccount: subAccount,
		fromEmail:  fromEmail,
		fromName:   fromName,
	}
}

func (m *Client) APIKey() string {
	return m.apiKey
}

func (m *Client) SubAccount() string {
	return m.subAccount
}

func (m *Client) SendMessage(html, subject, toEmail, toName string, tags []string) ([]*SendResponse, error) {
	return m.SendMessageWithAttachments(html, subject, toEmail, toName, tags, nil)
}

func (m *Client) SendMessageWithAttachments(html, subject, toEmail, toName string,
	tags []string, attachments []*MessageAttachment) ([]*SendResponse, error) {
	requestData, err := getSendRequestData(m.apiKey, html, subject, m.fromEmail, m.fromName, toEmail,
		toName, m.subAccount, tags, attachments)
	if err != nil {
		return nil, err
	}

	response, err := sendRequest(SEND_LOCATION, requestData)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (m *Client) SendMessageWithReader(html, subject, toEmail, toName string,
	tags []string, fname string, r io.Reader) ([]*SendResponse, error) {
	att, err := AttachmentFromReader(fname, r)
	if err != nil {
		return nil, err
	}
	return m.SendMessageWithAttachments(html, subject, toEmail, toName, tags, []*MessageAttachment{att})
}

func AttachmentFromReader(fname string, r io.Reader) (*MessageAttachment, error) {
	var (
		buf = getBuffer()
		enc = base64.NewEncoder(base64.RawStdEncoding, buf)
	)
	defer putBuffer(buf)

	if _, err := io.Copy(enc, r); err != nil {
		return nil, err
	}
	enc.Close()

	return &MessageAttachment{
		Type:    mime.TypeByExtension(filepath.Ext(fname)),
		Name:    fname,
		Content: buf.String(),
	}, nil
}

func sendRequest(loc, requestData string) ([]*SendResponse, error) {
	resp, err := http.Post(MANDRILL_LOCATION+loc, "application/json", strings.NewReader(requestData))
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		var r struct {
			Code    int    `json:"code"`
			Message string `json:"message"`
		}
		if err := json.NewDecoder(resp.Body).Decode(&r); err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("error(%d): %s", r.Code, r.Message)

	}

	var s []*SendResponse
	if err := json.NewDecoder(resp.Body).Decode(&s); err != nil {
		return nil, err
	}

	return s, nil
}
