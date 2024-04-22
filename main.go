/*
이 서버는 간단한 TCP 서버를 구현합니다.
클라이언트의 연결을 수락하고 각 연결에 대해 처리합니다.
*/

package main

import (
	"log/slog"
	"net"
)

const defaultListenAddr = ":5001"

// Config는 서버 설정을 나타냅니다.
type Config struct {
	ListenAddr string // 서버가 리스닝할 주소
}

// Server는 TCP 서버를 나타냅니다.
type Server struct {
	Config              // 서버 설정
	ln     net.Listener // 네트워크 연결 수신 대기
}

// NewServer는 새로운 서버 인스턴스를 생성합니다.
func NewServer(cfg Config) *Server {
	if len(cfg.ListenAddr) == 0 {
		cfg.ListenAddr = defaultListenAddr
	}
	return &Server{
		Config: cfg,
	}
}

// Start는 서버를 시작합니다.
func (s *Server) Start() error {
	ln, err := net.Listen("tcp", s.ListenAddr)
	if err != nil {
		return err
	}
	s.ln = ln

	return s.acceptLoop()
}

// acceptLoop는 클라이언트의 연결을 수락하는 무한 루프입니다.
func (s *Server) acceptLoop() error {
	for {
		conn, err := s.ln.Accept()
		if err != nil {
			slog.Error("accep error", "err", err)
			continue
		}
		go s.handleConn(conn)
	}

}

// handleConn은 클라이언트 연결을 처리합니다.
func (s *Server) handleConn(conn net.Conn) {

}

func main() {

}

// https://youtu.be/LMrxfWB6sbQ?si=Z6txQb2xBSojIbBI
