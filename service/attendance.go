package service

import (
	"announce/config"
	"announce/constant"
	"announce/helper"
	"encoding/json"
	"fmt"
	"math/rand"
	"strings"
)

func Attendance() (*string, string, error) {
	pickInt := rand.Intn(len(constant.TextType))
	log.Log("Random number: ", pickInt+1)

	randomTextType := constant.TextTypeString[pickInt+1]
	log.Log("Random text type: ", randomTextType)
	
	var (
		url  strings.Builder
		word strings.Builder
	)

	url.WriteString(config.NinjaApiBaseUrl)

	switch constant.TextType[randomTextType] {
	case constant.TextType["Quotes"]:
		url.WriteString(config.NinjaApiQuotesUrl)
		code, resAPI, err := helper.CreateHttpReq(url.String(), "GET", "", config.NinjaApiKey, "")
		if err != nil {
			log.Error("Error get ninja API: ", err.Error())
			return nil, "", err
		} else {
			if code < 200 || code > 299 {
				log.Error("Error get ninja API: ", string(resAPI))
				return nil, "", fmt.Errorf("error get ninja API")
			} else {
				var res []map[string]any
				if err := json.Unmarshal(resAPI, &res); err != nil {
					log.Error("Error unmarshal ninja API: ", err.Error())
					return nil, "", err
				}

				if _, ok := res[0]["quote"].(string); !ok {
					log.Error("Error get quote : ", res)
					return nil, "", fmt.Errorf("error get quote")
				}

				if _, ok := res[0]["author"].(string); !ok {
					log.Error("Error get author quote : ", res)
					return nil, "", fmt.Errorf("error get author quote")
				}

				word.WriteString(fmt.Sprintf("%s - %s", res[0]["author"].(string), res[0]["quote"].(string)))
			}
		}
	case constant.TextType["Jokes"]:
		url.WriteString(config.NinjaApiJokesUrl)
		code, resAPI, err := helper.CreateHttpReq(url.String(), "GET", "", config.NinjaApiKey, "")
		if err != nil {
			log.Error("Error get ninja API: ", err.Error())
			return nil, "", err
		} else {
			if code < 200 || code > 299 {
				log.Error("Error get ninja API: ", string(resAPI))
				return nil, "", fmt.Errorf("error get ninja API")
			} else {
				var res []map[string]any
				if err := json.Unmarshal(resAPI, &res); err != nil {
					log.Error("Error unmarshal ninja API: ", err.Error())
					return nil, "", err
				}

				if _, ok := res[0]["joke"].(string); !ok {
					log.Error("Error get joke : ", res)
					return nil, "", fmt.Errorf("error get joke")
				}

				word.WriteString(res[0]["joke"].(string))
			}
		}
	case constant.TextType["Riddle"]:
		url.WriteString(config.NinjaApiRiddleUrl)
		code, resAPI, err := helper.CreateHttpReq(url.String(), "GET", "", config.NinjaApiKey, "")
		if err != nil {
			log.Error("Error get ninja API: ", err.Error())
			return nil, "", err
		} else {
			if code < 200 || code > 299 {
				log.Error("Error get ninja API: ", string(resAPI))
				return nil, "", fmt.Errorf("error get ninja API")
			} else {
				var res []map[string]any
				if err := json.Unmarshal(resAPI, &res); err != nil {
					log.Error("Error unmarshal ninja API: ", err.Error())
					return nil, "", err
				}

				if _, ok := res[0]["question"].(string); !ok {
					log.Error("Error get riddle question : ", res)
					return nil, "", fmt.Errorf("error get riddle question")
				}

				if _, ok := res[0]["answer"].(string); !ok {
					log.Error("Error get riddle answer : ", res)
					return nil, "", fmt.Errorf("error get riddle answer")
				}

				word.WriteString(fmt.Sprintf("%s %s", res[0]["question"].(string), res[0]["answer"].(string)))
			}
		}
	case constant.TextType["Trivia"]:
		url.WriteString(config.NinjaApiTriviaUrl)
		code, resAPI, err := helper.CreateHttpReq(url.String(), "GET", "", config.NinjaApiKey, "")
		if err != nil {
			log.Error("Error get ninja API: ", err.Error())
			return nil, "", err
		} else {
			if code < 200 || code > 299 {
				log.Error("Error get ninja API: ", string(resAPI))
				return nil, "", fmt.Errorf("error get ninja API")
			} else {
				var res []map[string]any
				if err := json.Unmarshal(resAPI, &res); err != nil {
					log.Error("Error unmarshal ninja API: ", err.Error())
					return nil, "", err
				}

				if _, ok := res[0]["question"].(string); !ok {
					log.Error("Error get riddle question : ", res)
					return nil, "", fmt.Errorf("error get riddle question")
				}

				if _, ok := res[0]["answer"].(string); !ok {
					log.Error("Error get riddle answer : ", res)
					return nil, "", fmt.Errorf("error get riddle answer")
				}

				word.WriteString(fmt.Sprintf("%s, %s", res[0]["question"].(string), res[0]["answer"].(string)))
			}
		}
	case constant.TextType["Advise"]:
		url.WriteString(config.NinjaApiAdviceUrl)
		code, resAPI, err := helper.CreateHttpReq(url.String(), "GET", "", config.NinjaApiKey, "")
		if err != nil {
			log.Error("Error get ninja API: ", err.Error())
			return nil, "", err
		} else {
			if code < 200 || code > 299 {
				log.Error("Error get ninja API: ", string(resAPI))
				return nil, "", fmt.Errorf("error get ninja API")
			} else {
				var res map[string]any
				if err := json.Unmarshal(resAPI, &res); err != nil {
					log.Error("Error unmarshal ninja API: ", err.Error())
					return nil, "", err
				}

				if _, ok := res["advice"].(string); !ok {
					log.Error("Error get advice : ", res)
					return nil, "", fmt.Errorf("error get advice")
				}

				word.WriteString(res["advice"].(string))
			}
		}
	case constant.TextType["Fun Fact"]:
		url.WriteString(config.NinjaApiFactUrl)
		code, resAPI, err := helper.CreateHttpReq(url.String(), "GET", "", config.NinjaApiKey, "")
		if err != nil {
			log.Error("Error get ninja API: ", err.Error())
			return nil, "", err
		} else {
			if code < 200 || code > 299 {
				log.Error("Error get ninja API: ", string(resAPI))
				return nil, "", fmt.Errorf("error get ninja API")
			} else {
				var res []map[string]any
				if err := json.Unmarshal(resAPI, &res); err != nil {
					log.Error("Error unmarshal ninja API: ", err.Error())
					return nil, "", err
				}

				if _, ok := res[0]["fact"].(string); !ok {
					log.Error("Error get fact : ", res)
					return nil, "", fmt.Errorf("error get fact")
				}

				word.WriteString(res[0]["fact"].(string))
			}
		}
	default:
		log.Error("Invalid text type: ", randomTextType)
		return nil, "", fmt.Errorf("invalid text type: %s", randomTextType)
	}

	sayWord := word.String()
	return &sayWord, randomTextType, nil
}

func HandleAttendance()  {

	if word, textType, err := Attendance(); err != nil {
		if err := helper.HandleSendToSpace(constant.Notif["ATTENDANCE"], "🔔 [ATTENDANCE] 🔔 \n", "No Words Today"); err != nil {
			log.Error("Error send notif : ", err.Error())
		}
		return
	}else {

		var emote strings.Builder
		switch constant.TextType[textType] {
		case constant.TextType["Quotes"]:
			emote.WriteString("💬")
		case constant.TextType["Jokes"]:
			emote.WriteString("🤡")
		case constant.TextType["Riddle"]:
			emote.WriteString("🧩")
		case constant.TextType["Trivia"]:
			emote.WriteString("🧠")
		case constant.TextType["Advise"]:
			emote.WriteString("💡")
		case constant.TextType["Fun Fact"]:
			emote.WriteString("🧪")
		}
		if err := helper.HandleSendToSpace(constant.Notif["ATTENDANCE"], fmt.Sprintf("%s [ATTENDANCE] - [%s] %s\n", emote.String(), strings.ToUpper(textType), emote.String()), *word); err != nil {
			log.Error("Error send notif : ", err.Error())
		}
	}
}
