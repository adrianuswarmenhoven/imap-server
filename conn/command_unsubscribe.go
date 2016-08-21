package conn

func cmdUnsubscribe(args commandArgs, c *Conn) {
	if !c.assertAuthenticated(args.ID()) {
		return
	}
	c.writeResponse(args.ID(), "OK UNSUBSCRIBE completed")
}
