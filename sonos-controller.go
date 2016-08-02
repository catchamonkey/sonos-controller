package main

import (
	"bytes"
	"flag"
	"fmt"
	"log"
	"net/http"
)

var sonosIP string
var pauseCommand bool
var playCommand bool

func init() {
	flag.StringVar(&sonosIP, "sonos-ip", "0.0.0.0", "The IP Address of the target Sonos device")
	flag.BoolVar(&pauseCommand, "pause", false, "Invoke the pause Command")
	flag.BoolVar(&playCommand, "play", false, "Invoke the Play Command")
}

func main() {
	flag.Parse()
	if pauseCommand {
		pause()
	}
	if playCommand {
		play()
	}
}

func pause() {
	log.Println("Pausing")
	body := "<s:Envelope xmlns:s=\"http://schemas.xmlsoap.org/soap/envelope/\" s:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\"><s:Body><u:Pause xmlns:u=\"urn:schemas-upnp-org:service:AVTransport:1\"><InstanceID>0</InstanceID><Speed>1</Speed></u:Pause></s:Body></s:Envelope>"
	client := &http.Client{}
	// build a new request, but not doing the POST yet
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%v:1400/MediaRenderer/AVTransport/Control", sonosIP), bytes.NewBuffer([]byte(body)))

	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("SOAPACTION", "urn:schemas-upnp-org:service:AVTransport:1#Pause")
	_, err = client.Do(req)
	if err != nil {
		log.Println(err)
	}
}

func play() {
	log.Println("Playing")
	body := "<s:Envelope xmlns:s=\"http://schemas.xmlsoap.org/soap/envelope/\" s:encodingStyle=\"http://schemas.xmlsoap.org/soap/encoding/\"><s:Body><u:Play xmlns:u=\"urn:schemas-upnp-org:service:AVTransport:1\"><InstanceID>0</InstanceID><Speed>1</Speed></u:Play></s:Body></s:Envelope>"
	client := &http.Client{}
	// build a new request, but not doing the POST yet
	req, err := http.NewRequest("POST", fmt.Sprintf("http://%v:1400/MediaRenderer/AVTransport/Control", sonosIP), bytes.NewBuffer([]byte(body)))

	if err != nil {
		log.Println(err)
	}

	req.Header.Add("Content-Type", "text/xml; charset=utf-8")
	req.Header.Add("SOAPACTION", "urn:schemas-upnp-org:service:AVTransport:1#Play")
	_, err = client.Do(req)
	if err != nil {
		log.Println(err)
	}
}
