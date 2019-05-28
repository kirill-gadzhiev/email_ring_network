package main

import (
	"encoding/json"
	"fmt"
	"github.com/Equanox/gotron"
	"os"
)

const (
	MESSAGE_RECEIVED          = "MESSAGE_RECEIVED_GOTRON_EVENT"
	MESSAGE_SEND 	          = "MESSAGE_SEND_GOTRON_EVENT"
	SET_USER 		          = "SET_USER_GOTRON_EVENT"
	SET_USER_RESPONSE         = "SET_USER_RESPONSE_GOTRON_EVENT"

	SET_OLD_LETTERS			  = "SET_OLD_LETTERS_GOTRON_EVENT"

	NETWORK_STATUS            = "NETWORK_STATUS_GOTRON_EVENT"
	INTERFACE_READY		      = "INTERFACE_READY_GOTRON_EVENT"
	CLOSE_WINDOW              = "CLOSE_WINDOW_GOTRON_EVENT"

	COM_PORTS_SETTINGS        = "COM_PORTS_SETTINGS_GOTRON_EVENT"
	COM_PORTS_SETTINGS_CHANGE = "COM_PORTS_SETTINGS_CHANGE_GOTRON_EVENT"
)

type User struct {
	Email string `json:"email"`
}

type UserEvent struct {
	Event string `json:"event"`
	User User `json:"user"`
}

type Letter struct {
	ID string `json:"id"`
	Author string `json:"author"`
	Responder string `json:"responder"`
	Message string `json:"message"`
	Date int64 `json:"date"`
	CheckedSubEvent bool `json:"checkedSubEvent"`
	Checked bool `json:"checked"`
}

type LetterEvent struct {
	Event string `json:"event"`
	Letter Letter `json:"letter"`
}

type OldLettersEvent struct {
	*gotron.Event
	Letters []Letter `json:"letters"`
}

type NetworkStatus struct {
	Connection bool `json:"connection"`
	AvailableUsers []User `json:"availableUsers"`
}

type NetworkStatusEvent struct {
	*gotron.Event
	NetworkStatus NetworkStatus `json:"networkStatus"`
}

type SetUserResponseEvent struct {
	*gotron.Event
	Status bool `json:"status"`
}

type MessageReceivedEvent struct {
	*gotron.Event
	Letter Letter `json:"letter"`
}

type ComPort struct {
	Name string `json:"name"`
	Speed int `json:"speed"`
}

type ComPortsSettings struct {
	In ComPort `json:"inCom"`
	Out ComPort `json:"outCom"`
	Ports []string `json:"ports"`
}

type ComPortsSettingsEvent struct {
	*gotron.Event
	ComPortsSettings ComPortsSettings `json:"comPortsSettings"`
}

type SomeEvent interface {
	GetEvent() string
}

func (e UserEvent) GetEvent() string {
	return e.Event
}
func (e LetterEvent) GetEvent() string {
	return e.Event
}
func (e NetworkStatusEvent) GetEvent() string {
	return e.Event.Event
}
func (e ComPortsSettingsEvent) GetEvent() string {
	return e.Event.Event
}
func (e MessageReceivedEvent) GetEvent() string {
	return e.Event.Event
}
func (e SetUserResponseEvent) GetEvent() string {
	return e.Event.Event
}
func (e OldLettersEvent) GetEvent() string {
	return e.Event.Event
}



func sendNetworkStatusEvent(connection bool, users []string, window *gotron.BrowserWindow) {
	availableUsers := []User{}
	for _, user := range users {
		availableUsers = append(availableUsers, User{Email:user})
	}
	networkStatus := NetworkStatus{Connection: connection, AvailableUsers: availableUsers}
	networkStatusEvent := &NetworkStatusEvent{Event: &gotron.Event{Event:NETWORK_STATUS}, NetworkStatus: networkStatus}
	window.Send(networkStatusEvent)
}

func sendLetterReceivedEvent(letter Letter, window *gotron.BrowserWindow) {
	messageReceivedEvent := &MessageReceivedEvent{Event: &gotron.Event{Event:MESSAGE_RECEIVED}, Letter: letter}
	window.Send(messageReceivedEvent)
}

func sendComPortsSettingsEvent(settings ComPortsSettings, window *gotron.BrowserWindow) {
	messageReceivedEvent := &ComPortsSettingsEvent{
		&gotron.Event{Event:COM_PORTS_SETTINGS},
		settings,
	}
	window.Send(messageReceivedEvent)
}


func initializeInterface(core2interface <-chan SomeEvent, interface2core chan<- SomeEvent, availableUsers []string, frontPort *DataLinkLayer) {
	//defer frontPort.terminate()
	// Create a new browser window instance
	window, err := gotron.New("interface/build")
	if err != nil {
		panic(err)
	}

	// Alter default window size and window title.
	window.WindowOptions.Width = 800
	window.WindowOptions.Height = 600
	window.WindowOptions.Title = "3kMail"
	window.WindowOptions.MinWidth = 600
	window.WindowOptions.MinHeight = 400
	// window.WindowOptions.TitleBarStyle =  "customButtonsOnHover"

	currentUser := ""

	window.On(&gotron.Event{Event:SET_USER}, func(data []byte) {
		userEvent := UserEvent{}
		err := json.Unmarshal(data, &userEvent)
		if err != nil {
			fmt.Println(err)
			return
		}
		interface2core <- userEvent
		fmt.Println("User email: ", userEvent.User.Email)

		setUserResponse := SetUserResponseEvent{Event: &gotron.Event{Event:SET_USER_RESPONSE}, Status:true}
		for _, user := range availableUsers {
			if userEvent.User.Email == user {
				setUserResponse.Status = false
				break
			}
		}

		if setUserResponse.Status {
			currentUser = userEvent.User.Email
		}
		window.Send(setUserResponse)
	})

	window.On(&gotron.Event{Event:MESSAGE_SEND}, func(data []byte) {
		letterEvent := LetterEvent{}
		err := json.Unmarshal(data, &letterEvent)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Letter message: ", letterEvent.Letter.Message)
		err = CreateOrUpdateUserLetter(currentUser, letterEvent.Letter)
		if err != nil {
			fmt.Println(err)
		}
		frontPort.sendLetter(letterEvent.Letter)
	})

	window.On(&gotron.Event{Event: CLOSE_WINDOW}, func(data []byte) {
		frontPort.terminate()
		os.Exit(0)
	})

	window.On(&gotron.Event{Event:INTERFACE_READY}, func(data []byte) {
		sendNetworkStatusEvent( false, availableUsers, window)
		settings := ComPortsSettings{
			In:ComPort{"COM1", 115200},
			Out:ComPort{"COM2", 115200},
			Ports:[]string{"COM1", "COM2", "COM3"},
		}
		sendComPortsSettingsEvent(settings, window)
	})

	window.On(&gotron.Event{Event:COM_PORTS_SETTINGS_CHANGE}, func(data []byte) {
		event := ComPortsSettingsEvent{}
		err := json.Unmarshal(data, &event)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println("Here we can change com ports settings", event.ComPortsSettings)
	})

	// когда уже имеем всех юзеров, отправляем доступных


	done, err := window.Start()
	if err != nil {
		panic(err)
	}


	// Open dev tools must be used after window.Start
	//window.OpenDevTools()

	for event := range core2interface {
		if event.GetEvent() == NETWORK_STATUS {
			castedEvent, ok := event.(NetworkStatusEvent)
			if !ok {
				continue
			}
			window.Send(castedEvent)
			continue
		}
		if event.GetEvent() == MESSAGE_RECEIVED {
			fmt.Println("message event", event)
			castedEvent, ok := event.(MessageReceivedEvent)
			if !ok {
				continue
			}
			err = CreateOrUpdateUserLetter(currentUser, castedEvent.Letter)
			if err != nil {
				fmt.Println(err)
			}
			window.Send(castedEvent)
			continue
		}
		if event.GetEvent() == SET_OLD_LETTERS {
			fmt.Println("old letters event", event)
			castedEvent, ok := event.(OldLettersEvent)
			fmt.Println("old letters event casted", event)
			if !ok {
				continue
			}
			window.Send(castedEvent)
			continue
		}
	}

	// Wait for the application to close
	<-done
}
