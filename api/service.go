package bot

import (
	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
	"github.com/twilio/twilio-go/twiml"
)

var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
	Username: envTwilioAccountSID(),
	Password: envTwilioAuthToken(),
})

func (app *Config) twilioSendMessage(phoneNumber string, message string) (string, error) {
	params := &twilioApi.CreateMessageParams{}
	params.SetTo(phoneNumber)
	params.SetFrom(envTwilioPhoneNumber())
	params.SetBody(message)

	resp, err := client.Api.CreateMessage(params)

	if err != nil {
		return "", err
	}

	return *resp.Sid, nil
}

func (app *Config) twilioReplyMessage(message string) (string, error) {
	tm := &twiml.MessagingMessage{
		Body: message,
	}

	twimlResult, err := twiml.Messages([]twiml.Element{tm})
	if err != nil {
		return "", err
	}

	return twimlResult, nil
}
