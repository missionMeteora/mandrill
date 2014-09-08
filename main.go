package main

import(
	"net/http"
	"log"
	"strings"
	"encoding/json"
	"io/ioutil"
)

const (
	API_KEY string = "Pc4PtpXLpGQr0iHjzUPCjg"
	//SEND_LOCATION string = "https://mandrillapp.com/api/1.0/messages/send.json"
)

func main(){
	log.Println("Word!")
	SendMessage()
	log.Println("Done!")

}

func SendMessage() error {
	request := &SendRequest{
		Key: 		API_KEY,			
		Message: 	&Message{
						Html: 				"<p>Hello!</p>",
						Text: 				"Hello!",
						Subject: 			"Test Subject",
						FromEmail: 			"test@meteora.co",
						FromName: 			"Mr. Test",
						To: 				getMessageTo("josh@meteora.co", "Josh", "to"),
						Important: 			false,
						TrackOpens: 		false,
						TrackClicks: 		false,
						AutoText: 			true,
						AutoHtml: 			true,
						InlineCss: 			true,
						UrlStripQs: 		false,
						PreserveRecipients: false,
						ViewContentLink: 	false, 
						Merge: 				true,
						Tags: 				[]string {"Test"},
						SubAccount: 		"customer-123",
					},
		Async: 		false,
	}
	
	requestJson, err := json.Marshal(request)
	if err != nil {
		log.Println("ERROR", err)
		return err
	}

	resp, err := http.Post("https://mandrillapp.com/api/1.0/messages/send.json", "application/json", strings.NewReader(string(requestJson)))
	if err != nil {
		log.Println("ERROR", err)
		return err
	}

/*
	if resp.StatusCode != 200 {
		return nil
	}
*/

	responseData, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		log.Println("ERROR", err)
		return err
	}

	log.Println(string(responseData))
	return nil
}

func getMessageTo(email, name, messageType string) []*MessageTo {
	var returnMap []*MessageTo 
	messageTo := &MessageTo{
		Email: 	email,
		Name: 	name,
		Type: 	messageType,
	}
	returnMap = append(returnMap, messageTo)
	return returnMap
}

func getMessageGlobalMergeVars(name, content string) []*MessageMergeItem {
	var returnMap []*MessageMergeItem 
	mergeItem := &MessageMergeItem{
		Name: 	name,
		Content: content,
	}
	returnMap = append(returnMap, mergeItem)
	return returnMap
}




///////




/**
 *	SendRequest struct
 *	==============
 *	Key 		Api key
 *	Message 	The information on the message to send
 *	Async 		Enable a background sending mode that is optimized for bulk sending. In async mode, messages/send will immediately return a status 
 *				of "queued" for every recipient. To handle rejections when sending in async mode, set up a webhook for the 'reject' event. 
 *				Defaults to false for messages with no more than 10 recipients; messages with more than 10 recipients are always sent asynchronously, 
 *				regardless of the value of async.
 *	IpPool 		The name of the dedicated ip pool that should be used to send the message. If you do not have any dedicated IPs,
 *				this parameter has no effect. If you specify a pool that does not exist, your default pool will be used instead.
 *	SendAt 		When this message should be sent as a UTC timestamp in YYYY-MM-DD HH:MM:SS format. If you specify a time in the past,
 *				the message will be sent immediately. An additional fee applies for scheduled email, and this feature is only available to accounts
 *				with a positive balance.
 */
type SendRequest struct {
	Key 		string 		`json:"key"`
	Message 	*Message 	`json:"message"`
	Async 		bool 		`json:"async"`
	IpPool 		string 		`json:"ip_pool"` 
	SendAt 		string 		`json:"send_at"`
}

/**
 *	Message struct
 *	==============		
 *	Html 						The full HTML content to be sent
 *	Text 						Optional full text content to be sent
 *	Subject 					The message subject
 *	FromEmail 					The sender email address
 *	FromName  					Optional from name to be used
 *	To 							An array of recipient information
 *	Headers 					Optional extra headers to add to the message (most headers are allowed)
 *	Important 					Whether or not this message is important, and should be delivered ahead of non-important messages
 *	TrackOpens 					Whether or not to turn on open tracking for the message
 *	TrackClicks 				Whether or not to turn on click tracking for the message
 *	AutoText 					Whether or not to automatically generate a text part for messages that are not given text
 *	AutoHtml 					Whether or not to automatically generate an HTML part for messages that are not given HTML
 *	InlineCss 					Whether or not to automatically inline all CSS styles provided in the message HTML - only for HTML documents less than 256KB in size
 *	UrlStripQs 					Whether or not to strip the query string from URLs when aggregating tracked URL data
 *	PreserveRecipients 			Whether or not to expose all recipients in to "To" header for each email
 *	ViewContentLink 			Set to false to remove content logging for sensitive emails
 *	BccAddress 					An optional address to receive an exact copy of each recipient's email
 *	TrackingDomain 				A custom domain to use for tracking opens and clicks instead of mandrillapp.com
 *	SigningDomain 				A custom domain to use for SPF/DKIM signing instead of mandrill (for "via" or "on behalf of" in email clients)
 *	ReturnPathDomain 			A custom domain to use for the messages's return-path
 *	Merge 						Whether to evaluate merge tags in the message. Will automatically be set to true if either merge_vars or global_merge_vars are provided.
 *	GlobalMergeVars 			Global merge variables to use for all recipients. You can override these per recipient.
 *	MergeVars 					Per-recipient merge variables, which override global merge variables with the same name.
 *	Tags 						An array of string to tag the message with. Stats are accumulated using tags. (Do not start tag with an underscore, reserved for Manrill internal use)
 *	SubAccount 					The unique id of a subaccount for this message - must already exist or will fail with an error
 *	GoogleAnalyticsDomains 		An array of strings indicating for which any matching URLs will automatically have Google Analytics parameters appended to their query string automatically.
 *	GoogleAnalyticsCampaign 	Optional string indicating the value to set for the utm_campaign tracking parameter. If this isn't provided the email's from address will be used instead.
 *	MetaData 					Metadata an associative array of user metadata. Mandrill will store this metadata and make it available for retrieval.
 *								In addition, you can select up to 10 metadata fields to index and make searchable using the Mandrill search api.
 *	RecipientMeteData 			Per-recipient metadata that will override the global values specified in the metadata parameter.
 *	Attachments 				An array of supported attachments to add to the message
 *	Images 						An array of embedded images to add to the message
 */
type Message struct {
	Html 						string 							`json:"html"`
	Text 						string 							`json:"text"`
	Subject 					string 							`json:"subject"`
	FromEmail 					string 							`json:"from_email"`
	FromName  					string 							`json:"from_name"`
	To 							[]*MessageTo 					`json:"to"`
	Headers 					*MessageHeaders 				`json:"headers"`
	Important 					bool 							`json:"important"`
	TrackOpens 					bool 							`json:"track_opens"`
	TrackClicks 				bool 							`json:"track_clicks"`
	AutoText 					bool							`json:"auto_text"`
	AutoHtml 					bool 							`json:"auto_html"`
	InlineCss 					bool 							`json:"inline_css"`
	UrlStripQs 					bool 							`json:"url_strip_qs"`
	PreserveRecipients 			bool 							`json:"preserve_recipients"`
	ViewContentLink 			bool 							`json:"view_content_link"`
	BccAddress  				string 							`json:"bcc_address"`
	TrackingDomain 				string 							`json:"tracking_domain"`
	SigningDomain 				string 							`json:"signing_domain"`
	ReturnPathDomain 			string 							`json:"return_path_domain"`
	Merge 						bool 							`json:"merge"`
	GlobalMergeVars 			[]*MessageMergeItem 			`json:"global_merge_vars"`
	MergeVars 					[]*MessageMergeWrapper			`json:"merge_vars"`
	Tags 						[]string 						`json:"tags"`
	SubAccount 					string 							`json:"subaccount"`
	GoogleAnalyticsDomains 		[]string 						`json:"google_analytics_domains"`
	GoogleAnalyticsCampaign 	string 							`json:"google_analytics_campaign"`
	MetaData 					*MessageMetaData				`json:"metadata`
	RecipientMetaData 			[]*MessageRecipientMetaData	 	`json:"recipient_metadata"`
	Attachments 				[]*MessageAttachment 			`json:"attachments"`
	Images 						[]*MessageAttachment 			`json:"images"`
}

type MessageTo struct {
	Email 		string `json:"email"`
	Name 		string `json:"name"`
	Type 		string `json:"type"`
}

type MessageHeaders struct {
	ReplyTo 	string `json:"Reply-To"`
}

type MessageMergeItem struct {
	Name 		string `json:"name"`
	Content 	string `json:"content"`
}

type MessageMergeWrapper struct {
	Recipient 	string `json:"rcpt"`
	Vars 		[]*MessageMergeItem `json:"vars"`
}

type MessageMetaData struct {
	Website 		string `json:"website"`
}

type MessageRecipientMetaData struct {
	Recipient 		string 								`json:"rcpt"`
	Values 			*MessageRecipientMetaDataValues		`json:"values"`
}

type MessageRecipientMetaDataValues struct {
	UserId 		int 	`json:"user_id"`
}

type MessageAttachment struct {
	Type  		string `json:"type"`
	Name 		string `json:"name"`
	Content 	string `json:"content"`
}