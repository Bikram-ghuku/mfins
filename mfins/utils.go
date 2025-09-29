package main

import "strings"

func fmt_notice(content NoticeElement) NoticeElement {
	content.MessageBody = fmt_notice_body(content.MessageBody)
	if content.SerialNo == 0 {
    	content.SerialNo = content.MessageId
	}
	return content
}

func fmt_notice_body(body string) string {
	body = strings.ReplaceAll(body, "<br />", "\\n")
	body = strings.ReplaceAll(body, "&nbsp;", " ")
	body = strings.ReplaceAll(body, "&ldquo;", "\\\"")
	body = strings.ReplaceAll(body, "&rdquo;", "\\\"")
	body = strings.ReplaceAll(body, "&lsquo;", "\\'")
	body = strings.ReplaceAll(body, "&rsquo;", "\\'")
	return  body
}