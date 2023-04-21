# transcriber-bot

Convert a message received via WhatsApp into a text (Text-to-speech)

## Usage and documentation

Create environment variables in your .env file
```sh
TWILIO_ACCOUNT_SID=YOUR_SID
TWILIO_AUTH_TOKEN=YOUR_TOKEN
OPEN_AI_APY_KEY=YOUR_OPEN_AI_KEY
OPEN_AI_URL=https://api.openai.com/v1/audio/transcriptions
```

Use below command to run locally
```sh
make run
```
And for you to test it integrated with WhatsApp, you will need to expose this local address, to an address that is accessible via the internet.

To do this, you can use some tools like:
- [ngrok](https://ngrok.com/)
- [localtunnel](https://theboroer.github.io/localtunnel-www/)

localtunnel is simpler, and for this test scenario it is sufficient. To get the accessible DNS from the internet just run:
```sh
lt -p 8080
```

For more details, access this [medium link]() to the complete step-by-step article on how the code was built.

## License

Distributed under MIT License, please see license file in code for more details.
