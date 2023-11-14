package http

import (
	"log"
	"net"
)

const (
	S_LISTENING = 0x00000
	S_OFF       = 0x00001
)

const (
	PROT_TCP   = PROT_TCPv4
	PROT_TCPv4 = "tcp4"
	PROT_TCPv6 = "tcp6"
	PROT_UDP   = PROT_UDPv4
	PROT_UDPv4 = "udp4"
	PROT_UDPv6 = "udp6"
)

type Server struct {
	m_listener net.Listener
	m_protocol string

	Status int
}

func (self Server) StartListening(prot string, addr string) {
	log.Printf("Server is starting to listen at address '%s' using '%s'...\n", addr, prot)

	self.Status = S_LISTENING
	listener, err := net.Listen(prot, addr)

	if err != nil {
		self.StopListening()
		log.Println(err)
		return
	}
	defer self.StopListening()

	self.m_listener = listener
	self.m_protocol = prot

	for {
		conn, err := self.m_listener.Accept()
		if err != nil {
			log.Println(err)
			conn.Close()
			continue
		}
		go ProcessConnection(conn)
	}
}

func (self Server) StopListening() {
	defer self.m_listener.Close()

	log.Printf("Server is stop listening...\n")
	self.Status = S_OFF
	self.m_protocol = ""
}

func ProcessConnection(conn net.Conn) {
	defer conn.Close()

	data := make([]byte, 1024*4)
	n, err := conn.Read(data)
	if n == 0 || err != nil {
		log.Println("Read error:", err)
	}

	source := string(data)
	log.Println(conn.RemoteAddr(), "is connecting to me!")
	log.Println("Client is sending:", source)

	message := "Hello, I am a server and I recieved your message!"
	conn.Write([]byte(message))
}
