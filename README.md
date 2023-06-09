# ![SMTP mock server written on Golang. Mimic any SMTP server behavior for your test environment with fake SMTP server](https://repository-images.githubusercontent.com/401721985/848bc1dd-fc35-4d78-8bd9-0ac3430270d8)

`smtpmock` is lightweight configurable multithreaded fake SMTP server written in Go. It meets the minimum requirements specified by [RFC 2821](https://datatracker.ietf.org/doc/html/rfc2821) & [RFC 5321](https://datatracker.ietf.org/doc/html/rfc5321). Allows to mimic any SMTP server behavior for your test environment and even more 🚀

## Table of Contents

- [](#)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Requirements](#requirements)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Inside of Golang ecosystem](#inside-of-golang-ecosystem)
      - [Configuring](#configuring)
      - [Manipulation with server](#manipulation-with-server)
    - [Inside of Ruby ecosystem](#inside-of-ruby-ecosystem)
      - [Example of usage](#example-of-usage)
    - [Inside of any ecosystem](#inside-of-any-ecosystem)
      - [Configuring with command line arguments](#configuring-with-command-line-arguments)
      - [Other options](#other-options)
      - [Stopping server](#stopping-server)
    - [Implemented SMTP commands](#implemented-smtp-commands)
  - [Contributing](#contributing)
  - [License](#license)
  - [Code of Conduct](#code-of-conduct)
  - [Credits](#credits)
  - [Versioning](#versioning)

## Features

- Configurable multithreaded RFC compatible SMTP server
- Implements the minimum command set, responds to commands and adds a valid received header to messages as specified in [RFC 2821](https://datatracker.ietf.org/doc/html/rfc2821) & [RFC 5321](https://datatracker.ietf.org/doc/html/rfc5321)
- Ability to configure behavior for each SMTP command
- Comes with default settings out of the box, configure only what you need
- Ability to override previous SMTP commands
- Fail fast scenario (ability to close client session for case when command was inconsistent or failed)
- Multiple receivers (ability to configure multiple `RCPT TO` commands receiving during one session)
- Multiple message receiving (ability to configure multiple message receiving during one session)
- Mock-server activity logger
- Ability to do graceful/force shutdown of SMTP mock server
- No authentication support
- Zero runtime dependencies
- Ability to access to server messages
- Simple and intuitive DSL
- Ability to run server as binary with command line arguments

## Usage

- [](#)
  - [Table of Contents](#table-of-contents)
  - [Features](#features)
  - [Requirements](#requirements)
  - [Installation](#installation)
  - [Usage](#usage)
    - [Inside of Golang ecosystem](#inside-of-golang-ecosystem)
      - [Configuring](#configuring)
      - [Manipulation with server](#manipulation-with-server)
    - [Inside of Ruby ecosystem](#inside-of-ruby-ecosystem)
      - [Example of usage](#example-of-usage)
    - [Inside of any ecosystem](#inside-of-any-ecosystem)
      - [Configuring with command line arguments](#configuring-with-command-line-arguments)
      - [Other options](#other-options)
      - [Stopping server](#stopping-server)
    - [Implemented SMTP commands](#implemented-smtp-commands)
  - [Contributing](#contributing)
  - [License](#license)
  - [Code of Conduct](#code-of-conduct)
  - [Credits](#credits)
  - [Versioning](#versioning)

### Inside of Golang ecosystem

You have to create your SMTP mock server using `smtpmock.New()` and `smtpmock.ConfigurationAttr{}` to start interaction with it.

#### Configuring

`smtpmock` is SMTP server for test environment with configurable behavior. It comes with default settings out of the box. But you can override any default behavior if you need.

```go
smtpmock.ConfigurationAttr{

  // Customizing server behavior
  // ---------------------------------------------------------------------
  // Host address where smtpmock will run, it's equal to "127.0.0.1" by default
  HostAddress:                   "[::]",

  // Port number on which the server will bind. If it not specified, it will be
  // assigned dynamically after server.Start() by default
  PortNumber:                    2525,

  // Enables/disables log to stdout. It's equal to false by default
  LogToStdout:                   true,

  // Enables/disables log server activity. It's equal to false by default
  LogServerActivity:             true,

  // Ability to specify session timeout. It's equal to 30 seconds by default
  SessionTimeout:                42,

  // Ability to specify graceful shutdown timeout. It's equal to 1 second by default
  ShutdownTimeout:               5,


  // Customizing SMTP command handlers behavior
  // ---------------------------------------------------------------------
  // Ability to configure fail fast scenario. It means that server will
  // close client session for case when command was inconsistent or failed.
  // It's equal to false by default
  IsCmdFailFast:                 true,

  // Ability to configure multiple RCPT TO command receiving during one session.
  // It means that server will handle and save all RCPT TO command request-response
  // pairs until receive successful response and next SMTP command has been passed.
  // Please note, by default will be saved only one, the last RCPT TO command
  // request-response pair. It's equal to false by default
  MultipleRcptto:                true,

  // Ability to configure multiple message receiving during one session. It means that server
  // will create another message during current SMTP session in case when RSET
  // command has been used after successful commands chain: MAIL FROM, RCPT TO, DATA.
  // Please note, by default RSET command flushes current message in any case.
  // It's equal to false by default
  MultipleMessageReceiving:      true,

  // Ability to specify blacklisted HELO domains. It's equal to empty []string
  BlacklistedHeloDomains:        []string{"example1.com", "example2.com", "localhost"},

  // Ability to specify blacklisted MAIL FROM emails. It's equal to empty []string
  BlacklistedMailfromEmails:     []string{"bot@olo.com", "robot@molo.com"},

  // Ability to specify blacklisted RCPT TO emails. It's equal to empty []string
  BlacklistedRcpttoEmails:       []string{"blacklisted@olo.com", "blacklisted@molo.com"},

  // Ability to specify not registered (non-existent) RCPT TO emails.
  // It's equal to empty []string
  NotRegisteredEmails:           []string{"nobody@olo.com", "non-existent@email.com"},

  // Ability to specify HELO response delay in seconds. It runs immediately,
  // equals to 0 seconds by default
  ResponseDelayHelo:             2,

  // Ability to specify MAIL FROM response delay in seconds. It runs immediately,
  // equals to 0 seconds by default
  ResponseDelayMailfrom:         2,

  // Ability to specify RCPT TO response delay in seconds. It runs immediately,
  // equals to 0 seconds by default
  ResponseDelayRcptto:           2,

  // Ability to specify DATA response delay in seconds. It runs immediately,
  // equals to 0 seconds by default
  ResponseDelayData:             2,

  // Ability to specify message response delay in seconds. It runs immediately,
  // equals to 0 seconds by default
  ResponseDelayMessage:          2,

  // Ability to specify RSET response delay in seconds. It runs immediately,
  // equals to 0 seconds by default
  ResponseDelayRset:             2,

  // Ability to specify QUIT response delay in seconds. It runs immediately,
  // equals to 0 seconds by default
  ResponseDelayQuit:             2,

  // Ability to specify message body size limit. It's equal to 10485760 bytes (10MB) by default
  MsgSizeLimit:                  5,


  // Customizing SMTP command handler messages context
  // ---------------------------------------------------------------------
  // Custom server greeting message. Base on defaultGreetingMsg by default
  MsgGreeting:                   "msgGreeting",

  // Custom invalid command message. Based on defaultInvalidCmdMsg by default
  MsgInvalidCmd:                 "msgInvalidCmd",

  // Custom invalid command HELO sequence message.
  // Based on defaultInvalidCmdHeloSequenceMsg by default
  MsgInvalidCmdHeloSequence:     "msgInvalidCmdHeloSequence",

  // Custom invalid command HELO argument message.
  // Based on defaultInvalidCmdHeloArgMsg by default
  MsgInvalidCmdHeloArg:          "msgInvalidCmdHeloArg",

  // Custom HELO blacklisted domain message. Based on defaultQuitMsg by default
  MsgHeloBlacklistedDomain:      "msgHeloBlacklistedDomain",

  // Custom HELO received message. Based on defaultReceivedMsg by default
  MsgHeloReceived:               "msgHeloReceived",

  // Custom invalid command MAIL FROM sequence message.
  // Based on defaultInvalidCmdMailfromSequenceMsg by default
  MsgInvalidCmdMailfromSequence: "msgInvalidCmdMailfromSequence",

  // Custom invalid command MAIL FROM argument message.
  // Based on defaultInvalidCmdMailfromArgMsg by default
  MsgInvalidCmdMailfromArg:      "msgInvalidCmdMailfromArg",

  // Custom MAIL FROM blacklisted email message. Based on defaultQuitMsg by default
  MsgMailfromBlacklistedEmail:   "msgMailfromBlacklistedEmail",

  // Custom MAIL FROM received message. Based on defaultReceivedMsg by default
  MsgMailfromReceived:           "msgMailfromReceived",

  // Custom invalid command RCPT TO sequence message.
  // Based on defaultInvalidCmdRcpttoSequenceMsg by default
  MsgInvalidCmdRcpttoSequence:   "msgInvalidCmdRcpttoSequence",

  // Custom invalid command RCPT TO argument message.
  // Based on defaultInvalidCmdRcpttoArgMsg by default
  MsgInvalidCmdRcpttoArg:        "msgInvalidCmdRcpttoArg",

  // Custom RCPT TO not registered email message.
  // Based on defaultNotRegisteredRcpttoEmailMsg by default
  MsgRcpttoNotRegisteredEmail:   "msgRcpttoNotRegisteredEmail",

  // Custom RCPT TO blacklisted email message. Based on defaultQuitMsg by default
  MsgRcpttoBlacklistedEmail:     "msgRcpttoBlacklistedEmail",

  // Custom RCPT TO received message. Based on defaultReceivedMsg by default
  MsgRcpttoReceived:             "msgRcpttoReceived",

  // Custom invalid command DATA sequence message.
  // Based on defaultInvalidCmdDataSequenceMsg by default
  MsgInvalidCmdDataSequence:     "msgInvalidCmdDataSequence",

  // Custom DATA received message. Based on defaultReadyForReceiveMsg by default
  MsgDataReceived:               "msgDataReceived",

  // Custom size is too big message. Based on defaultMsgSizeIsTooBigMsg by default
  MsgMsgSizeIsTooBig:            "msgMsgSizeIsTooBig",

  // Custom received message body message. Based on defaultReceivedMsg by default
  MsgMsgReceived:                "msgMsgReceived",

  // Custom invalid command RSET sequence message.
  // Based on defaultInvalidCmdHeloSequenceMsg by default
  MsgInvalidCmdRsetSequence:     "msgInvalidCmdRsetSequence",

  // Custom invalid command RSET message. Based on defaultInvalidCmdMsg by default
  MsgInvalidCmdRsetArg:           "msgInvalidCmdRsetArg",

  // Custom RSET received message. Based on defaultOkMsg by default
  MsgRsetReceived:               "msgRsetReceived",

  // Custom quit command message. Based on defaultQuitMsg by default
  MsgQuitCmd:                    "msgQuitCmd",

  // Custom NOOP command message. Based on defaultNoopMsg (NOOP) by default
  MsgNoopCmd:                    "msgQuitCmd",
}
```

#### Manipulation with server

```go
package main

import (
  "fmt"
  "net"
  "net/smtp"

  smtpmock "github.com/mocktools/go-smtp-mock/v2"
)

func main() {
  // You can pass empty smtpmock.ConfigurationAttr{}. It means that smtpmock will use default settings
  server := smtpmock.New(smtpmock.ConfigurationAttr{
    LogToStdout:       true,
    LogServerActivity: true,
  })

  // To start server use Start() method
  if err := server.Start(); err != nil {
    fmt.Println(err)
  }

  // Server's port will be assigned dynamically after server.Start()
  // for case when portNumber wasn't specified
  hostAddress, portNumber := "127.0.0.1", server.PortNumber()

  // Possible SMTP-client stuff for iteration with mock server
  address := fmt.Sprintf("%s:%d", hostAddress, portNumber)
  timeout := time.Duration(2) * time.Second

  connection, _ := net.DialTimeout("tcp", address, timeout)
  client, _ := smtp.NewClient(connection, hostAddress)
  client.Hello("example.com")
  client.Quit()
  client.Close()

  // Each result of SMTP session will be saved as message.
  // To get access to server messages use Messages() method
  server.Messages()

  // To stop the server use Stop() method. Please note, smtpmock uses graceful shutdown.
  // It means that smtpmock will end all sessions after client responses or by session
  // timeouts immediately.
  if err := server.Stop(); err != nil {
    fmt.Println(err)
  }
}
```

Code from example above will produce next output to stdout:

```code
INFO: 2021/11/30 22:07:30.554827 SMTP mock server started on port: 2525
INFO: 2021/11/30 22:07:30.554961 SMTP session started
INFO: 2021/11/30 22:07:30.554998 SMTP response: 220 Welcome
INFO: 2021/11/30 22:07:30.555059 SMTP request: EHLO example.com
INFO: 2021/11/30 22:07:30.555648 SMTP response: 250 Received
INFO: 2021/11/30 22:07:30.555686 SMTP request: QUIT
INFO: 2021/11/30 22:07:30.555722 SMTP response: 221 Closing connection
INFO: 2021/11/30 22:07:30.555732 SMTP session finished
WARNING: 2021/11/30 22:07:30.555801 SMTP mock server is in the shutdown mode and won't accept new connections
INFO: 2021/11/30 22:07:30.555808 SMTP mock server was stopped successfully
```

#### Stopping server

`smtpmock` accepts 3 shutdown signals: `SIGINT`, `SIGQUIT`, `SIGTERM`.

### Implemented SMTP commands

| id | Command | Sequenceable |  Available args | Example of usage |
| --- | --- | --- | --- | --- |
| `1` | `HELO` | no | `domain name`, `localhost`, `ip address`, `[ip address]` | `HELO example.com` |
| `1` | `EHLO` | no | `domain name`, `localhost`, `ip address`, `[ip address]` | `EHLO example.com` |
| `2` | `MAIL FROM` | can be used after command with id `1` and greater | `email address`, `<email address>` | `MAIL FROM: user@domain.com` |
| `3` | `RCPT TO` | can be used after command with id `2` and greater | `email address`, `<email address>` | `RCPT TO: user@domain.com` |
| `4` | `DATA` | can be used after command with id `3` | - | `DATA` |
| `5` | `RSET` | can be used after command with id `1` and greater | - | `RSET` |
| `6` | `QUIT` | no | - | `QUIT` |
| `7` | `NOOP` | no | - | `NOOP` |

Please note in case when same command used more the one time during same session all saved data upper this command will be erased.
