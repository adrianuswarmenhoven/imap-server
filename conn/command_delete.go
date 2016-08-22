package conn

func cmdDelete(args commandArgs, c *Conn) {
	if !c.assertAuthenticated(args.ID()) {
		return
	}
	err := c.User.DeleteMailboxByName(args.Arg(0))
	if err != nil {
		c.writeResponse(args.ID(), "NO DELETE Failure: can't delete mailbox with that name")
		return
	}
	c.writeResponse(args.ID(), "OK DELETE Completed")
}
