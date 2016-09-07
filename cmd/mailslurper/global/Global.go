// Copyright 2013-2016 Adam Presley. All rights reserved
// Use of this source code is governed by the MIT license
// that can be found in the LICENSE file.

package global

import "github.com/mailslurper/libmailslurper/storage"

const (
	// Version of the MailSlurper Server application
	SERVER_VERSION string = "1.11.1"
	DEBUG_ASSETS   bool   = false
)

var Database storage.IStorage
