package adapter

type Messenger interface {
	Send(message string) string
}

const(
	FbNote = "send to facebook: "
	TwNote = "send to twitter: "
)

// social api
type Facebook struct {
}

func (f Facebook) PostToFacebook(msg string) string {
	return "send to facebook: " + msg
}

type Twitter struct {
}

func (f Twitter) Twit(msg string) string {
	return "send to twitter: " + msg
}

// adapters
type FacebookAdapter struct {
	Messenger Facebook
}

func (a FacebookAdapter) Send(msg string) string {
	return a.Messenger.PostToFacebook(msg)
}

type TwitterAdapter struct {
	Messenger Twitter
}

func (a TwitterAdapter) Send(msg string) string {
	return a.Messenger.Twit(msg)
}
