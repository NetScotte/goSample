package mysmtp

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"net/smtp"
	"os"
	"strings"
)

func SendMail(toAddr, cc, subject, msg string, attachments []map[string]string) (err error) {
	mailServer := "smtp.sina.com"
	mailPort := 25
	mailFrom := "abc@sina.com"
	mailPasswd := "abc"
	boundry := "CUSTOM"

	buf := bytes.NewBuffer([]byte{})
	headers := make(map[string]string)
	headers["From"] = mailFrom
	headers["To"] = toAddr
	headers["Cc"] = cc
	headers["Subject"] = "=?UTF-8?B?" + base64.StdEncoding.EncodeToString([]byte(subject)) + "?="
	if len(attachments) > 0 {
		headers["Content-Type"] = "multipart/mixed;boundary=" + boundry
		headers["Mime-Version"] = "1.0"
	} else {
		headers["Content-Type"] = "text/html; charset=UTF-8"
	}

	for key, value := range headers {
		buf.WriteString(key + ":" + value + "\r\n")
	}
	buf.WriteString("\r\n")

	if len(attachments) > 0 {
		buf.WriteString("--" + boundry + "\r\n")
		buf.WriteString("Content-Type: text/plain; charset=utf-8\r\n")
		buf.WriteString("\r\n")
		buf.WriteString(msg)
		buf.WriteString("\r\n")

		for _, attachment := range attachments {
			fileName, ok := attachment["filename"]
			if !ok {
				err = fmt.Errorf("附件必须指定文件名:filename")
				return
			}
			contentType, ok := attachment["contentType"]
			if !ok {
				err = fmt.Errorf("附件必须指定文件类型:contentType")
				return
			}
			filePath, ok := attachment["filePath"]
			if !ok {
				err = fmt.Errorf("附件必须指定文件路径:filePath")
				return
			}

			buf.WriteString("--" + boundry + "\r\n")
			buf.WriteString("Content-Transfer-Encoding:base64\r\n")
			buf.WriteString("Content-Disposition:attachment\r\n")
			buf.WriteString(fmt.Sprintf("Content-Type:%s;name=\"%s\"\r\n",
				contentType, "=?UTF-8?B?"+base64.StdEncoding.EncodeToString([]byte(fileName))+"?="))
			buf.WriteString("\r\n")
			fileObj, aerr := os.Open(filePath)
			if aerr != nil {
				err = aerr
				return
			}
			fileContent, aerr := ioutil.ReadAll(fileObj)
			if aerr != nil {
				err = aerr
				return
			}
			base64FileContent := []byte{}
			base64.StdEncoding.Encode(base64FileContent, fileContent)
			for index := 0; index < len(base64FileContent); index++ {
				buf.WriteByte(base64FileContent[index])
				if (index+1)%76 == 0 {
					buf.WriteString("\r\n")
				}
			}
			buf.WriteString("\r\n")
		}

	} else {
		buf.WriteString(msg)
		buf.WriteString("\r\n")
	}

	// Set up authentication information.
	auth := smtp.PlainAuth("", mailFrom, mailPasswd, mailServer)
	err = smtp.SendMail(fmt.Sprintf("%s:%d", mailServer, mailPort), auth, mailFrom, strings.Split(toAddr, ";"), buf.Bytes())
	return
}
