package helper

import (
	"announce/config"
	"announce/constant"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CardWithLink struct {
	CardsV2 []CardV2 `json:"cardsV2"`
}

type CardV2 struct {
	Cards *MessageCard `json:"card"`
}

type Card struct {
	Cards []*MessageCard `json:"cards"`
}

type MessageCard struct {
	Header   *CardHeader    `json:"header"`
	Sections []*CardSection `json:"sections"`
}

type CardHeader struct {
	Title    string `json:"title"`
	Subtitle string `json:"subtitle"`
}

type CardSection struct {
	Widgets []CardWidget `json:"widgets,omitempty"`
}

type CardWidget struct {
	TextParagraph *TextParagraph `json:"textParagraph,omitempty"`
	ButtonList    *ButtonList    `json:"buttonList,omitempty"`
}

type TextParagraph struct {
	Text string `json:"text"`
}

// Button structure
type ButtonList struct {
	Buttons []Button `json:"buttons"`
}

type Button struct {
	Text    string   `json:"text"`
	OnClick *OnClick `json:"onClick,omitempty"`
}

type OnClick struct {
	OpenLink *OpenLink `json:"openLink,omitempty"`
}

type OpenLink struct {
	URL string `json:"url"`
}

func CreateCard(flag, msg string) *Card {
	card := &Card{
		Cards: []*MessageCard{
			{
				Header: &CardHeader{
					Title: flag,
				},
				Sections: []*CardSection{
					{
						Widgets: []CardWidget{
							{
								TextParagraph: &TextParagraph{
									Text: msg,
								},
							},
						},
					},
				},
			},
		},
	}

	return card
}

func CreateCardWithButton(flag, msg string) *CardWithLink {
	card := &CardWithLink{
		CardsV2: []CardV2{
			{
				Cards: &MessageCard{
					Header: &CardHeader{
						Title:    flag,
					},
					Sections: []*CardSection{
						{
							Widgets: []CardWidget{
								{
									TextParagraph: &TextParagraph{
										Text: msg,
									},
								},
							},
						},
						{
							Widgets: []CardWidget{
								{
									ButtonList: &ButtonList{
										Buttons: []Button{
											{
												Text: "Absence Here",
												OnClick: &OnClick{
													OpenLink: &OpenLink{
														URL: config.KyAttendanceUrl,
													},
												},
											},
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
	return card
}

func HandleSendToSpace(notifType int, flag, msg string) error {
	// notif := struct {
	// 	Text string `json:"text"`
	// }{
	// 	Text: flag + " " + msg,
	// }

	var (
		card    *Card
		cardV2  *CardWithLink
		payload []byte
		err     error
	)

	switch notifType {
	case constant.Notif["ATTENDANCE"]:
		cardV2 = CreateCardWithButton(flag, msg)

		payload, err = json.Marshal(cardV2)
		if err != nil {
			return err
		}

		log.Log("Payload : ", string(payload))
	case constant.Notif["UNIFORM"]:
		card = CreateCard(flag, msg)
		payload, err = json.Marshal(card)
		if err != nil {
			return err
		}
	default:
		log.Error("Invalid notif type: ", notifType)
		return fmt.Errorf("invalid notif type: %d", notifType)
	}

	resp, errSend := http.Post(config.SpaceNotif, "application/json", bytes.NewBuffer(payload))

	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response body: %v", err)
	}

	// code, res, errSend := CreateHttpReq(config.SpaceNotif, "POST", "", "", string(payload))
	log.Log("Code send notif : ", resp.StatusCode)
	if errSend != nil {
		log.Error("Error send notif : ", errSend.Error())
		return errSend
	} else {
		if resp.StatusCode < 200 || resp.StatusCode > 299 {
			log.Log("Error send notif :", string(res))
		} else {
			log.Log("Success send notif : ", string(res))
		}
	}

	resp.Body.Close()
	return nil
}
