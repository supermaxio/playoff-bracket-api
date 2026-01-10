package types

import "time"

type MessageDB struct {
	Username   string `bson:"username"`
	Message    string `bson:"message"`
	DatePosted int64  `bson:"date_posted"`
}

type MessagePost struct {
	Username   string `bson:"username"`
	Message    string `bson:"message"`
	DatePosted string `bson:"date_posted"`
}

type MessageResponse struct {
	Username   string `bson:"username"`
	Message    string `bson:"message"`
	DatePosted string `bson:"date_posted"`
}

func (u *MessageDB) Response() (messageResponse MessageResponse) {
	messageResponse.Username = u.Username
	messageResponse.Message = u.Message
	// convert unix timestamp to Date formatted string
	messageResponse.DatePosted = UnixToDateString(u.DatePosted)

	return
}

func (u *MessagePost) Create() (messageDB MessageDB, err error) {
	messageDB.Username = u.Username
	messageDB.Message = u.Message
	// convert unix timestamp to Date formatted string
	messageDB.DatePosted, err = DateStringToUnix(u.DatePosted)
	if err != nil {
		return MessageDB{}, err
	}

	return
}

func UnixToDateString(unixTimestamp int64) string {
	return time.Unix(unixTimestamp, 0).Format("2006-01-02 15:04:05")
}

func DateStringToUnix(dateString string) (int64, error) {
	layout := "2006-01-02 15:04:05"
	t, err := time.Parse(layout, dateString)
	if err != nil {
		return 0, err
	}
	return t.Unix(), nil
}