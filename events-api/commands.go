package main

//CommandID is a type of commands. Wrappes int.
type CommandID int

const (
	addComment CommandID = iota
	editComment
	deleteComment
	addLike
	removeLike
	addReply
	removeReply
)

// Command is a command which processes by Node.
type Command struct {
	Cmd CommandID
}

// GetCommandID returns command's ID
func (c *Command) GetCommandID() CommandID {
	return c.Cmd
}

// NodeCommandControlLink represents a command to Node for act with LinkControl
type NodeCommandControlLink struct {
	NodeCommand
	ctrl *LinkControl
}

// NodeCommandTopic represents a command to Node for act with topics
type NodeCommandTopic struct {
	NodeCommand
	active    *LinkActive
	topicName string
}

// NodeCommandActiveLink representsa a command to Node for act with LinkActive
type NodeCommandActiveLink struct {
	NodeCommand
	active *LinkActive
}

// NodeCommandSendMessage represents a command to Node to send a message
type NodeCommandSendMessage struct {
	NodeCommand
	msg string
}

type cmdActiveLink struct {
	cmd commandActiveLink
	msg []byte
	err string
}

type commandActiveLink int

const (
	quitLinkActive commandActiveLink = iota
	errorReadingActive
	sendMessageActive
	registerTopicActive
)

func (c commandActiveLink) String() string {
	switch c {
	case 0:
		return "quitLinkActive"
	case 1:
		return "errorReadingActive"
	case 2:
		return "sendMessage"
	case 3:
		return "registerTopicActive"
	default:
		return "unknown"
	}
}

type cmdContrlLink struct {
	cmd commandContrlLink
	err string
}

type commandContrlLink int

const (
	quitControlLink commandContrlLink = iota
	errorControlLinkAccept
	errorControlLinkRead
	errorControlLinkHandler

//	errorFromActiveLink
)

func (c commandContrlLink) String() string {
	switch c {
	case 0:
		return "quitControlLink"
	case 1:
		return "errorControlLinkAccept"
	case 2:
		return "errorControlLinkRead"
	case 3:
		return "errorControlLinkHandler"
	default:
		return "unknown"
	}
}
