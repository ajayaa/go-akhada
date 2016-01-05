package main

import (
	"log"
	"net"
)

type node struct {
	member_ship_list  member
	address           net.addr
	received_messages queue
}

/*
send_message is used to send a message to a remote address
*/
func (n *node) send_message(msg *message, addr net.Addr) {
}

/*
process_message processes a message from queue and takes
appropriate action depending on type of message.
*/
func (n *node) process_message(msg *message) {
}

/*
inialize is called when a node is being introduced to the
cluster.
*/
func (n *node) initialize(seed net.Addr) {
}
