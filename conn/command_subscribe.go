package conn

func cmdSubscribe(args commandArgs, c *Conn) {
	if !c.assertAuthenticated(args.ID()) {
		return
	}
	c.writeResponse(args.ID(), "OK SUBSCRIBE completed")
}
