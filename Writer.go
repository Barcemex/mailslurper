// Copyright 2013-3014 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package mailslurper

import (
	"net"
	"time"

	"github.com/adampresley/webframework/logging2"
)

/*
An SMTPWriter is a simple object for writing commands and responses
to a client connected on a TCP socket
*/
type SMTPWriter struct {
	Connection net.Conn

	logger logging2.ILogger
}

/*
SayGoodbye tells a client that we are done communicating. This sends
a 221 response. It returns true/false for success and a string
with any response.
*/
func (smtpWriter *SMTPWriter) SayGoodbye() error {
	return smtpWriter.SendResponse(SMTP_CLOSING_MESSAGE)
}

/*
SayHello sends a hello message to a new client. The SMTP protocol
dictates that you must be polite. :)
*/
func (smtpWriter *SMTPWriter) SayHello() error {
	err := smtpWriter.SendResponse(SMTP_WELCOME_MESSAGE)
	if err != nil {
		return err
	}

	smtpWriter.logger.Infof("Reading data from client connection...")
	return nil
}

/*
SendDataResponse is a function to send a DATA response message
*/
func (smtpWriter *SMTPWriter) SendDataResponse() error {
	return smtpWriter.SendResponse(SMTP_DATA_RESPONSE_MESSAGE)
}

/*
SendResponse sends a response to a client connection. It returns true/false for success and a string
with any response.
*/
func (smtpWriter *SMTPWriter) SendResponse(response string) error {
	var err error

	if err = smtpWriter.Connection.SetWriteDeadline(time.Now().Add(time.Second * 2)); err != nil {
		smtpWriter.logger.Errorf("Problem setting write deadline: %s", err.Error())
	}

	_, err = smtpWriter.Connection.Write([]byte(string(response + SMTP_CRLF)))
	return err
}

/*
SendHELOResponse sends a HELO message to a client
*/
func (smtpWriter *SMTPWriter) SendHELOResponse() error {
	return smtpWriter.SendResponse(SMTP_HELLO_RESPONSE_MESSAGE)
}

/*
SendOkResponse sends an OK to a client
*/
func (smtpWriter *SMTPWriter) SendOkResponse() error {
	return smtpWriter.SendResponse(SMTP_OK_MESSAGE)
}
