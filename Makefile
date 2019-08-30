.PHONY: all shell tcp_listener udp_listener clean

all: shell tcp_listener udp_listener

shell: 
	cd cmd/shell; make build

shell-windows:
	cd cmd/shell; make windows

shell-linux:
	cd cmd/shell; make linux

shell-darwin:
	cd cmd/shell; make darwin

tcp_listener: 
	cd test/cmd/tcp_listener; make build

tcp_listener-windows:
	cd test/cmd/tcp_listener; make windows

tcp_listener-linux:
	cd test/cmd/tcp_listener; make linux

tcp_listener-darwin:
	cd test/cmd/tcp_listener; make darwin

udp_listener: 
	cd test/cmd/udp_listener; make build

udp_listener-windows:
	cd test/cmd/udp_listener; make windows

udp_listener-linux:
	cd test/cmd/udp_listener; make linux

udp_listener-darwin: 
	cd test/cmd/udp_listener; make darwin

clean: clean-shell clean-tcp_listener clean-udp_listener

clean-shell:
	cd cmd/shell; make clean

clean-tcp_listener:
	cd test/cmd/tcp_listener; make clean

clean-udp_listener:
	cd test/cmd/udp_listener; make clean
