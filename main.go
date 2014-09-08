package mandrill

import(
	"net/http"
	"strings"
	"encoding/json"
	"io/ioutil"
)

const (
	MANDRILL_LOCATION string = "https://mandrillapp.com/api/1.0/"
	SEND_LOCATION string = "messages/send.json"
)

func New(apiKey, subAccount, fromEmail, fromName string) *Client {
	return &Client{
		apiKey: apiKey,
		subAccount: subAccount,
		fromEmail: fromEmail,
		fromName: fromName,
	}
}

func (m *Client) SendMessage(html, subject, toEmail, toName string, tags []string) ([]*SendResponse, error) {
	requestData, err := getSendRequestData(m.apiKey, html, subject, m.fromEmail, m.fromName, toEmail, toName, m.subAccount, tags)
	if err != nil {
		return nil, err
	}

	response, err := sendRequest(SEND_LOCATION, requestData);
	if err != nil {
		return nil, err
	}

	return response, nil
}

func sendRequest(loc, requestData string) ([]*SendResponse, error) {
	var sendResponse []*SendResponse
	resp, err := http.Post(MANDRILL_LOCATION + loc, "application/json", strings.NewReader(requestData))
	if err != nil {
		return nil, err
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		return nil, err
	}

	if err := json.Unmarshal(responseData, &sendResponse); err != nil {
		return nil, err
	}

	return sendResponse, nil 
}

