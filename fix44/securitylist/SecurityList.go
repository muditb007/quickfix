//Package securitylist msg type = y.
package securitylist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a SecurityList wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//SecurityReqID is a required field for SecurityList.
func (m Message) SecurityReqID() (*field.SecurityReqIDField, errors.MessageRejectError) {
	f := &field.SecurityReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from SecurityList.
func (m Message) GetSecurityReqID(f *field.SecurityReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseID is a required field for SecurityList.
func (m Message) SecurityResponseID() (*field.SecurityResponseIDField, errors.MessageRejectError) {
	f := &field.SecurityResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseID reads a SecurityResponseID from SecurityList.
func (m Message) GetSecurityResponseID(f *field.SecurityResponseIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityRequestResult is a required field for SecurityList.
func (m Message) SecurityRequestResult() (*field.SecurityRequestResultField, errors.MessageRejectError) {
	f := &field.SecurityRequestResultField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityRequestResult reads a SecurityRequestResult from SecurityList.
func (m Message) GetSecurityRequestResult(f *field.SecurityRequestResultField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoRelatedSym is a non-required field for SecurityList.
func (m Message) TotNoRelatedSym() (*field.TotNoRelatedSymField, errors.MessageRejectError) {
	f := &field.TotNoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoRelatedSym reads a TotNoRelatedSym from SecurityList.
func (m Message) GetTotNoRelatedSym(f *field.TotNoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for SecurityList.
func (m Message) LastFragment() (*field.LastFragmentField, errors.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from SecurityList.
func (m Message) GetLastFragment(f *field.LastFragmentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for SecurityList.
func (m Message) NoRelatedSym() (*field.NoRelatedSymField, errors.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from SecurityList.
func (m Message) GetNoRelatedSym(f *field.NoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds SecurityList messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for SecurityList.
func Builder(
	securityreqid *field.SecurityReqIDField,
	securityresponseid *field.SecurityResponseIDField,
	securityrequestresult *field.SecurityRequestResultField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = quickfix.NewMessageBuilder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header().Set(field.NewMsgType("y"))
	builder.Body().Set(securityreqid)
	builder.Body().Set(securityresponseid)
	builder.Body().Set(securityrequestresult)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX44, "y", r
}