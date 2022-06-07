package utcp

import "testing"

/********************************************************
* @author: Ihc
* @date: 2022/6/7 0007 11:30
* @version: 1.0
* @description:
*********************************************************/

var tcpServer *TCPServer

func TestNewTCPServerWithConf(t *testing.T) {
	defaultConfigure := GetDefaultConfigure()
	if defaultConfigure.Host != "localhost" {
		t.Fatal("GetDefaultConfigure Err")
	}
	tcpServer = NewTCPServerWithConf(defaultConfigure)
	if tcpServer == nil {
		t.Fatal("NewTCPServerWithConf Err")
	}
}

func TestTCPServer_LocalAddr(t *testing.T) {
	if tcpAddr := tcpServer.LocalAddr(); tcpAddr != "127.0.0.1:9090" {
		t.Fatal("LocalAddr Err")
	}
}

func TestTCPServer_Start(t *testing.T) {
	go tcpServer.Start()
	if tcpServer.running {
		t.Fatal("Start Err")
	}
	tcpCli := GetDefaultTCPClient()
	tcpCli.Send([]byte("Send Message"))
}
