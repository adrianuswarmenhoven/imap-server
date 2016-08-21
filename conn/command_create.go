package conn

func cmdCreate(args commandArgs, c *Conn) {
	if !c.assertAuthenticated(args.ID()) {
		return
	}
	c.writeResponse(args.ID(), "OK CREATE Completed")
}
