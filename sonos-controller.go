package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var sonosIP string
var command string

var commands = map[string]func(){
	"play":     play,
	"pause":    pause,
	"next":     next,
	"stop":     stop,
	"previous": previous,
}

func init() {
	flag.StringVar(&sonosIP, "sonos-ip", "0.0.0.0", "The IP Address of the target Sonos device")
	flag.StringVar(&command, "command", "play", "The command to execute, must be one of play, pause, next, stop, previous")
}

func main() {
	flag.Parse()
	if commands[command] != nil {
		commands[command]()
	} else {
		fmt.Println("Invalid command provided, see -help")
	}
}

func pause() {
	exec("Pause")
}

func play() {
	exec("Play")
}

func next() {
	exec("Next")
}

func stop() {
	exec("Stop")
}

func previous() {
	exec("Previous")
}

func exec(cmd string) {
	log.Println("Executing", cmd)
	body := fmt.Sprintf("<s:Envelope xmlns:s=\"http://schemas.xmlsoap.org/soap/envelope/\" s:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\"><s:Body><u:%v xmlns:u=\"urn:schemas-upnp-org:service:AVTransport:1\"><InstanceID>0</InstanceID><Speed>1</Speed></u:%v></s:Body></s:Envelope>", cmd, cmd)
	client := &http.Client{}
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%v:1400/MediaRenderer/AVTransport/Control", sonosIP), bytes.NewBuffer([]byte(body)))

	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("SOAPACTION", fmt.Sprintf("urn:schemas-upnp-org:service:AVTransport:1#%v", cmd))
	_, err = client.Do(req)
	if err != nil {
		log.Println(err)
	}
}
