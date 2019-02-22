package proxyserver

import (
	"io"
	"log"
	"net"
	"errors"
)

type Server struct {
	LocalAddr          string
	RemoteAddrs        []string
	remoteConns        []*net.Conn
	currentRemoteIndex int
}

func (s *Server) Start() error {
	log.Printf("Creating remote connections for local address %s", s.LocalAddr)
	for _, remoteAddr := range s.RemoteAddrs {
		s.addRemoteConnection(remoteAddr)
	}

	if len(s.remoteConns) == 0 {
		return errors.New("no available remote connections")
	}

	log.Printf("Create listener on local address %s", s.LocalAddr)
	listener, err := net.Listen("tcp", s.LocalAddr)
	if err != nil {
		return err
	}

	log.Println("Run listener infinite loop")
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go s.handleClient(conn)
	}
}

func (s *Server) getNextRemoteConnection() *net.Conn {
	s.currentRemoteIndex ++
	if s.currentRemoteIndex > len(s.remoteConns)-1 {
		s.currentRemoteIndex = 0
	}

	log.Printf("Get next remote connection with id %d", s.currentRemoteIndex)

	return s.remoteConns[s.currentRemoteIndex]
}

func (s *Server) addRemoteConnection(remoteAddr string) {
	log.Printf("Create and add connection to remote address %s", remoteAddr)

	remoteConn, err := net.Dial("tcp", remoteAddr)
	if err != nil {
		log.Printf("Can not connect to address %s %s", remoteAddr, err.Error())
		return
	}

	s.remoteConns = append(s.remoteConns, &remoteConn)
}

func (s *Server) handleClient(conn net.Conn) {
	log.Printf("Handle connection from new client %s", conn.RemoteAddr())
	remoteConn := *s.getNextRemoteConnection()

	go s.copyData(conn, remoteConn)
	s.copyData(remoteConn, conn)
}

func (s *Server) copyData(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Panic("Can not copy data. " + err.Error())
	}
}
