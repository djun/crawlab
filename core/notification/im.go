package notification

import (
	"errors"
	"fmt"
	"github.com/apex/log"
	"github.com/crawlab-team/crawlab/core/models/models/v2"
	"github.com/crawlab-team/crawlab/trace"
	"github.com/imroc/req"
	"regexp"
	"strings"
)

type ResBody struct {
	ErrCode int    `json:"errcode"`
	ErrMsg  string `json:"errmsg"`
}

func SendIMNotification(ch *models.NotificationChannelV2, title, content string) error {
	// TODO: compatibility with different IM providers
	switch ch.Provider {
	case ChannelIMProviderLark:
		return sendImLark(ch, title, content)
	case ChannelIMProviderSlack:
		return sendImSlack(ch, title, content)
	case ChannelIMProviderDingtalk:
		return sendImDingTalk(ch, title, content)
	case ChannelIMProviderWechatWork:
		return sendImWechatWork(ch, title, content)
	case ChannelIMProviderTelegram:
		return sendImTelegram(ch, title, content)
	}

	// request header
	header := req.Header{
		"Content-Type": "application/json; charset=utf-8",
	}

	// request data
	data := req.Param{
		"msgtype": "markdown",
		"markdown": req.Param{
			"title":   title,
			"text":    content,
			"content": content,
		},
		"at": req.Param{
			"atMobiles": []string{},
			"isAtAll":   false,
		},
		"text": content,
	}
	if strings.Contains(strings.ToLower(ch.WebhookUrl), "feishu") {
		data = req.Param{
			"msg_type": "text",
			"content": req.Param{
				"text": content,
			},
		}
	}

	// perform request
	res, err := req.Post(ch.WebhookUrl, header, req.BodyJSON(&data))
	if err != nil {
		return trace.TraceError(err)
	}

	// parse response
	var resBody ResBody
	if err := res.ToJSON(&resBody); err != nil {
		return trace.TraceError(err)
	}

	// validate response code
	if resBody.ErrCode != 0 {
		return errors.New(resBody.ErrMsg)
	}

	return nil
}

func getIMRequestHeader() req.Header {
	return req.Header{
		"Content-Type": "application/json; charset=utf-8",
	}
}

func performIMRequest[T any](webhookUrl string, data req.Param) (resBody T, err error) {
	// perform request
	res, err := req.Post(webhookUrl, getIMRequestHeader(), req.BodyJSON(&data))
	if err != nil {
		log.Errorf("IM request error: %v", err)
		return resBody, err
	}

	// parse response
	if err := res.ToJSON(&resBody); err != nil {
		log.Warnf("Parsing IM response error: %v", err)
		resText, err := res.ToString()
		if err != nil {
			log.Warnf("Converting response to string error: %v", err)
			return resBody, err
		}
		log.Infof("IM response: %s", resText)
		return resBody, nil
	}

	return resBody, nil
}

func convertMarkdownToSlack(markdown string) string {
	// Convert bold text
	reBold := regexp.MustCompile(`\*\*(.*?)\*\*`)
	slack := reBold.ReplaceAllString(markdown, `*$1*`)

	// Convert italic text
	reItalic := regexp.MustCompile(`\*(.*?)\*`)
	slack = reItalic.ReplaceAllString(slack, `_$1_`)

	// Convert links
	reLink := regexp.MustCompile(`\[(.*?)\]\((.*?)\)`)
	slack = reLink.ReplaceAllString(slack, `<$2|$1>`)

	// Convert inline code
	reInlineCode := regexp.MustCompile("`(.*?)`")
	slack = reInlineCode.ReplaceAllString(slack, "`$1`")

	// Convert unordered list
	slack = strings.ReplaceAll(slack, "- ", "• ")

	// Convert ordered list
	reOrderedList := regexp.MustCompile(`^\d+\. `)
	slack = reOrderedList.ReplaceAllStringFunc(slack, func(s string) string {
		return strings.Replace(s, ". ", ". ", 1)
	})

	// Convert blockquote
	reBlockquote := regexp.MustCompile(`^> (.*)`)
	slack = reBlockquote.ReplaceAllString(slack, `> $1`)

	return slack
}

func convertMarkdownToTelegram(markdownText string) string {
	// Combined regex to handle bold and italic
	re := regexp.MustCompile(`(?m)(\*\*)(.*)(\*\*)|(__)(.*)(__)|(\*)(.*)(\*)|(_)(.*)(_)`)
	markdownText = re.ReplaceAllStringFunc(markdownText, func(match string) string {
		groups := re.FindStringSubmatch(match)
		if groups[1] != "" || groups[4] != "" {
			// Handle bold
			return "*" + match[2:len(match)-2] + "*"
		} else if groups[6] != "" || groups[9] != "" {
			// Handle italic
			return "_" + match[1:len(match)-1] + "_"
		} else {
			// No match
			return match
		}
	})

	// Convert unordered list
	re = regexp.MustCompile(`(?m)^- (.*)`)
	markdownText = re.ReplaceAllString(markdownText, "• $1")

	// Escape characters
	escapeChars := []string{"#", "-", "."}
	for _, c := range escapeChars {
		markdownText = strings.ReplaceAll(markdownText, c, "\\"+c)
	}

	return markdownText
}

func sendImLark(ch *models.NotificationChannelV2, title, content string) error {
	data := req.Param{
		"msg_type": "interactive",
		"card": req.Param{
			"header": req.Param{
				"title": req.Param{
					"tag":     "plain_text",
					"content": title,
				},
			},
			"elements": []req.Param{
				{
					"tag":     "markdown",
					"content": content,
				},
			},
		},
	}
	resBody, err := performIMRequest[ResBody](ch.WebhookUrl, data)
	if err != nil {
		return err
	}
	if resBody.ErrCode != 0 {
		return errors.New(resBody.ErrMsg)
	}
	return nil
}

func sendImSlack(ch *models.NotificationChannelV2, title, content string) error {
	data := req.Param{
		"blocks": []req.Param{
			{"type": "header", "text": req.Param{"type": "plain_text", "text": title}},
			{"type": "section", "text": req.Param{"type": "mrkdwn", "text": convertMarkdownToSlack(content)}},
		},
	}
	resBody, err := performIMRequest[ResBody](ch.WebhookUrl, data)
	if err != nil {
		return err
	}
	if resBody.ErrCode != 0 {
		return errors.New(resBody.ErrMsg)
	}
	return nil
}

func sendImDingTalk(ch *models.NotificationChannelV2, title string, content string) error {
	data := req.Param{
		"msgtype": "markdown",
		"markdown": req.Param{
			"title": title,
			"text":  fmt.Sprintf("# %s\n\n%s", title, content),
		},
	}
	resBody, err := performIMRequest[ResBody](ch.WebhookUrl, data)
	if err != nil {
		return err
	}
	if resBody.ErrCode != 0 {
		return errors.New(resBody.ErrMsg)
	}
	return nil
}

func sendImWechatWork(ch *models.NotificationChannelV2, title string, content string) error {
	data := req.Param{
		"msgtype": "markdown",
		"markdown": req.Param{
			"content": fmt.Sprintf("# %s\n\n%s", title, content),
		},
	}
	resBody, err := performIMRequest[ResBody](ch.WebhookUrl, data)
	if err != nil {
		return err
	}
	if resBody.ErrCode != 0 {
		return errors.New(resBody.ErrMsg)
	}
	return nil
}

func sendImTelegram(ch *models.NotificationChannelV2, title string, content string) error {
	type ResBody struct {
		Ok          bool   `json:"ok"`
		Description string `json:"description"`
	}

	// chat id
	chatId := ch.TelegramChatId
	if !strings.HasPrefix("@", ch.TelegramChatId) {
		chatId = fmt.Sprintf("@%s", ch.TelegramChatId)
	}

	// webhook url
	webhookUrl := fmt.Sprintf("https://api.telegram.org/bot%s/sendMessage", ch.TelegramBotToken)

	// original Markdown text
	text := fmt.Sprintf("**%s**\n\n%s", title, content)

	// convert to Telegram MarkdownV2
	text = convertMarkdownToTelegram(text)

	// request data
	data := req.Param{
		"chat_id":    chatId,
		"text":       text,
		"parse_mode": "MarkdownV2",
	}

	// perform request
	resBody, err := performIMRequest[ResBody](webhookUrl, data)
	if err != nil {
		return err
	}
	if !resBody.Ok {
		return errors.New(resBody.Description)
	}
	return nil
}
