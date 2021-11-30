package smtpmock

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewConfiguration(t *testing.T) {
	t.Run("creates new configuration with default settings", func(t *testing.T) {
		buildedConfiguration := NewConfiguration(ConfigurationAttr{})

		assert.Equal(t, DefaultHostAddress, buildedConfiguration.hostAddress)
		assert.Equal(t, DefaultPortNuber, buildedConfiguration.portNumber)
		assert.False(t, buildedConfiguration.logToStdout)
		assert.False(t, buildedConfiguration.isCmdFailFast)
		assert.False(t, buildedConfiguration.logServerActivity)
		assert.Equal(t, DefaultGreetingMsg, buildedConfiguration.msgGreeting)
		assert.Equal(t, DefaultInvalidCmdMsg, buildedConfiguration.msgInvalidCmd)
		assert.Equal(t, DefaultQuitMsg, buildedConfiguration.msgQuitCmd)
		assert.Equal(t, DefaultSessionTimeout, buildedConfiguration.sessionTimeout)

		assert.Equal(t, DefaultInvalidCmdHeloSequenceMsg, buildedConfiguration.msgInvalidCmdHeloSequence)
		assert.Equal(t, DefaultInvalidCmdHeloArgMsg, buildedConfiguration.msgInvalidCmdHeloArg)
		assert.Equal(t, DefaultQuitMsg, buildedConfiguration.msgHeloBlacklistedDomain)
		assert.Equal(t, DefaultReceivedMsg, buildedConfiguration.msgHeloReceived)

		assert.Equal(t, DefaultInvalidCmdMailfromSequenceMsg, buildedConfiguration.msgInvalidCmdMailfromSequence)
		assert.Equal(t, DefaultInvalidCmdMailfromArgMsg, buildedConfiguration.msgInvalidCmdMailfromArg)
		assert.Equal(t, DefaultQuitMsg, buildedConfiguration.msgMailfromBlacklistedEmail)
		assert.Equal(t, DefaultReceivedMsg, buildedConfiguration.msgMailfromReceived)

		assert.Equal(t, DefaultInvalidCmdRcpttoSequenceMsg, buildedConfiguration.msgInvalidCmdRcpttoSequence)
		assert.Equal(t, DefaultInvalidCmdRcpttoArgMsg, buildedConfiguration.msgInvalidCmdRcpttoArg)
		assert.Equal(t, DefaultQuitMsg, buildedConfiguration.msgRcpttoBlacklistedEmail)
		assert.Equal(t, DefaultNotRegistredRcpttoEmailMsg, buildedConfiguration.msgRcpttoNotRegisteredEmail)
		assert.Equal(t, DefaultReceivedMsg, buildedConfiguration.msgRcpttoReceived)

		assert.Equal(t, DefaultInvalidCmdDataSequenceMsg, buildedConfiguration.msgInvalidCmdDataSequence)
		assert.Equal(t, DefaultReadyForReceiveMsg, buildedConfiguration.msgDataReceived)

		assert.Equal(t, fmt.Sprintf(DefaultMsgSizeIsTooBigMsg+" %d bytes", DefaultMessageSizeLimit), buildedConfiguration.msgMsgSizeIsTooBig)
		assert.Equal(t, DefaultReceivedMsg, buildedConfiguration.msgMsgReceived)
		assert.Equal(t, DefaultMessageSizeLimit, buildedConfiguration.msqSizeLimit)

		assert.Empty(t, buildedConfiguration.blacklistedHeloDomains)
		assert.Empty(t, buildedConfiguration.blacklistedMailfromEmails)
		assert.Empty(t, buildedConfiguration.blacklistedRcpttoEmails)
		assert.Empty(t, buildedConfiguration.notRegisteredEmails)
	})

	t.Run("creates new configuration with custom settings", func(t *testing.T) {
		configAttr := ConfigurationAttr{
			hostAddress:                   "hostAddress",
			portNumber:                    25,
			logToStdout:                   true,
			logServerActivity:             true,
			isCmdFailFast:                 true,
			msgGreeting:                   "msgGreeting",
			msgInvalidCmd:                 "msgInvalidCmd",
			msgQuitCmd:                    "msgQuitCmd",
			msgInvalidCmdHeloSequence:     "msgInvalidCmdHeloSequence",
			msgInvalidCmdHeloArg:          "msgInvalidCmdHeloArg",
			msgHeloBlacklistedDomain:      "msgHeloBlacklistedDomain",
			msgHeloReceived:               "msgHeloReceived",
			msgInvalidCmdMailfromSequence: "msgInvalidCmdMailfromSequence",
			msgInvalidCmdMailfromArg:      "msgInvalidCmdMailfromArg",
			msgMailfromBlacklistedEmail:   "msgMailfromBlacklistedEmail",
			msgMailfromReceived:           "msgMailfromReceived",
			msgInvalidCmdRcpttoSequence:   "msgInvalidCmdRcpttoSequence",
			msgInvalidCmdRcpttoArg:        "msgInvalidCmdRcpttoArg",
			msgRcpttoNotRegisteredEmail:   "msgRcpttoNotRegisteredEmail",
			msgRcpttoBlacklistedEmail:     "msgRcpttoBlacklistedEmail",
			msgRcpttoReceived:             "msgRcpttoReceived",
			msgInvalidCmdDataSequence:     "msgInvalidCmdDataSequence",
			msgDataReceived:               "msgDataReceived",
			msgMsgSizeIsTooBig:            EmptyString,
			msgMsgReceived:                "msgMsgReceived",
			blacklistedHeloDomains:        []string{},
			blacklistedMailfromEmails:     []string{},
			notRegisteredEmails:           []string{},
			blacklistedRcpttoEmails:       []string{},
			msqSizeLimit:                  42,
			sessionTimeout:                120,
		}
		buildedConfiguration := NewConfiguration(configAttr)

		assert.Equal(t, configAttr.hostAddress, buildedConfiguration.hostAddress)
		assert.Equal(t, configAttr.portNumber, buildedConfiguration.portNumber)
		assert.Equal(t, configAttr.logToStdout, buildedConfiguration.logToStdout)
		assert.Equal(t, configAttr.isCmdFailFast, buildedConfiguration.isCmdFailFast)
		assert.Equal(t, configAttr.logServerActivity, buildedConfiguration.logServerActivity)
		assert.Equal(t, configAttr.msgGreeting, buildedConfiguration.msgGreeting)
		assert.Equal(t, configAttr.msgInvalidCmd, buildedConfiguration.msgInvalidCmd)
		assert.Equal(t, configAttr.msgQuitCmd, buildedConfiguration.msgQuitCmd)
		assert.Equal(t, configAttr.sessionTimeout, buildedConfiguration.sessionTimeout)

		assert.Equal(t, configAttr.msgInvalidCmdHeloSequence, buildedConfiguration.msgInvalidCmdHeloSequence)
		assert.Equal(t, configAttr.msgInvalidCmdHeloArg, buildedConfiguration.msgInvalidCmdHeloArg)
		assert.Equal(t, configAttr.msgHeloBlacklistedDomain, buildedConfiguration.msgHeloBlacklistedDomain)
		assert.Equal(t, configAttr.msgHeloReceived, buildedConfiguration.msgHeloReceived)

		assert.Equal(t, configAttr.msgInvalidCmdMailfromSequence, buildedConfiguration.msgInvalidCmdMailfromSequence)
		assert.Equal(t, configAttr.msgInvalidCmdMailfromArg, buildedConfiguration.msgInvalidCmdMailfromArg)
		assert.Equal(t, configAttr.msgMailfromBlacklistedEmail, buildedConfiguration.msgMailfromBlacklistedEmail)
		assert.Equal(t, configAttr.msgMailfromReceived, buildedConfiguration.msgMailfromReceived)

		assert.Equal(t, configAttr.msgInvalidCmdRcpttoSequence, buildedConfiguration.msgInvalidCmdRcpttoSequence)
		assert.Equal(t, configAttr.msgInvalidCmdRcpttoArg, buildedConfiguration.msgInvalidCmdRcpttoArg)
		assert.Equal(t, configAttr.msgRcpttoBlacklistedEmail, buildedConfiguration.msgRcpttoBlacklistedEmail)
		assert.Equal(t, configAttr.msgRcpttoNotRegisteredEmail, buildedConfiguration.msgRcpttoNotRegisteredEmail)
		assert.Equal(t, configAttr.msgRcpttoReceived, buildedConfiguration.msgRcpttoReceived)

		assert.Equal(t, configAttr.msgInvalidCmdDataSequence, buildedConfiguration.msgInvalidCmdDataSequence)
		assert.Equal(t, configAttr.msgDataReceived, buildedConfiguration.msgDataReceived)

		assert.Equal(t, fmt.Sprintf(DefaultMsgSizeIsTooBigMsg+" %d bytes", configAttr.msqSizeLimit), buildedConfiguration.msgMsgSizeIsTooBig)
		assert.Equal(t, configAttr.msgMsgReceived, buildedConfiguration.msgMsgReceived)
		assert.Equal(t, configAttr.msqSizeLimit, buildedConfiguration.msqSizeLimit)

		assert.Equal(t, configAttr.blacklistedHeloDomains, buildedConfiguration.blacklistedHeloDomains)
		assert.Equal(t, configAttr.blacklistedMailfromEmails, buildedConfiguration.blacklistedMailfromEmails)
		assert.Equal(t, configAttr.blacklistedRcpttoEmails, buildedConfiguration.blacklistedRcpttoEmails)
		assert.Equal(t, configAttr.notRegisteredEmails, buildedConfiguration.notRegisteredEmails)
	})
}

func TestConfigurationAttrAssignDefaultValues(t *testing.T) {
	t.Run("assignes default values", func(t *testing.T) {
		configurationAttr := new(ConfigurationAttr)
		configurationAttr.assignDefaultValues()

		assert.Equal(t, DefaultHostAddress, configurationAttr.hostAddress)
		assert.Equal(t, DefaultPortNuber, configurationAttr.portNumber)
		assert.Equal(t, DefaultGreetingMsg, configurationAttr.msgGreeting)
		assert.Equal(t, DefaultInvalidCmdMsg, configurationAttr.msgInvalidCmd)
		assert.Equal(t, DefaultQuitMsg, configurationAttr.msgQuitCmd)
		assert.Equal(t, DefaultSessionTimeout, configurationAttr.sessionTimeout)

		assert.Equal(t, DefaultInvalidCmdHeloSequenceMsg, configurationAttr.msgInvalidCmdHeloSequence)
		assert.Equal(t, DefaultInvalidCmdHeloArgMsg, configurationAttr.msgInvalidCmdHeloArg)
		assert.Equal(t, DefaultQuitMsg, configurationAttr.msgHeloBlacklistedDomain)
		assert.Equal(t, DefaultReceivedMsg, configurationAttr.msgHeloReceived)

		assert.Equal(t, DefaultInvalidCmdMailfromSequenceMsg, configurationAttr.msgInvalidCmdMailfromSequence)
		assert.Equal(t, DefaultInvalidCmdMailfromArgMsg, configurationAttr.msgInvalidCmdMailfromArg)
		assert.Equal(t, DefaultQuitMsg, configurationAttr.msgMailfromBlacklistedEmail)
		assert.Equal(t, DefaultReceivedMsg, configurationAttr.msgMailfromReceived)

		assert.Equal(t, DefaultInvalidCmdRcpttoSequenceMsg, configurationAttr.msgInvalidCmdRcpttoSequence)
		assert.Equal(t, DefaultInvalidCmdRcpttoArgMsg, configurationAttr.msgInvalidCmdRcpttoArg)
		assert.Equal(t, DefaultQuitMsg, configurationAttr.msgRcpttoBlacklistedEmail)
		assert.Equal(t, DefaultNotRegistredRcpttoEmailMsg, configurationAttr.msgRcpttoNotRegisteredEmail)
		assert.Equal(t, DefaultReceivedMsg, configurationAttr.msgRcpttoReceived)

		assert.Equal(t, DefaultInvalidCmdDataSequenceMsg, configurationAttr.msgInvalidCmdDataSequence)
		assert.Equal(t, DefaultReadyForReceiveMsg, configurationAttr.msgDataReceived)

		assert.Equal(t, fmt.Sprintf(DefaultMsgSizeIsTooBigMsg+" %d bytes", DefaultMessageSizeLimit), configurationAttr.msgMsgSizeIsTooBig)
		assert.Equal(t, DefaultReceivedMsg, configurationAttr.msgMsgReceived)
		assert.Equal(t, DefaultMessageSizeLimit, configurationAttr.msqSizeLimit)
	})
}