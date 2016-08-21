package conn

func cmdCreate(args commandArgs, c *Conn) {
	if !c.assertAuthenticated(args.ID()) {
		return
	}
	_, err := c.User.CreateMailboxByName(args.Arg(0))
	if err != nil {
		c.writeResponse(args.ID(), "NO create failure: can't create mailbox with that name")
		return
	}
	c.writeResponse(args.ID(), "OK CREATE Completed")
}
